// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: account/v1alpha1/account.proto

package accountv1alpha1

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

// Validate checks the field values on SignupRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignupRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignupRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignupRequestMultiError, or
// nil if none found.
func (m *SignupRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SignupRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateEmail(m.GetAdminEmailAddress()); err != nil {
		err = SignupRequestValidationError{
			field:  "AdminEmailAddress",
			reason: "value must be a valid email address",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetTenantName()) > 50 {
		err := SignupRequestValidationError{
			field:  "TenantName",
			reason: "value length must be at most 50 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if utf8.RuneCountInString(m.GetTenantSlug()) > 50 {
		err := SignupRequestValidationError{
			field:  "TenantSlug",
			reason: "value length must be at most 50 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return SignupRequestMultiError(errors)
	}
	return nil
}

func (m *SignupRequest) _validateHostname(host string) error {
	s := strings.ToLower(strings.TrimSuffix(host, "."))

	if len(host) > 253 {
		return errors.New("hostname cannot exceed 253 characters")
	}

	for _, part := range strings.Split(s, ".") {
		if l := len(part); l == 0 || l > 63 {
			return errors.New("hostname part must be non-empty and cannot exceed 63 characters")
		}

		if part[0] == '-' {
			return errors.New("hostname parts cannot begin with hyphens")
		}

		if part[len(part)-1] == '-' {
			return errors.New("hostname parts cannot end with hyphens")
		}

		for _, r := range part {
			if (r < 'a' || r > 'z') && (r < '0' || r > '9') && r != '-' {
				return fmt.Errorf("hostname parts can only contain alphanumeric characters or hyphens, got %q", string(r))
			}
		}
	}

	return nil
}

func (m *SignupRequest) _validateEmail(addr string) error {
	a, err := mail.ParseAddress(addr)
	if err != nil {
		return err
	}
	addr = a.Address

	if len(addr) > 254 {
		return errors.New("email addresses cannot exceed 254 characters")
	}

	parts := strings.SplitN(addr, "@", 2)

	if len(parts[0]) > 64 {
		return errors.New("email address local phrase cannot exceed 64 characters")
	}

	return m._validateHostname(parts[1])
}

// SignupRequestMultiError is an error wrapping multiple validation errors
// returned by SignupRequest.ValidateAll() if the designated constraints
// aren't met.
type SignupRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignupRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignupRequestMultiError) AllErrors() []error { return m }

// SignupRequestValidationError is the validation error returned by
// SignupRequest.Validate if the designated constraints aren't met.
type SignupRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignupRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignupRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignupRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignupRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignupRequestValidationError) ErrorName() string { return "SignupRequestValidationError" }

// Error satisfies the builtin error interface
func (e SignupRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignupRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignupRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignupRequestValidationError{}

// Validate checks the field values on SignupResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *SignupResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SignupResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in SignupResponseMultiError,
// or nil if none found.
func (m *SignupResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *SignupResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return SignupResponseMultiError(errors)
	}
	return nil
}

// SignupResponseMultiError is an error wrapping multiple validation errors
// returned by SignupResponse.ValidateAll() if the designated constraints
// aren't met.
type SignupResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SignupResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SignupResponseMultiError) AllErrors() []error { return m }

// SignupResponseValidationError is the validation error returned by
// SignupResponse.Validate if the designated constraints aren't met.
type SignupResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SignupResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SignupResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SignupResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SignupResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SignupResponseValidationError) ErrorName() string { return "SignupResponseValidationError" }

// Error satisfies the builtin error interface
func (e SignupResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSignupResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SignupResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SignupResponseValidationError{}

// Validate checks the field values on GetDeviceIdRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeviceIdRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeviceIdRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeviceIdRequestMultiError, or nil if none found.
func (m *GetDeviceIdRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeviceIdRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetDeviceIdRequestMultiError(errors)
	}
	return nil
}

