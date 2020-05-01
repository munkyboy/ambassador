// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/config/bootstrap/v2/bootstrap.proto

package envoy_config_bootstrap_v2

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

// define the regex for a UUID once up-front
var _bootstrap_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Bootstrap with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Bootstrap) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetNode()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "Node",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetStaticResources()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "StaticResources",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetDynamicResources()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "DynamicResources",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetClusterManager()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "ClusterManager",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetHdsConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "HdsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for FlagsPath

	for idx, item := range m.GetStatsSinks() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return BootstrapValidationError{
						field:  fmt.Sprintf("StatsSinks[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	{
		tmp := m.GetStatsConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "StatsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	if d := m.GetStatsFlushInterval(); d != nil {
		dur, err := types.DurationFromProto(d)
		if err != nil {
			return BootstrapValidationError{
				field:  "StatsFlushInterval",
				reason: "value is not a valid duration",
				cause:  err,
			}
		}

		lt := time.Duration(300*time.Second + 0*time.Nanosecond)
		gte := time.Duration(0*time.Second + 1000000*time.Nanosecond)

		if dur < gte || dur >= lt {
			return BootstrapValidationError{
				field:  "StatsFlushInterval",
				reason: "value must be inside range [1ms, 5m0s)",
			}
		}

	}

	{
		tmp := m.GetWatchdog()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "Watchdog",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetTracing()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "Tracing",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetRuntime()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "Runtime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetLayeredRuntime()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "LayeredRuntime",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetAdmin()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "Admin",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetOverloadManager()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "OverloadManager",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for EnableDispatcherStats

	// no validation rules for HeaderPrefix

	{
		tmp := m.GetStatsServerVersionOverride()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return BootstrapValidationError{
					field:  "StatsServerVersionOverride",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	// no validation rules for UseTcpForDnsLookups

	return nil
}

// BootstrapValidationError is the validation error returned by
// Bootstrap.Validate if the designated constraints aren't met.
type BootstrapValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BootstrapValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BootstrapValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BootstrapValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BootstrapValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BootstrapValidationError) ErrorName() string { return "BootstrapValidationError" }

// Error satisfies the builtin error interface
func (e BootstrapValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBootstrap.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BootstrapValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BootstrapValidationError{}

// Validate checks the field values on Admin with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Admin) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for AccessLogPath

	// no validation rules for ProfilePath

	{
		tmp := m.GetAddress()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return AdminValidationError{
					field:  "Address",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	for idx, item := range m.GetSocketOptions() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return AdminValidationError{
						field:  fmt.Sprintf("SocketOptions[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// AdminValidationError is the validation error returned by Admin.Validate if
// the designated constraints aren't met.
type AdminValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e AdminValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e AdminValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e AdminValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e AdminValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e AdminValidationError) ErrorName() string { return "AdminValidationError" }

// Error satisfies the builtin error interface
func (e AdminValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sAdmin.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = AdminValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = AdminValidationError{}

// Validate checks the field values on ClusterManager with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ClusterManager) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for LocalClusterName

	{
		tmp := m.GetOutlierDetection()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ClusterManagerValidationError{
					field:  "OutlierDetection",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetUpstreamBindConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ClusterManagerValidationError{
					field:  "UpstreamBindConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetLoadStatsConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ClusterManagerValidationError{
					field:  "LoadStatsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ClusterManagerValidationError is the validation error returned by
// ClusterManager.Validate if the designated constraints aren't met.
type ClusterManagerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterManagerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterManagerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterManagerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterManagerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterManagerValidationError) ErrorName() string { return "ClusterManagerValidationError" }

// Error satisfies the builtin error interface
func (e ClusterManagerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterManager.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterManagerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterManagerValidationError{}

// Validate checks the field values on Watchdog with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Watchdog) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetMissTimeout()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return WatchdogValidationError{
					field:  "MissTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetMegamissTimeout()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return WatchdogValidationError{
					field:  "MegamissTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetKillTimeout()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return WatchdogValidationError{
					field:  "KillTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetMultikillTimeout()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return WatchdogValidationError{
					field:  "MultikillTimeout",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// WatchdogValidationError is the validation error returned by
// Watchdog.Validate if the designated constraints aren't met.
type WatchdogValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e WatchdogValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e WatchdogValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e WatchdogValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e WatchdogValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e WatchdogValidationError) ErrorName() string { return "WatchdogValidationError" }

// Error satisfies the builtin error interface
func (e WatchdogValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sWatchdog.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = WatchdogValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = WatchdogValidationError{}

// Validate checks the field values on Runtime with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Runtime) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SymlinkRoot

	// no validation rules for Subdirectory

	// no validation rules for OverrideSubdirectory

	{
		tmp := m.GetBase()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RuntimeValidationError{
					field:  "Base",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// RuntimeValidationError is the validation error returned by Runtime.Validate
// if the designated constraints aren't met.
type RuntimeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuntimeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuntimeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuntimeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuntimeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuntimeValidationError) ErrorName() string { return "RuntimeValidationError" }

// Error satisfies the builtin error interface
func (e RuntimeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntime.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuntimeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuntimeValidationError{}

// Validate checks the field values on RuntimeLayer with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *RuntimeLayer) Validate() error {
	if m == nil {
		return nil
	}

	if len(m.GetName()) < 1 {
		return RuntimeLayerValidationError{
			field:  "Name",
			reason: "value length must be at least 1 bytes",
		}
	}

	switch m.LayerSpecifier.(type) {

	case *RuntimeLayer_StaticLayer:

		{
			tmp := m.GetStaticLayer()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return RuntimeLayerValidationError{
						field:  "StaticLayer",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *RuntimeLayer_DiskLayer_:

		{
			tmp := m.GetDiskLayer()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return RuntimeLayerValidationError{
						field:  "DiskLayer",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *RuntimeLayer_AdminLayer_:

		{
			tmp := m.GetAdminLayer()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return RuntimeLayerValidationError{
						field:  "AdminLayer",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	case *RuntimeLayer_RtdsLayer_:

		{
			tmp := m.GetRtdsLayer()

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return RuntimeLayerValidationError{
						field:  "RtdsLayer",
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	default:
		return RuntimeLayerValidationError{
			field:  "LayerSpecifier",
			reason: "value is required",
		}

	}

	return nil
}

// RuntimeLayerValidationError is the validation error returned by
// RuntimeLayer.Validate if the designated constraints aren't met.
type RuntimeLayerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuntimeLayerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuntimeLayerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuntimeLayerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuntimeLayerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuntimeLayerValidationError) ErrorName() string { return "RuntimeLayerValidationError" }

// Error satisfies the builtin error interface
func (e RuntimeLayerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntimeLayer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuntimeLayerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuntimeLayerValidationError{}

// Validate checks the field values on LayeredRuntime with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *LayeredRuntime) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetLayers() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return LayeredRuntimeValidationError{
						field:  fmt.Sprintf("Layers[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// LayeredRuntimeValidationError is the validation error returned by
// LayeredRuntime.Validate if the designated constraints aren't met.
type LayeredRuntimeValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LayeredRuntimeValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LayeredRuntimeValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LayeredRuntimeValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LayeredRuntimeValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LayeredRuntimeValidationError) ErrorName() string { return "LayeredRuntimeValidationError" }

// Error satisfies the builtin error interface
func (e LayeredRuntimeValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLayeredRuntime.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LayeredRuntimeValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LayeredRuntimeValidationError{}

// Validate checks the field values on Bootstrap_StaticResources with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Bootstrap_StaticResources) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetListeners() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return Bootstrap_StaticResourcesValidationError{
						field:  fmt.Sprintf("Listeners[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	for idx, item := range m.GetClusters() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return Bootstrap_StaticResourcesValidationError{
						field:  fmt.Sprintf("Clusters[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	for idx, item := range m.GetSecrets() {
		_, _ = idx, item

		{
			tmp := item

			if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

				if err := v.Validate(); err != nil {
					return Bootstrap_StaticResourcesValidationError{
						field:  fmt.Sprintf("Secrets[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					}
				}
			}
		}

	}

	return nil
}

// Bootstrap_StaticResourcesValidationError is the validation error returned by
// Bootstrap_StaticResources.Validate if the designated constraints aren't met.
type Bootstrap_StaticResourcesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Bootstrap_StaticResourcesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Bootstrap_StaticResourcesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Bootstrap_StaticResourcesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Bootstrap_StaticResourcesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Bootstrap_StaticResourcesValidationError) ErrorName() string {
	return "Bootstrap_StaticResourcesValidationError"
}

// Error satisfies the builtin error interface
func (e Bootstrap_StaticResourcesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBootstrap_StaticResources.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Bootstrap_StaticResourcesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Bootstrap_StaticResourcesValidationError{}

// Validate checks the field values on Bootstrap_DynamicResources with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *Bootstrap_DynamicResources) Validate() error {
	if m == nil {
		return nil
	}

	{
		tmp := m.GetLdsConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return Bootstrap_DynamicResourcesValidationError{
					field:  "LdsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetCdsConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return Bootstrap_DynamicResourcesValidationError{
					field:  "CdsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	{
		tmp := m.GetAdsConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return Bootstrap_DynamicResourcesValidationError{
					field:  "AdsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// Bootstrap_DynamicResourcesValidationError is the validation error returned
// by Bootstrap_DynamicResources.Validate if the designated constraints aren't met.
type Bootstrap_DynamicResourcesValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e Bootstrap_DynamicResourcesValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e Bootstrap_DynamicResourcesValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e Bootstrap_DynamicResourcesValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e Bootstrap_DynamicResourcesValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e Bootstrap_DynamicResourcesValidationError) ErrorName() string {
	return "Bootstrap_DynamicResourcesValidationError"
}

// Error satisfies the builtin error interface
func (e Bootstrap_DynamicResourcesValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBootstrap_DynamicResources.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = Bootstrap_DynamicResourcesValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = Bootstrap_DynamicResourcesValidationError{}

// Validate checks the field values on ClusterManager_OutlierDetection with the
// rules defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *ClusterManager_OutlierDetection) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for EventLogPath

	{
		tmp := m.GetEventService()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return ClusterManager_OutlierDetectionValidationError{
					field:  "EventService",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// ClusterManager_OutlierDetectionValidationError is the validation error
// returned by ClusterManager_OutlierDetection.Validate if the designated
// constraints aren't met.
type ClusterManager_OutlierDetectionValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClusterManager_OutlierDetectionValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClusterManager_OutlierDetectionValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClusterManager_OutlierDetectionValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClusterManager_OutlierDetectionValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClusterManager_OutlierDetectionValidationError) ErrorName() string {
	return "ClusterManager_OutlierDetectionValidationError"
}

// Error satisfies the builtin error interface
func (e ClusterManager_OutlierDetectionValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClusterManager_OutlierDetection.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClusterManager_OutlierDetectionValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClusterManager_OutlierDetectionValidationError{}

// Validate checks the field values on RuntimeLayer_DiskLayer with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RuntimeLayer_DiskLayer) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for SymlinkRoot

	// no validation rules for Subdirectory

	// no validation rules for AppendServiceCluster

	return nil
}

// RuntimeLayer_DiskLayerValidationError is the validation error returned by
// RuntimeLayer_DiskLayer.Validate if the designated constraints aren't met.
type RuntimeLayer_DiskLayerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuntimeLayer_DiskLayerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuntimeLayer_DiskLayerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuntimeLayer_DiskLayerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuntimeLayer_DiskLayerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuntimeLayer_DiskLayerValidationError) ErrorName() string {
	return "RuntimeLayer_DiskLayerValidationError"
}

// Error satisfies the builtin error interface
func (e RuntimeLayer_DiskLayerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntimeLayer_DiskLayer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuntimeLayer_DiskLayerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuntimeLayer_DiskLayerValidationError{}

// Validate checks the field values on RuntimeLayer_AdminLayer with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RuntimeLayer_AdminLayer) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// RuntimeLayer_AdminLayerValidationError is the validation error returned by
// RuntimeLayer_AdminLayer.Validate if the designated constraints aren't met.
type RuntimeLayer_AdminLayerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuntimeLayer_AdminLayerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuntimeLayer_AdminLayerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuntimeLayer_AdminLayerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuntimeLayer_AdminLayerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuntimeLayer_AdminLayerValidationError) ErrorName() string {
	return "RuntimeLayer_AdminLayerValidationError"
}

// Error satisfies the builtin error interface
func (e RuntimeLayer_AdminLayerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntimeLayer_AdminLayer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuntimeLayer_AdminLayerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuntimeLayer_AdminLayerValidationError{}

// Validate checks the field values on RuntimeLayer_RtdsLayer with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *RuntimeLayer_RtdsLayer) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Name

	{
		tmp := m.GetRtdsConfig()

		if v, ok := interface{}(tmp).(interface{ Validate() error }); ok {

			if err := v.Validate(); err != nil {
				return RuntimeLayer_RtdsLayerValidationError{
					field:  "RtdsConfig",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}
	}

	return nil
}

// RuntimeLayer_RtdsLayerValidationError is the validation error returned by
// RuntimeLayer_RtdsLayer.Validate if the designated constraints aren't met.
type RuntimeLayer_RtdsLayerValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RuntimeLayer_RtdsLayerValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RuntimeLayer_RtdsLayerValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RuntimeLayer_RtdsLayerValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RuntimeLayer_RtdsLayerValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RuntimeLayer_RtdsLayerValidationError) ErrorName() string {
	return "RuntimeLayer_RtdsLayerValidationError"
}

// Error satisfies the builtin error interface
func (e RuntimeLayer_RtdsLayerValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRuntimeLayer_RtdsLayer.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RuntimeLayer_RtdsLayerValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RuntimeLayer_RtdsLayerValidationError{}
