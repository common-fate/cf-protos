// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: team/v1alpha1/assume_role.proto

package teamv1alpha1

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

// Validate checks the field values on AssumeRoleSignatureRoleAccount with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AssumeRoleSignatureRoleAccount) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AssumeRoleSignatureRoleAccount with
// the rules defined in the proto definition for this message. If any rules
// are violated, the result is a list of violation errors wrapped in
// AssumeRoleSignatureRoleAccountMultiError, or nil if none found.
func (m *AssumeRoleSignatureRoleAccount) ValidateAll() error {
	return m.validate(true)
}

func (m *AssumeRoleSignatureRoleAccount) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Role

	// no validation rules for Account

	if len(errors) > 0 {
		return AssumeRoleSignatureRoleAccountMultiError(errors)
	}
	return nil
}

// AssumeRoleSignatureRoleAccountMultiError is an error wrapping multiple
// validation errors returned by AssumeRoleSignatureRoleAccount.ValidateAll()
// if the designated constraints aren't met.
type AssumeRoleSignatureRoleAccountMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AssumeRoleSignatureRoleAccountMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AssumeRoleSignatureRoleAccountMultiError) AllErrors() []error { return m }

// AssumeRoleSignatureRoleAccountValidationError is the validation error
// returned by AssumeRoleSignatureRoleAccount.Validate if the designated
// constraints aren't met.
type AssumeRoleSignatureRoleAccountValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AssumeRoleSignatureRoleAccountValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AssumeRoleSignatureRoleAccountValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AssumeRoleSignatureRoleAccountValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AssumeRoleSignatureRoleAccountValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AssumeRoleSignatureRoleAccountValidationError) ErrorName() string {
	return "AssumeRoleSignatureRoleAccountValidationError"
}

// Error satisfies the builtin error interface
func (e AssumeRoleSignatureRoleAccountValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAssumeRoleSignatureRoleAccount.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AssumeRoleSignatureRoleAccountValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AssumeRoleSignatureRoleAccountValidationError{}

// Validate checks the field values on AssumeRoleSignatureTimestamp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AssumeRoleSignatureTimestamp) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AssumeRoleSignatureTimestamp with the
// rules defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AssumeRoleSignatureTimestampMultiError, or nil if none found.
func (m *AssumeRoleSignatureTimestamp) ValidateAll() error {
	return m.validate(true)
}

func (m *AssumeRoleSignatureTimestamp) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for InnerDigest

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AssumeRoleSignatureTimestampValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AssumeRoleSignatureTimestampValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AssumeRoleSignatureTimestampValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AssumeRoleSignatureTimestampMultiError(errors)
	}
	return nil
}

// AssumeRoleSignatureTimestampMultiError is an error wrapping multiple
// validation errors returned by AssumeRoleSignatureTimestamp.ValidateAll() if
// the designated constraints aren't met.
type AssumeRoleSignatureTimestampMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AssumeRoleSignatureTimestampMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AssumeRoleSignatureTimestampMultiError) AllErrors() []error { return m }

// AssumeRoleSignatureTimestampValidationError is the validation error returned
// by AssumeRoleSignatureTimestamp.Validate if the designated constraints
// aren't met.
type AssumeRoleSignatureTimestampValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AssumeRoleSignatureTimestampValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AssumeRoleSignatureTimestampValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AssumeRoleSignatureTimestampValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AssumeRoleSignatureTimestampValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AssumeRoleSignatureTimestampValidationError) ErrorName() string {
	return "AssumeRoleSignatureTimestampValidationError"
}

// Error satisfies the builtin error interface
func (e AssumeRoleSignatureTimestampValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAssumeRoleSignatureTimestamp.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AssumeRoleSignatureTimestampValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AssumeRoleSignatureTimestampValidationError{}

// Validate checks the field values on AssumeRoleSignature with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *AssumeRoleSignature) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AssumeRoleSignature with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// AssumeRoleSignatureMultiError, or nil if none found.
func (m *AssumeRoleSignature) ValidateAll() error {
	return m.validate(true)
}

func (m *AssumeRoleSignature) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for InnerDigest

	// no validation rules for CertificateFingerprint

	if len(errors) > 0 {
		return AssumeRoleSignatureMultiError(errors)
	}
	return nil
}

// AssumeRoleSignatureMultiError is an error wrapping multiple validation
// errors returned by AssumeRoleSignature.ValidateAll() if the designated
// constraints aren't met.
type AssumeRoleSignatureMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AssumeRoleSignatureMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AssumeRoleSignatureMultiError) AllErrors() []error { return m }

// AssumeRoleSignatureValidationError is the validation error returned by
// AssumeRoleSignature.Validate if the designated constraints aren't met.
type AssumeRoleSignatureValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AssumeRoleSignatureValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AssumeRoleSignatureValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AssumeRoleSignatureValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AssumeRoleSignatureValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AssumeRoleSignatureValidationError) ErrorName() string {
	return "AssumeRoleSignatureValidationError"
}

// Error satisfies the builtin error interface
func (e AssumeRoleSignatureValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAssumeRoleSignature.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AssumeRoleSignatureValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AssumeRoleSignatureValidationError{}