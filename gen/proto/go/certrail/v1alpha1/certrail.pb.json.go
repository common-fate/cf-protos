// Code generated by protoc-gen-go-json. DO NOT EDIT.
// source: certrail/v1alpha1/certrail.proto

package certrailv1alpha1

import (
	"google.golang.org/protobuf/encoding/protojson"
)

// MarshalJSON implements json.Marshaler
func (msg *ApproveConfigPayload) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ApproveConfigPayload) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetEntryIndexByMerkleHashRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetEntryIndexByMerkleHashRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetEntryIndexByMerkleHashResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetEntryIndexByMerkleHashResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *RoleAccessRequestPayload) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *RoleAccessRequestPayload) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *ApprovedRoleAccessRequestPayload) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *ApprovedRoleAccessRequestPayload) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *Window) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *Window) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *Expires) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *Expires) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *DeclinedRoleAccessRequestPayload) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *DeclinedRoleAccessRequestPayload) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *CancelledRoleAccessRequestPayload) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *CancelledRoleAccessRequestPayload) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *IssueCertificatePayload) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *IssueCertificatePayload) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *RevokeCertificatePayload) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *RevokeCertificatePayload) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *IssueSessionCredentialsPayload) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *IssueSessionCredentialsPayload) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *OktaAccessWorkflowInitiatedPayload) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *OktaAccessWorkflowInitiatedPayload) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *Envelope) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *Envelope) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *Payload) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *Payload) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *StoreRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *StoreRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *StoreResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *StoreResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetEntriesRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetEntriesRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetEntriesResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetEntriesResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *IncludedEnvelope) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *IncludedEnvelope) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetLatestSignedLogRootRequest) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetLatestSignedLogRootRequest) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *GetLatestSignedLogRootResponse) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *GetLatestSignedLogRootResponse) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}

// MarshalJSON implements json.Marshaler
func (msg *LogRoot) MarshalJSON() ([]byte, error) {
	return protojson.MarshalOptions{
		UseEnumNumbers:  false,
		EmitUnpopulated: false,
		UseProtoNames:   false,
	}.Marshal(msg)
}

// UnmarshalJSON implements json.Unmarshaler
func (msg *LogRoot) UnmarshalJSON(b []byte) error {
	return protojson.UnmarshalOptions{
		DiscardUnknown: false,
	}.Unmarshal(b, msg)
}
