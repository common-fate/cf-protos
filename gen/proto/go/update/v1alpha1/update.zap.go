// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: update/v1alpha1/update.proto

package updatev1alpha1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	go_uber_org_zap_zapcore "go.uber.org/zap/zapcore"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *CheckForUpdatesRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "version" // field version = 1
	enc.AddString(keyName, m.Version)

	return nil
}

func (m *CheckForUpdatesResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "update_required" // field update_required = 1
	enc.AddBool(keyName, m.UpdateRequired)

	keyName = "message" // field message = 2
	enc.AddString(keyName, m.Message)

	return nil
}