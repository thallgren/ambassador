// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/filter/http/grpc_stats/v2alpha/config.proto

package envoy_config_filter_http_grpc_stats_v2alpha

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
)

// Validate checks the field values on FilterConfig with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *FilterConfig) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for EmitFilterState

	return nil
}

// FilterConfigValidationError is the validation error returned by
// FilterConfig.Validate if the designated constraints aren't met.
type FilterConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e FilterConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e FilterConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e FilterConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e FilterConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e FilterConfigValidationError) ErrorName() string { return "FilterConfigValidationError" }

// Error satisfies the builtin error interface
func (e FilterConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sFilterConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = FilterConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = FilterConfigValidationError{}
