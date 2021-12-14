// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: certrail/v1alpha1/glog.proto

package certrailv1alpha1

import (
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
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
		mi := &file_certrail_v1alpha1_glog_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigPayload) ProtoMessage() {}

func (x *UpdateConfigPayload) ProtoReflect() protoreflect.Message {
	mi := &file_certrail_v1alpha1_glog_proto_msgTypes[0]
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
	return file_certrail_v1alpha1_glog_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateConfigPayload) GetConfigSha256() []byte {
	if x != nil {
		return x.ConfigSha256
	}
	return nil
}

type IssueCertificatePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificate []byte `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
}

func (x *IssueCertificatePayload) Reset() {
	*x = IssueCertificatePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certrail_v1alpha1_glog_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IssueCertificatePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IssueCertificatePayload) ProtoMessage() {}

func (x *IssueCertificatePayload) ProtoReflect() protoreflect.Message {
	mi := &file_certrail_v1alpha1_glog_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IssueCertificatePayload.ProtoReflect.Descriptor instead.
func (*IssueCertificatePayload) Descriptor() ([]byte, []int) {
	return file_certrail_v1alpha1_glog_proto_rawDescGZIP(), []int{1}
}

func (x *IssueCertificatePayload) GetCertificate() []byte {
	if x != nil {
		return x.Certificate
	}
	return nil
}

type RevokeCertificatePayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Certificate []byte `protobuf:"bytes,1,opt,name=certificate,proto3" json:"certificate,omitempty"`
	// the certificate of the actor who is performing the revoking
	RevokedBy []byte `protobuf:"bytes,2,opt,name=revoked_by,json=revokedBy,proto3" json:"revoked_by,omitempty"`
}

func (x *RevokeCertificatePayload) Reset() {
	*x = RevokeCertificatePayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certrail_v1alpha1_glog_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RevokeCertificatePayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RevokeCertificatePayload) ProtoMessage() {}

func (x *RevokeCertificatePayload) ProtoReflect() protoreflect.Message {
	mi := &file_certrail_v1alpha1_glog_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RevokeCertificatePayload.ProtoReflect.Descriptor instead.
func (*RevokeCertificatePayload) Descriptor() ([]byte, []int) {
	return file_certrail_v1alpha1_glog_proto_rawDescGZIP(), []int{2}
}

func (x *RevokeCertificatePayload) GetCertificate() []byte {
	if x != nil {
		return x.Certificate
	}
	return nil
}

func (x *RevokeCertificatePayload) GetRevokedBy() []byte {
	if x != nil {
		return x.RevokedBy
	}
	return nil
}

type Envelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Payload   *Payload `protobuf:"bytes,1,opt,name=payload,proto3" json:"payload,omitempty"`
	Signature []byte   `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (x *Envelope) Reset() {
	*x = Envelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certrail_v1alpha1_glog_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Envelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Envelope) ProtoMessage() {}

func (x *Envelope) ProtoReflect() protoreflect.Message {
	mi := &file_certrail_v1alpha1_glog_proto_msgTypes[3]
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
	return file_certrail_v1alpha1_glog_proto_rawDescGZIP(), []int{3}
}

