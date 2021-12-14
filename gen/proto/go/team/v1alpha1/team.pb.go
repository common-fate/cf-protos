// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.1
// source: team/v1alpha1/team.proto

package teamv1alpha1

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

type ListMembersRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *ListMembersRequest) Reset() {
	*x = ListMembersRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_v1alpha1_team_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMembersRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMembersRequest) ProtoMessage() {}

func (x *ListMembersRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_v1alpha1_team_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMembersRequest.ProtoReflect.Descriptor instead.
func (*ListMembersRequest) Descriptor() ([]byte, []int) {
	return file_team_v1alpha1_team_proto_rawDescGZIP(), []int{0}
}

type ListMembersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Members []*Member `protobuf:"bytes,1,rep,name=members,proto3" json:"members,omitempty"`
}

func (x *ListMembersResponse) Reset() {
	*x = ListMembersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_v1alpha1_team_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListMembersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListMembersResponse) ProtoMessage() {}

func (x *ListMembersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_team_v1alpha1_team_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListMembersResponse.ProtoReflect.Descriptor instead.
func (*ListMembersResponse) Descriptor() ([]byte, []int) {
	return file_team_v1alpha1_team_proto_rawDescGZIP(), []int{1}
}

func (x *ListMembersResponse) GetMembers() []*Member {
	if x != nil {
		return x.Members
	}
	return nil
}

type Member struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EmailAddress string `protobuf:"bytes,1,opt,name=email_address,json=emailAddress,proto3" json:"email_address,omitempty"`
	IsAdmin      bool   `protobuf:"varint,2,opt,name=is_admin,json=isAdmin,proto3" json:"is_admin,omitempty"`
}

func (x *Member) Reset() {
	*x = Member{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_v1alpha1_team_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Member) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Member) ProtoMessage() {}

func (x *Member) ProtoReflect() protoreflect.Message {
	mi := &file_team_v1alpha1_team_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Member.ProtoReflect.Descriptor instead.
func (*Member) Descriptor() ([]byte, []int) {
	return file_team_v1alpha1_team_proto_rawDescGZIP(), []int{2}
}

func (x *Member) GetEmailAddress() string {
	if x != nil {
		return x.EmailAddress
	}
	return ""
}

func (x *Member) GetIsAdmin() bool {
	if x != nil {
		return x.IsAdmin
	}
	return false
}

type UpdateConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *v1alpha1.Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *UpdateConfigRequest) Reset() {
	*x = UpdateConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_v1alpha1_team_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigRequest) ProtoMessage() {}

func (x *UpdateConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_v1alpha1_team_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigRequest.ProtoReflect.Descriptor instead.
func (*UpdateConfigRequest) Descriptor() ([]byte, []int) {
	return file_team_v1alpha1_team_proto_rawDescGZIP(), []int{3}
}

func (x *UpdateConfigRequest) GetConfig() *v1alpha1.Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type UpdateConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Sha256Hash []byte `protobuf:"bytes,1,opt,name=sha256_hash,json=sha256Hash,proto3" json:"sha256_hash,omitempty"`
}

func (x *UpdateConfigResponse) Reset() {
	*x = UpdateConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_v1alpha1_team_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateConfigResponse) ProtoMessage() {}

func (x *UpdateConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_team_v1alpha1_team_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateConfigResponse.ProtoReflect.Descriptor instead.
func (*UpdateConfigResponse) Descriptor() ([]byte, []int) {
	return file_team_v1alpha1_team_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateConfigResponse) GetSha256Hash() []byte {
	if x != nil {
		return x.Sha256Hash
	}
	return nil
}

type GetConfigRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetConfigRequest) Reset() {
	*x = GetConfigRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_v1alpha1_team_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigRequest) ProtoMessage() {}

func (x *GetConfigRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_v1alpha1_team_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigRequest.ProtoReflect.Descriptor instead.
func (*GetConfigRequest) Descriptor() ([]byte, []int) {
	return file_team_v1alpha1_team_proto_rawDescGZIP(), []int{5}
}

type GetConfigResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Config *v1alpha1.Config `protobuf:"bytes,1,opt,name=config,proto3" json:"config,omitempty"`
}

func (x *GetConfigResponse) Reset() {
	*x = GetConfigResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_v1alpha1_team_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetConfigResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetConfigResponse) ProtoMessage() {}

