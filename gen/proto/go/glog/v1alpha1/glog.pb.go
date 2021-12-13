// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: glog/v1alpha1/glog.proto

package glogv1alpha1

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

type UpdateConfigPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ConfigSha256 []byte `protobuf:"bytes,1,opt,name=config_sha256,json=configSha256,proto3" json:"config_sha256,omitempty"`
}

func (x *UpdateConfigPayload) Reset() {
	*x = UpdateConfigPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_glog_v1alpha1_glog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigPayload) ProtoMessage() {}

func (x *UpdateConfigPayload) ProtoReflect() protoreflect.Message {
	mi := &file_glog_v1alpha1_glog_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigPayload.ProtoReflect.Descriptor instead.
func (*UpdateConfigPayload) Descriptor() ([]byte, []int) {
	return file_glog_v1alpha1_glog_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateConfigPayload) GetConfigSha256() []byte {
	if x != nil {
		return x.ConfigSha256
	}
	return nil
}

type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload   *TimestampedPayload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature []byte              `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_glog_v1alpha1_glog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_glog_v1alpha1_glog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Envelope.ProtoReflect.Descriptor instead.
func (*Envelope) Descriptor() ([]byte, []int) {
	return file_glog_v1alpha1_glog_proto_rawDescGZIP(), []int{1}
}

func (x *Envelope) GetPayload() *TimestampedPayload {
	if x != nil {
		return x.Payload
	}
	return nil
}

func (x *Envelope) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type TimestampedPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Contents:
	//	*TimestampedPayload_UpdateConfig
	Contents  isTimestampedPayload_Contents `protobuf_oneof:"contents"`
	Timestamp *timestamppb.Timestamp        `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *TimestampedPayload) Reset() {
	*x = TimestampedPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_glog_v1alpha1_glog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TimestampedPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TimestampedPayload) ProtoMessage() {}

