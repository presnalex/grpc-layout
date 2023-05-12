// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: grpc_layout_proto/ping_pong.proto

package hello

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _ping_pong_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on RequestMsg with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *RequestMsg) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Msg

	return nil
}

// RequestMsgValidationError is the validation error returned by
// RequestMsg.Validate if the designated constraints aren't met.
type RequestMsgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RequestMsgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RequestMsgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RequestMsgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RequestMsgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RequestMsgValidationError) ErrorName() string { return "RequestMsgValidationError" }

// Error satisfies the builtin error interface
func (e RequestMsgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRequestMsg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RequestMsgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RequestMsgValidationError{}

// Validate checks the field values on ResponseMsg with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ResponseMsg) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Msg

	return nil
}

// ResponseMsgValidationError is the validation error returned by
// ResponseMsg.Validate if the designated constraints aren't met.
type ResponseMsgValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ResponseMsgValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ResponseMsgValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ResponseMsgValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ResponseMsgValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ResponseMsgValidationError) ErrorName() string { return "ResponseMsgValidationError" }

// Error satisfies the builtin error interface
func (e ResponseMsgValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sResponseMsg.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ResponseMsgValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ResponseMsgValidationError{}
