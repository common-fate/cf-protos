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

// Validate checks the field values on EnrolRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EnrolRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EnrolRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EnrolRequestMultiError, or
// nil if none found.
func (m *EnrolRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *EnrolRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for EnrolToken

	switch m.Enrollment.(type) {

	case *EnrolRequest_Aws:

		if all {
			switch v := interface{}(m.GetAws()).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, EnrolRequestValidationError{
						field:  "Aws",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, EnrolRequestValidationError{
						field:  "Aws",
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(m.GetAws()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return EnrolRequestValidationError{
					field:  "Aws",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if len(errors) > 0 {
		return EnrolRequestMultiError(errors)
	}
	return nil
}

// EnrolRequestMultiError is an error wrapping multiple validation errors
// returned by EnrolRequest.ValidateAll() if the designated constraints aren't met.
type EnrolRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnrolRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnrolRequestMultiError) AllErrors() []error { return m }

// EnrolRequestValidationError is the validation error returned by
// EnrolRequest.Validate if the designated constraints aren't met.
type EnrolRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnrolRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnrolRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnrolRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnrolRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnrolRequestValidationError) ErrorName() string { return "EnrolRequestValidationError" }

// Error satisfies the builtin error interface
func (e EnrolRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnrolRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnrolRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnrolRequestValidationError{}

// Validate checks the field values on EnrolResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *EnrolResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on EnrolResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in EnrolResponseMultiError, or
// nil if none found.
func (m *EnrolResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *EnrolResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EnrolResponseMultiError(errors)
	}
	return nil
}

// EnrolResponseMultiError is an error wrapping multiple validation errors
// returned by EnrolResponse.ValidateAll() if the designated constraints
// aren't met.
type EnrolResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EnrolResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EnrolResponseMultiError) AllErrors() []error { return m }

// EnrolResponseValidationError is the validation error returned by
// EnrolResponse.Validate if the designated constraints aren't met.
type EnrolResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EnrolResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EnrolResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EnrolResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EnrolResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EnrolResponseValidationError) ErrorName() string { return "EnrolResponseValidationError" }

// Error satisfies the builtin error interface
func (e EnrolResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEnrolResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EnrolResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EnrolResponseValidationError{}

// Validate checks the field values on AWSEnrollment with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AWSEnrollment) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AWSEnrollment with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AWSEnrollmentMultiError, or
// nil if none found.
func (m *AWSEnrollment) ValidateAll() error {
	return m.validate(true)
}

func (m *AWSEnrollment) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetIdentity()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AWSEnrollmentValidationError{
					field:  "Identity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AWSEnrollmentValidationError{
					field:  "Identity",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetIdentity()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AWSEnrollmentValidationError{
				field:  "Identity",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetOrganization()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AWSEnrollmentValidationError{
					field:  "Organization",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AWSEnrollmentValidationError{
					field:  "Organization",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetOrganization()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AWSEnrollmentValidationError{
				field:  "Organization",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return AWSEnrollmentMultiError(errors)
	}
	return nil
}

// AWSEnrollmentMultiError is an error wrapping multiple validation errors
// returned by AWSEnrollment.ValidateAll() if the designated constraints
// aren't met.
type AWSEnrollmentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AWSEnrollmentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AWSEnrollmentMultiError) AllErrors() []error { return m }

// AWSEnrollmentValidationError is the validation error returned by
// AWSEnrollment.Validate if the designated constraints aren't met.
type AWSEnrollmentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AWSEnrollmentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AWSEnrollmentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AWSEnrollmentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AWSEnrollmentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AWSEnrollmentValidationError) ErrorName() string { return "AWSEnrollmentValidationError" }

// Error satisfies the builtin error interface
func (e AWSEnrollmentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAWSEnrollment.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AWSEnrollmentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AWSEnrollmentValidationError{}

// Validate checks the field values on AWSProof with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *AWSProof) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on AWSProof with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in AWSProofMultiError, or nil
// if none found.
func (m *AWSProof) ValidateAll() error {
	return m.validate(true)
}

func (m *AWSProof) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Signature

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, AWSProofValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, AWSProofValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return AWSProofValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for SecurityToken

	if len(errors) > 0 {
		return AWSProofMultiError(errors)
	}
	return nil
}

// AWSProofMultiError is an error wrapping multiple validation errors returned
// by AWSProof.ValidateAll() if the designated constraints aren't met.
type AWSProofMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m AWSProofMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m AWSProofMultiError) AllErrors() []error { return m }

// AWSProofValidationError is the validation error returned by
// AWSProof.Validate if the designated constraints aren't met.
type AWSProofValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AWSProofValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AWSProofValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AWSProofValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AWSProofValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AWSProofValidationError) ErrorName() string { return "AWSProofValidationError" }

// Error satisfies the builtin error interface
func (e AWSProofValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAWSProof.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AWSProofValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AWSProofValidationError{}
