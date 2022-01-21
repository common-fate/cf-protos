// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package teamv1alpha1

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

// TeamServiceClient is the client API for TeamService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TeamServiceClient interface {
	ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error)
	ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error)
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error)
	EnrolProvider(ctx context.Context, in *EnrolProviderRequest, opts ...grpc.CallOption) (*EnrolProviderResponse, error)
	ListProviders(ctx context.Context, in *ListProvidersRequest, opts ...grpc.CallOption) (*ListProvidersResponse, error)
	// DeleteProvider removes a provider from a Granted team.
	DeleteProvider(ctx context.Context, in *DeleteProviderRequest, opts ...grpc.CallOption) (*DeleteProviderResponse, error)
	// GetStatus returns the overall state of a team's Granted deployments and whether any
	// actions are required from an administrator.
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	GetProvider(ctx context.Context, in *GetProviderRequest, opts ...grpc.CallOption) (*GetProviderResponse, error)
	// GetAllProviderDetails returns details about a team's providers including all accounts and
	// access handlers associated with the provider.
	GetAllProviderDetails(ctx context.Context, in *GetAllProviderDetailsRequest, opts ...grpc.CallOption) (*GetAllProviderDetailsResponse, error)
	// GetAllProviderChecksum is used by clients to determine whether their local cache of provider
	// details requires an update.
	GetAllProviderChecksum(ctx context.Context, in *GetAllProviderChecksumRequest, opts ...grpc.CallOption) (*GetAllProviderChecksumResponse, error)
	// GetAccessHandlersForProvider lists the Access Handlers associated with a provider
	GetAccessHandlersForProvider(ctx context.Context, in *GetAccessHandlersForProviderRequest, opts ...grpc.CallOption) (*GetAccessHandlersForProviderResponse, error)
	// AddAccessHandler registers a new Access Handler for a provider
	AddAccessHandler(ctx context.Context, in *AddAccessHandlerRequest, opts ...grpc.CallOption) (*AddAccessHandlerResponse, error)
	// DeleteAccessHandler deletes an Access Handler from a provider
	DeleteAccessHandler(ctx context.Context, in *DeleteAccessHandlerRequest, opts ...grpc.CallOption) (*DeleteAccessHandlerResponse, error)
}

type teamServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTeamServiceClient(cc grpc.ClientConnInterface) TeamServiceClient {
	return &teamServiceClient{cc}
}

