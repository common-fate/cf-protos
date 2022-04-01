// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: provider/v1alpha1/provider.proto

package providerv1alpha1

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

// Validate checks the field values on GetProviderConfigByDigestRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetProviderConfigByDigestRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProviderConfigByDigestRequest with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// GetProviderConfigByDigestRequestMultiError, or nil if none found.
func (m *GetProviderConfigByDigestRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProviderConfigByDigestRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sha256Digest

	if len(errors) > 0 {
		return GetProviderConfigByDigestRequestMultiError(errors)
	}
	return nil
}

// GetProviderConfigByDigestRequestMultiError is an error wrapping multiple
// validation errors returned by
// GetProviderConfigByDigestRequest.ValidateAll() if the designated
// constraints aren't met.
type GetProviderConfigByDigestRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProviderConfigByDigestRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProviderConfigByDigestRequestMultiError) AllErrors() []error { return m }

// GetProviderConfigByDigestRequestValidationError is the validation error
// returned by GetProviderConfigByDigestRequest.Validate if the designated
// constraints aren't met.
type GetProviderConfigByDigestRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProviderConfigByDigestRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProviderConfigByDigestRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProviderConfigByDigestRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProviderConfigByDigestRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProviderConfigByDigestRequestValidationError) ErrorName() string {
	return "GetProviderConfigByDigestRequestValidationError"
}

// Error satisfies the builtin error interface
func (e GetProviderConfigByDigestRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProviderConfigByDigestRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProviderConfigByDigestRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProviderConfigByDigestRequestValidationError{}

// Validate checks the field values on GetProviderConfigByDigestResponse with
// the rules defined in the proto definition for this message. If any rules
// are violated, the first error encountered is returned, or nil if there are
// no violations.
func (m *GetProviderConfigByDigestResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetProviderConfigByDigestResponse
// with the rules defined in the proto definition for this message. If any
// rules are violated, the result is a list of violation errors wrapped in
// GetProviderConfigByDigestResponseMultiError, or nil if none found.
func (m *GetProviderConfigByDigestResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *GetProviderConfigByDigestResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetProvider()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetProviderConfigByDigestResponseValidationError{
					field:  "Provider",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetProviderConfigByDigestResponseValidationError{
					field:  "Provider",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetProvider()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetProviderConfigByDigestResponseValidationError{
				field:  "Provider",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetProviderConfigByDigestResponseMultiError(errors)
	}
	return nil
}

// GetProviderConfigByDigestResponseMultiError is an error wrapping multiple
// validation errors returned by
// GetProviderConfigByDigestResponse.ValidateAll() if the designated
// constraints aren't met.
type GetProviderConfigByDigestResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetProviderConfigByDigestResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetProviderConfigByDigestResponseMultiError) AllErrors() []error { return m }

// GetProviderConfigByDigestResponseValidationError is the validation error
// returned by GetProviderConfigByDigestResponse.Validate if the designated
// constraints aren't met.
type GetProviderConfigByDigestResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetProviderConfigByDigestResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetProviderConfigByDigestResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetProviderConfigByDigestResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetProviderConfigByDigestResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetProviderConfigByDigestResponseValidationError) ErrorName() string {
	return "GetProviderConfigByDigestResponseValidationError"
}

// Error satisfies the builtin error interface
func (e GetProviderConfigByDigestResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetProviderConfigByDigestResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetProviderConfigByDigestResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetProviderConfigByDigestResponseValidationError{}