func (x *Envelope) GetPayload() *Payload {
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

type Payload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp *timestamppb.Timestamp `protobuf:"bytes,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Types that are assignable to Contents:
	//	*Payload_UpdateConfig
	//	*Payload_IssueCertificate
	//	*Payload_RevokeCertificate
	Contents isPayload_Contents `protobuf_oneof:"contents"`
}

func (x *Payload) Reset() {
	*x = Payload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certrail_v1alpha1_glog_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Payload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Payload) ProtoMessage() {}

func (x *Payload) ProtoReflect() protoreflect.Message {
	mi := &file_certrail_v1alpha1_glog_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Payload.ProtoReflect.Descriptor instead.
func (*Payload) Descriptor() ([]byte, []int) {
	return file_certrail_v1alpha1_glog_proto_rawDescGZIP(), []int{4}
}

func (x *Payload) GetTimestamp() *timestamppb.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (m *Payload) GetContents() isPayload_Contents {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (x *Payload) GetUpdateConfig() *UpdateConfigPayload {
	if x, ok := x.GetContents().(*Payload_UpdateConfig); ok {
		return x.UpdateConfig
	}
	return nil
}

func (x *Payload) GetIssueCertificate() *IssueCertificatePayload {
	if x, ok := x.GetContents().(*Payload_IssueCertificate); ok {
		return x.IssueCertificate
	}
	return nil
}

func (x *Payload) GetRevokeCertificate() *RevokeCertificatePayload {
	if x, ok := x.GetContents().(*Payload_RevokeCertificate); ok {
		return x.RevokeCertificate
	}
	return nil
}

type isPayload_Contents interface {
	isPayload_Contents()
}

type Payload_UpdateConfig struct {
	UpdateConfig *UpdateConfigPayload `protobuf:"bytes,2,opt,name=update_config,json=updateConfig,proto3,oneof"`
}

type Payload_IssueCertificate struct {
	IssueCertificate *IssueCertificatePayload `protobuf:"bytes,3,opt,name=issue_certificate,json=issueCertificate,proto3,oneof"`
}

type Payload_RevokeCertificate struct {
	RevokeCertificate *RevokeCertificatePayload `protobuf:"bytes,4,opt,name=revoke_certificate,json=revokeCertificate,proto3,oneof"`
}

func (*Payload_UpdateConfig) isPayload_Contents() {}

func (*Payload_IssueCertificate) isPayload_Contents() {}

func (*Payload_RevokeCertificate) isPayload_Contents() {}

type StoreRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Envelope *Envelope `protobuf:"bytes,1,opt,name=envelope,proto3" json:"envelope,omitempty"`
}

func (x *StoreRequest) Reset() {
	*x = StoreRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certrail_v1alpha1_glog_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreRequest) ProtoMessage() {}

func (x *StoreRequest) ProtoReflect() protoreflect.Message {
	mi := &file_certrail_v1alpha1_glog_proto_msgTypes[5]
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
	return file_certrail_v1alpha1_glog_proto_rawDescGZIP(), []int{5}
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
		mi := &file_certrail_v1alpha1_glog_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StoreResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StoreResponse) ProtoMessage() {}

func (x *StoreResponse) ProtoReflect() protoreflect.Message {
	mi := &file_certrail_v1alpha1_glog_proto_msgTypes[6]
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
	return file_certrail_v1alpha1_glog_proto_rawDescGZIP(), []int{6}
}

type GetEntriesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StartIndex int64 `protobuf:"varint,1,opt,name=start_index,json=startIndex,proto3" json:"start_index,omitempty"`
	Count      int64 `protobuf:"varint,2,opt,name=count,proto3" json:"count,omitempty"`
}

func (x *GetEntriesRequest) Reset() {
	*x = GetEntriesRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certrail_v1alpha1_glog_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntriesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntriesRequest) ProtoMessage() {}

func (x *GetEntriesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_certrail_v1alpha1_glog_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntriesRequest.ProtoReflect.Descriptor instead.
func (*GetEntriesRequest) Descriptor() ([]byte, []int) {
	return file_certrail_v1alpha1_glog_proto_rawDescGZIP(), []int{7}
}

func (x *GetEntriesRequest) GetStartIndex() int64 {
	if x != nil {
		return x.StartIndex
	}
	return 0
}

func (x *GetEntriesRequest) GetCount() int64 {
	if x != nil {
		return x.Count
	}
	return 0
}

type GetEntriesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Envelopes []*IncludedEnvelope `protobuf:"bytes,1,rep,name=envelopes,proto3" json:"envelopes,omitempty"`
}

func (x *GetEntriesResponse) Reset() {
	*x = GetEntriesResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certrail_v1alpha1_glog_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEntriesResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEntriesResponse) ProtoMessage() {}

func (x *GetEntriesResponse) ProtoReflect() protoreflect.Message {
	mi := &file_certrail_v1alpha1_glog_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEntriesResponse.ProtoReflect.Descriptor instead.
func (*GetEntriesResponse) Descriptor() ([]byte, []int) {
	return file_certrail_v1alpha1_glog_proto_rawDescGZIP(), []int{8}
}

func (x *GetEntriesResponse) GetEnvelopes() []*IncludedEnvelope {
	if x != nil {
		return x.Envelopes
	}
	return nil
}

type IncludedEnvelope struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Envelope *Envelope `protobuf:"bytes,1,opt,name=envelope,proto3" json:"envelope,omitempty"`
}

func (x *IncludedEnvelope) Reset() {
	*x = IncludedEnvelope{}
	if protoimpl.UnsafeEnabled {
		mi := &file_certrail_v1alpha1_glog_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IncludedEnvelope) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IncludedEnvelope) ProtoMessage() {}

func (x *IncludedEnvelope) ProtoReflect() protoreflect.Message {
	mi := &file_certrail_v1alpha1_glog_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IncludedEnvelope.ProtoReflect.Descriptor instead.
func (*IncludedEnvelope) Descriptor() ([]byte, []int) {
	return file_certrail_v1alpha1_glog_proto_rawDescGZIP(), []int{9}
}

func (x *IncludedEnvelope) GetEnvelope() *Envelope {
	if x != nil {
		return x.Envelope
	}
	return nil
}

var File_certrail_v1alpha1_glog_proto protoreflect.FileDescriptor

var file_certrail_v1alpha1_glog_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2f, 0x67, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x11,
	0x63, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x17, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x2f, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x13, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x12, 0x2c, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x73, 0x68, 0x61,
	0x32, 0x35, 0x36, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x42, 0x07, 0xfa, 0x42, 0x04, 0x7a, 0x02,
	0x68, 0x20, 0x52, 0x0c, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x53, 0x68, 0x61, 0x32, 0x35, 0x36,
	0x22, 0x3b, 0x0a, 0x17, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0b, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x22, 0x5b, 0x0a,
	0x18, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x65, 0x72,
	0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0b,
	0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x72,
	0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x5f, 0x62, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x09, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x64, 0x42, 0x79, 0x22, 0x68, 0x0a, 0x08, 0x45, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x72, 0x61,
	0x69, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x50, 0x61, 0x79, 0x6c,
	0x6f, 0x61, 0x64, 0x52, 0x07, 0x70, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12, 0x26, 0x0a, 0x09,
	0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x42,
	0x08, 0xfa, 0x42, 0x05, 0x7a, 0x03, 0x18, 0x80, 0x02, 0x52, 0x09, 0x73, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x22, 0xd7, 0x02, 0x0a, 0x07, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x4d, 0x0a, 0x0d, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x26, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52, 0x0c, 0x75, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x59, 0x0a, 0x11, 0x69, 0x73, 0x73,
	0x75, 0x65, 0x5f, 0x63, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x2a, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x43, 0x65,
	0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64,
	0x48, 0x00, 0x52, 0x10, 0x69, 0x73, 0x73, 0x75, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69,
	0x63, 0x61, 0x74, 0x65, 0x12, 0x5c, 0x0a, 0x12, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x5f, 0x63,
	0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x2b, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x52, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69,
	0x66, 0x69, 0x63, 0x61, 0x74, 0x65, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x48, 0x00, 0x52,
	0x11, 0x72, 0x65, 0x76, 0x6f, 0x6b, 0x65, 0x43, 0x65, 0x72, 0x74, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x65, 0x42, 0x0a, 0x0a, 0x08, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x73, 0x22, 0x47,
	0x0a, 0x0c, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x37,
	0x0a, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c,
	0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x08, 0x65,
	0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x22, 0x0f, 0x0a, 0x0d, 0x53, 0x74, 0x6f, 0x72, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x4a, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x45,
	0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a,
	0x0b, 0x73, 0x74, 0x61, 0x72, 0x74, 0x5f, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x72, 0x74, 0x49, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x14,
	0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x57, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69,
	0x65, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x09, 0x65, 0x6e,
	0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23, 0x2e,
	0x63, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f,
	0x70, 0x65, 0x52, 0x09, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x73, 0x22, 0x4b, 0x0a,
	0x10, 0x49, 0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70,
	0x65, 0x12, 0x37, 0x0a, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65,
	0x52, 0x08, 0x65, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x32, 0xb8, 0x01, 0x0a, 0x0f, 0x43,
	0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4a,
	0x0a, 0x05, 0x53, 0x74, 0x6f, 0x72, 0x65, 0x12, 0x1f, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x72, 0x61,
	0x69, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x6f, 0x72,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x72,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x53, 0x74, 0x6f,
	0x72, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x59, 0x0a, 0x0a, 0x47, 0x65,
	0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x12, 0x24, 0x2e, 0x63, 0x65, 0x72, 0x74, 0x72,
	0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74,
	0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x25,
	0x2e, 0x63, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68,
	0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x69, 0x65, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xd9, 0x01, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x65,
	0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42,
	0x09, 0x47, 0x6c, 0x6f, 0x67, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x50, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d,
	0x66, 0x61, 0x74, 0x65, 0x2f, 0x63, 0x66, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67,
	0x65, 0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x63, 0x65, 0x72, 0x74,
	0x72, 0x61, 0x69, 0x6c, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x63, 0x65,
	0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02,
	0x03, 0x43, 0x58, 0x58, 0xaa, 0x02, 0x11, 0x43, 0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x2e,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x11, 0x43, 0x65, 0x72, 0x74, 0x72,
	0x61, 0x69, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x1d, 0x43,
	0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x43,
	0x65, 0x72, 0x74, 0x72, 0x61, 0x69, 0x6c, 0x3a, 0x3a, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_certrail_v1alpha1_glog_proto_rawDescOnce sync.Once
	file_certrail_v1alpha1_glog_proto_rawDescData = file_certrail_v1alpha1_glog_proto_rawDesc
)

func file_certrail_v1alpha1_glog_proto_rawDescGZIP() []byte {
	file_certrail_v1alpha1_glog_proto_rawDescOnce.Do(func() {
		file_certrail_v1alpha1_glog_proto_rawDescData = protoimpl.X.CompressGZIP(file_certrail_v1alpha1_glog_proto_rawDescData)
	})
	return file_certrail_v1alpha1_glog_proto_rawDescData
}

var file_certrail_v1alpha1_glog_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_certrail_v1alpha1_glog_proto_goTypes = []interface{}{
	(*UpdateConfigPayload)(nil),      // 0: certrail.v1alpha1.UpdateConfigPayload
	(*IssueCertificatePayload)(nil),  // 1: certrail.v1alpha1.IssueCertificatePayload
	(*RevokeCertificatePayload)(nil), // 2: certrail.v1alpha1.RevokeCertificatePayload
	(*Envelope)(nil),                 // 3: certrail.v1alpha1.Envelope
	(*Payload)(nil),                  // 4: certrail.v1alpha1.Payload
	(*StoreRequest)(nil),             // 5: certrail.v1alpha1.StoreRequest
	(*StoreResponse)(nil),            // 6: certrail.v1alpha1.StoreResponse
	(*GetEntriesRequest)(nil),        // 7: certrail.v1alpha1.GetEntriesRequest
	(*GetEntriesResponse)(nil),       // 8: certrail.v1alpha1.GetEntriesResponse
	(*IncludedEnvelope)(nil),         // 9: certrail.v1alpha1.IncludedEnvelope
	(*timestamppb.Timestamp)(nil),    // 10: google.protobuf.Timestamp
}
var file_certrail_v1alpha1_glog_proto_depIdxs = []int32{
	4,  // 0: certrail.v1alpha1.Envelope.payload:type_name -> certrail.v1alpha1.Payload
	10, // 1: certrail.v1alpha1.Payload.timestamp:type_name -> google.protobuf.Timestamp
	0,  // 2: certrail.v1alpha1.Payload.update_config:type_name -> certrail.v1alpha1.UpdateConfigPayload
	1,  // 3: certrail.v1alpha1.Payload.issue_certificate:type_name -> certrail.v1alpha1.IssueCertificatePayload
	2,  // 4: certrail.v1alpha1.Payload.revoke_certificate:type_name -> certrail.v1alpha1.RevokeCertificatePayload
	3,  // 5: certrail.v1alpha1.StoreRequest.envelope:type_name -> certrail.v1alpha1.Envelope
	9,  // 6: certrail.v1alpha1.GetEntriesResponse.envelopes:type_name -> certrail.v1alpha1.IncludedEnvelope
	3,  // 7: certrail.v1alpha1.IncludedEnvelope.envelope:type_name -> certrail.v1alpha1.Envelope
	5,  // 8: certrail.v1alpha1.CertrailService.Store:input_type -> certrail.v1alpha1.StoreRequest
	7,  // 9: certrail.v1alpha1.CertrailService.GetEntries:input_type -> certrail.v1alpha1.GetEntriesRequest
	6,  // 10: certrail.v1alpha1.CertrailService.Store:output_type -> certrail.v1alpha1.StoreResponse
	8,  // 11: certrail.v1alpha1.CertrailService.GetEntries:output_type -> certrail.v1alpha1.GetEntriesResponse
	10, // [10:12] is the sub-list for method output_type
	8,  // [8:10] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_certrail_v1alpha1_glog_proto_init() }
func file_certrail_v1alpha1_glog_proto_init() {
	if File_certrail_v1alpha1_glog_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_certrail_v1alpha1_glog_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_certrail_v1alpha1_glog_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IssueCertificatePayload); i {
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
		file_certrail_v1alpha1_glog_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RevokeCertificatePayload); i {
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
		file_certrail_v1alpha1_glog_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
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
		file_certrail_v1alpha1_glog_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Payload); i {
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
		file_certrail_v1alpha1_glog_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
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
		file_certrail_v1alpha1_glog_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
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
		file_certrail_v1alpha1_glog_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntriesRequest); i {
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
		file_certrail_v1alpha1_glog_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEntriesResponse); i {
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
		file_certrail_v1alpha1_glog_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IncludedEnvelope); i {
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
	file_certrail_v1alpha1_glog_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*Payload_UpdateConfig)(nil),
		(*Payload_IssueCertificate)(nil),
		(*Payload_RevokeCertificate)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_certrail_v1alpha1_glog_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_certrail_v1alpha1_glog_proto_goTypes,
		DependencyIndexes: file_certrail_v1alpha1_glog_proto_depIdxs,
		MessageInfos:      file_certrail_v1alpha1_glog_proto_msgTypes,
	}.Build()
	File_certrail_v1alpha1_glog_proto = out.File
	file_certrail_v1alpha1_glog_proto_rawDesc = nil
	file_certrail_v1alpha1_glog_proto_goTypes = nil
	file_certrail_v1alpha1_glog_proto_depIdxs = nil
}
