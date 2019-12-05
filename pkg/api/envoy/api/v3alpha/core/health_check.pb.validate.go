// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/api/v3alpha/core/health_check.proto

package envoy_api_v3alpha_core

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

	"github.com/gogo/protobuf/types"

	envoy_type_v3alpha "github.com/datawire/ambassador/pkg/api/envoy/type/v3alpha"
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
	_ = types.DynamicAny{}

	_ = envoy_type_v3alpha.CodecClientType(0)
)

// Validate checks the field values on HealthCheck with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *HealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	if m.GetTimeout() == nil {
		return HealthCheckValidationError{
			field:  "Timeout",
			reason: "value is required",
		}
	}

	if d := m.GetTimeout(); d != nil {
		dur, err := types.DurationFromProto(d)
		if err != nil {
			return HealthCheckValidationError{
				field:  "Timeout",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return HealthCheckValidationError{
				field:  "Timeout",
				reason: "value must be greater than 0s",
			}
		}

	}

	if m.GetInterval() == nil {
		return HealthCheckValidationError{
			field:  "Interval",
			reason: "value is required",
		}
	}

	if d := m.GetInterval(); d != nil {
		dur, err := types.DurationFromProto(d)
		if err != nil {
			return HealthCheckValidationError{
				field:  "Interval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return HealthCheckValidationError{
				field:  "Interval",
				reason: "value must be greater than 0s",
			}
		}

	}

	{
		tmp := m.GetInitialJitter()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "InitialJitter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetIntervalJitter()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "IntervalJitter",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for IntervalJitterPercent

	{
		tmp := m.GetUnhealthyThreshold()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "UnhealthyThreshold",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetHealthyThreshold()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "HealthyThreshold",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetAltPort()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "AltPort",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetReuseConnection()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return HealthCheckValidationError{
					field:  "ReuseConnection",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if d := m.GetNoTrafficInterval(); d != nil {
		dur, err := types.DurationFromProto(d)
		if err != nil {
			return HealthCheckValidationError{
				field:  "NoTrafficInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return HealthCheckValidationError{
				field:  "NoTrafficInterval",
				reason: "value must be greater than 0s",
			}
		}

	}

	if d := m.GetUnhealthyInterval(); d != nil {
		dur, err := types.DurationFromProto(d)
		if err != nil {
			return HealthCheckValidationError{
				field:  "UnhealthyInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return HealthCheckValidationError{
				field:  "UnhealthyInterval",
				reason: "value must be greater than 0s",
			}
		}

	}

	if d := m.GetUnhealthyEdgeInterval(); d != nil {
		dur, err := types.DurationFromProto(d)
		if err != nil {
			return HealthCheckValidationError{
				field:  "UnhealthyEdgeInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return HealthCheckValidationError{
				field:  "UnhealthyEdgeInterval",
				reason: "value must be greater than 0s",
			}
		}

	}

	if d := m.GetHealthyEdgeInterval(); d != nil {
		dur, err := types.DurationFromProto(d)
		if err != nil {
			return HealthCheckValidationError{
				field:  "HealthyEdgeInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		gt := time.Duration(0*time.Second + 0*time.Nanosecond)

		if dur <= gt {
			return HealthCheckValidationError{
				field:  "HealthyEdgeInterval",
				reason: "value must be greater than 0s",
			}
		}

	}

	// no validation rules for EventLogPath

	// no validation rules for AlwaysLogHealthCheckFailures

	switch m.HealthChecker.(type) {

	case *HealthCheck_HttpHealthCheck_:

		{
			tmp := m.GetHttpHealthCheck()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return HealthCheckValidationError{
						field:  "HttpHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *HealthCheck_TcpHealthCheck_:

		{
			tmp := m.GetTcpHealthCheck()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return HealthCheckValidationError{
						field:  "TcpHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *HealthCheck_GrpcHealthCheck_:

		{
			tmp := m.GetGrpcHealthCheck()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return HealthCheckValidationError{
						field:  "GrpcHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *HealthCheck_CustomHealthCheck_:

		{
			tmp := m.GetCustomHealthCheck()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return HealthCheckValidationError{
						field:  "CustomHealthCheck",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	default:
		return HealthCheckValidationError{
			field:  "HealthChecker",
			reason: "value is required",
		}

	}

	return nil
}

// HealthCheckValidationError is the validation error returned by
// HealthCheck.Validate if the designated constraints aren't met.
type HealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckValidationError) ErrorName() string { return "HealthCheckValidationError" }

// Error satisfies the builtin error interface
func (e HealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckValidationError{}

// Validate checks the field values on HealthCheck_Payload with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheck_Payload) Validate() error {
	if m == nil {
		return nil
	}

	switch m.Payload.(type) {

	case *HealthCheck_Payload_Text:

		if len(m.GetText()) < 1 {
			return HealthCheck_PayloadValidationError{
				field:  "Text",
				reason: "value length must be at least 1 bytes",
			}
		}

	case *HealthCheck_Payload_Binary:
		// no validation rules for Binary

	default:
		return HealthCheck_PayloadValidationError{
			field:  "Payload",
			reason: "value is required",
		}

	}

	return nil
}

// HealthCheck_PayloadValidationError is the validation error returned by
// HealthCheck_Payload.Validate if the designated constraints aren't met.
type HealthCheck_PayloadValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_PayloadValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_PayloadValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_PayloadValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_PayloadValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_PayloadValidationError) ErrorName() string {
	return "HealthCheck_PayloadValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_PayloadValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_Payload.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_PayloadValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_PayloadValidationError{}

// Validate checks the field values on HealthCheck_HttpHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheck_HttpHealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Host

	if len(m.GetPath()) < 1 {
		return HealthCheck_HttpHealthCheckValidationError{
			field:  "Path",
			reason: "value length must be at least 1 bytes",
		}
	}

	{
		tmp := m.GetSend()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return HealthCheck_HttpHealthCheckValidationError{
					field:  "Send",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetReceive()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return HealthCheck_HttpHealthCheckValidationError{
					field:  "Receive",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for ServiceName

	if len(m.GetRequestHeadersToAdd()) > 1000 {
		return HealthCheck_HttpHealthCheckValidationError{
			field:  "RequestHeadersToAdd",
			reason: "value must contain no more than 1000 item(s)",
		}
	}

	for idx, item := range m.GetRequestHeadersToAdd() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return HealthCheck_HttpHealthCheckValidationError{
						field:  fmt.Sprintf("RequestHeadersToAdd[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	for idx, item := range m.GetExpectedStatuses() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return HealthCheck_HttpHealthCheckValidationError{
						field:  fmt.Sprintf("ExpectedStatuses[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	if _, ok := envoy_type_v3alpha.CodecClientType_name[int32(m.GetCodecClientType())]; !ok {
		return HealthCheck_HttpHealthCheckValidationError{
			field:  "CodecClientType",
			reason: "value must be one of the defined enum values",
		}
	}

	return nil
}

// HealthCheck_HttpHealthCheckValidationError is the validation error returned
// by HealthCheck_HttpHealthCheck.Validate if the designated constraints
// aren't met.
type HealthCheck_HttpHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_HttpHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_HttpHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_HttpHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_HttpHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_HttpHealthCheckValidationError) ErrorName() string {
	return "HealthCheck_HttpHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_HttpHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_HttpHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_HttpHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_HttpHealthCheckValidationError{}

// Validate checks the field values on HealthCheck_TcpHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheck_TcpHealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetSend()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return HealthCheck_TcpHealthCheckValidationError{
					field:  "Send",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for idx, item := range m.GetReceive() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return HealthCheck_TcpHealthCheckValidationError{
						field:  fmt.Sprintf("Receive[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// HealthCheck_TcpHealthCheckValidationError is the validation error returned
// by HealthCheck_TcpHealthCheck.Validate if the designated constraints aren't met.
type HealthCheck_TcpHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_TcpHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_TcpHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_TcpHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_TcpHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_TcpHealthCheckValidationError) ErrorName() string {
	return "HealthCheck_TcpHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_TcpHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_TcpHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_TcpHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_TcpHealthCheckValidationError{}

// Validate checks the field values on HealthCheck_RedisHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheck_RedisHealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Key

	return nil
}

// HealthCheck_RedisHealthCheckValidationError is the validation error returned
// by HealthCheck_RedisHealthCheck.Validate if the designated constraints
// aren't met.
type HealthCheck_RedisHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_RedisHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_RedisHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_RedisHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_RedisHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_RedisHealthCheckValidationError) ErrorName() string {
	return "HealthCheck_RedisHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_RedisHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_RedisHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_RedisHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_RedisHealthCheckValidationError{}

// Validate checks the field values on HealthCheck_GrpcHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheck_GrpcHealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for ServiceName

	// no validation rules for Authority

	return nil
}

// HealthCheck_GrpcHealthCheckValidationError is the validation error returned
// by HealthCheck_GrpcHealthCheck.Validate if the designated constraints
// aren't met.
type HealthCheck_GrpcHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_GrpcHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_GrpcHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_GrpcHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_GrpcHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_GrpcHealthCheckValidationError) ErrorName() string {
	return "HealthCheck_GrpcHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_GrpcHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_GrpcHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_GrpcHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_GrpcHealthCheckValidationError{}

// Validate checks the field values on HealthCheck_CustomHealthCheck with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheck_CustomHealthCheck) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return HealthCheck_CustomHealthCheckValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	switch m.ConfigType.(type) {

	case *HealthCheck_CustomHealthCheck_TypedConfig:

		{
			tmp := m.GetTypedConfig()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return HealthCheck_CustomHealthCheckValidationError{
						field:  "TypedConfig",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// HealthCheck_CustomHealthCheckValidationError is the validation error
// returned by HealthCheck_CustomHealthCheck.Validate if the designated
// constraints aren't met.
type HealthCheck_CustomHealthCheckValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheck_CustomHealthCheckValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheck_CustomHealthCheckValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheck_CustomHealthCheckValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheck_CustomHealthCheckValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheck_CustomHealthCheckValidationError) ErrorName() string {
	return "HealthCheck_CustomHealthCheckValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheck_CustomHealthCheckValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheck_CustomHealthCheck.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheck_CustomHealthCheckValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheck_CustomHealthCheckValidationError{}
