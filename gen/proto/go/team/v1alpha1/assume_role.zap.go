// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: team/v1alpha1/assume_role.proto

package teamv1alpha1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	go_uber_org_zap_zapcore "go.uber.org/zap/zapcore"
	github_com_golang_protobuf_ptypes "github.com/golang/protobuf/ptypes"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *AssumeRole) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "role" // field role = 1
	enc.AddString(keyName, m.Role)

	keyName = "account" // field account = 2
	enc.AddString(keyName, m.Account)

	keyName = "certificate_fingerprint" // field certificate_fingerprint = 3
	enc.AddByteString(keyName, m.CertificateFingerprint)

	keyName = "timestamp" // field timestamp = 4
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.Timestamp); err == nil {
		enc.AddTime(keyName, t)
	}

	return nil
}
