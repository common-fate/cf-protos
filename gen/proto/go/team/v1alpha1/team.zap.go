// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: team/v1alpha1/team.proto

package teamv1alpha1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "google.golang.org/protobuf/types/known/structpb"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/common-fate/gconfig/gen/gconfig/v1alpha1"
	_ "github.com/common-fate/cf-protos/gen/proto/go/certrail/v1alpha1"
	go_uber_org_zap_zapcore "go.uber.org/zap/zapcore"
	github_com_golang_protobuf_ptypes "github.com/golang/protobuf/ptypes"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *RequestAccessRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "provider_id" // field provider_id = 1
	enc.AddString(keyName, m.ProviderId)

	keyName = "provider_type" // field provider_type = 2
	enc.AddString(keyName, m.ProviderType)

	keyName = "principal" // field principal = 3
	enc.AddString(keyName, m.Principal)

	keyName = "start" // field start = 4
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.Start); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "end" // field end = 5
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.End); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "claims" // field claims = 6
	if m.Claims != nil {
		var vv interface{} = m.Claims
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *RequestAccessResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	keyName = "status" // field status = 2
	enc.AddString(keyName, m.Status.String())

	return nil
}

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

	keyName = "provision_strategy" // field provision_strategy = 6
	enc.AddString(keyName, m.ProvisionStrategy)

	keyName = "expiry_duration" // field expiry_duration = 7
	enc.AddString(keyName, m.ExpiryDuration)

	keyName = "provider_id" // field provider_id = 8
	enc.AddString(keyName, m.ProviderId)

	keyName = "provider_type" // field provider_type = 9
	enc.AddString(keyName, m.ProviderType)

	keyName = "principal" // field principal = 10
	enc.AddString(keyName, m.Principal)

	keyName = "start" // field start = 11
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.Start); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "end" // field end = 12
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.End); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "requested_by" // field requested_by = 13
	enc.AddString(keyName, m.RequestedBy)

	return nil
}

func (m *ReviewRequestRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "request_id" // field request_id = 1
	enc.AddString(keyName, m.RequestId)

	keyName = "decision" // field decision = 2
	enc.AddString(keyName, m.Decision.String())

	return nil
}

