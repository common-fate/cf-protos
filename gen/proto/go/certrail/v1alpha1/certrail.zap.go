// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: certrail/v1alpha1/certrail.proto

package certrailv1alpha1

import (
	fmt "fmt"
	math "math"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "google.golang.org/protobuf/types/known/durationpb"
	go_uber_org_zap_zapcore "go.uber.org/zap/zapcore"
	github_com_golang_protobuf_ptypes "github.com/golang/protobuf/ptypes"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (m *ApproveConfigPayload) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "config_sha256" // field config_sha256 = 1
	enc.AddByteString(keyName, m.ConfigSha256)

	keyName = "approved_by" // field approved_by = 2
	enc.AddByteString(keyName, m.ApprovedBy)

	return nil
}

func (m *RoleAccessRequestPayload) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "role" // field role = 1
	enc.AddString(keyName, m.Role)

	keyName = "provider" // field provider = 2
	enc.AddString(keyName, m.Provider)

	keyName = "account" // field account = 3
	enc.AddString(keyName, m.Account)

	keyName = "session_duration" // field session_duration = 4
	if d, err := github_com_golang_protobuf_ptypes.Duration(m.SessionDuration); err == nil {
		enc.AddDuration(keyName, d)
	}

	keyName = "reason" // field reason = 5
	enc.AddString(keyName, m.Reason)

	keyName = "requested_by" // field requested_by = 6
	enc.AddByteString(keyName, m.RequestedBy)

	return nil
}

func (m *ApprovedRoleAccessRequestPayload) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "role_access_request_certrail_index" // field role_access_request_certrail_index = 1
	enc.AddInt64(keyName, m.RoleAccessRequestCertrailIndex)

	keyName = "approved_by" // field approved_by = 2
	enc.AddByteString(keyName, m.ApprovedBy)

	return nil
}

func (m *DeclinedRoleAccessRequestPayload) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "role_access_request_certrail_index" // field role_access_request_certrail_index = 1
	enc.AddInt64(keyName, m.RoleAccessRequestCertrailIndex)

	keyName = "declined_by" // field declined_by = 2
	enc.AddByteString(keyName, m.DeclinedBy)

	return nil
}

func (m *CancelledRoleAccessRequestPayload) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "role_access_request_certrail_index" // field role_access_request_certrail_index = 1
	enc.AddInt64(keyName, m.RoleAccessRequestCertrailIndex)

	keyName = "cancelled_by" // field cancelled_by = 2
	enc.AddByteString(keyName, m.CancelledBy)

	return nil
}

func (m *IssueCertificatePayload) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "certificate" // field certificate = 1
	enc.AddByteString(keyName, m.Certificate)

	return nil
}

func (m *RevokeCertificatePayload) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "certificate" // field certificate = 1
	enc.AddByteString(keyName, m.Certificate)

	keyName = "revoked_by" // field revoked_by = 2
	enc.AddByteString(keyName, m.RevokedBy)

	return nil
}

func (m *IssueSessionCredentialsPayload) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "id" // field id = 1
	enc.AddString(keyName, m.Id)

	keyName = "user_certificate" // field user_certificate = 2
	enc.AddByteString(keyName, m.UserCertificate)

	keyName = "user_supplied_reason" // field user_supplied_reason = 3
	enc.AddString(keyName, m.UserSuppliedReason)

	return nil
}

func (m *Envelope) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "payload" // field payload = 1
	if m.Payload != nil {
		var vv interface{} = m.Payload
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	keyName = "signature" // field signature = 2
	enc.AddByteString(keyName, m.Signature)

	return nil
}

