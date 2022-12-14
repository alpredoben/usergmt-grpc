// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.12.4
// source: proto/usermgmt.proto

package proto

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

// UserManagementServicesClient is the client API for UserManagementServices service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type UserManagementServicesClient interface {
	CreateNewUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error)
	GetUsers(ctx context.Context, in *GetUsersParamRequest, opts ...grpc.CallOption) (*UserListResponses, error)
}

type userManagementServicesClient struct {
	cc grpc.ClientConnInterface
}

func NewUserManagementServicesClient(cc grpc.ClientConnInterface) UserManagementServicesClient {
	return &userManagementServicesClient{cc}
}

func (c *userManagementServicesClient) CreateNewUser(ctx context.Context, in *UserRequest, opts ...grpc.CallOption) (*UserResponse, error) {
	out := new(UserResponse)
	err := c.cc.Invoke(ctx, "/usermgmt.UserManagementServices/CreateNewUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *userManagementServicesClient) GetUsers(ctx context.Context, in *GetUsersParamRequest, opts ...grpc.CallOption) (*UserListResponses, error) {
	out := new(UserListResponses)
	err := c.cc.Invoke(ctx, "/usermgmt.UserManagementServices/GetUsers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// UserManagementServicesServer is the server API for UserManagementServices service.
// All implementations must embed UnimplementedUserManagementServicesServer
// for forward compatibility
type UserManagementServicesServer interface {
	CreateNewUser(context.Context, *UserRequest) (*UserResponse, error)
	GetUsers(context.Context, *GetUsersParamRequest) (*UserListResponses, error)
	mustEmbedUnimplementedUserManagementServicesServer()
}

// UnimplementedUserManagementServicesServer must be embedded to have forward compatible implementations.
type UnimplementedUserManagementServicesServer struct {
}

func (UnimplementedUserManagementServicesServer) CreateNewUser(context.Context, *UserRequest) (*UserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateNewUser not implemented")
}
func (UnimplementedUserManagementServicesServer) GetUsers(context.Context, *GetUsersParamRequest) (*UserListResponses, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUsers not implemented")
}
func (UnimplementedUserManagementServicesServer) mustEmbedUnimplementedUserManagementServicesServer() {
}

// UnsafeUserManagementServicesServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to UserManagementServicesServer will
// result in compilation errors.
type UnsafeUserManagementServicesServer interface {
	mustEmbedUnimplementedUserManagementServicesServer()
}

func RegisterUserManagementServicesServer(s grpc.ServiceRegistrar, srv UserManagementServicesServer) {
	s.RegisterService(&UserManagementServices_ServiceDesc, srv)
}

func _UserManagementServices_CreateNewUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServicesServer).CreateNewUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgmt.UserManagementServices/CreateNewUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServicesServer).CreateNewUser(ctx, req.(*UserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _UserManagementServices_GetUsers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUsersParamRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(UserManagementServicesServer).GetUsers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/usermgmt.UserManagementServices/GetUsers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(UserManagementServicesServer).GetUsers(ctx, req.(*GetUsersParamRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// UserManagementServices_ServiceDesc is the grpc.ServiceDesc for UserManagementServices service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var UserManagementServices_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "usermgmt.UserManagementServices",
	HandlerType: (*UserManagementServicesServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateNewUser",
			Handler:    _UserManagementServices_CreateNewUser_Handler,
		},
		{
			MethodName: "GetUsers",
			Handler:    _UserManagementServices_GetUsers_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/usermgmt.proto",
}
