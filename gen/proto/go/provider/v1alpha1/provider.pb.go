// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: provider/v1alpha1/provider.proto

package providerv1alpha1

import (
	v1alpha1 "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
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

type GetProviderConfigByDigestRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sha256Digest []byte `protobuf:"bytes,1,opt,name=sha256_digest,json=sha256Digest,proto3" json:"sha256_digest,omitempty"`
}

func (x *GetProviderConfigByDigestRequest) Reset() {
	*x = GetProviderConfigByDigestRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_v1alpha1_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProviderConfigByDigestRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProviderConfigByDigestRequest) ProtoMessage() {}

func (x *GetProviderConfigByDigestRequest) ProtoReflect() protoreflect.Message {
	mi := &file_provider_v1alpha1_provider_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProviderConfigByDigestRequest.ProtoReflect.Descriptor instead.
func (*GetProviderConfigByDigestRequest) Descriptor() ([]byte, []int) {
	return file_provider_v1alpha1_provider_proto_rawDescGZIP(), []int{0}
}

func (x *GetProviderConfigByDigestRequest) GetSha256Digest() []byte {
	if x != nil {
		return x.Sha256Digest
	}
	return nil
}

type GetProviderConfigByDigestResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Provider *v1alpha1.Provider `protobuf:"bytes,1,opt,name=provider,proto3" json:"provider,omitempty"`
}

func (x *GetProviderConfigByDigestResponse) Reset() {
	*x = GetProviderConfigByDigestResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_v1alpha1_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProviderConfigByDigestResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProviderConfigByDigestResponse) ProtoMessage() {}

func (x *GetProviderConfigByDigestResponse) ProtoReflect() protoreflect.Message {
	mi := &file_provider_v1alpha1_provider_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProviderConfigByDigestResponse.ProtoReflect.Descriptor instead.
func (*GetProviderConfigByDigestResponse) Descriptor() ([]byte, []int) {
	return file_provider_v1alpha1_provider_proto_rawDescGZIP(), []int{1}
}

func (x *GetProviderConfigByDigestResponse) GetProvider() *v1alpha1.Provider {
	if x != nil {
		return x.Provider
	}
	return nil
}

var File_provider_v1alpha1_provider_proto protoreflect.FileDescriptor

var file_provider_v1alpha1_provider_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e, 0x67, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2f, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x47, 0x0a, 0x20, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x44, 0x69, 0x67, 0x65,
	0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x73, 0x68, 0x61,
	0x32, 0x35, 0x36, 0x5f, 0x64, 0x69, 0x67, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0c, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x22, 0x5b,
	0x0a, 0x21, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x79, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65,
	0x72, 0x52, 0x08, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x32, 0x9a, 0x01, 0x0a, 0x0f,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x86, 0x01, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x12, 0x33, 0x2e,
	0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x43, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x42, 0x79, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x34, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x42, 0x79, 0x44, 0x69, 0x67, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xdd, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d,
	0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x42, 0x0d, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x66, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f,
	0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x11, 0x50, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca,
	0x02, 0x11, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xe2, 0x02, 0x1d, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x56,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64,
	0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_provider_v1alpha1_provider_proto_rawDescOnce sync.Once
	file_provider_v1alpha1_provider_proto_rawDescData = file_provider_v1alpha1_provider_proto_rawDesc
)

func file_provider_v1alpha1_provider_proto_rawDescGZIP() []byte {
	file_provider_v1alpha1_provider_proto_rawDescOnce.Do(func() {
		file_provider_v1alpha1_provider_proto_rawDescData = protoimpl.X.CompressGZIP(file_provider_v1alpha1_provider_proto_rawDescData)
	})
	return file_provider_v1alpha1_provider_proto_rawDescData
}

var file_provider_v1alpha1_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_provider_v1alpha1_provider_proto_goTypes = []interface{}{
	(*GetProviderConfigByDigestRequest)(nil),  // 0: provider.v1alpha1.GetProviderConfigByDigestRequest
	(*GetProviderConfigByDigestResponse)(nil), // 1: provider.v1alpha1.GetProviderConfigByDigestResponse
	(*v1alpha1.Provider)(nil),                 // 2: gconfig.v1alpha1.Provider
}
var file_provider_v1alpha1_provider_proto_depIdxs = []int32{
	2, // 0: provider.v1alpha1.GetProviderConfigByDigestResponse.provider:type_name -> gconfig.v1alpha1.Provider
	0, // 1: provider.v1alpha1.ProviderService.GetProviderConfigByDigest:input_type -> provider.v1alpha1.GetProviderConfigByDigestRequest
	1, // 2: provider.v1alpha1.ProviderService.GetProviderConfigByDigest:output_type -> provider.v1alpha1.GetProviderConfigByDigestResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_provider_v1alpha1_provider_proto_init() }
func file_provider_v1alpha1_provider_proto_init() {
	if File_provider_v1alpha1_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_provider_v1alpha1_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProviderConfigByDigestRequest); i {
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
		file_provider_v1alpha1_provider_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProviderConfigByDigestResponse); i {
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
			RawDescriptor: file_provider_v1alpha1_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_provider_v1alpha1_provider_proto_goTypes,
		DependencyIndexes: file_provider_v1alpha1_provider_proto_depIdxs,
		MessageInfos:      file_provider_v1alpha1_provider_proto_msgTypes,
	}.Build()
	File_provider_v1alpha1_provider_proto = out.File
	file_provider_v1alpha1_provider_proto_rawDesc = nil
	file_provider_v1alpha1_provider_proto_goTypes = nil
	file_provider_v1alpha1_provider_proto_depIdxs = nil
}
