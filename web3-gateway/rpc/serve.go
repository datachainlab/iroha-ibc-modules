package rpc

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/hyperledger/burrow/core"
	"github.com/hyperledger/burrow/logging/logconfig"
	"github.com/hyperledger/burrow/process"
	"github.com/hyperledger/burrow/rpc/lib/server"
	"github.com/hyperledger/burrow/rpc/web3"
	"google.golang.org/grpc"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/acm"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/config"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/evm"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/api"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/postgres"
)

func Serve(cfg *config.Config) error {
	listener, err := process.ListenerFromAddress(
		fmt.Sprintf("%s:%v", cfg.Gateway.Rpc.Host, cfg.Gateway.Rpc.Port),
	)
	if err != nil {
		return err
	}

	accountDB, err := acm.NewMemDB()
	if err != nil {
		return err
	}

	keyStore := acm.NewKeyStore()
	accountState := acm.NewAccountState(accountDB)
	for _, account := range cfg.Accounts {
		if err = accountState.Add(account.ID, account.PrivateKey); err != nil {
			return err
		}
		if err = keyStore.Set(account.ID, account.PrivateKey); err != nil {
			return err
		}
	}

	grpConn, err := grpc.Dial(
		fmt.Sprintf("%s:%v", cfg.Iroha.Api.Host, cfg.Iroha.Api.Port),
		// TODO configurable
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)
	if err != nil {
		return err
	}

	irohaApiClient := api.NewClient(
		grpConn,
		api.CommandTimeout(cfg.Iroha.Api.CommandTimeout),
		api.QueryTimeout(cfg.Iroha.Api.QueryTimeout),
	)

	irohaDBClient, dbConn, err := postgres.NewClient(
		cfg.Iroha.Database.Postgres.User,
		cfg.Iroha.Database.Postgres.Password,
		cfg.Iroha.Database.Postgres.Host,
		cfg.Iroha.Database.Postgres.Port,
		cfg.Iroha.Database.Postgres.Database,
	)
	if err != nil {
		return err
	}

	irohaClient := iroha.New(irohaApiClient, irohaDBClient)

	// TODO configurable
	logger, err := logconfig.New().NewLogger()
	if err != nil {
		return err
	}

	evm.RegisterCallContext(irohaApiClient, irohaDBClient, cfg.EVM.Querier, keyStore)

	web3Server := web3.NewServer(
		NewEthService(
			accountState,
			keyStore,
			irohaClient,
			logger,
			cfg.EVM.Querier,
		),
	)

	srv, err := server.StartHTTPServer(
		listener,
		web3Server,
		logger.WithScope("Web3"),
	)
	if err != nil {
		return err
	}

	return trapSignal(func() error {
		ctx, cancel := context.WithTimeout(context.Background(), core.ServerShutdownTimeout)
		defer cancel()

		if err := srv.Shutdown(ctx); err != nil {
			return err
		}

		if err := grpConn.Close(); err != nil {
			return err
		}

		if err := dbConn.Close(); err != nil {
			return err
		}

		return nil
	})
}

func trapSignal(shutdown func() error) error {
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)
	for {
		select {
		case <-shutdownCh:
			return shutdown()
		}
	}
}
