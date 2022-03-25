// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: sso/v1alpha1/sso.proto

package ssov1alpha1

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

type SSOMode int32

const (
	SSOMode_SSO_MODE_UNSPECIFIED SSOMode = 0
	SSOMode_SSO_MODE_DISABLED    SSOMode = 1
	SSOMode_SSO_MODE_SAML        SSOMode = 2
)

// Enum value maps for SSOMode.
var (
	SSOMode_name = map[int32]string{
		0: "SSO_MODE_UNSPECIFIED",
		1: "SSO_MODE_DISABLED",
		2: "SSO_MODE_SAML",
	}
	SSOMode_value = map[string]int32{
		"SSO_MODE_UNSPECIFIED": 0,
		"SSO_MODE_DISABLED":    1,
		"SSO_MODE_SAML":        2,
	}
)

func (x SSOMode) Enum() *SSOMode {
	p := new(SSOMode)
	*p = x
	return p
}

func (x SSOMode) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SSOMode) Descriptor() protoreflect.EnumDescriptor {
	return file_sso_v1alpha1_sso_proto_enumTypes[0].Descriptor()
}

func (SSOMode) Type() protoreflect.EnumType {
	return &file_sso_v1alpha1_sso_proto_enumTypes[0]
}

func (x SSOMode) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SSOMode.Descriptor instead.
func (SSOMode) EnumDescriptor() ([]byte, []int) {
	return file_sso_v1alpha1_sso_proto_rawDescGZIP(), []int{0}
}

type SetupSAMLRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetupSAMLRequest) Reset() {
	*x = SetupSAMLRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_v1alpha1_sso_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupSAMLRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupSAMLRequest) ProtoMessage() {}

func (x *SetupSAMLRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_v1alpha1_sso_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupSAMLRequest.ProtoReflect.Descriptor instead.
func (*SetupSAMLRequest) Descriptor() ([]byte, []int) {
	return file_sso_v1alpha1_sso_proto_rawDescGZIP(), []int{0}
}

type SetupSAMLResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetadataUrl string `protobuf:"bytes,1,opt,name=metadata_url,json=metadataUrl,proto3" json:"metadata_url,omitempty"`
}

func (x *SetupSAMLResponse) Reset() {
	*x = SetupSAMLResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_v1alpha1_sso_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetupSAMLResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetupSAMLResponse) ProtoMessage() {}

