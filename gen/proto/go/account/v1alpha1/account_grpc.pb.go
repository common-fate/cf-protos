// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package accountv1alpha1

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

// AccountServiceClient is the client API for AccountService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type AccountServiceClient interface {
	Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error)
	InviteMember(ctx context.Context, in *InviteMemberRequest, opts ...grpc.CallOption) (*InviteMemberResponse, error)
	RemoveMember(ctx context.Context, in *RemoveMemberRequest, opts ...grpc.CallOption) (*RemoveMemberResponse, error)
	CheckForUpdates(ctx context.Context, in *CheckForUpdatesRequest, opts ...grpc.CallOption) (*CheckForUpdatesResponse, error)
	GetDeviceId(ctx context.Context, in *GetDeviceIdRequest, opts ...grpc.CallOption) (*GetDeviceIdResponse, error)
	// Authenticated informs the metadata service that the client has successfully
	// completed the login flow with their access broker
	Authenticated(ctx context.Context, in *AuthenticatedRequest, opts ...grpc.CallOption) (*AuthenticatedResponse, error)
}

type accountServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewAccountServiceClient(cc grpc.ClientConnInterface) AccountServiceClient {
	return &accountServiceClient{cc}
}

func (c *accountServiceClient) Signup(ctx context.Context, in *SignupRequest, opts ...grpc.CallOption) (*SignupResponse, error) {
	out := new(SignupResponse)
	err := c.cc.Invoke(ctx, "/account.v1alpha1.AccountService/Signup", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) InviteMember(ctx context.Context, in *InviteMemberRequest, opts ...grpc.CallOption) (*InviteMemberResponse, error) {
	out := new(InviteMemberResponse)
	err := c.cc.Invoke(ctx, "/account.v1alpha1.AccountService/InviteMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) RemoveMember(ctx context.Context, in *RemoveMemberRequest, opts ...grpc.CallOption) (*RemoveMemberResponse, error) {
	out := new(RemoveMemberResponse)
	err := c.cc.Invoke(ctx, "/account.v1alpha1.AccountService/RemoveMember", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) CheckForUpdates(ctx context.Context, in *CheckForUpdatesRequest, opts ...grpc.CallOption) (*CheckForUpdatesResponse, error) {
	out := new(CheckForUpdatesResponse)
	err := c.cc.Invoke(ctx, "/account.v1alpha1.AccountService/CheckForUpdates", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) GetDeviceId(ctx context.Context, in *GetDeviceIdRequest, opts ...grpc.CallOption) (*GetDeviceIdResponse, error) {
	out := new(GetDeviceIdResponse)
	err := c.cc.Invoke(ctx, "/account.v1alpha1.AccountService/GetDeviceId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *accountServiceClient) Authenticated(ctx context.Context, in *AuthenticatedRequest, opts ...grpc.CallOption) (*AuthenticatedResponse, error) {
	out := new(AuthenticatedResponse)
	err := c.cc.Invoke(ctx, "/account.v1alpha1.AccountService/Authenticated", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// AccountServiceServer is the server API for AccountService service.
// All implementations should embed UnimplementedAccountServiceServer
// for forward compatibility
type AccountServiceServer interface {
	Signup(context.Context, *SignupRequest) (*SignupResponse, error)
	InviteMember(context.Context, *InviteMemberRequest) (*InviteMemberResponse, error)
	RemoveMember(context.Context, *RemoveMemberRequest) (*RemoveMemberResponse, error)
	CheckForUpdates(context.Context, *CheckForUpdatesRequest) (*CheckForUpdatesResponse, error)
	GetDeviceId(context.Context, *GetDeviceIdRequest) (*GetDeviceIdResponse, error)
	// Authenticated informs the metadata service that the client has successfully
	// completed the login flow with their access broker
	Authenticated(context.Context, *AuthenticatedRequest) (*AuthenticatedResponse, error)
}

// UnimplementedAccountServiceServer should be embedded to have forward compatible implementations.
type UnimplementedAccountServiceServer struct {
}

func (UnimplementedAccountServiceServer) Signup(context.Context, *SignupRequest) (*SignupResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Signup not implemented")
}
func (UnimplementedAccountServiceServer) InviteMember(context.Context, *InviteMemberRequest) (*InviteMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method InviteMember not implemented")
}
func (UnimplementedAccountServiceServer) RemoveMember(context.Context, *RemoveMemberRequest) (*RemoveMemberResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveMember not implemented")
}
func (UnimplementedAccountServiceServer) CheckForUpdates(context.Context, *CheckForUpdatesRequest) (*CheckForUpdatesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckForUpdates not implemented")
}
func (UnimplementedAccountServiceServer) GetDeviceId(context.Context, *GetDeviceIdRequest) (*GetDeviceIdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDeviceId not implemented")
}
func (UnimplementedAccountServiceServer) Authenticated(context.Context, *AuthenticatedRequest) (*AuthenticatedResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Authenticated not implemented")
}

// UnsafeAccountServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to AccountServiceServer will
// result in compilation errors.
type UnsafeAccountServiceServer interface {
	mustEmbedUnimplementedAccountServiceServer()
}

func RegisterAccountServiceServer(s grpc.ServiceRegistrar, srv AccountServiceServer) {
	s.RegisterService(&AccountService_ServiceDesc, srv)
}

func _AccountService_Signup_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SignupRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Signup(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.v1alpha1.AccountService/Signup",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Signup(ctx, req.(*SignupRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_InviteMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InviteMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).InviteMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.v1alpha1.AccountService/InviteMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).InviteMember(ctx, req.(*InviteMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_RemoveMember_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveMemberRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).RemoveMember(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.v1alpha1.AccountService/RemoveMember",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).RemoveMember(ctx, req.(*RemoveMemberRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_CheckForUpdates_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CheckForUpdatesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).CheckForUpdates(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.v1alpha1.AccountService/CheckForUpdates",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).CheckForUpdates(ctx, req.(*CheckForUpdatesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_GetDeviceId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetDeviceIdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).GetDeviceId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.v1alpha1.AccountService/GetDeviceId",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).GetDeviceId(ctx, req.(*GetDeviceIdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _AccountService_Authenticated_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AuthenticatedRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountServiceServer).Authenticated(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/account.v1alpha1.AccountService/Authenticated",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountServiceServer).Authenticated(ctx, req.(*AuthenticatedRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// AccountService_ServiceDesc is the grpc.ServiceDesc for AccountService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var AccountService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "account.v1alpha1.AccountService",
	HandlerType: (*AccountServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Signup",
			Handler:    _AccountService_Signup_Handler,
		},
		{
			MethodName: "InviteMember",
			Handler:    _AccountService_InviteMember_Handler,
		},
		{
			MethodName: "RemoveMember",
			Handler:    _AccountService_RemoveMember_Handler,
		},
		{
			MethodName: "CheckForUpdates",
			Handler:    _AccountService_CheckForUpdates_Handler,
		},
		{
			MethodName: "GetDeviceId",
			Handler:    _AccountService_GetDeviceId_Handler,
		},
		{
			MethodName: "Authenticated",
			Handler:    _AccountService_Authenticated_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "account/v1alpha1/account.proto",
}
