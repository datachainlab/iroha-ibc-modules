package api

import (
	"time"

	"google.golang.org/grpc"

	"github.com/datachainlab/iroha-ibc-modules/iroha-go/command"
	"github.com/datachainlab/iroha-ibc-modules/iroha-go/query"
)

type ApiClient interface {
	command.CommandClient
	query.QueryClient

	SetCommandTimeout(time.Duration)
	SetQueryTimeout(time.Duration)
	SetGRPCOptions(opts ...grpc.CallOption)
}

type client struct {
	command.CommandClient
	query.QueryClient

	commandTimeout time.Duration
	queryTimeout   time.Duration
	grpcOpts       []grpc.CallOption
}

func (c client) SetCommandTimeout(timeout time.Duration) {
	c.commandTimeout = timeout
}

func (c client) SetQueryTimeout(timeout time.Duration) {
	c.queryTimeout = timeout
}

func (c client) SetGRPCOptions(opts ...grpc.CallOption) {
	c.grpcOpts = opts
}

type ClientOption func(ApiClient)

func CommandTimeout(t time.Duration) ClientOption {
	return func(c ApiClient) {
		c.SetCommandTimeout(time.Millisecond * t)
	}
}

func QueryTimeout(t time.Duration) ClientOption {
	return func(c ApiClient) {
		c.SetQueryTimeout(time.Millisecond * t)
	}
}

func GRPCOptions(opts ...grpc.CallOption) ClientOption {
	return func(c ApiClient) {
		c.SetGRPCOptions(opts...)
	}
}

func NewClient(conn *grpc.ClientConn, opts ...ClientOption) ApiClient {
	c := &client{
		commandTimeout: time.Millisecond * 5000,
		queryTimeout:   time.Millisecond * 5000,
		grpcOpts:       nil,
	}

	for _, opt := range opts {
		opt(c)
	}

	c.CommandClient = command.New(conn, c.commandTimeout, c.grpcOpts...)
	c.QueryClient = query.New(conn, c.queryTimeout, c.grpcOpts...)

	return c
}
