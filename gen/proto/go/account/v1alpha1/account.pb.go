// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: account/v1alpha1/account.proto

package accountv1alpha1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type SignupRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AdminEmailAddress string `protobuf:"bytes,1,opt,name=admin_email_address,json=adminEmailAddress,proto3" json:"admin_email_address,omitempty"`
	TenantName        string `protobuf:"bytes,2,opt,name=tenant_name,json=tenantName,proto3" json:"tenant_name,omitempty"`
	TenantSlug        string `protobuf:"bytes,3,opt,name=tenant_slug,json=tenantSlug,proto3" json:"tenant_slug,omitempty"`
}

func (x *SignupRequest) Reset() {
	*x = SignupRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_v1alpha1_account_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupRequest) ProtoMessage() {}

func (x *SignupRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1alpha1_account_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupRequest.ProtoReflect.Descriptor instead.
func (*SignupRequest) Descriptor() ([]byte, []int) {
	return file_account_v1alpha1_account_proto_rawDescGZIP(), []int{0}
}

func (x *SignupRequest) GetAdminEmailAddress() string {
	if x != nil {
		return x.AdminEmailAddress
	}
	return ""
}

func (x *SignupRequest) GetTenantName() string {
	if x != nil {
		return x.TenantName
	}
	return ""
}

func (x *SignupRequest) GetTenantSlug() string {
	if x != nil {
		return x.TenantSlug
	}
	return ""
}

type SignupResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SignupResponse) Reset() {
	*x = SignupResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_v1alpha1_account_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SignupResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignupResponse) ProtoMessage() {}

func (x *SignupResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1alpha1_account_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignupResponse.ProtoReflect.Descriptor instead.
func (*SignupResponse) Descriptor() ([]byte, []int) {
	return file_account_v1alpha1_account_proto_rawDescGZIP(), []int{1}
}

type GetDeviceIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetDeviceIdRequest) Reset() {
	*x = GetDeviceIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_v1alpha1_account_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceIdRequest) ProtoMessage() {}

func (x *GetDeviceIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1alpha1_account_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceIdRequest.ProtoReflect.Descriptor instead.
func (*GetDeviceIdRequest) Descriptor() ([]byte, []int) {
	return file_account_v1alpha1_account_proto_rawDescGZIP(), []int{2}
}

type GetDeviceIdResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetDeviceIdResponse) Reset() {
	*x = GetDeviceIdResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_v1alpha1_account_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDeviceIdResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDeviceIdResponse) ProtoMessage() {}

func (x *GetDeviceIdResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1alpha1_account_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDeviceIdResponse.ProtoReflect.Descriptor instead.
func (*GetDeviceIdResponse) Descriptor() ([]byte, []int) {
	return file_account_v1alpha1_account_proto_rawDescGZIP(), []int{3}
}

func (x *GetDeviceIdResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type AuthenticatedRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthenticatedRequest) Reset() {
	*x = AuthenticatedRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_v1alpha1_account_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticatedRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticatedRequest) ProtoMessage() {}

func (x *AuthenticatedRequest) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1alpha1_account_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticatedRequest.ProtoReflect.Descriptor instead.
func (*AuthenticatedRequest) Descriptor() ([]byte, []int) {
	return file_account_v1alpha1_account_proto_rawDescGZIP(), []int{4}
}

type AuthenticatedResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *AuthenticatedResponse) Reset() {
	*x = AuthenticatedResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_account_v1alpha1_account_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AuthenticatedResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AuthenticatedResponse) ProtoMessage() {}