func (c *teamServiceClient) ListMembers(ctx context.Context, in *ListMembersRequest, opts ...grpc.CallOption) (*ListMembersResponse, error) {
	out := new(ListMembersResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/ListMembers", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) ListRoles(ctx context.Context, in *ListRolesRequest, opts ...grpc.CallOption) (*ListRolesResponse, error) {
	out := new(ListRolesResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/ListRoles", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error) {
	out := new(UpdateConfigResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) EnrolProvider(ctx context.Context, in *EnrolProviderRequest, opts ...grpc.CallOption) (*EnrolProviderResponse, error) {
	out := new(EnrolProviderResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/EnrolProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) ListProviders(ctx context.Context, in *ListProvidersRequest, opts ...grpc.CallOption) (*ListProvidersResponse, error) {
	out := new(ListProvidersResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/ListProviders", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) DeleteProvider(ctx context.Context, in *DeleteProviderRequest, opts ...grpc.CallOption) (*DeleteProviderResponse, error) {
	out := new(DeleteProviderResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/DeleteProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error) {
	out := new(GetStatusResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/GetStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetProvider(ctx context.Context, in *GetProviderRequest, opts ...grpc.CallOption) (*GetProviderResponse, error) {
	out := new(GetProviderResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/GetProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetAllProviderDetails(ctx context.Context, in *GetAllProviderDetailsRequest, opts ...grpc.CallOption) (*GetAllProviderDetailsResponse, error) {
	out := new(GetAllProviderDetailsResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/GetAllProviderDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetAllProviderChecksum(ctx context.Context, in *GetAllProviderChecksumRequest, opts ...grpc.CallOption) (*GetAllProviderChecksumResponse, error) {
	out := new(GetAllProviderChecksumResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/GetAllProviderChecksum", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetAccessHandlersForProvider(ctx context.Context, in *GetAccessHandlersForProviderRequest, opts ...grpc.CallOption) (*GetAccessHandlersForProviderResponse, error) {
	out := new(GetAccessHandlersForProviderResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/GetAccessHandlersForProvider", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) AddAccessHandler(ctx context.Context, in *AddAccessHandlerRequest, opts ...grpc.CallOption) (*AddAccessHandlerResponse, error) {
	out := new(AddAccessHandlerResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/AddAccessHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) DeleteAccessHandler(ctx context.Context, in *DeleteAccessHandlerRequest, opts ...grpc.CallOption) (*DeleteAccessHandlerResponse, error) {
	out := new(DeleteAccessHandlerResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/DeleteAccessHandler", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TeamServiceServer is the server API for TeamService service.
// All implementations should embed UnimplementedTeamServiceServer
// for forward compatibility
type TeamServiceServer interface {
	ListMembers(context.Context, *ListMembersRequest) (*ListMembersResponse, error)
	ListRoles(context.Context, *ListRolesRequest) (*ListRolesResponse, error)
	UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error)
	EnrolProvider(context.Context, *EnrolProviderRequest) (*EnrolProviderResponse, error)
	ListProviders(context.Context, *ListProvidersRequest) (*ListProvidersResponse, error)
	// DeleteProvider removes a provider from a Granted team.
	DeleteProvider(context.Context, *DeleteProviderRequest) (*DeleteProviderResponse, error)
	// GetStatus returns the overall state of a team's Granted deployments and whether any
	// actions are required from an administrator.
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	GetProvider(context.Context, *GetProviderRequest) (*GetProviderResponse, error)
	// GetAllProviderDetails returns details about a team's providers including all accounts and
	// access handlers associated with the provider.
	GetAllProviderDetails(context.Context, *GetAllProviderDetailsRequest) (*GetAllProviderDetailsResponse, error)
	// GetAllProviderChecksum is used by clients to determine whether their local cache of provider
	// details requires an update.
	GetAllProviderChecksum(context.Context, *GetAllProviderChecksumRequest) (*GetAllProviderChecksumResponse, error)
	// GetAccessHandlersForProvider lists the Access Handlers associated with a provider
	GetAccessHandlersForProvider(context.Context, *GetAccessHandlersForProviderRequest) (*GetAccessHandlersForProviderResponse, error)
	// AddAccessHandler registers a new Access Handler for a provider
	AddAccessHandler(context.Context, *AddAccessHandlerRequest) (*AddAccessHandlerResponse, error)
	// DeleteAccessHandler deletes an Access Handler from a provider
	DeleteAccessHandler(context.Context, *DeleteAccessHandlerRequest) (*DeleteAccessHandlerResponse, error)
}

// UnimplementedTeamServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTeamServiceServer struct {
}

func (UnimplementedTeamServiceServer) ListMembers(context.Context, *ListMembersRequest) (*ListMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMembers not implemented")
}
func (UnimplementedTeamServiceServer) ListRoles(context.Context, *ListRolesRequest) (*ListRolesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRoles not implemented")
}
func (UnimplementedTeamServiceServer) UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedTeamServiceServer) EnrolProvider(context.Context, *EnrolProviderRequest) (*EnrolProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EnrolProvider not implemented")
}
func (UnimplementedTeamServiceServer) ListProviders(context.Context, *ListProvidersRequest) (*ListProvidersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListProviders not implemented")
}
func (UnimplementedTeamServiceServer) DeleteProvider(context.Context, *DeleteProviderRequest) (*DeleteProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteProvider not implemented")
}
func (UnimplementedTeamServiceServer) GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetStatus not implemented")
}
func (UnimplementedTeamServiceServer) GetProvider(context.Context, *GetProviderRequest) (*GetProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProvider not implemented")
}
func (UnimplementedTeamServiceServer) GetAllProviderDetails(context.Context, *GetAllProviderDetailsRequest) (*GetAllProviderDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProviderDetails not implemented")
}
func (UnimplementedTeamServiceServer) GetAllProviderChecksum(context.Context, *GetAllProviderChecksumRequest) (*GetAllProviderChecksumResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllProviderChecksum not implemented")
}
func (UnimplementedTeamServiceServer) GetAccessHandlersForProvider(context.Context, *GetAccessHandlersForProviderRequest) (*GetAccessHandlersForProviderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAccessHandlersForProvider not implemented")
}
func (UnimplementedTeamServiceServer) AddAccessHandler(context.Context, *AddAccessHandlerRequest) (*AddAccessHandlerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddAccessHandler not implemented")
}
func (UnimplementedTeamServiceServer) DeleteAccessHandler(context.Context, *DeleteAccessHandlerRequest) (*DeleteAccessHandlerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteAccessHandler not implemented")
}

// UnsafeTeamServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TeamServiceServer will
// result in compilation errors.
type UnsafeTeamServiceServer interface {
	mustEmbedUnimplementedTeamServiceServer()
}

func RegisterTeamServiceServer(s grpc.ServiceRegistrar, srv TeamServiceServer) {
	s.RegisterService(&TeamService_ServiceDesc, srv)
}

func _TeamService_ListMembers_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListMembersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).ListMembers(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/ListMembers",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).ListMembers(ctx, req.(*ListMembersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_ListRoles_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRolesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).ListRoles(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/ListRoles",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).ListRoles(ctx, req.(*ListRolesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_UpdateConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).UpdateConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/UpdateConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).UpdateConfig(ctx, req.(*UpdateConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_EnrolProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EnrolProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).EnrolProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/EnrolProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).EnrolProvider(ctx, req.(*EnrolProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_ListProviders_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListProvidersRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).ListProviders(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/ListProviders",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).ListProviders(ctx, req.(*ListProvidersRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_DeleteProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).DeleteProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/DeleteProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).DeleteProvider(ctx, req.(*DeleteProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/GetStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetStatus(ctx, req.(*GetStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/GetProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetProvider(ctx, req.(*GetProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetAllProviderDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProviderDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetAllProviderDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/GetAllProviderDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetAllProviderDetails(ctx, req.(*GetAllProviderDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetAllProviderChecksum_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAllProviderChecksumRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetAllProviderChecksum(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/GetAllProviderChecksum",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetAllProviderChecksum(ctx, req.(*GetAllProviderChecksumRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetAccessHandlersForProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAccessHandlersForProviderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetAccessHandlersForProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/GetAccessHandlersForProvider",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetAccessHandlersForProvider(ctx, req.(*GetAccessHandlersForProviderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_AddAccessHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddAccessHandlerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).AddAccessHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/AddAccessHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).AddAccessHandler(ctx, req.(*AddAccessHandlerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_DeleteAccessHandler_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteAccessHandlerRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).DeleteAccessHandler(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/DeleteAccessHandler",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).DeleteAccessHandler(ctx, req.(*DeleteAccessHandlerRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TeamService_ServiceDesc is the grpc.ServiceDesc for TeamService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TeamService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "team.v1alpha1.TeamService",
	HandlerType: (*TeamServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListMembers",
			Handler:    _TeamService_ListMembers_Handler,
		},
		{
			MethodName: "ListRoles",
			Handler:    _TeamService_ListRoles_Handler,
		},
		{
			MethodName: "UpdateConfig",
			Handler:    _TeamService_UpdateConfig_Handler,
		},
		{
			MethodName: "EnrolProvider",
			Handler:    _TeamService_EnrolProvider_Handler,
		},
		{
			MethodName: "ListProviders",
			Handler:    _TeamService_ListProviders_Handler,
		},
		{
			MethodName: "DeleteProvider",
			Handler:    _TeamService_DeleteProvider_Handler,
		},
		{
			MethodName: "GetStatus",
			Handler:    _TeamService_GetStatus_Handler,
		},
		{
			MethodName: "GetProvider",
			Handler:    _TeamService_GetProvider_Handler,
		},
		{
			MethodName: "GetAllProviderDetails",
			Handler:    _TeamService_GetAllProviderDetails_Handler,
		},
		{
			MethodName: "GetAllProviderChecksum",
			Handler:    _TeamService_GetAllProviderChecksum_Handler,
		},
		{
			MethodName: "GetAccessHandlersForProvider",
			Handler:    _TeamService_GetAccessHandlersForProvider_Handler,
		},
		{
			MethodName: "AddAccessHandler",
			Handler:    _TeamService_AddAccessHandler_Handler,
		},
		{
			MethodName: "DeleteAccessHandler",
			Handler:    _TeamService_DeleteAccessHandler_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "team/v1alpha1/team.proto",
}
