// Package sdk provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.4.1 DO NOT EDIT.
package sdk

import (
	"time"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// Defines values for AppLog.
const (
	Kafka AppLog = "kafka"
	None  AppLog = "none"
)

// Defines values for TemplateParamDataType.
const (
	TemplateParamDataTypeDate   TemplateParamDataType = "date"
	TemplateParamDataTypeNumber TemplateParamDataType = "number"
	TemplateParamDataTypeSecret TemplateParamDataType = "secret"
	TemplateParamDataTypeString TemplateParamDataType = "string"
	TemplateParamDataTypeTime   TemplateParamDataType = "time"
)

// App defines model for app.
type App struct {
	// ApiType Wasm API type
	ApiType *string `json:"api_type,omitempty"`

	// Binary Binary ID
	Binary *int64 `json:"binary,omitempty"`

	// Comment Description of the binary
	Comment *string `json:"comment,omitempty"`

	// Debug Switch on logging for 30 minutes (switched off by default)
	Debug *bool `json:"debug,omitempty"`

	// DebugUntil When debugging finishes
	DebugUntil *time.Time `json:"debug_until,omitempty"`

	// Env Environment variables
	Env *map[string]string `json:"env,omitempty"`

	// Log Logging channel (by default - kafka, which allows exploring logs with API)
	Log *AppLog `json:"log"`

	// Name App name
	Name *string `json:"name,omitempty"`

	// Networks Networks
	Networks *[]string `json:"networks,omitempty"`

	// Plan Plan name
	Plan *string `json:"plan,omitempty"`

	// RspHeaders Extra headers to add to the response
	RspHeaders *map[string]string `json:"rsp_headers,omitempty"`

	// Secrets Application secrets
	Secrets *map[string]string `json:"secrets,omitempty"`

	// Status Status code:<br>0 - draft (inactive)<br>1 - enabled<br>2 - disabled<br>3 - hourly call limit exceeded<br>4 - daily call limit exceeded<br>5 - suspended
	Status *int `json:"status,omitempty"`

	// Template Template ID
	Template *int64 `json:"template,omitempty"`

	// Url App URL
	Url *string `json:"url,omitempty"`
}

// AppLog Logging channel (by default - kafka, which allows exploring logs with API)
type AppLog string

// AppShort defines model for app_short.
type AppShort struct {
	// ApiType Wasm API type
	ApiType string `json:"api_type"`

	// Binary Binary ID
	Binary int64 `json:"binary"`

	// Comment Description of the binary
	Comment *string `json:"comment,omitempty"`

	// Debug Switch on logging for 30 minutes (switched off by default)
	Debug *bool `json:"debug,omitempty"`

	// DebugUntil When debugging finishes
	DebugUntil *time.Time `json:"debug_until,omitempty"`

	// Id App ID
	Id int64 `json:"id"`

	// Name App name
	Name string `json:"name"`

	// Networks Networks
	Networks *[]string `json:"networks,omitempty"`

	// Plan Application plan
	Plan string `json:"plan"`

	// Status Status code:<br>0 - draft (inactive)<br>1 - enabled<br>2 - disabled<br>3 - hourly call limit exceeded<br>4 - daily call limit exceeded<br>5 - suspended
	Status int `json:"status"`

	// Template Template ID
	Template *int64 `json:"template,omitempty"`

	// UpgradeableTo ID of the binary the app can be upgraded to
	UpgradeableTo *int64 `json:"upgradeable_to,omitempty"`

	// Url App URL
	Url *string `json:"url,omitempty"`
}

// AppShortAdmin defines model for app_short_admin.
type AppShortAdmin struct {
	// Embedded struct due to allOf(#/components/schemas/app_short)
	AppShort `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// ClientId Client ID
	ClientId int64 `json:"client_id"`

	// ResellerId Reseller ID
	ResellerId *int64 `json:"reseller_id"`
}

// Binary defines model for binary.
type Binary struct {
	// ApiType Wasm API type
	ApiType string `json:"api_type"`

	// Errors Compilation errors
	Errors *string `json:"errors,omitempty"`

	// Id Binary ID
	Id int64 `json:"id"`

	// Source Source language:<br>0 - unknown<br>1 - Rust<br>2 - JavaScript
	Source int `json:"source"`

	// Status Status code:<br>0 - pending<br>1 - compiled<br>2 - compilation failed (errors available)<br>3 - compilation failed (errors not available)<br>4 - resulting binary exceeded the limit<br>5 - unsupported source language
	Status int `json:"status"`

	// UnrefSince Not used since (UTC)
	UnrefSince *string `json:"unref_since,omitempty"`
}

// BinaryShort defines model for binary_short.
type BinaryShort struct {
	// ApiType Wasm API type
	ApiType string `json:"api_type"`

	// Id Binary ID
	Id int64 `json:"id"`

	// Status Status code:<br>0 - pending<br>1 - compiled<br>2 - compilation failed (errors available)<br>3 - compilation failed (errors not available)<br>4 - resulting binary exceeded the limit<br>5 - unsupported source language
	Status int `json:"status"`

	// UnrefSince Not used since (UTC)
	UnrefSince *string `json:"unref_since,omitempty"`
}

// CallStats Edge app call statistics
type CallStats struct {
	// CountByStatus Count by status
	CountByStatus []CountByStatus `json:"count_by_status"`

	// Time Beginning ot reporting slot
	Time time.Time `json:"time"`
}

// Client defines model for client.
type Client struct {
	// AppCount Actual allowed number of apps
	AppCount int `json:"app_count"`

	// AppLimit Max allowed number of apps
	AppLimit int `json:"app_limit"`

	// DailyConsumption Actual number of calls for all apps during the current day (UTC)
	DailyConsumption int `json:"daily_consumption"`

	// DailyLimit Max allowed calls for all apps during a day (UTC)
	DailyLimit int `json:"daily_limit"`

	// HourlyConsumption Actual number of calls for all apps during the current hour
	HourlyConsumption int `json:"hourly_consumption"`

	// HourlyLimit Max allowed calls for all apps during an hour
	HourlyLimit int `json:"hourly_limit"`

	// MonthlyConsumption Actual number of calls for all apps during the current calendar month (UTC)
	MonthlyConsumption int `json:"monthly_consumption"`

	// Networks List of enabled networks
	Networks []Network `json:"networks"`

	// Plan Plan name
	Plan string `json:"plan"`

	// Status Status code:<br>1 - enabled<br>2 - disabled<br>3 - hourly call limit exceeded<br>4 - daily call limit exceeded<br>5 - suspended
	Status int `json:"status"`
}

// ClientUpdate defines model for clientUpdate.
type ClientUpdate struct {
	// Networks List of available networks
	Networks []string `json:"networks"`

	// Plan Plan name
	Plan string `json:"plan"`
}

// CountByStatus defines model for count_by_status.
type CountByStatus struct {
	// Count Number of app calls
	Count int `json:"count"`

	// Status HTTP status
	Status int `json:"status"`
}

// DurationStats Edge app execution duration statistics
type DurationStats struct {
	// Avg Average duration in usec
	Avg int64 `json:"avg"`

	// Max Max duration in usec
	Max int64 `json:"max"`

	// Median Median (50% percentile) duration in usec
	Median int64 `json:"median"`

	// Min Min duration in usec
	Min int64 `json:"min"`

	// Perc75 75% percentile duration in usec
	Perc75 int64 `json:"perc75"`

	// Perc90 90% percentile duration in usec
	Perc90 int64 `json:"perc90"`

	// Time Beginning ot reporting slot
	Time time.Time `json:"time"`
}

// Error defines model for error.
type Error struct {
	// Error Error message
	Error string `json:"error"`
}

// Group defines model for group.
type Group struct {
	Clients []GroupMember `json:"clients"`

	// DefaultFlag Add new members to the group by default
	DefaultFlag bool `json:"default_flag"`

	// Name Group name
	Name      string               `json:"name"`
	Templates []TemplateShortAdmin `json:"templates"`
}

// GroupMember defines model for group_member.
type GroupMember struct {
	// Id Client ID
	Id int64 `json:"id"`
}

// GroupShort defines model for group_short.
type GroupShort struct {
	// DefaultFlag Add new members to the group by default
	DefaultFlag bool `json:"default_flag"`

	// Id Group ID
	Id int64 `json:"id"`

	// Name Group name
	Name string `json:"name"`
}

// Log defines model for log.
type Log struct {
	// AppName Name of the application
	AppName *string `json:"app_name,omitempty"`

	// ClientIp Client IP
	ClientIp *string `json:"client_ip,omitempty"`

	// Edge Edge name
	Edge *string `json:"edge,omitempty"`

	// Id Id of the log
	Id *string `json:"id,omitempty"`

	// Log Log message
	Log *string `json:"log,omitempty"`

	// Timestamp Timestamp of a log in RFC3339 format
	Timestamp *time.Time `json:"timestamp,omitempty"`
}

// Network defines model for network.
type Network struct {
	// IsDefault Is network is default
	IsDefault bool `json:"is_default"`

	// Name Network name
	Name string `json:"name"`
}

// Plan defines model for plan.
type Plan struct {
	// BillingPlanId Billing plan ID
	BillingPlanId *int `json:"billing_plan_id,omitempty"`

	// MaxDuration Max duration in msec
	MaxDuration int `json:"max_duration"`

	// MaxSubrequests Max number of external network requests (0 means disabled)
	MaxSubrequests int `json:"max_subrequests"`

	// MemLimit Max memory in bytes
	MemLimit int `json:"mem_limit"`

	// ResellerId Reseller plan ID
	ResellerId int64 `json:"reseller_id"`
}

// PrebillingConsumptionData defines model for prebillingConsumptionData.
type PrebillingConsumptionData struct {
	ClientId    *int                `json:"client_id,omitempty"`
	Date        *openapi_types.Date `json:"date,omitempty"`
	MetricName  *string             `json:"metric_name,omitempty"`
	MetricUnit  *string             `json:"metric_unit,omitempty"`
	MetricValue *int                `json:"metric_value,omitempty"`
	Subproduct  *string             `json:"subproduct,omitempty"`
}

// PrebillingConsumptionPagination defines model for prebillingConsumptionPagination.
type PrebillingConsumptionPagination struct {
	CurrentPage *int32 `json:"current_page,omitempty"`
	Limit       *int32 `json:"limit,omitempty"`
	TotalPages  *int32 `json:"total_pages,omitempty"`
}

// PrebillingConsumptionResponse defines model for prebillingConsumptionResponse.
type PrebillingConsumptionResponse struct {
	Data       *[]PrebillingConsumptionData     `json:"data,omitempty"`
	Pagination *PrebillingConsumptionPagination `json:"pagination,omitempty"`
}

// PrebillingMetric defines model for prebillingMetric.
type PrebillingMetric struct {
	NumberOfCalls   *int     `json:"number_of_calls,omitempty"`
	TotalRuntimeMms *float32 `json:"total_runtime_mms,omitempty"`
}

// PrebillingObject defines model for prebillingObject.
type PrebillingObject struct {
	BillingPlan *string           `json:"billing_plan,omitempty"`
	Client      *int              `json:"client,omitempty"`
	Metrics     *PrebillingMetric `json:"metrics,omitempty"`
	Region      *string           `json:"region,omitempty"`
}

// PrebillingResponse defines model for prebillingResponse.
type PrebillingResponse = []PrebillingObject

// Secret defines model for secret.
type Secret struct {
	// Comment Description of the secret
	Comment *string `json:"comment"`

	// EffectiveFrom The UNIX timestamp when the secret becomes effective.
	EffectiveFrom *int32 `json:"effective_from"`

	// Name The unique name of the secret.
	Name string `json:"name"`

	// Value The value of the secret.
	Value string `json:"value"`
}

// SecretShort defines model for secret_short.
type SecretShort struct {
	// Name The unique name of the secret.
	Name string `json:"name"`
}

// Template defines model for template.
type Template struct {
	// ApiType Wasm API type
	ApiType *string `json:"api_type,omitempty"`

	// BinaryId Binary ID
	BinaryId int64 `json:"binary_id"`

	// LongDescr Long description of the template
	LongDescr *string `json:"long_descr,omitempty"`

	// Name Name of the template
	Name string `json:"name"`

	// Owned Is the template owned by user?
	Owned bool `json:"owned"`

	// Params Parameters
	Params []TemplateParam `json:"params"`

	// ShortDescr Short description of the template
	ShortDescr *string `json:"short_descr,omitempty"`
}

// TemplateParam defines model for template_param.
type TemplateParam struct {
	// DataType Parameter type
	DataType TemplateParamDataType `json:"data_type"`

	// Descr Parameter description
	Descr *string `json:"descr,omitempty"`

	// Mandatory Is this field mandatory?
	Mandatory bool `json:"mandatory"`

	// Name Parameter name
	Name string `json:"name"`
}

// TemplateParamDataType Parameter type
type TemplateParamDataType string

// TemplateShort defines model for template_short.
type TemplateShort struct {
	// ApiType Wasm API type
	ApiType string `json:"api_type"`

	// Id Template ID
	Id int64 `json:"id"`

	// LongDescr Long description of the template
	LongDescr *string `json:"long_descr,omitempty"`

	// Name Name of the template
	Name string `json:"name"`

	// Owned Is the template owned by user?
	Owned bool `json:"owned"`

	// ShortDescr Short description of the template
	ShortDescr *string `json:"short_descr,omitempty"`
}

// TemplateShortAdmin defines model for template_short_admin.
type TemplateShortAdmin struct {
	// Embedded struct due to allOf(#/components/schemas/template_short)
	TemplateShort `yaml:",inline"`
	// Embedded fields due to inline allOf schema
	// ClientId Client ID
	ClientId int64 `json:"client_id"`

	// Groups Groups the template is shared to
	Groups []int64 `json:"groups"`

	// ResellerId Reseller ID
	ResellerId *int64 `json:"reseller_id"`
}