func (x *GetConfigResponse) ProtoReflect() protoreflect.Message {
	mi := &file_team_v1alpha1_team_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetConfigResponse.ProtoReflect.Descriptor instead.
func (*GetConfigResponse) Descriptor() ([]byte, []int) {
	return file_team_v1alpha1_team_proto_rawDescGZIP(), []int{6}
}

func (x *GetConfigResponse) GetConfig() *v1alpha1.Config {
	if x != nil {
		return x.Config
	}
	return nil
}

type EnrolAWSProviderIntervention struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProviderId        string `protobuf:"bytes,1,opt,name=provider_id,json=providerId,proto3" json:"provider_id,omitempty"`
	AwsAccountId      string `protobuf:"bytes,2,opt,name=aws_account_id,json=awsAccountId,proto3" json:"aws_account_id,omitempty"`
	CloudformationUrl string `protobuf:"bytes,3,opt,name=cloudformation_url,json=cloudformationUrl,proto3" json:"cloudformation_url,omitempty"`
}

func (x *EnrolAWSProviderIntervention) Reset() {
	*x = EnrolAWSProviderIntervention{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_v1alpha1_team_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EnrolAWSProviderIntervention) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EnrolAWSProviderIntervention) ProtoMessage() {}

func (x *EnrolAWSProviderIntervention) ProtoReflect() protoreflect.Message {
	mi := &file_team_v1alpha1_team_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EnrolAWSProviderIntervention.ProtoReflect.Descriptor instead.
func (*EnrolAWSProviderIntervention) Descriptor() ([]byte, []int) {
	return file_team_v1alpha1_team_proto_rawDescGZIP(), []int{7}
}

func (x *EnrolAWSProviderIntervention) GetProviderId() string {
	if x != nil {
		return x.ProviderId
	}
	return ""
}

func (x *EnrolAWSProviderIntervention) GetAwsAccountId() string {
	if x != nil {
		return x.AwsAccountId
	}
	return ""
}

func (x *EnrolAWSProviderIntervention) GetCloudformationUrl() string {
	if x != nil {
		return x.CloudformationUrl
	}
	return ""
}

type GetInterventionsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetInterventionsRequest) Reset() {
	*x = GetInterventionsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_v1alpha1_team_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInterventionsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInterventionsRequest) ProtoMessage() {}

func (x *GetInterventionsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_team_v1alpha1_team_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInterventionsRequest.ProtoReflect.Descriptor instead.
func (*GetInterventionsRequest) Descriptor() ([]byte, []int) {
	return file_team_v1alpha1_team_proto_rawDescGZIP(), []int{8}
}

// Interventions are steps which we require the user to complete manually so that Granted works properly.
type Intervention struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//	*Intervention_EnrolAwsProvider
	Data isIntervention_Data `protobuf_oneof:"data"`
}

func (x *Intervention) Reset() {
	*x = Intervention{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_v1alpha1_team_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Intervention) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Intervention) ProtoMessage() {}

func (x *Intervention) ProtoReflect() protoreflect.Message {
	mi := &file_team_v1alpha1_team_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Intervention.ProtoReflect.Descriptor instead.
func (*Intervention) Descriptor() ([]byte, []int) {
	return file_team_v1alpha1_team_proto_rawDescGZIP(), []int{9}
}

func (m *Intervention) GetData() isIntervention_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *Intervention) GetEnrolAwsProvider() *EnrolAWSProviderIntervention {
	if x, ok := x.GetData().(*Intervention_EnrolAwsProvider); ok {
		return x.EnrolAwsProvider
	}
	return nil
}

type isIntervention_Data interface {
	isIntervention_Data()
}

type Intervention_EnrolAwsProvider struct {
	EnrolAwsProvider *EnrolAWSProviderIntervention `protobuf:"bytes,1,opt,name=enrol_aws_provider,json=enrolAwsProvider,proto3,oneof"`
}

func (*Intervention_EnrolAwsProvider) isIntervention_Data() {}

type GetInterventionsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Interventions []*Intervention `protobuf:"bytes,1,rep,name=interventions,proto3" json:"interventions,omitempty"`
}

func (x *GetInterventionsResponse) Reset() {
	*x = GetInterventionsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_team_v1alpha1_team_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetInterventionsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetInterventionsResponse) ProtoMessage() {}

