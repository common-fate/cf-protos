// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: provider/v1alpha1/provider.proto

package providerv1alpha1

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

func (m *StoreAssumeRoleReasonRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "requestor_mail" // field requestor_mail = 1
	enc.AddString(keyName, m.RequestorMail)

	keyName = "customer_role_id" // field customer_role_id = 2
	enc.AddString(keyName, m.CustomerRoleId)

	keyName = "customer_account_id" // field customer_account_id = 3
	enc.AddString(keyName, m.CustomerAccountId)

	keyName = "role_arn" // field role_arn = 4
	enc.AddString(keyName, m.RoleArn)

	keyName = "reason" // field reason = 5
	enc.AddString(keyName, m.Reason)

	return nil
}

func (m *StoreAssumeRoleReasonResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	return nil
}

func (m *GetProviderConfigByDigestRequest) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
	var keyName string
	_ = keyName

	if m == nil {
		return nil
	}

	keyName = "sha256_digest" // field sha256_digest = 1
	enc.AddByteString(keyName, m.Sha256Digest)

	return nil
}

func (m *GetProviderConfigByDigestResponse) MarshalLogObject(enc go_uber_org_zap_zapcore.ObjectEncoder) error {
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