func (m *Payload) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "timestamp" // field timestamp = 1
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.Timestamp); err == nil {
		enc.AddTime(keyName, t)
	}

	keyName = "approve_config" // field approve_config = 2
	if ov, ok := m.GetContents().(*Payload_ApproveConfig); ok {
		_ = ov
		if ov.ApproveConfig != nil {
			var vv interface{} = ov.ApproveConfig
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	keyName = "issue_certificate" // field issue_certificate = 3
	if ov, ok := m.GetContents().(*Payload_IssueCertificate); ok {
		_ = ov
		if ov.IssueCertificate != nil {
			var vv interface{} = ov.IssueCertificate
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	keyName = "revoke_certificate" // field revoke_certificate = 4
	if ov, ok := m.GetContents().(*Payload_RevokeCertificate); ok {
		_ = ov
		if ov.RevokeCertificate != nil {
			var vv interface{} = ov.RevokeCertificate
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	keyName = "issue_session_credentials" // field issue_session_credentials = 5
	if ov, ok := m.GetContents().(*Payload_IssueSessionCredentials); ok {
		_ = ov
		if ov.IssueSessionCredentials != nil {
			var vv interface{} = ov.IssueSessionCredentials
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	keyName = "role_access_request" // field role_access_request = 6
	if ov, ok := m.GetContents().(*Payload_RoleAccessRequest); ok {
		_ = ov
		if ov.RoleAccessRequest != nil {
			var vv interface{} = ov.RoleAccessRequest
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	keyName = "approved_role_access_request" // field approved_role_access_request = 7
	if ov, ok := m.GetContents().(*Payload_ApprovedRoleAccessRequest); ok {
		_ = ov
		if ov.ApprovedRoleAccessRequest != nil {
			var vv interface{} = ov.ApprovedRoleAccessRequest
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	keyName = "declined_role_access_request" // field declined_role_access_request = 8
	if ov, ok := m.GetContents().(*Payload_DeclinedRoleAccessRequest); ok {
		_ = ov
		if ov.DeclinedRoleAccessRequest != nil {
			var vv interface{} = ov.DeclinedRoleAccessRequest
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	keyName = "cancelled_role_access_request" // field cancelled_role_access_request = 9
	if ov, ok := m.GetContents().(*Payload_CancelledRoleAccessRequest); ok {
		_ = ov
		if ov.CancelledRoleAccessRequest != nil {
			var vv interface{} = ov.CancelledRoleAccessRequest
			if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
				enc.AddObject(keyName, marshaler)
			}
		}
	}

	return nil
}

func (m *StoreRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "envelope" // field envelope = 1
	if m.Envelope != nil {
		var vv interface{} = m.Envelope
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *StoreResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "role_access_request_id" // field role_access_request_id = 1
	if ov, ok := m.GetContents().(*StoreResponse_RoleAccessRequestId); ok {
		_ = ov
		enc.AddString(keyName, ov.RoleAccessRequestId)
	}

	return nil
}

func (m *GetEntriesRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "start_index" // field start_index = 1
	enc.AddInt64(keyName, m.StartIndex)

	keyName = "count" // field count = 2
	enc.AddInt64(keyName, m.Count)

	return nil
}

func (m *GetEntriesResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "envelopes" // field envelopes = 1
	enc.AddArray(keyName, go_uber_org_zap_zapcore.ArrayMarshalerFunc(func(aenc go_uber_org_zap_zapcore.ArrayEncoder) error {
		for _, rv := range m.Envelopes {
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

func (m *IncludedEnvelope) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "envelope" // field envelope = 1
	if m.Envelope != nil {
		var vv interface{} = m.Envelope
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *GetLatestSignedLogRootRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *GetLatestSignedLogRootResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "log_root" // field log_root = 1
	if m.LogRoot != nil {
		var vv interface{} = m.LogRoot
		if marshaler, ok := vv.(go_uber_org_zap_zapcore.ObjectMarshaler); ok {
			enc.AddObject(keyName, marshaler)
		}
	}

	return nil
}

func (m *LogRoot) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "tree_size" // field tree_size = 1
	enc.AddUint64(keyName, m.TreeSize)

	keyName = "root_hash" // field root_hash = 2
	enc.AddByteString(keyName, m.RootHash)

	keyName = "timestamp" // field timestamp = 3
	if t, err := github_com_golang_protobuf_ptypes.Timestamp(m.Timestamp); err == nil {
		enc.AddTime(keyName, t)
	}

	return nil
}
