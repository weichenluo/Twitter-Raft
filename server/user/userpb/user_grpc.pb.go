// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: user.proto

package userpb

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

// UserServiceClient is the client API for UserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserServiceClient interface {
	Add(ctx context.Context, in *AddUserParameters, opts ...grpc.CallOption) (*User, error)
	GetUserByNamePasswrod(ctx context.Context, in *LoginDetails, opts ...grpc.CallOption) (*User, error)
	FollowUser(ctx context.Context, in *FollowerParameters, opts ...grpc.CallOption) (*Status, error)
	GetUserFollowingByName(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*Login, error)
}

type userServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewUserServiceClient(cc grpc.ClientConnInterface) UserServiceClient {
	return &userServiceClient{cc}
}

func (c *userServiceClient) Add(ctx context.Context, in *AddUserParameters, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/userpb.UserService/Add", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserByNamePasswrod(ctx context.Context, in *LoginDetails, opts ...grpc.CallOption) (*User, error) {
	out := new(User)
	err := c.cc.Invoke(ctx, "/userpb.UserService/GetUserByNamePasswrod", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) FollowUser(ctx context.Context, in *FollowerParameters, opts ...grpc.CallOption) (*Status, error) {
	out := new(Status)
	err := c.cc.Invoke(ctx, "/userpb.UserService/FollowUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userServiceClient) GetUserFollowingByName(ctx context.Context, in *UserName, opts ...grpc.CallOption) (*Login, error) {
	out := new(Login)
	err := c.cc.Invoke(ctx, "/userpb.UserService/GetUserFollowingByName", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserServiceServer is the server API for UserService service.
// All implementations must embed UnimplementedUserServiceServer
// for forward compatibility
type UserServiceServer interface {
	Add(context.Context, *AddUserParameters) (*User, error)
	GetUserByNamePasswrod(context.Context, *LoginDetails) (*User, error)
	FollowUser(context.Context, *FollowerParameters) (*Status, error)
	GetUserFollowingByName(context.Context, *UserName) (*Login, error)
	mustEmbedUnimplementedUserServiceServer()
}

// UnimplementedUserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedUserServiceServer struct {
}

func (UnimplementedUserServiceServer) Add(context.Context, *AddUserParameters) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Add not implemented")
}
func (UnimplementedUserServiceServer) GetUserByNamePasswrod(context.Context, *LoginDetails) (*User, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserByNamePasswrod not implemented")
}
func (UnimplementedUserServiceServer) FollowUser(context.Context, *FollowerParameters) (*Status, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FollowUser not implemented")
}
func (UnimplementedUserServiceServer) GetUserFollowingByName(context.Context, *UserName) (*Login, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserFollowingByName not implemented")
}
func (UnimplementedUserServiceServer) mustEmbedUnimplementedUserServiceServer() {}

// UnsafeUserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserServiceServer will
// result in compilation errors.
type UnsafeUserServiceServer interface {
	mustEmbedUnimplementedUserServiceServer()
}

func RegisterUserServiceServer(s grpc.ServiceRegistrar, srv UserServiceServer) {
	s.RegisterService(&UserService_ServiceDesc, srv)
}

func _UserService_Add_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddUserParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).Add(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.UserService/Add",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).Add(ctx, req.(*AddUserParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserByNamePasswrod_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginDetails)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserByNamePasswrod(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.UserService/GetUserByNamePasswrod",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserByNamePasswrod(ctx, req.(*LoginDetails))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_FollowUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FollowerParameters)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).FollowUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.UserService/FollowUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).FollowUser(ctx, req.(*FollowerParameters))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserService_GetUserFollowingByName_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserName)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserServiceServer).GetUserFollowingByName(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/userpb.UserService/GetUserFollowingByName",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserServiceServer).GetUserFollowingByName(ctx, req.(*UserName))
	}
	return interceptor(ctx, in, info, handler)
}

// UserService_ServiceDesc is the grpc.ServiceDesc for UserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "userpb.UserService",
	HandlerType: (*UserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Add",
			Handler:    _UserService_Add_Handler,
		},
		{
			MethodName: "GetUserByNamePasswrod",
			Handler:    _UserService_GetUserByNamePasswrod_Handler,
		},
		{
			MethodName: "FollowUser",
			Handler:    _UserService_FollowUser_Handler,
		},
		{
			MethodName: "GetUserFollowingByName",
			Handler:    _UserService_GetUserFollowingByName_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "user.proto",
}