// GetDeviceIdRequestMultiError is an error wrapping multiple validation errors
// returned by GetDeviceIdRequest.ValidateAll() if the designated constraints
// aren't met.
type GetDeviceIdRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeviceIdRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeviceIdRequestMultiError) AllErrors() []error { return m }

// GetDeviceIdRequestValidationError is the validation error returned by
// GetDeviceIdRequest.Validate if the designated constraints aren't met.
type GetDeviceIdRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeviceIdRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeviceIdRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeviceIdRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeviceIdRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeviceIdRequestValidationError) ErrorName() string {
	return "GetDeviceIdRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeviceIdRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeviceIdRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeviceIdRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeviceIdRequestValidationError{}

// Validate checks the field values on GetDeviceIdResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *GetDeviceIdResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetDeviceIdResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetDeviceIdResponseMultiError, or nil if none found.
func (m *GetDeviceIdResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetDeviceIdResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	if len(errors) > 0 {
		return GetDeviceIdResponseMultiError(errors)
	}
	return nil
}

// GetDeviceIdResponseMultiError is an error wrapping multiple validation
// errors returned by GetDeviceIdResponse.ValidateAll() if the designated
// constraints aren't met.
type GetDeviceIdResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetDeviceIdResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetDeviceIdResponseMultiError) AllErrors() []error { return m }

// GetDeviceIdResponseValidationError is the validation error returned by
// GetDeviceIdResponse.Validate if the designated constraints aren't met.
type GetDeviceIdResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetDeviceIdResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetDeviceIdResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetDeviceIdResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetDeviceIdResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetDeviceIdResponseValidationError) ErrorName() string {
	return "GetDeviceIdResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetDeviceIdResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetDeviceIdResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetDeviceIdResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetDeviceIdResponseValidationError{}

// Validate checks the field values on AuthenticatedRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AuthenticatedRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthenticatedRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthenticatedRequestMultiError, or nil if none found.
func (m *AuthenticatedRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthenticatedRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AuthenticatedRequestMultiError(errors)
	}
	return nil
}

// AuthenticatedRequestMultiError is an error wrapping multiple validation
// errors returned by AuthenticatedRequest.ValidateAll() if the designated
// constraints aren't met.
type AuthenticatedRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthenticatedRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthenticatedRequestMultiError) AllErrors() []error { return m }

// AuthenticatedRequestValidationError is the validation error returned by
// AuthenticatedRequest.Validate if the designated constraints aren't met.
type AuthenticatedRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthenticatedRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthenticatedRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthenticatedRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthenticatedRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthenticatedRequestValidationError) ErrorName() string {
	return "AuthenticatedRequestValidationError"
}

// Error satisfies the builtin error interface
func (e AuthenticatedRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthenticatedRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthenticatedRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthenticatedRequestValidationError{}

// Validate checks the field values on AuthenticatedResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AuthenticatedResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AuthenticatedResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AuthenticatedResponseMultiError, or nil if none found.
func (m *AuthenticatedResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *AuthenticatedResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return AuthenticatedResponseMultiError(errors)
	}
	return nil
}

// AuthenticatedResponseMultiError is an error wrapping multiple validation
// errors returned by AuthenticatedResponse.ValidateAll() if the designated
// constraints aren't met.
type AuthenticatedResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AuthenticatedResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AuthenticatedResponseMultiError) AllErrors() []error { return m }

// AuthenticatedResponseValidationError is the validation error returned by
// AuthenticatedResponse.Validate if the designated constraints aren't met.
type AuthenticatedResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AuthenticatedResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AuthenticatedResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AuthenticatedResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AuthenticatedResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AuthenticatedResponseValidationError) ErrorName() string {
	return "AuthenticatedResponseValidationError"
}

// Error satisfies the builtin error interface
func (e AuthenticatedResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAuthenticatedResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AuthenticatedResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AuthenticatedResponseValidationError{}
