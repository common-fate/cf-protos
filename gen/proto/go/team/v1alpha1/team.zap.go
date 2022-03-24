// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: team/v1alpha1/team.proto

package teamv1alpha1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	_ "github.com/common-fate/cf-protos/gen/proto/go/certrail/v1alpha1"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	go_uber_org_zap_zapcore "go.uber.org/zap/zapcore"
	github_com_golang_protobuf_ptypes "github.com/golang/protobuf/ptypes"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *ListRoleAccessRequestsRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "include_expired" // field include_expired = 1
	enc.AddBool(keyName, m.IncludeExpired)

	return nil
}

func (m *ListRoleAccessRequestsResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "requests" // field requests = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Requests {
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

func (m *RoleAccessRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "request" // field request = 1
	if m.Request != nil {
		var vv interface{} = m.Request
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	keyName = "status" // field status = 2
	enc.AddString(keyName, m.Status.String())

	keyName = "timestamp" // field timestamp = 3
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.Timestamp); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "id" // field id = 4
	enc.AddString(keyName, m.Id)

	keyName = "trillian_merkle_hash" // field trillian_merkle_hash = 5
	enc.AddByteString(keyName, m.TrillianMerkleHash)

	return nil
}

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

func (m *Account) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "provider_id" // field provider_id = 1
	enc.AddString(keyName, m.ProviderId)

	keyName = "account_id" // field account_id = 2
	enc.AddString(keyName, m.AccountId)

	keyName = "alias" // field alias = 3
	enc.AddString(keyName, m.Alias)

	return nil
}

func (m *UpdateAccountsRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "accounts" // field accounts = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Accounts {
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

func (m *UpdateAccountsResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
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

func (m *ListRolesRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *ListRolesResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "roles" // field roles = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Roles {
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

func (m *IsAdminUserRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *IsAdminUserResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "is_admin" // field is_admin = 1
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

	keyName = "aws_sso" // field aws_sso = 4
	if ov, ok := m.GetProvider().(*EnrolProviderRequest_AwsSso); ok {
		_ = ov
		if ov.AwsSso != nil {
			var vv interface{} = ov.AwsSso
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

func (m *EnrolAWSSSOProvider) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
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

	keyName = "deployment_url" // field deployment_url = 2
	enc.AddString(keyName, m.DeploymentUrl)

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

	keyName = "pending_enrollments_count" // field pending_enrollments_count = 2
	enc.AddInt32(keyName, m.PendingEnrollmentsCount)

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

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	keyName = "name" // field name = 2
	enc.AddString(keyName, m.Name)

	keyName = "deployment_url" // field deployment_url = 3
	enc.AddString(keyName, m.DeploymentUrl)

	keyName = "enrollment_token" // field enrollment_token = 4
	enc.AddString(keyName, m.EnrollmentToken)

	keyName = "expires_at" // field expires_at = 5
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.ExpiresAt); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "aws" // field aws = 6
	if ov, ok := m.GetDetails().(*ProviderEnrollment_Aws); ok {
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

func (m *AWSProviderEnrollment) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "account_id" // field account_id = 1
	enc.AddString(keyName, m.AccountId)

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

func (m *GetProviderRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	return nil
}

func (m *GetProviderResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "provider" // field provider = 1
	if m.Provider != nil {
		var vv interface{} = m.Provider
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

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

	keyName = "org_management_account_id" // field org_management_account_id = 1
	enc.AddString(keyName, m.OrgManagementAccountId)

	return nil
}

func (m *GetAllProviderDetailsRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *GetAllProviderDetailsResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "providers" // field providers = 1
	if m.Providers != nil {
		var vv interface{} = m.Providers
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	keyName = "sha256_checksum" // field sha256_checksum = 2
	enc.AddByteString(keyName, m.Sha256Checksum)

	return nil
}

func (m *GetAllProviderChecksumRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *GetAllProviderChecksumResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "sha256_checksum" // field sha256_checksum = 1
	enc.AddByteString(keyName, m.Sha256Checksum)

	return nil
}

func (m *GetAccessHandlersForProviderRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "provider_id" // field provider_id = 1
	enc.AddString(keyName, m.ProviderId)

	return nil
}

func (m *GetAccessHandlersForProviderResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "access_handlers" // field access_handlers = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.AccessHandlers {
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

func (m *AddAccessHandlerRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "provider_id" // field provider_id = 1
	enc.AddString(keyName, m.ProviderId)

	keyName = "access_handler_url" // field access_handler_url = 2
	enc.AddString(keyName, m.AccessHandlerUrl)

	return nil
}

func (m *AddAccessHandlerResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *DeleteAccessHandlerRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "provider_id" // field provider_id = 1
	enc.AddString(keyName, m.ProviderId)

	keyName = "access_handler_url" // field access_handler_url = 2
	enc.AddString(keyName, m.AccessHandlerUrl)

	return nil
}

func (m *DeleteAccessHandlerResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *UpdateCISettingsRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "enabled" // field enabled = 1
	enc.AddBool(keyName, m.Enabled)

	keyName = "repository_url" // field repository_url = 2
	enc.AddString(keyName, m.RepositoryUrl)

	return nil
}

func (m *UpdateCISettingsResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *ConnectSlackRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "bot_access_token" // field bot_access_token = 1
	enc.AddString(keyName, m.BotAccessToken)

	keyName = "channel_id" // field channel_id = 2
	enc.AddString(keyName, m.ChannelId)

	return nil
}

func (m *ConnectSlackResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}
