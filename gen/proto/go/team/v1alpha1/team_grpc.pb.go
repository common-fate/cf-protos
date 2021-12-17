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
	UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error)
	// GetConfig returns the latest approved config
	GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error)
	GetConfigByHash(ctx context.Context, in *GetConfigByHashRequest, opts ...grpc.CallOption) (*GetConfigByHashResponse, error)
	EnrolProvider(ctx context.Context, in *EnrolProviderRequest, opts ...grpc.CallOption) (*EnrolProviderResponse, error)
	ListProviders(ctx context.Context, in *ListProvidersRequest, opts ...grpc.CallOption) (*ListProvidersResponse, error)
	// DeleteProvider removes a provider from a Granted team.
	DeleteProvider(ctx context.Context, in *DeleteProviderRequest, opts ...grpc.CallOption) (*DeleteProviderResponse, error)
	// GetStatus returns the overall state of a team's Granted deployments and whether any
	// actions are required from an administrator.
	GetStatus(ctx context.Context, in *GetStatusRequest, opts ...grpc.CallOption) (*GetStatusResponse, error)
	GetProvider(ctx context.Context, in *GetProviderRequest, opts ...grpc.CallOption) (*GetProviderResponse, error)
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

func (c *teamServiceClient) UpdateConfig(ctx context.Context, in *UpdateConfigRequest, opts ...grpc.CallOption) (*UpdateConfigResponse, error) {
	out := new(UpdateConfigResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/UpdateConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetConfig(ctx context.Context, in *GetConfigRequest, opts ...grpc.CallOption) (*GetConfigResponse, error) {
	out := new(GetConfigResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/GetConfig", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *teamServiceClient) GetConfigByHash(ctx context.Context, in *GetConfigByHashRequest, opts ...grpc.CallOption) (*GetConfigByHashResponse, error) {
	out := new(GetConfigByHashResponse)
	err := c.cc.Invoke(ctx, "/team.v1alpha1.TeamService/GetConfigByHash", in, out, opts...)
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

// TeamServiceServer is the server API for TeamService service.
// All implementations should embed UnimplementedTeamServiceServer
// for forward compatibility
type TeamServiceServer interface {
	ListMembers(context.Context, *ListMembersRequest) (*ListMembersResponse, error)
	UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error)
	// GetConfig returns the latest approved config
	GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error)
	GetConfigByHash(context.Context, *GetConfigByHashRequest) (*GetConfigByHashResponse, error)
	EnrolProvider(context.Context, *EnrolProviderRequest) (*EnrolProviderResponse, error)
	ListProviders(context.Context, *ListProvidersRequest) (*ListProvidersResponse, error)
	// DeleteProvider removes a provider from a Granted team.
	DeleteProvider(context.Context, *DeleteProviderRequest) (*DeleteProviderResponse, error)
	// GetStatus returns the overall state of a team's Granted deployments and whether any
	// actions are required from an administrator.
	GetStatus(context.Context, *GetStatusRequest) (*GetStatusResponse, error)
	GetProvider(context.Context, *GetProviderRequest) (*GetProviderResponse, error)
}

// UnimplementedTeamServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTeamServiceServer struct {
}

func (UnimplementedTeamServiceServer) ListMembers(context.Context, *ListMembersRequest) (*ListMembersResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListMembers not implemented")
}
func (UnimplementedTeamServiceServer) UpdateConfig(context.Context, *UpdateConfigRequest) (*UpdateConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateConfig not implemented")
}
func (UnimplementedTeamServiceServer) GetConfig(context.Context, *GetConfigRequest) (*GetConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfig not implemented")
}
func (UnimplementedTeamServiceServer) GetConfigByHash(context.Context, *GetConfigByHashRequest) (*GetConfigByHashResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConfigByHash not implemented")
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

func _TeamService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetConfig(ctx, req.(*GetConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _TeamService_GetConfigByHash_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetConfigByHashRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TeamServiceServer).GetConfigByHash(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/team.v1alpha1.TeamService/GetConfigByHash",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TeamServiceServer).GetConfigByHash(ctx, req.(*GetConfigByHashRequest))
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
			MethodName: "UpdateConfig",
			Handler:    _TeamService_UpdateConfig_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _TeamService_GetConfig_Handler,
		},
		{
			MethodName: "GetConfigByHash",
			Handler:    _TeamService_GetConfigByHash_Handler,
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
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "team/v1alpha1/team.proto",
}