func (x *GetInterventionsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_team_v1alpha1_team_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetInterventionsResponse.ProtoReflect.Descriptor instead.
func (*GetInterventionsResponse) Descriptor() ([]byte, []int) {
	return file_team_v1alpha1_team_proto_rawDescGZIP(), []int{10}
}

func (x *GetInterventionsResponse) GetInterventions() []*Intervention {
	if x != nil {
		return x.Interventions
	}
	return nil
}

var File_team_v1alpha1_team_proto protoreflect.FileDescriptor

var file_team_v1alpha1_team_proto_rawDesc = []byte{
	0x0a, 0x18, 0x74, 0x65, 0x61, 0x6d, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f,
	0x74, 0x65, 0x61, 0x6d, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x74, 0x65, 0x61, 0x6d,
	0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x1a, 0x1e, 0x67, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2f, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2f, 0x67, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x14, 0x0a, 0x12, 0x4c, 0x69, 0x73,
	0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22,
	0x46, 0x0a, 0x13, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2f, 0x0a, 0x07, 0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x52, 0x07,
	0x6d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x22, 0x48, 0x0a, 0x06, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x23, 0x0a, 0x0d, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x5f, 0x61, 0x64, 0x64, 0x72, 0x65,
	0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x61, 0x64, 0x6d,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69, 0x73, 0x41, 0x64, 0x6d, 0x69,
	0x6e, 0x22, 0x47, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x30, 0x0a, 0x06, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x37, 0x0a, 0x14, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x5f, 0x68, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x48,
	0x61, 0x73, 0x68, 0x22, 0x12, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x45, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x06,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x67,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x06, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x22, 0x94,
	0x01, 0x0a, 0x1c, 0x45, 0x6e, 0x72, 0x6f, 0x6c, 0x41, 0x57, 0x53, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x64, 0x65, 0x72, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x1f, 0x0a, 0x0b, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x24, 0x0a, 0x0e, 0x61, 0x77, 0x73, 0x5f, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x61, 0x77, 0x73, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x2d, 0x0a, 0x12, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x75, 0x72, 0x6c, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x11, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x55, 0x72, 0x6c, 0x22, 0x19, 0x0a, 0x17, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0x73, 0x0a, 0x0c, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x5b, 0x0a, 0x12, 0x65, 0x6e, 0x72, 0x6f, 0x6c, 0x5f, 0x61, 0x77, 0x73, 0x5f, 0x70, 0x72,
	0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x74,
	0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x45, 0x6e, 0x72,
	0x6f, 0x6c, 0x41, 0x57, 0x53, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x49, 0x6e, 0x74,
	0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x10, 0x65, 0x6e, 0x72,
	0x6f, 0x6c, 0x41, 0x77, 0x73, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x42, 0x06, 0x0a,
	0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0x5d, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65,
	0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x41, 0x0a, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65,
	0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x32, 0xf1, 0x02, 0x0a, 0x0b, 0x54, 0x65, 0x61, 0x6d, 0x53, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x54, 0x0a, 0x0b, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62,
	0x65, 0x72, 0x73, 0x12, 0x21, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70,
	0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65, 0x72, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x22, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31,
	0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x4d, 0x65, 0x6d, 0x62, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x57, 0x0a, 0x0c, 0x55, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22, 0x2e, 0x74, 0x65, 0x61,
	0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x23,
	0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x55,
	0x70, 0x64, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x4e, 0x0a, 0x09, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x12, 0x1f, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x20, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61,
	0x31, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x63, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76,
	0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x26, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x76,
	0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e, 0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72,
	0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x27, 0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x2e,
	0x47, 0x65, 0x74, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x76, 0x65, 0x6e, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0xbd, 0x01, 0x0a, 0x11, 0x63, 0x6f, 0x6d,
	0x2e, 0x74, 0x65, 0x61, 0x6d, 0x2e, 0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x42, 0x09,
	0x54, 0x65, 0x61, 0x6d, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x48, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2d, 0x66,
	0x61, 0x74, 0x65, 0x2f, 0x63, 0x66, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x67, 0x65,
	0x6e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x2f, 0x74, 0x65, 0x61, 0x6d, 0x2f,
	0x76, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x3b, 0x74, 0x65, 0x61, 0x6d, 0x76, 0x31, 0x61,
	0x6c, 0x70, 0x68, 0x61, 0x31, 0xa2, 0x02, 0x03, 0x54, 0x58, 0x58, 0xaa, 0x02, 0x0d, 0x54, 0x65,
	0x61, 0x6d, 0x2e, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xca, 0x02, 0x0d, 0x54, 0x65,
	0x61, 0x6d, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0xe2, 0x02, 0x19, 0x54, 0x65,
	0x61, 0x6d, 0x5c, 0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x0e, 0x54, 0x65, 0x61, 0x6d, 0x3a, 0x3a,
	0x56, 0x31, 0x61, 0x6c, 0x70, 0x68, 0x61, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_team_v1alpha1_team_proto_rawDescOnce sync.Once
	file_team_v1alpha1_team_proto_rawDescData = file_team_v1alpha1_team_proto_rawDesc
)

func file_team_v1alpha1_team_proto_rawDescGZIP() []byte {
	file_team_v1alpha1_team_proto_rawDescOnce.Do(func() {
		file_team_v1alpha1_team_proto_rawDescData = protoimpl.X.CompressGZIP(file_team_v1alpha1_team_proto_rawDescData)
	})
	return file_team_v1alpha1_team_proto_rawDescData
}

var file_team_v1alpha1_team_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_team_v1alpha1_team_proto_goTypes = []interface{}{
	(*ListMembersRequest)(nil),           // 0: team.v1alpha1.ListMembersRequest
	(*ListMembersResponse)(nil),          // 1: team.v1alpha1.ListMembersResponse
	(*Member)(nil),                       // 2: team.v1alpha1.Member
	(*UpdateConfigRequest)(nil),          // 3: team.v1alpha1.UpdateConfigRequest
	(*UpdateConfigResponse)(nil),         // 4: team.v1alpha1.UpdateConfigResponse
	(*GetConfigRequest)(nil),             // 5: team.v1alpha1.GetConfigRequest
	(*GetConfigResponse)(nil),            // 6: team.v1alpha1.GetConfigResponse
	(*EnrolAWSProviderIntervention)(nil), // 7: team.v1alpha1.EnrolAWSProviderIntervention
	(*GetInterventionsRequest)(nil),      // 8: team.v1alpha1.GetInterventionsRequest
	(*Intervention)(nil),                 // 9: team.v1alpha1.Intervention
	(*GetInterventionsResponse)(nil),     // 10: team.v1alpha1.GetInterventionsResponse
	(*v1alpha1.Config)(nil),              // 11: gconfig.v1alpha1.Config
}
var file_team_v1alpha1_team_proto_depIdxs = []int32{
	2,  // 0: team.v1alpha1.ListMembersResponse.members:type_name -> team.v1alpha1.Member
	11, // 1: team.v1alpha1.UpdateConfigRequest.config:type_name -> gconfig.v1alpha1.Config
	11, // 2: team.v1alpha1.GetConfigResponse.config:type_name -> gconfig.v1alpha1.Config
	7,  // 3: team.v1alpha1.Intervention.enrol_aws_provider:type_name -> team.v1alpha1.EnrolAWSProviderIntervention
	9,  // 4: team.v1alpha1.GetInterventionsResponse.interventions:type_name -> team.v1alpha1.Intervention
	0,  // 5: team.v1alpha1.TeamService.ListMembers:input_type -> team.v1alpha1.ListMembersRequest
	3,  // 6: team.v1alpha1.TeamService.UpdateConfig:input_type -> team.v1alpha1.UpdateConfigRequest
	5,  // 7: team.v1alpha1.TeamService.GetConfig:input_type -> team.v1alpha1.GetConfigRequest
	8,  // 8: team.v1alpha1.TeamService.GetInterventions:input_type -> team.v1alpha1.GetInterventionsRequest
	1,  // 9: team.v1alpha1.TeamService.ListMembers:output_type -> team.v1alpha1.ListMembersResponse
	4,  // 10: team.v1alpha1.TeamService.UpdateConfig:output_type -> team.v1alpha1.UpdateConfigResponse
	6,  // 11: team.v1alpha1.TeamService.GetConfig:output_type -> team.v1alpha1.GetConfigResponse
	10, // 12: team.v1alpha1.TeamService.GetInterventions:output_type -> team.v1alpha1.GetInterventionsResponse
	9,  // [9:13] is the sub-list for method output_type
	5,  // [5:9] is the sub-list for method input_type
	5,  // [5:5] is the sub-list for extension type_name
	5,  // [5:5] is the sub-list for extension extendee
	0,  // [0:5] is the sub-list for field type_name
}

func init() { file_team_v1alpha1_team_proto_init() }
func file_team_v1alpha1_team_proto_init() {
	if File_team_v1alpha1_team_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_team_v1alpha1_team_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMembersRequest); i {
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
		file_team_v1alpha1_team_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListMembersResponse); i {
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
		file_team_v1alpha1_team_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Member); i {
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
		file_team_v1alpha1_team_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigRequest); i {
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
		file_team_v1alpha1_team_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateConfigResponse); i {
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
		file_team_v1alpha1_team_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigRequest); i {
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
		file_team_v1alpha1_team_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetConfigResponse); i {
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
		file_team_v1alpha1_team_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EnrolAWSProviderIntervention); i {
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
		file_team_v1alpha1_team_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInterventionsRequest); i {
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
		file_team_v1alpha1_team_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Intervention); i {
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
		file_team_v1alpha1_team_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetInterventionsResponse); i {
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
	file_team_v1alpha1_team_proto_msgTypes[9].OneofWrappers = []interface{}{
		(*Intervention_EnrolAwsProvider)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_team_v1alpha1_team_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_team_v1alpha1_team_proto_goTypes,
		DependencyIndexes: file_team_v1alpha1_team_proto_depIdxs,
		MessageInfos:      file_team_v1alpha1_team_proto_msgTypes,
	}.Build()
	File_team_v1alpha1_team_proto = out.File
	file_team_v1alpha1_team_proto_rawDesc = nil
	file_team_v1alpha1_team_proto_goTypes = nil
	file_team_v1alpha1_team_proto_depIdxs = nil
}
