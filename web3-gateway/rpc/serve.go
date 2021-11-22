package rpc

import (
	"fmt"
	"net/http"

	"github.com/hyperledger/burrow/logging/logconfig"
	"github.com/hyperledger/burrow/process"
	"github.com/hyperledger/burrow/rpc/lib/server"
	"github.com/hyperledger/burrow/rpc/web3"

	"github.com/datachainlab/iroha-ibc-modules/web3-gateway/config"
)

func Serve(cfg *config.Config) (*http.Server, error) {
	listener, err := process.ListenerFromAddress(
		fmt.Sprintf("%s:%v", cfg.Gateway.Rpc.Host, cfg.Gateway.Rpc.Port),
	)
	if err != nil {
		return nil, err
	}

	// TODO configurable
	logConf := logconfig.New()
	logger, err := logConf.NewLogger()
	if err != nil {
		return nil, err
	}

	srv, err := server.StartHTTPServer(listener, web3.NewServer(NewEthService()), logger)
	if err != nil {
		return nil, err
	}

	return srv, err
}
