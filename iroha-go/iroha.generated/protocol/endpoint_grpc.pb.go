// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package protocol

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// CommandServiceV1Client is the client API for CommandServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CommandServiceV1Client interface {
	Torii(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ListTorii(ctx context.Context, in *TxList, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Status(ctx context.Context, in *TxStatusRequest, opts ...grpc.CallOption) (*ToriiResponse, error)
	StatusStream(ctx context.Context, in *TxStatusRequest, opts ...grpc.CallOption) (CommandServiceV1_StatusStreamClient, error)
}

type commandServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewCommandServiceV1Client(cc grpc.ClientConnInterface) CommandServiceV1Client {
	return &commandServiceV1Client{cc}
}

func (c *commandServiceV1Client) Torii(ctx context.Context, in *Transaction, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/iroha.protocol.CommandService_v1/Torii", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceV1Client) ListTorii(ctx context.Context, in *TxList, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/iroha.protocol.CommandService_v1/ListTorii", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceV1Client) Status(ctx context.Context, in *TxStatusRequest, opts ...grpc.CallOption) (*ToriiResponse, error) {
	out := new(ToriiResponse)
	err := c.cc.Invoke(ctx, "/iroha.protocol.CommandService_v1/Status", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *commandServiceV1Client) StatusStream(ctx context.Context, in *TxStatusRequest, opts ...grpc.CallOption) (CommandServiceV1_StatusStreamClient, error) {
	stream, err := c.cc.NewStream(ctx, &CommandServiceV1_ServiceDesc.Streams[0], "/iroha.protocol.CommandService_v1/StatusStream", opts...)
	if err != nil {
		return nil, err
	}
	x := &commandServiceV1StatusStreamClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type CommandServiceV1_StatusStreamClient interface {
	Recv() (*ToriiResponse, error)
	grpc.ClientStream
}

type commandServiceV1StatusStreamClient struct {
	grpc.ClientStream
}

func (x *commandServiceV1StatusStreamClient) Recv() (*ToriiResponse, error) {
	m := new(ToriiResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// CommandServiceV1Server is the server API for CommandServiceV1 service.
// All implementations should embed UnimplementedCommandServiceV1Server
// for forward compatibility
type CommandServiceV1Server interface {
	Torii(context.Context, *Transaction) (*emptypb.Empty, error)
	ListTorii(context.Context, *TxList) (*emptypb.Empty, error)
	Status(context.Context, *TxStatusRequest) (*ToriiResponse, error)
	StatusStream(*TxStatusRequest, CommandServiceV1_StatusStreamServer) error
}

// UnimplementedCommandServiceV1Server should be embedded to have forward compatible implementations.
type UnimplementedCommandServiceV1Server struct {
}

func (UnimplementedCommandServiceV1Server) Torii(context.Context, *Transaction) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Torii not implemented")
}
func (UnimplementedCommandServiceV1Server) ListTorii(context.Context, *TxList) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTorii not implemented")
}
func (UnimplementedCommandServiceV1Server) Status(context.Context, *TxStatusRequest) (*ToriiResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Status not implemented")
}
func (UnimplementedCommandServiceV1Server) StatusStream(*TxStatusRequest, CommandServiceV1_StatusStreamServer) error {
	return status.Errorf(codes.Unimplemented, "method StatusStream not implemented")
}

// UnsafeCommandServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CommandServiceV1Server will
// result in compilation errors.
type UnsafeCommandServiceV1Server interface {
	mustEmbedUnimplementedCommandServiceV1Server()
}

func RegisterCommandServiceV1Server(s grpc.ServiceRegistrar, srv CommandServiceV1Server) {
	s.RegisterService(&CommandServiceV1_ServiceDesc, srv)
}

func _CommandServiceV1_Torii_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Transaction)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceV1Server).Torii(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iroha.protocol.CommandService_v1/Torii",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceV1Server).Torii(ctx, req.(*Transaction))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandServiceV1_ListTorii_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxList)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceV1Server).ListTorii(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iroha.protocol.CommandService_v1/ListTorii",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceV1Server).ListTorii(ctx, req.(*TxList))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandServiceV1_Status_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TxStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CommandServiceV1Server).Status(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iroha.protocol.CommandService_v1/Status",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CommandServiceV1Server).Status(ctx, req.(*TxStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CommandServiceV1_StatusStream_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(TxStatusRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(CommandServiceV1Server).StatusStream(m, &commandServiceV1StatusStreamServer{stream})
}

type CommandServiceV1_StatusStreamServer interface {
	Send(*ToriiResponse) error
	grpc.ServerStream
}

type commandServiceV1StatusStreamServer struct {
	grpc.ServerStream
}

