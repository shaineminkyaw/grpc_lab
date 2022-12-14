// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: login_service.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// LoginServiceClient is the client API for LoginService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type LoginServiceClient interface {
	Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (LoginService_LoginClient, error)
}

type loginServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewLoginServiceClient(cc grpc.ClientConnInterface) LoginServiceClient {
	return &loginServiceClient{cc}
}

func (c *loginServiceClient) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (LoginService_LoginClient, error) {
	stream, err := c.cc.NewStream(ctx, &LoginService_ServiceDesc.Streams[0], "/login.LoginService/Login", opts...)
	if err != nil {
		return nil, err
	}
	x := &loginServiceLoginClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type LoginService_LoginClient interface {
	Recv() (*LoginResponse, error)
	grpc.ClientStream
}

type loginServiceLoginClient struct {
	grpc.ClientStream
}

func (x *loginServiceLoginClient) Recv() (*LoginResponse, error) {
	m := new(LoginResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// LoginServiceServer is the server API for LoginService service.
// All implementations must embed UnimplementedLoginServiceServer
// for forward compatibility
type LoginServiceServer interface {
	Login(*LoginRequest, LoginService_LoginServer) error
	mustEmbedUnimplementedLoginServiceServer()
}

// UnimplementedLoginServiceServer must be embedded to have forward compatible implementations.
type UnimplementedLoginServiceServer struct {
}

func (UnimplementedLoginServiceServer) Login(*LoginRequest, LoginService_LoginServer) error {
	return status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedLoginServiceServer) mustEmbedUnimplementedLoginServiceServer() {}

// UnsafeLoginServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to LoginServiceServer will
// result in compilation errors.
type UnsafeLoginServiceServer interface {
	mustEmbedUnimplementedLoginServiceServer()
}

func RegisterLoginServiceServer(s grpc.ServiceRegistrar, srv LoginServiceServer) {
	s.RegisterService(&LoginService_ServiceDesc, srv)
}

func _LoginService_Login_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(LoginRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(LoginServiceServer).Login(m, &loginServiceLoginServer{stream})
}

type LoginService_LoginServer interface {
	Send(*LoginResponse) error
	grpc.ServerStream
}

type loginServiceLoginServer struct {
	grpc.ServerStream
}

func (x *loginServiceLoginServer) Send(m *LoginResponse) error {
	return x.ServerStream.SendMsg(m)
}

// LoginService_ServiceDesc is the grpc.ServiceDesc for LoginService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var LoginService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "login.LoginService",
	HandlerType: (*LoginServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "Login",
			Handler:       _LoginService_Login_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "login_service.proto",
}
