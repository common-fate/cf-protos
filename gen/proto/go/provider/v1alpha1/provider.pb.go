// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: provider/v1alpha1/provider.proto

package providerv1alpha1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type EnrolRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EnrolToken string `protobuf:"bytes,1,opt,name=enrol_token,json=enrolToken,proto3" json:"enrol_token,omitempty"`
	Csr        []byte `protobuf:"bytes,2,opt,name=csr,proto3" json:"csr,omitempty"`
	// Types that are assignable to Enrollment:
	//	*EnrolRequest_Aws
	Enrollment isEnrolRequest_Enrollment `protobuf_oneof:"enrollment"`
}

func (x *EnrolRequest) Reset() {
	*x = EnrolRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_v1alpha1_provider_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrolRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrolRequest) ProtoMessage() {}

func (x *EnrolRequest) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EnrolRequest.ProtoReflect.Descriptor instead.
func (*EnrolRequest) Descriptor() ([]byte, []int) {
	return file_provider_v1alpha1_provider_proto_rawDescGZIP(), []int{0}
}

func (x *EnrolRequest) GetEnrolToken() string {
	if x != nil {
		return x.EnrolToken
	}
	return ""
}

func (x *EnrolRequest) GetCsr() []byte {
	if x != nil {
		return x.Csr
	}
	return nil
}

func (m *EnrolRequest) GetEnrollment() isEnrolRequest_Enrollment {
	if m != nil {
		return m.Enrollment
	}
	return nil
}

func (x *EnrolRequest) GetAws() *AWSEnrollment {
	if x, ok := x.GetEnrollment().(*EnrolRequest_Aws); ok {
		return x.Aws
	}
	return nil
}

type isEnrolRequest_Enrollment interface {
	isEnrolRequest_Enrollment()
}

type EnrolRequest_Aws struct {
	Aws *AWSEnrollment `protobuf:"bytes,3,opt,name=aws,proto3,oneof"`
}

func (*EnrolRequest_Aws) isEnrolRequest_Enrollment() {}

type EnrolResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *EnrolResponse) Reset() {
	*x = EnrolResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_v1alpha1_provider_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrolResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrolResponse) ProtoMessage() {}

func (x *EnrolResponse) ProtoReflect() protoreflect.Message {
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

// Deprecated: Use EnrolResponse.ProtoReflect.Descriptor instead.
func (*EnrolResponse) Descriptor() ([]byte, []int) {
	return file_provider_v1alpha1_provider_proto_rawDescGZIP(), []int{1}
}

type AWSEnrollment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Identity     *AWSProof `protobuf:"bytes,1,opt,name=identity,proto3" json:"identity,omitempty"`
	Organization *AWSProof `protobuf:"bytes,2,opt,name=organization,proto3" json:"organization,omitempty"`
}

func (x *AWSEnrollment) Reset() {
	*x = AWSEnrollment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_v1alpha1_provider_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSEnrollment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSEnrollment) ProtoMessage() {}

func (x *AWSEnrollment) ProtoReflect() protoreflect.Message {
	mi := &file_provider_v1alpha1_provider_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSEnrollment.ProtoReflect.Descriptor instead.
func (*AWSEnrollment) Descriptor() ([]byte, []int) {
	return file_provider_v1alpha1_provider_proto_rawDescGZIP(), []int{2}
}

func (x *AWSEnrollment) GetIdentity() *AWSProof {
	if x != nil {
		return x.Identity
	}
	return nil
}

func (x *AWSEnrollment) GetOrganization() *AWSProof {
	if x != nil {
		return x.Organization
	}
	return nil
}

type AWSProof struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Signature     string                 `protobuf:"bytes,1,opt,name=signature,proto3" json:"signature,omitempty"`
	Timestamp     *timestamppb.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	SecurityToken string                 `protobuf:"bytes,3,opt,name=security_token,json=securityToken,proto3" json:"security_token,omitempty"`
}

func (x *AWSProof) Reset() {
	*x = AWSProof{}
	if protoimpl.UnsafeEnabled {
		mi := &file_provider_v1alpha1_provider_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AWSProof) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AWSProof) ProtoMessage() {}

