// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: certrail/v1alpha1/glog.proto

package certrailv1alpha1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on UpdateConfigPayload with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *UpdateConfigPayload) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on UpdateConfigPayload with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// UpdateConfigPayloadMultiError, or nil if none found.
func (m *UpdateConfigPayload) ValidateAll() error {
	return m.validate(true)
}

func (m *UpdateConfigPayload) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(m.GetConfigSha256()) != 32 {
		err := UpdateConfigPayloadValidationError{
			field:  "ConfigSha256",
			reason: "value length must be 32 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return UpdateConfigPayloadMultiError(errors)
	}
	return nil
}

// UpdateConfigPayloadMultiError is an error wrapping multiple validation
// errors returned by UpdateConfigPayload.ValidateAll() if the designated
// constraints aren't met.
type UpdateConfigPayloadMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m UpdateConfigPayloadMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m UpdateConfigPayloadMultiError) AllErrors() []error { return m }

// UpdateConfigPayloadValidationError is the validation error returned by
// UpdateConfigPayload.Validate if the designated constraints aren't met.
type UpdateConfigPayloadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateConfigPayloadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateConfigPayloadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateConfigPayloadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateConfigPayloadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateConfigPayloadValidationError) ErrorName() string {
	return "UpdateConfigPayloadValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateConfigPayloadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateConfigPayload.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateConfigPayloadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateConfigPayloadValidationError{}

// Validate checks the field values on IssueCertificatePayload with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *IssueCertificatePayload) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IssueCertificatePayload with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IssueCertificatePayloadMultiError, or nil if none found.
func (m *IssueCertificatePayload) ValidateAll() error {
	return m.validate(true)
}

func (m *IssueCertificatePayload) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Certificate

	if len(errors) > 0 {
		return IssueCertificatePayloadMultiError(errors)
	}
	return nil
}

// IssueCertificatePayloadMultiError is an error wrapping multiple validation
// errors returned by IssueCertificatePayload.ValidateAll() if the designated
// constraints aren't met.
type IssueCertificatePayloadMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IssueCertificatePayloadMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IssueCertificatePayloadMultiError) AllErrors() []error { return m }

// IssueCertificatePayloadValidationError is the validation error returned by
// IssueCertificatePayload.Validate if the designated constraints aren't met.
type IssueCertificatePayloadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IssueCertificatePayloadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IssueCertificatePayloadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IssueCertificatePayloadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IssueCertificatePayloadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IssueCertificatePayloadValidationError) ErrorName() string {
	return "IssueCertificatePayloadValidationError"
}

// Error satisfies the builtin error interface
func (e IssueCertificatePayloadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIssueCertificatePayload.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IssueCertificatePayloadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IssueCertificatePayloadValidationError{}

// Validate checks the field values on RevokeCertificatePayload with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *RevokeCertificatePayload) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RevokeCertificatePayload with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RevokeCertificatePayloadMultiError, or nil if none found.
func (m *RevokeCertificatePayload) ValidateAll() error {
	return m.validate(true)
}

func (m *RevokeCertificatePayload) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Certificate

	// no validation rules for RevokedBy

	if len(errors) > 0 {
		return RevokeCertificatePayloadMultiError(errors)
	}
	return nil
}

// RevokeCertificatePayloadMultiError is an error wrapping multiple validation
// errors returned by RevokeCertificatePayload.ValidateAll() if the designated
// constraints aren't met.
type RevokeCertificatePayloadMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RevokeCertificatePayloadMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RevokeCertificatePayloadMultiError) AllErrors() []error { return m }

// RevokeCertificatePayloadValidationError is the validation error returned by
// RevokeCertificatePayload.Validate if the designated constraints aren't met.
type RevokeCertificatePayloadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RevokeCertificatePayloadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RevokeCertificatePayloadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RevokeCertificatePayloadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RevokeCertificatePayloadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RevokeCertificatePayloadValidationError) ErrorName() string {
	return "RevokeCertificatePayloadValidationError"
}