func (x *TimestampedPayload) ProtoReflect() protoreflect.Message {
	mi := &file_glog_v1alpha1_glog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TimestampedPayload.ProtoReflect.Descriptor instead.
func (*TimestampedPayload) Descriptor() ([]byte, []int) {
	return file_glog_v1alpha1_glog_proto_rawDescGZIP(), []int{2}
}

func (m *TimestampedPayload) GetContents() isTimestampedPayload_Contents {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (x *TimestampedPayload) GetUpdateConfig() *UpdateConfigPayload {
	if x, ok := x.GetContents().(*TimestampedPayload_UpdateConfig); ok {
		return x.UpdateConfig
	}
	return nil
}

func (x *TimestampedPayload) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

type isTimestampedPayload_Contents interface {
	isTimestampedPayload_Contents()
}

type TimestampedPayload_UpdateConfig struct {
	UpdateConfig *UpdateConfigPayload `protobuf:"bytes,1,opt,name=update_config,json=updateConfig,proto3,oneof"`
}

func (*TimestampedPayload_UpdateConfig) isTimestampedPayload_Contents() {}

type StoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Envelope *Envelope `protobuf:"bytes,1,opt,name=envelope,proto3" json:"envelope,omitempty"`
}

func (x *StoreRequest) Reset() {
	*x = StoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_glog_v1alpha1_glog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreRequest) ProtoMessage() {}

func (x *StoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_glog_v1alpha1_glog_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreRequest.ProtoReflect.Descriptor instead.
func (*StoreRequest) Descriptor() ([]byte, []int) {
	return file_glog_v1alpha1_glog_proto_rawDescGZIP(), []int{3}
}

func (x *StoreRequest) GetEnvelope() *Envelope {
	if x != nil {
		return x.Envelope
	}
	return nil
}

type StoreResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *StoreResponse) Reset() {
	*x = StoreResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_glog_v1alpha1_glog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreResponse) ProtoMessage() {}

func (x *StoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_glog_v1alpha1_glog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StoreResponse.ProtoReflect.Descriptor instead.
func (*StoreResponse) Descriptor() ([]byte, []int) {
	return file_glog_v1alpha1_glog_proto_rawDescGZIP(), []int{4}
}

var File_glog_v1alpha1_glog_proto protoreflect.FileDescriptor

var file_glog_v1alpha1_glog_proto_rawDesc = []byte{
	0x0a, 0x18, 0x67, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x67, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x67, 0x6c, 0x6f, 0x67,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3a, 0x0a, 0x13, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x68, 0x61, 0x32,
	0x35, 0x36, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x53, 0x68, 0x61, 0x32, 0x35, 0x36, 0x22, 0x65, 0x0a, 0x08, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x12, 0x3b, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x67, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x65, 0x64, 0x50,
	0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x1c, 0x0a, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x22, 0xa5, 0x01,
	0x0a, 0x12, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x65, 0x64, 0x50, 0x61, 0x79,
	0x6c, 0x6f, 0x61, 0x64, 0x12, 0x49, 0x0a, 0x0d, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x5f, 0x63,
	0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x67, 0x6c,
	0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61,
	0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48,
	0x00, 0x52, 0x0c, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12,
	0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x42, 0x0a, 0x0a, 0x08, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x43, 0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x33, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x67, 0x6c, 0x6f, 0x67, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x52, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x32, 0x51, 0x0a, 0x0b, 0x47,
	0x6c, 0x6f, 0x67, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x05, 0x53, 0x74,
	0x6f, 0x72, 0x65, 0x12, 0x1b, 0x2e, 0x67, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x1c, 0x2e, 0x67, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xbd,
	0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d, 0x2e, 0x67, 0x6c, 0x6f, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x42, 0x09, 0x47, 0x6c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50,
	0x01, 0x5a, 0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x66, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f,
	0x2f, 0x67, 0x6c, 0x6f, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x67,
	0x6c, 0x6f, 0x67, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x47, 0x58,
	0x58, 0xaa, 0x02, 0x0d, 0x47, 0x6c, 0x6f, 0x67, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xca, 0x02, 0x0d, 0x47, 0x6c, 0x6f, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0xe2, 0x02, 0x19, 0x47, 0x6c, 0x6f, 0x67, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e,
	0x47, 0x6c, 0x6f, 0x67, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_glog_v1alpha1_glog_proto_rawDescOnce sync.Once
	file_glog_v1alpha1_glog_proto_rawDescData = file_glog_v1alpha1_glog_proto_rawDesc
)

func file_glog_v1alpha1_glog_proto_rawDescGZIP() []byte {
	file_glog_v1alpha1_glog_proto_rawDescOnce.Do(func() {
		file_glog_v1alpha1_glog_proto_rawDescData = protoimpl.X.CompressGZIP(file_glog_v1alpha1_glog_proto_rawDescData)
	})
	return file_glog_v1alpha1_glog_proto_rawDescData
}

var file_glog_v1alpha1_glog_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_glog_v1alpha1_glog_proto_goTypes = []interface{}{
	(*UpdateConfigPayload)(nil),   // 0: glog.v1alpha1.UpdateConfigPayload
	(*Envelope)(nil),              // 1: glog.v1alpha1.Envelope
	(*TimestampedPayload)(nil),    // 2: glog.v1alpha1.TimestampedPayload
	(*StoreRequest)(nil),          // 3: glog.v1alpha1.StoreRequest
	(*StoreResponse)(nil),         // 4: glog.v1alpha1.StoreResponse
	(*timestamppb.Timestamp)(nil), // 5: google.protobuf.Timestamp
}
var file_glog_v1alpha1_glog_proto_depIdxs = []int32{
	2, // 0: glog.v1alpha1.Envelope.payload:type_name -> glog.v1alpha1.TimestampedPayload
	0, // 1: glog.v1alpha1.TimestampedPayload.update_config:type_name -> glog.v1alpha1.UpdateConfigPayload
	5, // 2: glog.v1alpha1.TimestampedPayload.timestamp:type_name -> google.protobuf.Timestamp
	1, // 3: glog.v1alpha1.StoreRequest.envelope:type_name -> glog.v1alpha1.Envelope
	3, // 4: glog.v1alpha1.GlogService.Store:input_type -> glog.v1alpha1.StoreRequest
	4, // 5: glog.v1alpha1.GlogService.Store:output_type -> glog.v1alpha1.StoreResponse
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_glog_v1alpha1_glog_proto_init() }
func file_glog_v1alpha1_glog_proto_init() {
	if File_glog_v1alpha1_glog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_glog_v1alpha1_glog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigPayload); i {
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
		file_glog_v1alpha1_glog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Envelope); i {
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
		file_glog_v1alpha1_glog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TimestampedPayload); i {
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
		file_glog_v1alpha1_glog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreRequest); i {
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
		file_glog_v1alpha1_glog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StoreResponse); i {
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
	file_glog_v1alpha1_glog_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*TimestampedPayload_UpdateConfig)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_glog_v1alpha1_glog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_glog_v1alpha1_glog_proto_goTypes,
		DependencyIndexes: file_glog_v1alpha1_glog_proto_depIdxs,
		MessageInfos:      file_glog_v1alpha1_glog_proto_msgTypes,
	}.Build()
	File_glog_v1alpha1_glog_proto = out.File
	file_glog_v1alpha1_glog_proto_rawDesc = nil
	file_glog_v1alpha1_glog_proto_goTypes = nil
	file_glog_v1alpha1_glog_proto_depIdxs = nil
}
