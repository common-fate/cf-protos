// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: accesshandler/v1alpha1/accesshandler.proto

package accesshandlerv1alpha1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	go_uber_org_zap_zapcore "go.uber.org/zap/zapcore"
	github_com_golang_protobuf_ptypes "github.com/golang/protobuf/ptypes"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *Grant) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	keyName = "provider_id" // field provider_id = 2
	enc.AddString(keyName, m.ProviderId)

	keyName = "provider_type" // field provider_type = 3
	enc.AddString(keyName, m.ProviderType)

	keyName = "principal" // field principal = 4
	enc.AddString(keyName, m.Principal)

	keyName = "start" // field start = 5
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.Start); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "end" // field end = 6
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.End); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "user" // field user = 7
	enc.AddString(keyName, m.User)

	keyName = "claims" // field claims = 8
	if m.Claims != nil {
		var vv interface{} = m.Claims
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *NotifyGrantEventRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "grant" // field grant = 1
	if m.Grant != nil {
		var vv interface{} = m.Grant
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	keyName = "event" // field event = 2
	enc.AddString(keyName, m.Event.String())

	keyName = "data" // field data = 3
	if m.Data != nil {
		var vv interface{} = m.Data
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *NotifyGrantEventResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "event_id" // field event_id = 1
	enc.AddString(keyName, m.EventId)

	return nil
}