func (x *AWSProof) ProtoReflect() protoreflect.Message {
	mi := &file_provider_v1alpha1_provider_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AWSProof.ProtoReflect.Descriptor instead.
func (*AWSProof) Descriptor() ([]byte, []int) {
	return file_provider_v1alpha1_provider_proto_rawDescGZIP(), []int{3}
}

func (x *AWSProof) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *AWSProof) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *AWSProof) GetSecurityToken() string {
	if x != nil {
		return x.SecurityToken
	}
	return ""
}

var File_provider_v1alpha1_provider_proto protoreflect.FileDescriptor

var file_provider_v1alpha1_provider_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x11, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x85, 0x01, 0x0a, 0x0c, 0x45, 0x6e, 0x72, 0x6f, 0x6c,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x6e, 0x72, 0x6f, 0x6c,
	0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x6e,
	0x72, 0x6f, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x10, 0x0a, 0x03, 0x63, 0x73, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x03, 0x63, 0x73, 0x72, 0x12, 0x34, 0x0a, 0x03, 0x61, 0x77,
	0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64,
	0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x45,
	0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x03, 0x61, 0x77, 0x73,
	0x42, 0x0c, 0x0a, 0x0a, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x0f,
	0x0a, 0x0d, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0x89, 0x01, 0x0a, 0x0d, 0x41, 0x57, 0x53, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x6c, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x37, 0x0a, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x50, 0x72, 0x6f, 0x6f, 0x66,
	0x52, 0x08, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x3f, 0x0a, 0x0c, 0x6f, 0x72,
	0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x41, 0x57, 0x53, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x0c, 0x6f,
	0x72, 0x67, 0x61, 0x6e, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x89, 0x01, 0x0a, 0x08,
	0x41, 0x57, 0x53, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x69, 0x67,
	0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x12, 0x25, 0x0a, 0x0e, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69, 0x74, 0x79, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x73, 0x65, 0x63, 0x75, 0x72, 0x69,
	0x74, 0x79, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0x5d, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a, 0x0a, 0x05, 0x45, 0x6e,
	0x72, 0x6f, 0x6c, 0x12, 0x1f, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xdd, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x42, 0x0d, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x50, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x66, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x3b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58, 0xaa, 0x02, 0x11, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x64, 0x65, 0x72, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x11,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xe2, 0x02, 0x1d, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74,
	0x61, 0xea, 0x02, 0x12, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x3a, 0x3a, 0x56, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
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

var file_provider_v1alpha1_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_provider_v1alpha1_provider_proto_goTypes = []interface{}{
	(*EnrolRequest)(nil),          // 0: provider.v1alpha1.EnrolRequest
	(*EnrolResponse)(nil),         // 1: provider.v1alpha1.EnrolResponse
	(*AWSEnrollment)(nil),         // 2: provider.v1alpha1.AWSEnrollment
	(*AWSProof)(nil),              // 3: provider.v1alpha1.AWSProof
	(*timestamppb.Timestamp)(nil), // 4: google.protobuf.Timestamp
}
var file_provider_v1alpha1_provider_proto_depIdxs = []int32{
	2, // 0: provider.v1alpha1.EnrolRequest.aws:type_name -> provider.v1alpha1.AWSEnrollment
	3, // 1: provider.v1alpha1.AWSEnrollment.identity:type_name -> provider.v1alpha1.AWSProof
	3, // 2: provider.v1alpha1.AWSEnrollment.organization:type_name -> provider.v1alpha1.AWSProof
	4, // 3: provider.v1alpha1.AWSProof.timestamp:type_name -> google.protobuf.Timestamp
	0, // 4: provider.v1alpha1.ProviderService.Enrol:input_type -> provider.v1alpha1.EnrolRequest
	1, // 5: provider.v1alpha1.ProviderService.Enrol:output_type -> provider.v1alpha1.EnrolResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_provider_v1alpha1_provider_proto_init() }
func file_provider_v1alpha1_provider_proto_init() {
	if File_provider_v1alpha1_provider_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_provider_v1alpha1_provider_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrolRequest); i {
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
			switch v := v.(*EnrolResponse); i {
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
		file_provider_v1alpha1_provider_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSEnrollment); i {
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
		file_provider_v1alpha1_provider_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AWSProof); i {
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
	file_provider_v1alpha1_provider_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*EnrolRequest_Aws)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_provider_v1alpha1_provider_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
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