func (x *commandServiceV1StatusStreamServer) Send(m *ToriiResponse) error {
	return x.ServerStream.SendMsg(m)
}

// CommandServiceV1_ServiceDesc is the grpc.ServiceDesc for CommandServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CommandServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iroha.protocol.CommandService_v1",
	HandlerType: (*CommandServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Torii",
			Handler:    _CommandServiceV1_Torii_Handler,
		},
		{
			MethodName: "ListTorii",
			Handler:    _CommandServiceV1_ListTorii_Handler,
		},
		{
			MethodName: "Status",
			Handler:    _CommandServiceV1_Status_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StatusStream",
			Handler:       _CommandServiceV1_StatusStream_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "endpoint.proto",
}

// QueryServiceV1Client is the client API for QueryServiceV1 service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueryServiceV1Client interface {
	Find(ctx context.Context, in *Query, opts ...grpc.CallOption) (*QueryResponse, error)
	FetchCommits(ctx context.Context, in *BlocksQuery, opts ...grpc.CallOption) (QueryServiceV1_FetchCommitsClient, error)
}

type queryServiceV1Client struct {
	cc grpc.ClientConnInterface
}

func NewQueryServiceV1Client(cc grpc.ClientConnInterface) QueryServiceV1Client {
	return &queryServiceV1Client{cc}
}

func (c *queryServiceV1Client) Find(ctx context.Context, in *Query, opts ...grpc.CallOption) (*QueryResponse, error) {
	out := new(QueryResponse)
	err := c.cc.Invoke(ctx, "/iroha.protocol.QueryService_v1/Find", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryServiceV1Client) FetchCommits(ctx context.Context, in *BlocksQuery, opts ...grpc.CallOption) (QueryServiceV1_FetchCommitsClient, error) {
	stream, err := c.cc.NewStream(ctx, &QueryServiceV1_ServiceDesc.Streams[0], "/iroha.protocol.QueryService_v1/FetchCommits", opts...)
	if err != nil {
		return nil, err
	}
	x := &queryServiceV1FetchCommitsClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type QueryServiceV1_FetchCommitsClient interface {
	Recv() (*BlockQueryResponse, error)
	grpc.ClientStream
}

type queryServiceV1FetchCommitsClient struct {
	grpc.ClientStream
}

func (x *queryServiceV1FetchCommitsClient) Recv() (*BlockQueryResponse, error) {
	m := new(BlockQueryResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// QueryServiceV1Server is the server API for QueryServiceV1 service.
// All implementations should embed UnimplementedQueryServiceV1Server
// for forward compatibility
type QueryServiceV1Server interface {
	Find(context.Context, *Query) (*QueryResponse, error)
	FetchCommits(*BlocksQuery, QueryServiceV1_FetchCommitsServer) error
}

// UnimplementedQueryServiceV1Server should be embedded to have forward compatible implementations.
type UnimplementedQueryServiceV1Server struct {
}

func (UnimplementedQueryServiceV1Server) Find(context.Context, *Query) (*QueryResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Find not implemented")
}
func (UnimplementedQueryServiceV1Server) FetchCommits(*BlocksQuery, QueryServiceV1_FetchCommitsServer) error {
	return status.Errorf(codes.Unimplemented, "method FetchCommits not implemented")
}

// UnsafeQueryServiceV1Server may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueryServiceV1Server will
// result in compilation errors.
type UnsafeQueryServiceV1Server interface {
	mustEmbedUnimplementedQueryServiceV1Server()
}

func RegisterQueryServiceV1Server(s grpc.ServiceRegistrar, srv QueryServiceV1Server) {
	s.RegisterService(&QueryServiceV1_ServiceDesc, srv)
}

func _QueryServiceV1_Find_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Query)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServiceV1Server).Find(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/iroha.protocol.QueryService_v1/Find",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServiceV1Server).Find(ctx, req.(*Query))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueryServiceV1_FetchCommits_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(BlocksQuery)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(QueryServiceV1Server).FetchCommits(m, &queryServiceV1FetchCommitsServer{stream})
}

type QueryServiceV1_FetchCommitsServer interface {
	Send(*BlockQueryResponse) error
	grpc.ServerStream
}

type queryServiceV1FetchCommitsServer struct {
	grpc.ServerStream
}

func (x *queryServiceV1FetchCommitsServer) Send(m *BlockQueryResponse) error {
	return x.ServerStream.SendMsg(m)
}

// QueryServiceV1_ServiceDesc is the grpc.ServiceDesc for QueryServiceV1 service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueryServiceV1_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "iroha.protocol.QueryService_v1",
	HandlerType: (*QueryServiceV1Server)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Find",
			Handler:    _QueryServiceV1_Find_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "FetchCommits",
			Handler:       _QueryServiceV1_FetchCommits_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "endpoint.proto",
}