func (m *ReviewRequestResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "grant_id" // field grant_id = 1
	enc.AddString(keyName, m.GrantId)

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

func (m *EnrollProviderRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
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
	if ov, ok := m.GetProvider().(*EnrollProviderRequest_Aws); ok {
		_ = ov
		if ov.Aws != nil {
			var vv interface{} = ov.Aws
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	keyName = "aws_sso" // field aws_sso = 4
	if ov, ok := m.GetProvider().(*EnrollProviderRequest_AwsSso); ok {
		_ = ov
		if ov.AwsSso != nil {
			var vv interface{} = ov.AwsSso
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	keyName = "okta" // field okta = 5
	if ov, ok := m.GetProvider().(*EnrollProviderRequest_Okta); ok {
		_ = ov
		if ov.Okta != nil {
			var vv interface{} = ov.Okta
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	return nil
}

func (m *EnrollAWSProvider) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "account_id" // field account_id = 1
	enc.AddString(keyName, m.AccountId)

	return nil
}

func (m *EnrollAWSSSOProvider) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "account_id" // field account_id = 1
	enc.AddString(keyName, m.AccountId)

	return nil
}

func (m *EnrollOktaProvider) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "org_url" // field org_url = 1
	enc.AddString(keyName, m.OrgUrl)

	return nil
}

func (m *EnrollProviderResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
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

	keyName = "awssso" // field awssso = 7
	if ov, ok := m.GetDetails().(*ProviderEnrollment_Awssso); ok {
		_ = ov
		if ov.Awssso != nil {
			var vv interface{} = ov.Awssso
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	keyName = "okta" // field okta = 8
	if ov, ok := m.GetDetails().(*ProviderEnrollment_Okta); ok {
		_ = ov
		if ov.Okta != nil {
			var vv interface{} = ov.Okta
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

func (m *AWSSSOProviderEnrollment) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "account_id" // field account_id = 1
	enc.AddString(keyName, m.AccountId)

	return nil
}

func (m *OktaProviderEnrollment) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
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

	keyName = "awssso" // field awssso = 5
	if ov, ok := m.GetDetails().(*Provider_Awssso); ok {
		_ = ov
		if ov.Awssso != nil {
			var vv interface{} = ov.Awssso
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	keyName = "okta" // field okta = 6
	if ov, ok := m.GetDetails().(*Provider_Okta); ok {
		_ = ov
		if ov.Okta != nil {
			var vv interface{} = ov.Okta
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

func (m *AWSSSOProviderDetails) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "org_management_account_id" // field org_management_account_id = 1
	enc.AddString(keyName, m.OrgManagementAccountId)

	return nil
}

func (m *OktaProviderDetails) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "org_url" // field org_url = 1
	enc.AddString(keyName, m.OrgUrl)

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

func (m *ListAccessHandlersRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "provider_id" // field provider_id = 1
	enc.AddString(keyName, m.ProviderId)

	return nil
}

func (m *AccessHandler) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	keyName = "url" // field url = 2
	enc.AddString(keyName, m.Url)

	keyName = "provider_ids" // field provider_ids = 3
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.ProviderIds {
			_ = rv
			aenc.AppendString(rv)
		}
		return nil
	}))

	return nil
}

func (m *ListAccessHandlersResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
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

	keyName = "access_handler_id" // field access_handler_id = 1
	enc.AddString(keyName, m.AccessHandlerId)

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

func (m *SlackConnection) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "bot_access_token" // field bot_access_token = 1
	enc.AddString(keyName, m.BotAccessToken)

	keyName = "channel_id" // field channel_id = 2
	enc.AddString(keyName, m.ChannelId)

	keyName = "channel_name" // field channel_name = 3
	enc.AddString(keyName, m.ChannelName)

	keyName = "invited" // field invited = 4
	enc.AddBool(keyName, m.Invited)

	return nil
}

func (m *ListSlackConnectionsRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *ListSlackConnectionsResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "slack_connections" // field slack_connections = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.SlackConnections {
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

func (m *HasSlackConnectionRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *HasSlackConnectionResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "has_slack_connection" // field has_slack_connection = 1
	enc.AddBool(keyName, m.HasSlackConnection)

	return nil
}

func (m *GetSlackInstallURLRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *GetSlackInstallURLResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "install_url" // field install_url = 1
	enc.AddString(keyName, m.InstallUrl)

	return nil
}

func (m *UninstallSlackRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *UninstallSlackResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "success" // field success = 1
	enc.AddBool(keyName, m.Success)

	return nil
}

func (m *SlackChannelInviteTestRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "channel_id" // field channel_id = 2
	enc.AddString(keyName, m.ChannelId)

	return nil
}

func (m *SlackChannelInviteTestResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "success" // field success = 1
	enc.AddBool(keyName, m.Success)

	return nil
}

func (m *GetAccessHandlerDeploymentGuideRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "provider_id" // field provider_id = 1
	enc.AddString(keyName, m.ProviderId)

	keyName = "runtime" // field runtime = 2
	enc.AddString(keyName, m.Runtime.String())

	return nil
}

func (m *GetAccessHandlerDeploymentGuideResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "deployment_url" // field deployment_url = 1
	enc.AddString(keyName, m.DeploymentUrl)

	return nil
}

func (m *AWSRoleEnrollment) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "aws_role" // field aws_role = 1
	enc.AddString(keyName, m.AwsRole)

	return nil
}

func (m *SelfHostedEnrollment) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "certificate" // field certificate = 2
	enc.AddString(keyName, m.Certificate)

	return nil
}

func (m *EnrollAccessHandlerRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "access_handler_url" // field access_handler_url = 1
	enc.AddString(keyName, m.AccessHandlerUrl)

	keyName = "provider_ids" // field provider_ids = 2
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.ProviderIds {
			_ = rv
			aenc.AppendString(rv)
		}
		return nil
	}))

	keyName = "runtime" // field runtime = 3
	enc.AddString(keyName, m.Runtime.String())

	keyName = "aws_role" // field aws_role = 4
	if ov, ok := m.GetEnrollment().(*EnrollAccessHandlerRequest_AwsRole); ok {
		_ = ov
		if ov.AwsRole != nil {
			var vv interface{} = ov.AwsRole
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	keyName = "self_hosted" // field self_hosted = 5
	if ov, ok := m.GetEnrollment().(*EnrollAccessHandlerRequest_SelfHosted); ok {
		_ = ov
		if ov.SelfHosted != nil {
			var vv interface{} = ov.SelfHosted
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	return nil
}

func (m *EnrollAccessHandlerResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "access_handler_id" // field access_handler_id = 1
	enc.AddString(keyName, m.AccessHandlerId)

	return nil
}

func (m *QueryProviderInfo) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "provider_id" // field provider_id = 1
	enc.AddString(keyName, m.ProviderId)

	keyName = "provider_type" // field provider_type = 2
	enc.AddString(keyName, m.ProviderType)

	return nil
}

func (m *QueryAccessHandlersRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
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

func (m *QueryAccessHandlersResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "access_handler_urls" // field access_handler_urls = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.AccessHandlerUrls {
			_ = rv
			aenc.AppendString(rv)
		}
		return nil
	}))

	return nil
}