func (x *AuthenticatedResponse) ProtoReflect() protoreflect.Message {
	mi := &file_account_v1alpha1_account_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AuthenticatedResponse.ProtoReflect.Descriptor instead.
func (*AuthenticatedResponse) Descriptor() ([]byte, []int) {
	return file_account_v1alpha1_account_proto_rawDescGZIP(), []int{5}
}

var File_account_v1alpha1_account_proto protoreflect.FileDescriptor

var file_account_v1alpha1_account_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x10, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x9c, 0x01, 0x0a, 0x0d,
	0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37, 0x0a,
	0x13, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x5f, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64,
	0x72, 0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72,
	0x02, 0x60, 0x01, 0x52, 0x11, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x45, 0x6d, 0x61, 0x69, 0x6c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x28, 0x0a, 0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74,
	0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04,
	0x72, 0x02, 0x18, 0x32, 0x52, 0x0a, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x0b, 0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x5f, 0x73, 0x6c, 0x75, 0x67, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x72, 0x02, 0x18, 0x32, 0x52, 0x0a,
	0x74, 0x65, 0x6e, 0x61, 0x6e, 0x74, 0x53, 0x6c, 0x75, 0x67, 0x22, 0x10, 0x0a, 0x0e, 0x53, 0x69,
	0x67, 0x6e, 0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x14, 0x0a, 0x12,
	0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x22, 0x25, 0x0a, 0x13, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49,
	0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0x16, 0x0a, 0x14, 0x41, 0x75, 0x74,
	0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x22, 0x17, 0x0a, 0x15, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74,
	0x65, 0x64, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x9b, 0x02, 0x0a, 0x0e, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a,
	0x06, 0x53, 0x69, 0x67, 0x6e, 0x75, 0x70, 0x12, 0x1f, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x75,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x69, 0x67, 0x6e,
	0x75, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x5a, 0x0a, 0x0b, 0x47, 0x65,
	0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x24, 0x2e, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x25, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x49, 0x64, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x60, 0x0a, 0x0d, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e,
	0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x12, 0x26, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65,
	0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x65, 0x64,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd5, 0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d,
	0x2e, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x42, 0x0c, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x4e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x66, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x3b, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xa2, 0x02, 0x03, 0x41, 0x58, 0x58, 0xaa, 0x02, 0x10, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x10, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1c,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x11, 0x41,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_account_v1alpha1_account_proto_rawDescOnce sync.Once
	file_account_v1alpha1_account_proto_rawDescData = file_account_v1alpha1_account_proto_rawDesc
)

func file_account_v1alpha1_account_proto_rawDescGZIP() []byte {
	file_account_v1alpha1_account_proto_rawDescOnce.Do(func() {
		file_account_v1alpha1_account_proto_rawDescData = protoimpl.X.CompressGZIP(file_account_v1alpha1_account_proto_rawDescData)
	})
	return file_account_v1alpha1_account_proto_rawDescData
}

var file_account_v1alpha1_account_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_account_v1alpha1_account_proto_goTypes = []interface{}{
	(*SignupRequest)(nil),         // 0: account.v1alpha1.SignupRequest
	(*SignupResponse)(nil),        // 1: account.v1alpha1.SignupResponse
	(*GetDeviceIdRequest)(nil),    // 2: account.v1alpha1.GetDeviceIdRequest
	(*GetDeviceIdResponse)(nil),   // 3: account.v1alpha1.GetDeviceIdResponse
	(*AuthenticatedRequest)(nil),  // 4: account.v1alpha1.AuthenticatedRequest
	(*AuthenticatedResponse)(nil), // 5: account.v1alpha1.AuthenticatedResponse
}
var file_account_v1alpha1_account_proto_depIdxs = []int32{
	0, // 0: account.v1alpha1.AccountService.Signup:input_type -> account.v1alpha1.SignupRequest
	2, // 1: account.v1alpha1.AccountService.GetDeviceId:input_type -> account.v1alpha1.GetDeviceIdRequest
	4, // 2: account.v1alpha1.AccountService.Authenticated:input_type -> account.v1alpha1.AuthenticatedRequest
	1, // 3: account.v1alpha1.AccountService.Signup:output_type -> account.v1alpha1.SignupResponse
	3, // 4: account.v1alpha1.AccountService.GetDeviceId:output_type -> account.v1alpha1.GetDeviceIdResponse
	5, // 5: account.v1alpha1.AccountService.Authenticated:output_type -> account.v1alpha1.AuthenticatedResponse
	3, // [3:6] is the sub-list for method output_type
	0, // [0:3] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_account_v1alpha1_account_proto_init() }
func file_account_v1alpha1_account_proto_init() {
	if File_account_v1alpha1_account_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_account_v1alpha1_account_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_account_v1alpha1_account_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SignupResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_account_v1alpha1_account_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceIdRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_account_v1alpha1_account_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDeviceIdResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_account_v1alpha1_account_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticatedRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_account_v1alpha1_account_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AuthenticatedResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_account_v1alpha1_account_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_account_v1alpha1_account_proto_goTypes,
		DependencyIndexes: file_account_v1alpha1_account_proto_depIdxs,
		MessageInfos:      file_account_v1alpha1_account_proto_msgTypes,
	}.Build()
	File_account_v1alpha1_account_proto = out.File
	file_account_v1alpha1_account_proto_rawDesc = nil
	file_account_v1alpha1_account_proto_goTypes = nil
	file_account_v1alpha1_account_proto_depIdxs = nil
}