func (x *SetupSAMLResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_v1alpha1_sso_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetupSAMLResponse.ProtoReflect.Descriptor instead.
func (*SetupSAMLResponse) Descriptor() ([]byte, []int) {
	return file_sso_v1alpha1_sso_proto_rawDescGZIP(), []int{1}
}

func (x *SetupSAMLResponse) GetMetadataUrl() string {
	if x != nil {
		return x.MetadataUrl
	}
	return ""
}

type SetSAMLIdentityProviderMetadataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata string `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *SetSAMLIdentityProviderMetadataRequest) Reset() {
	*x = SetSAMLIdentityProviderMetadataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_v1alpha1_sso_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSAMLIdentityProviderMetadataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSAMLIdentityProviderMetadataRequest) ProtoMessage() {}

func (x *SetSAMLIdentityProviderMetadataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_v1alpha1_sso_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSAMLIdentityProviderMetadataRequest.ProtoReflect.Descriptor instead.
func (*SetSAMLIdentityProviderMetadataRequest) Descriptor() ([]byte, []int) {
	return file_sso_v1alpha1_sso_proto_rawDescGZIP(), []int{2}
}

func (x *SetSAMLIdentityProviderMetadataRequest) GetMetadata() string {
	if x != nil {
		return x.Metadata
	}
	return ""
}

type SetSAMLIdentityProviderMetadataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetSAMLIdentityProviderMetadataResponse) Reset() {
	*x = SetSAMLIdentityProviderMetadataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_v1alpha1_sso_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetSAMLIdentityProviderMetadataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetSAMLIdentityProviderMetadataResponse) ProtoMessage() {}

func (x *SetSAMLIdentityProviderMetadataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_v1alpha1_sso_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetSAMLIdentityProviderMetadataResponse.ProtoReflect.Descriptor instead.
func (*SetSAMLIdentityProviderMetadataResponse) Descriptor() ([]byte, []int) {
	return file_sso_v1alpha1_sso_proto_rawDescGZIP(), []int{3}
}

type ChangeModeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode SSOMode `protobuf:"varint,1,opt,name=mode,proto3,enum=sso.v1alpha1.SSOMode" json:"mode,omitempty"`
}

func (x *ChangeModeRequest) Reset() {
	*x = ChangeModeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_v1alpha1_sso_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeModeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeModeRequest) ProtoMessage() {}

func (x *ChangeModeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_v1alpha1_sso_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeModeRequest.ProtoReflect.Descriptor instead.
func (*ChangeModeRequest) Descriptor() ([]byte, []int) {
	return file_sso_v1alpha1_sso_proto_rawDescGZIP(), []int{4}
}

func (x *ChangeModeRequest) GetMode() SSOMode {
	if x != nil {
		return x.Mode
	}
	return SSOMode_SSO_MODE_UNSPECIFIED
}

type ChangeModeResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ChangeModeResponse) Reset() {
	*x = ChangeModeResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_v1alpha1_sso_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeModeResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeModeResponse) ProtoMessage() {}

func (x *ChangeModeResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_v1alpha1_sso_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeModeResponse.ProtoReflect.Descriptor instead.
func (*ChangeModeResponse) Descriptor() ([]byte, []int) {
	return file_sso_v1alpha1_sso_proto_rawDescGZIP(), []int{5}
}

type GetSettingsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetSettingsRequest) Reset() {
	*x = GetSettingsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_v1alpha1_sso_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingsRequest) ProtoMessage() {}

func (x *GetSettingsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sso_v1alpha1_sso_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingsRequest.ProtoReflect.Descriptor instead.
func (*GetSettingsRequest) Descriptor() ([]byte, []int) {
	return file_sso_v1alpha1_sso_proto_rawDescGZIP(), []int{6}
}

type GetSettingsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Mode        SSOMode `protobuf:"varint,1,opt,name=mode,proto3,enum=sso.v1alpha1.SSOMode" json:"mode,omitempty"`
	MetadataUrl string  `protobuf:"bytes,2,opt,name=metadata_url,json=metadataUrl,proto3" json:"metadata_url,omitempty"`
	IdpMetadata string  `protobuf:"bytes,3,opt,name=idp_metadata,json=idpMetadata,proto3" json:"idp_metadata,omitempty"`
}

func (x *GetSettingsResponse) Reset() {
	*x = GetSettingsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sso_v1alpha1_sso_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetSettingsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetSettingsResponse) ProtoMessage() {}

func (x *GetSettingsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_sso_v1alpha1_sso_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetSettingsResponse.ProtoReflect.Descriptor instead.
func (*GetSettingsResponse) Descriptor() ([]byte, []int) {
	return file_sso_v1alpha1_sso_proto_rawDescGZIP(), []int{7}
}

func (x *GetSettingsResponse) GetMode() SSOMode {
	if x != nil {
		return x.Mode
	}
	return SSOMode_SSO_MODE_UNSPECIFIED
}

func (x *GetSettingsResponse) GetMetadataUrl() string {
	if x != nil {
		return x.MetadataUrl
	}
	return ""
}

func (x *GetSettingsResponse) GetIdpMetadata() string {
	if x != nil {
		return x.IdpMetadata
	}
	return ""
}

var File_sso_v1alpha1_sso_proto protoreflect.FileDescriptor

var file_sso_v1alpha1_sso_proto_rawDesc = []byte{
	0x0a, 0x16, 0x73, 0x73, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x73,
	0x73, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22,
	0x12, 0x0a, 0x10, 0x53, 0x65, 0x74, 0x75, 0x70, 0x53, 0x41, 0x4d, 0x4c, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x22, 0x36, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x75, 0x70, 0x53, 0x41, 0x4d, 0x4c,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x72, 0x6c, 0x22, 0x44, 0x0a, 0x26, 0x53,
	0x65, 0x74, 0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x22, 0x29, 0x0a, 0x27, 0x53, 0x65, 0x74, 0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x3e, 0x0a, 0x11,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x29, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32,
	0x15, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53,
	0x53, 0x4f, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x22, 0x14, 0x0a, 0x12,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x14, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67,
	0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x86, 0x01, 0x0a, 0x13, 0x47, 0x65, 0x74,
	0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x29, 0x0a, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x15,
	0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x53,
	0x4f, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x04, 0x6d, 0x6f, 0x64, 0x65, 0x12, 0x21, 0x0a, 0x0c, 0x6d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x55, 0x72, 0x6c, 0x12, 0x21,
	0x0a, 0x0c, 0x69, 0x64, 0x70, 0x5f, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x69, 0x64, 0x70, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0x2a, 0x4d, 0x0a, 0x07, 0x53, 0x53, 0x4f, 0x4d, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x14,
	0x53, 0x53, 0x4f, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43, 0x49,
	0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x15, 0x0a, 0x11, 0x53, 0x53, 0x4f, 0x5f, 0x4d, 0x4f,
	0x44, 0x45, 0x5f, 0x44, 0x49, 0x53, 0x41, 0x42, 0x4c, 0x45, 0x44, 0x10, 0x01, 0x12, 0x11, 0x0a,
	0x0d, 0x53, 0x53, 0x4f, 0x5f, 0x4d, 0x4f, 0x44, 0x45, 0x5f, 0x53, 0x41, 0x4d, 0x4c, 0x10, 0x02,
	0x32, 0x90, 0x03, 0x0a, 0x0a, 0x53, 0x53, 0x4f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x4c, 0x0a, 0x09, 0x53, 0x65, 0x74, 0x75, 0x70, 0x53, 0x41, 0x4d, 0x4c, 0x12, 0x1e, 0x2e, 0x73,
	0x73, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x75,
	0x70, 0x53, 0x41, 0x4d, 0x4c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1f, 0x2e, 0x73,
	0x73, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x75,
	0x70, 0x53, 0x41, 0x4d, 0x4c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x52, 0x0a,
	0x0b, 0x47, 0x65, 0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x12, 0x20, 0x2e, 0x73,
	0x73, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x53,
	0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x21,
	0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x65, 0x74, 0x74, 0x69, 0x6e, 0x67, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x8e, 0x01, 0x0a, 0x1f, 0x53, 0x65, 0x74, 0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74,
	0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x34, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x41, 0x4d, 0x4c, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x35, 0x2e, 0x73, 0x73,
	0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x65, 0x74, 0x53, 0x41,
	0x4d, 0x4c, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4f, 0x0a, 0x0a, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x6f, 0x64, 0x65,
	0x12, 0x1f, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x73, 0x73, 0x6f, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x4d, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x42, 0xb5, 0x01, 0x0a, 0x10, 0x63, 0x6f, 0x6d, 0x2e, 0x73, 0x73, 0x6f, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x08, 0x53, 0x73, 0x6f, 0x50, 0x72, 0x6f,
	0x74, 0x6f, 0x50, 0x01, 0x5a, 0x46, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x66, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x2f, 0x73, 0x73, 0x6f, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x3b, 0x73, 0x73, 0x6f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x53,
	0x58, 0x58, 0xaa, 0x02, 0x0c, 0x53, 0x73, 0x6f, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xca, 0x02, 0x0c, 0x53, 0x73, 0x6f, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0xe2, 0x02, 0x18, 0x53, 0x73, 0x6f, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c,
	0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0d, 0x53, 0x73,
	0x6f, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_sso_v1alpha1_sso_proto_rawDescOnce sync.Once
	file_sso_v1alpha1_sso_proto_rawDescData = file_sso_v1alpha1_sso_proto_rawDesc
)

func file_sso_v1alpha1_sso_proto_rawDescGZIP() []byte {
	file_sso_v1alpha1_sso_proto_rawDescOnce.Do(func() {
		file_sso_v1alpha1_sso_proto_rawDescData = protoimpl.X.CompressGZIP(file_sso_v1alpha1_sso_proto_rawDescData)
	})
	return file_sso_v1alpha1_sso_proto_rawDescData
}

var file_sso_v1alpha1_sso_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_sso_v1alpha1_sso_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_sso_v1alpha1_sso_proto_goTypes = []interface{}{
	(SSOMode)(0),                                    // 0: sso.v1alpha1.SSOMode
	(*SetupSAMLRequest)(nil),                        // 1: sso.v1alpha1.SetupSAMLRequest
	(*SetupSAMLResponse)(nil),                       // 2: sso.v1alpha1.SetupSAMLResponse
	(*SetSAMLIdentityProviderMetadataRequest)(nil),  // 3: sso.v1alpha1.SetSAMLIdentityProviderMetadataRequest
	(*SetSAMLIdentityProviderMetadataResponse)(nil), // 4: sso.v1alpha1.SetSAMLIdentityProviderMetadataResponse
	(*ChangeModeRequest)(nil),                       // 5: sso.v1alpha1.ChangeModeRequest
	(*ChangeModeResponse)(nil),                      // 6: sso.v1alpha1.ChangeModeResponse
	(*GetSettingsRequest)(nil),                      // 7: sso.v1alpha1.GetSettingsRequest
	(*GetSettingsResponse)(nil),                     // 8: sso.v1alpha1.GetSettingsResponse
}
var file_sso_v1alpha1_sso_proto_depIdxs = []int32{
	0, // 0: sso.v1alpha1.ChangeModeRequest.mode:type_name -> sso.v1alpha1.SSOMode
	0, // 1: sso.v1alpha1.GetSettingsResponse.mode:type_name -> sso.v1alpha1.SSOMode
	1, // 2: sso.v1alpha1.SSOService.SetupSAML:input_type -> sso.v1alpha1.SetupSAMLRequest
	7, // 3: sso.v1alpha1.SSOService.GetSettings:input_type -> sso.v1alpha1.GetSettingsRequest
	3, // 4: sso.v1alpha1.SSOService.SetSAMLIdentityProviderMetadata:input_type -> sso.v1alpha1.SetSAMLIdentityProviderMetadataRequest
	5, // 5: sso.v1alpha1.SSOService.ChangeMode:input_type -> sso.v1alpha1.ChangeModeRequest
	2, // 6: sso.v1alpha1.SSOService.SetupSAML:output_type -> sso.v1alpha1.SetupSAMLResponse
	8, // 7: sso.v1alpha1.SSOService.GetSettings:output_type -> sso.v1alpha1.GetSettingsResponse
	4, // 8: sso.v1alpha1.SSOService.SetSAMLIdentityProviderMetadata:output_type -> sso.v1alpha1.SetSAMLIdentityProviderMetadataResponse
	6, // 9: sso.v1alpha1.SSOService.ChangeMode:output_type -> sso.v1alpha1.ChangeModeResponse
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_sso_v1alpha1_sso_proto_init() }
func file_sso_v1alpha1_sso_proto_init() {
	if File_sso_v1alpha1_sso_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sso_v1alpha1_sso_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupSAMLRequest); i {
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
		file_sso_v1alpha1_sso_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetupSAMLResponse); i {
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
		file_sso_v1alpha1_sso_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSAMLIdentityProviderMetadataRequest); i {
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
		file_sso_v1alpha1_sso_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetSAMLIdentityProviderMetadataResponse); i {
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
		file_sso_v1alpha1_sso_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeModeRequest); i {
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
		file_sso_v1alpha1_sso_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeModeResponse); i {
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
		file_sso_v1alpha1_sso_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingsRequest); i {
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
		file_sso_v1alpha1_sso_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetSettingsResponse); i {
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
			RawDescriptor: file_sso_v1alpha1_sso_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sso_v1alpha1_sso_proto_goTypes,
		DependencyIndexes: file_sso_v1alpha1_sso_proto_depIdxs,
		EnumInfos:         file_sso_v1alpha1_sso_proto_enumTypes,
		MessageInfos:      file_sso_v1alpha1_sso_proto_msgTypes,
	}.Build()
	File_sso_v1alpha1_sso_proto = out.File
	file_sso_v1alpha1_sso_proto_rawDesc = nil
	file_sso_v1alpha1_sso_proto_goTypes = nil
	file_sso_v1alpha1_sso_proto_depIdxs = nil
}
