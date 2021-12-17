// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: team/v1alpha1/team.proto

package teamv1alpha1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	go_uber_org_zap_zapcore "go.uber.org/zap/zapcore"
	github_com_golang_protobuf_ptypes "github.com/golang/protobuf/ptypes"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *ListMembersRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *ListMembersResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "members" // field members = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Members {
			_ = rv
			if rv != nil {
				var vv interface{} = rv
				if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
					aenc.AppendObject(marshaler)
				}
			}
		}
		return nil
	}))

	return nil
}

func (m *Member) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "email_address" // field email_address = 1
	enc.AddString(keyName, m.EmailAddress)

	keyName = "is_admin" // field is_admin = 2
	enc.AddBool(keyName, m.IsAdmin)

	return nil
}

func (m *UpdateConfigRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "config" // field config = 1
	if m.Config != nil {
		var vv interface{} = m.Config
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *UpdateConfigResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "sha256_hash" // field sha256_hash = 1
	enc.AddByteString(keyName, m.Sha256Hash)

	return nil
}

func (m *GetConfigRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *GetConfigResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "config" // field config = 1
	if m.Config != nil {
		var vv interface{} = m.Config
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *GetConfigByHashRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "sha256_hash" // field sha256_hash = 1
	enc.AddByteString(keyName, m.Sha256Hash)

	return nil
}

func (m *GetConfigByHashResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "config" // field config = 1
	if m.Config != nil {
		var vv interface{} = m.Config
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *EnrolProviderRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	keyName = "name" // field name = 2
	enc.AddString(keyName, m.Name)

	keyName = "aws" // field aws = 3
	if ov, ok := m.GetProvider().(*EnrolProviderRequest_Aws); ok {
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

func (m *EnrolAWSProvider) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "account_id" // field account_id = 1
	enc.AddString(keyName, m.AccountId)

	return nil
}

func (m *EnrolProviderResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "status" // field status = 1
	enc.AddString(keyName, m.Status.String())

	keyName = "deployment_url" // field deployment_url = 2
	enc.AddString(keyName, m.DeploymentUrl)

	keyName = "enrollment_token" // field enrollment_token = 3
	enc.AddString(keyName, m.EnrollmentToken)

	keyName = "expires_at" // field expires_at = 4
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.ExpiresAt); err == nil {
		enc.AddTime(keyName, t)
	}

	return nil
}

func (m *ListProvidersRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *ListProvidersResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "providers" // field providers = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Providers {
			_ = rv
			if rv != nil {
				var vv interface{} = rv
				if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
					aenc.AppendObject(marshaler)
				}
			}
		}
		return nil
	}))

	return nil
}

func (m *Provider) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	keyName = "name" // field name = 2
	enc.AddString(keyName, m.Name)

	keyName = "status" // field status = 3
	enc.AddString(keyName, m.Status.String())

	keyName = "aws" // field aws = 4
	if ov, ok := m.GetDetails().(*Provider_Aws); ok {
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

func (m *AWSProviderDetails) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "account_id" // field account_id = 1
	enc.AddString(keyName, m.AccountId)

	return nil
}

func (m *GetStatusRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *GetStatusResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "enrollments" // field enrollments = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Enrollments {
			_ = rv
			if rv != nil {
				var vv interface{} = rv
				if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
					aenc.AppendObject(marshaler)
				}
			}
		}
		return nil
	}))

	keyName = "expired_enrollments" // field expired_enrollments = 2
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.ExpiredEnrollments {
			_ = rv
			if rv != nil {
				var vv interface{} = rv
				if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
					aenc.AppendObject(marshaler)
				}
			}
		}
		return nil
	}))

	return nil
}

func (m *ProviderEnrollment) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "provider_id" // field provider_id = 1
	enc.AddString(keyName, m.ProviderId)

	keyName = "deployment_url" // field deployment_url = 2
	enc.AddString(keyName, m.DeploymentUrl)

	keyName = "enrollment_token" // field enrollment_token = 3
	enc.AddString(keyName, m.EnrollmentToken)

	keyName = "expires_at" // field expires_at = 4
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.ExpiresAt); err == nil {
		enc.AddTime(keyName, t)
	}

	return nil
}

func (m *DeleteProviderRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	return nil
}

func (m *DeleteProviderResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}
