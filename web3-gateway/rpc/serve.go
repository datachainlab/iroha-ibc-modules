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
	"google.golang.org/grpc"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/acm"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/config"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/api"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/iroha/db/postgres"
	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/keyring"
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

	keyStore := keyring.NewKeyStore()
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

	irohaDBTransactor, err := postgres.NewTransactor(
		cfg.Iroha.Database.Postgres.User,
		cfg.Iroha.Database.Postgres.Password,
		cfg.Iroha.Database.Postgres.Host,
		cfg.Iroha.Database.Postgres.Port,
		cfg.Iroha.Database.Postgres.Database,
	)
	if err != nil {
		return err
	}

	// TODO configurable
	logger, err := logconfig.New().NewLogger()
	if err != nil {
		return err
	}

	ethService := NewEthService(
		accountState,
		keyStore,
		irohaApiClient,
		irohaDBTransactor,
		logger,
		cfg.EVM.Querier,
	)

	web3Server := NewHTTPServer(ethService)

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

		if err := irohaDBTransactor.Close(); err != nil {
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