// Error satisfies the builtin error interface
func (e RevokeCertificatePayloadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRevokeCertificatePayload.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RevokeCertificatePayloadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RevokeCertificatePayloadValidationError{}

// Validate checks the field values on Envelope with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Envelope) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Envelope with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EnvelopeMultiError, or nil
// if none found.
func (m *Envelope) ValidateAll() error {
	return m.validate(true)
}

func (m *Envelope) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetPayload()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, EnvelopeValidationError{
					field:  "Payload",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, EnvelopeValidationError{
					field:  "Payload",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetPayload()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return EnvelopeValidationError{
				field:  "Payload",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(m.GetSignature()) > 256 {
		err := EnvelopeValidationError{
			field:  "Signature",
			reason: "value length must be at most 256 bytes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return EnvelopeMultiError(errors)
	}
	return nil
}

// EnvelopeMultiError is an error wrapping multiple validation errors returned
// by Envelope.ValidateAll() if the designated constraints aren't met.
type EnvelopeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnvelopeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnvelopeMultiError) AllErrors() []error { return m }

// EnvelopeValidationError is the validation error returned by
// Envelope.Validate if the designated constraints aren't met.
type EnvelopeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnvelopeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnvelopeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnvelopeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnvelopeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnvelopeValidationError) ErrorName() string { return "EnvelopeValidationError" }

// Error satisfies the builtin error interface
func (e EnvelopeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnvelope.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnvelopeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnvelopeValidationError{}

// Validate checks the field values on Payload with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Payload) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Payload with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in PayloadMultiError, or nil if none found.
func (m *Payload) ValidateAll() error {
	return m.validate(true)
}

func (m *Payload) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, PayloadValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, PayloadValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return PayloadValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	switch m.Contents.(type) {

	case *Payload_UpdateConfig:

		if all {
			switch v := interface{}(m.GetUpdateConfig()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PayloadValidationError{
						field:  "UpdateConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PayloadValidationError{
						field:  "UpdateConfig",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetUpdateConfig()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PayloadValidationError{
					field:  "UpdateConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Payload_IssueCertificate:

		if all {
			switch v := interface{}(m.GetIssueCertificate()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PayloadValidationError{
						field:  "IssueCertificate",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PayloadValidationError{
						field:  "IssueCertificate",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetIssueCertificate()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PayloadValidationError{
					field:  "IssueCertificate",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	case *Payload_RevokeCertificate:

		if all {
			switch v := interface{}(m.GetRevokeCertificate()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, PayloadValidationError{
						field:  "RevokeCertificate",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, PayloadValidationError{
						field:  "RevokeCertificate",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetRevokeCertificate()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return PayloadValidationError{
					field:  "RevokeCertificate",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return PayloadMultiError(errors)
	}
	return nil
}

// PayloadMultiError is an error wrapping multiple validation errors returned
// by Payload.ValidateAll() if the designated constraints aren't met.
type PayloadMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m PayloadMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m PayloadMultiError) AllErrors() []error { return m }

// PayloadValidationError is the validation error returned by Payload.Validate
// if the designated constraints aren't met.
type PayloadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e PayloadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e PayloadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e PayloadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e PayloadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e PayloadValidationError) ErrorName() string { return "PayloadValidationError" }

// Error satisfies the builtin error interface
func (e PayloadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sPayload.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = PayloadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = PayloadValidationError{}

// Validate checks the field values on StoreRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StoreRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StoreRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StoreRequestMultiError, or
// nil if none found.
func (m *StoreRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *StoreRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEnvelope()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, StoreRequestValidationError{
					field:  "Envelope",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, StoreRequestValidationError{
					field:  "Envelope",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEnvelope()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return StoreRequestValidationError{
				field:  "Envelope",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return StoreRequestMultiError(errors)
	}
	return nil
}

// StoreRequestMultiError is an error wrapping multiple validation errors
// returned by StoreRequest.ValidateAll() if the designated constraints aren't met.
type StoreRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StoreRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StoreRequestMultiError) AllErrors() []error { return m }

// StoreRequestValidationError is the validation error returned by
// StoreRequest.Validate if the designated constraints aren't met.
type StoreRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StoreRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StoreRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StoreRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StoreRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StoreRequestValidationError) ErrorName() string { return "StoreRequestValidationError" }

// Error satisfies the builtin error interface
func (e StoreRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStoreRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StoreRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StoreRequestValidationError{}

// Validate checks the field values on StoreResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *StoreResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on StoreResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in StoreResponseMultiError, or
// nil if none found.
func (m *StoreResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *StoreResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return StoreResponseMultiError(errors)
	}
	return nil
}

// StoreResponseMultiError is an error wrapping multiple validation errors
// returned by StoreResponse.ValidateAll() if the designated constraints
// aren't met.
type StoreResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m StoreResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m StoreResponseMultiError) AllErrors() []error { return m }

// StoreResponseValidationError is the validation error returned by
// StoreResponse.Validate if the designated constraints aren't met.
type StoreResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e StoreResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e StoreResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e StoreResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e StoreResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e StoreResponseValidationError) ErrorName() string { return "StoreResponseValidationError" }

// Error satisfies the builtin error interface
func (e StoreResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sStoreResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = StoreResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = StoreResponseValidationError{}

// Validate checks the field values on GetEntriesRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetEntriesRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetEntriesRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetEntriesRequestMultiError, or nil if none found.
func (m *GetEntriesRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetEntriesRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for StartIndex

	// no validation rules for Count

	if len(errors) > 0 {
		return GetEntriesRequestMultiError(errors)
	}
	return nil
}

// GetEntriesRequestMultiError is an error wrapping multiple validation errors
// returned by GetEntriesRequest.ValidateAll() if the designated constraints
// aren't met.
type GetEntriesRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetEntriesRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetEntriesRequestMultiError) AllErrors() []error { return m }

// GetEntriesRequestValidationError is the validation error returned by
// GetEntriesRequest.Validate if the designated constraints aren't met.
type GetEntriesRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEntriesRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEntriesRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEntriesRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEntriesRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEntriesRequestValidationError) ErrorName() string {
	return "GetEntriesRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetEntriesRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEntriesRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEntriesRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEntriesRequestValidationError{}

// Validate checks the field values on GetEntriesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetEntriesResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetEntriesResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetEntriesResponseMultiError, or nil if none found.
func (m *GetEntriesResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetEntriesResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetEnvelopes() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, GetEntriesResponseValidationError{
						field:  fmt.Sprintf("Envelopes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, GetEntriesResponseValidationError{
						field:  fmt.Sprintf("Envelopes[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return GetEntriesResponseValidationError{
					field:  fmt.Sprintf("Envelopes[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return GetEntriesResponseMultiError(errors)
	}
	return nil
}

// GetEntriesResponseMultiError is an error wrapping multiple validation errors
// returned by GetEntriesResponse.ValidateAll() if the designated constraints
// aren't met.
type GetEntriesResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetEntriesResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetEntriesResponseMultiError) AllErrors() []error { return m }

// GetEntriesResponseValidationError is the validation error returned by
// GetEntriesResponse.Validate if the designated constraints aren't met.
type GetEntriesResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetEntriesResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetEntriesResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetEntriesResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetEntriesResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetEntriesResponseValidationError) ErrorName() string {
	return "GetEntriesResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetEntriesResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetEntriesResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetEntriesResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetEntriesResponseValidationError{}

// Validate checks the field values on IncludedEnvelope with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *IncludedEnvelope) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on IncludedEnvelope with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// IncludedEnvelopeMultiError, or nil if none found.
func (m *IncludedEnvelope) ValidateAll() error {
	return m.validate(true)
}

func (m *IncludedEnvelope) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetEnvelope()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, IncludedEnvelopeValidationError{
					field:  "Envelope",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, IncludedEnvelopeValidationError{
					field:  "Envelope",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetEnvelope()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return IncludedEnvelopeValidationError{
				field:  "Envelope",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return IncludedEnvelopeMultiError(errors)
	}
	return nil
}

// IncludedEnvelopeMultiError is an error wrapping multiple validation errors
// returned by IncludedEnvelope.ValidateAll() if the designated constraints
// aren't met.
type IncludedEnvelopeMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IncludedEnvelopeMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IncludedEnvelopeMultiError) AllErrors() []error { return m }

// IncludedEnvelopeValidationError is the validation error returned by
// IncludedEnvelope.Validate if the designated constraints aren't met.
type IncludedEnvelopeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IncludedEnvelopeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IncludedEnvelopeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IncludedEnvelopeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IncludedEnvelopeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IncludedEnvelopeValidationError) ErrorName() string { return "IncludedEnvelopeValidationError" }

// Error satisfies the builtin error interface
func (e IncludedEnvelopeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIncludedEnvelope.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IncludedEnvelopeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IncludedEnvelopeValidationError{}
