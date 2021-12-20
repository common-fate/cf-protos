// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provider/v1alpha1/public.proto

package providerv1alpha1

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

func (m *EnrolRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "enrol_token" // field enrol_token = 1
	enc.AddString(keyName, m.EnrolToken)

	keyName = "csr" // field csr = 2
	enc.AddByteString(keyName, m.Csr)

	keyName = "access_handler_url" // field access_handler_url = 3
	enc.AddString(keyName, m.AccessHandlerUrl)

	keyName = "aws" // field aws = 4
	if ov, ok := m.GetProof().(*EnrolRequest_Aws); ok {
		_ = ov
		if ov.Aws != nil {
			var vv interface{} = ov.Aws
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	return nil
}

func (m *EnrolResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *AWSProof) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "identity" // field identity = 1
	if m.Identity != nil {
		var vv interface{} = m.Identity
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	keyName = "organization" // field organization = 2
	if m.Organization != nil {
		var vv interface{} = m.Organization
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *AWSSignature) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "signature" // field signature = 1
	enc.AddString(keyName, m.Signature)

	keyName = "timestamp" // field timestamp = 2
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.Timestamp); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "security_token" // field security_token = 3
	enc.AddString(keyName, m.SecurityToken)

	return nil
}

func (m *GetCertificateRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "csr" // field csr = 1
	enc.AddByteString(keyName, m.Csr)

	keyName = "tenant" // field tenant = 2
	enc.AddString(keyName, m.Tenant)

	keyName = "aws" // field aws = 3
	if ov, ok := m.GetProof().(*GetCertificateRequest_Aws); ok {
		_ = ov
		if ov.Aws != nil {
			var vv interface{} = ov.Aws
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	return nil
}

func (m *GetCertificateResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "certificate" // field certificate = 1
	enc.AddByteString(keyName, m.Certificate)

	return nil
}
