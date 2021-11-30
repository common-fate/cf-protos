// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package tenantv1alpha1

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

// TenantServiceClient is the client API for TenantService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type TenantServiceClient interface {
	GetTenantByPublicKey(ctx context.Context, in *GetTenantByPublicKeyRequest, opts ...grpc.CallOption) (*GetTenantByPublicKeyResponse, error)
}

type tenantServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewTenantServiceClient(cc grpc.ClientConnInterface) TenantServiceClient {
	return &tenantServiceClient{cc}
}

func (c *tenantServiceClient) GetTenantByPublicKey(ctx context.Context, in *GetTenantByPublicKeyRequest, opts ...grpc.CallOption) (*GetTenantByPublicKeyResponse, error) {
	out := new(GetTenantByPublicKeyResponse)
	err := c.cc.Invoke(ctx, "/tenant.v1alpha1.TenantService/GetTenantByPublicKey", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TenantServiceServer is the server API for TenantService service.
// All implementations should embed UnimplementedTenantServiceServer
// for forward compatibility
type TenantServiceServer interface {
	GetTenantByPublicKey(context.Context, *GetTenantByPublicKeyRequest) (*GetTenantByPublicKeyResponse, error)
}

// UnimplementedTenantServiceServer should be embedded to have forward compatible implementations.
type UnimplementedTenantServiceServer struct {
}

func (UnimplementedTenantServiceServer) GetTenantByPublicKey(context.Context, *GetTenantByPublicKeyRequest) (*GetTenantByPublicKeyResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTenantByPublicKey not implemented")
}

// UnsafeTenantServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to TenantServiceServer will
// result in compilation errors.
type UnsafeTenantServiceServer interface {
	mustEmbedUnimplementedTenantServiceServer()
}

func RegisterTenantServiceServer(s grpc.ServiceRegistrar, srv TenantServiceServer) {
	s.RegisterService(&TenantService_ServiceDesc, srv)
}

func _TenantService_GetTenantByPublicKey_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTenantByPublicKeyRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TenantServiceServer).GetTenantByPublicKey(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/tenant.v1alpha1.TenantService/GetTenantByPublicKey",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TenantServiceServer).GetTenantByPublicKey(ctx, req.(*GetTenantByPublicKeyRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// TenantService_ServiceDesc is the grpc.ServiceDesc for TenantService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var TenantService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "tenant.v1alpha1.TenantService",
	HandlerType: (*TenantServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTenantByPublicKey",
			Handler:    _TenantService_GetTenantByPublicKey_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "tenant/v1alpha1/tenant.proto",
}
