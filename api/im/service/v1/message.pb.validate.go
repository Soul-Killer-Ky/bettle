// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/im/service/v1/message.proto

package v1

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

// Validate checks the field values on Content with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Content) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Content with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ContentMultiError, or nil if none found.
func (m *Content) ValidateAll() error {
	return m.validate(true)
}

func (m *Content) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Type

	// no validation rules for Body

	if len(errors) > 0 {
		return ContentMultiError(errors)
	}

	return nil
}

// ContentMultiError is an error wrapping multiple validation errors returned
// by Content.ValidateAll() if the designated constraints aren't met.
type ContentMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ContentMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ContentMultiError) AllErrors() []error { return m }

// ContentValidationError is the validation error returned by Content.Validate
// if the designated constraints aren't met.
type ContentValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ContentValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ContentValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ContentValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ContentValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ContentValidationError) ErrorName() string { return "ContentValidationError" }

// Error satisfies the builtin error interface
func (e ContentValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sContent.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ContentValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ContentValidationError{}

// Validate checks the field values on ChatMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ChatMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ChatMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ChatMessageMultiError, or
// nil if none found.
func (m *ChatMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *ChatMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetContent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ChatMessageValidationError{
					field:  "Content",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatMessageValidationError{
					field:  "Content",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetContent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatMessageValidationError{
				field:  "Content",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for From

	// no validation rules for Sender

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ChatMessageValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ChatMessageValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ChatMessageValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ChatMessageMultiError(errors)
	}

	return nil
}

// ChatMessageMultiError is an error wrapping multiple validation errors
// returned by ChatMessage.ValidateAll() if the designated constraints aren't met.
type ChatMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ChatMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ChatMessageMultiError) AllErrors() []error { return m }

// ChatMessageValidationError is the validation error returned by
// ChatMessage.Validate if the designated constraints aren't met.
type ChatMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ChatMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ChatMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ChatMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ChatMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ChatMessageValidationError) ErrorName() string { return "ChatMessageValidationError" }

// Error satisfies the builtin error interface
func (e ChatMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sChatMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ChatMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ChatMessageValidationError{}

// Validate checks the field values on GroupMessage with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GroupMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GroupMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GroupMessageMultiError, or
// nil if none found.
func (m *GroupMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *GroupMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetContent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GroupMessageValidationError{
					field:  "Content",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GroupMessageValidationError{
					field:  "Content",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetContent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GroupMessageValidationError{
				field:  "Content",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for From

	// no validation rules for GroupId

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GroupMessageValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GroupMessageValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GroupMessageValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GroupMessageMultiError(errors)
	}

	return nil
}

// GroupMessageMultiError is an error wrapping multiple validation errors
// returned by GroupMessage.ValidateAll() if the designated constraints aren't met.
type GroupMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GroupMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GroupMessageMultiError) AllErrors() []error { return m }

// GroupMessageValidationError is the validation error returned by
// GroupMessage.Validate if the designated constraints aren't met.
type GroupMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GroupMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GroupMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GroupMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GroupMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GroupMessageValidationError) ErrorName() string { return "GroupMessageValidationError" }

// Error satisfies the builtin error interface
func (e GroupMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGroupMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GroupMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GroupMessageValidationError{}

// Validate checks the field values on BroadcastMessage with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *BroadcastMessage) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BroadcastMessage with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BroadcastMessageMultiError, or nil if none found.
func (m *BroadcastMessage) ValidateAll() error {
	return m.validate(true)
}

func (m *BroadcastMessage) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetContent()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BroadcastMessageValidationError{
					field:  "Content",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BroadcastMessageValidationError{
					field:  "Content",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetContent()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BroadcastMessageValidationError{
				field:  "Content",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for From

	if all {
		switch v := interface{}(m.GetTimestamp()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, BroadcastMessageValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, BroadcastMessageValidationError{
					field:  "Timestamp",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetTimestamp()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return BroadcastMessageValidationError{
				field:  "Timestamp",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return BroadcastMessageMultiError(errors)
	}

	return nil
}

// BroadcastMessageMultiError is an error wrapping multiple validation errors
// returned by BroadcastMessage.ValidateAll() if the designated constraints
// aren't met.
type BroadcastMessageMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BroadcastMessageMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BroadcastMessageMultiError) AllErrors() []error { return m }

// BroadcastMessageValidationError is the validation error returned by
// BroadcastMessage.Validate if the designated constraints aren't met.
type BroadcastMessageValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BroadcastMessageValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BroadcastMessageValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BroadcastMessageValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BroadcastMessageValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BroadcastMessageValidationError) ErrorName() string { return "BroadcastMessageValidationError" }

// Error satisfies the builtin error interface
func (e BroadcastMessageValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBroadcastMessage.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BroadcastMessageValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BroadcastMessageValidationError{}
