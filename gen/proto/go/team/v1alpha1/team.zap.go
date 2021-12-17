// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: team/v1alpha1/team.proto

package teamv1alpha1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	go_uber_org_zap_zapcore "go.uber.org/zap/zapcore"
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

func (m *EnrolAWSProviderIntervention) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "provider_id" // field provider_id = 1
	enc.AddString(keyName, m.ProviderId)

	keyName = "aws_account_id" // field aws_account_id = 2
	enc.AddString(keyName, m.AwsAccountId)

	keyName = "cloudformation_url" // field cloudformation_url = 3
	enc.AddString(keyName, m.CloudformationUrl)

	return nil
}

func (m *GetInterventionsRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *Intervention) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "enrol_aws_provider" // field enrol_aws_provider = 1
	if ov, ok := m.GetData().(*Intervention_EnrolAwsProvider); ok {
		_ = ov
		if ov.EnrolAwsProvider != nil {
			var vv interface{} = ov.EnrolAwsProvider
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	return nil
}

func (m *GetInterventionsResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "interventions" // field interventions = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Interventions {
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

func (m *EnrolProviderRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "aws" // field aws = 1
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

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	keyName = "account_id" // field account_id = 2
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

	return nil
}
