// Code generated by config model builder generator; DO NOT EDIT.

package model

import (
	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/config"
	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/provider/resources"
)

type DatabaseModel struct {
	Catalog                                 tfconfig.Variable `json:"catalog,omitempty"`
	Comment                                 tfconfig.Variable `json:"comment,omitempty"`
	DataRetentionTimeInDays                 tfconfig.Variable `json:"data_retention_time_in_days,omitempty"`
	DefaultDdlCollation                     tfconfig.Variable `json:"default_ddl_collation,omitempty"`
	DropPublicSchemaOnCreation              tfconfig.Variable `json:"drop_public_schema_on_creation,omitempty"`
	EnableConsoleOutput                     tfconfig.Variable `json:"enable_console_output,omitempty"`
	ExternalVolume                          tfconfig.Variable `json:"external_volume,omitempty"`
	FullyQualifiedName                      tfconfig.Variable `json:"fully_qualified_name,omitempty"`
	IsTransient                             tfconfig.Variable `json:"is_transient,omitempty"`
	LogLevel                                tfconfig.Variable `json:"log_level,omitempty"`
	MaxDataExtensionTimeInDays              tfconfig.Variable `json:"max_data_extension_time_in_days,omitempty"`
	Name                                    tfconfig.Variable `json:"name,omitempty"`
	QuotedIdentifiersIgnoreCase             tfconfig.Variable `json:"quoted_identifiers_ignore_case,omitempty"`
	ReplaceInvalidCharacters                tfconfig.Variable `json:"replace_invalid_characters,omitempty"`
	Replication                             tfconfig.Variable `json:"replication,omitempty"`
	StorageSerializationPolicy              tfconfig.Variable `json:"storage_serialization_policy,omitempty"`
	SuspendTaskAfterNumFailures             tfconfig.Variable `json:"suspend_task_after_num_failures,omitempty"`
	TaskAutoRetryAttempts                   tfconfig.Variable `json:"task_auto_retry_attempts,omitempty"`
	TraceLevel                              tfconfig.Variable `json:"trace_level,omitempty"`
	UserTaskManagedInitialWarehouseSize     tfconfig.Variable `json:"user_task_managed_initial_warehouse_size,omitempty"`
	UserTaskMinimumTriggerIntervalInSeconds tfconfig.Variable `json:"user_task_minimum_trigger_interval_in_seconds,omitempty"`
	UserTaskTimeoutMs                       tfconfig.Variable `json:"user_task_timeout_ms,omitempty"`

	*config.ResourceModelMeta
}

/////////////////////////////////////////////////
// Basic builders (resource name and required) //
/////////////////////////////////////////////////

func Database(
	resourceName string,
	name string,
) *DatabaseModel {
	d := &DatabaseModel{ResourceModelMeta: config.Meta(resourceName, resources.Database)}
	d.WithName(name)
	return d
}

func DatabaseWithDefaultMeta(
	name string,
) *DatabaseModel {
	d := &DatabaseModel{ResourceModelMeta: config.DefaultMeta(resources.Database)}
	d.WithName(name)
	return d
}

/////////////////////////////////
// below all the proper values //
/////////////////////////////////

func (d *DatabaseModel) WithCatalog(catalog string) *DatabaseModel {
	d.Catalog = tfconfig.StringVariable(catalog)
	return d
}

func (d *DatabaseModel) WithComment(comment string) *DatabaseModel {
	d.Comment = tfconfig.StringVariable(comment)
	return d
}

func (d *DatabaseModel) WithDataRetentionTimeInDays(dataRetentionTimeInDays int) *DatabaseModel {
	d.DataRetentionTimeInDays = tfconfig.IntegerVariable(dataRetentionTimeInDays)
	return d
}

func (d *DatabaseModel) WithDefaultDdlCollation(defaultDdlCollation string) *DatabaseModel {
	d.DefaultDdlCollation = tfconfig.StringVariable(defaultDdlCollation)
	return d
}

func (d *DatabaseModel) WithDropPublicSchemaOnCreation(dropPublicSchemaOnCreation bool) *DatabaseModel {
	d.DropPublicSchemaOnCreation = tfconfig.BoolVariable(dropPublicSchemaOnCreation)
	return d
}

func (d *DatabaseModel) WithEnableConsoleOutput(enableConsoleOutput bool) *DatabaseModel {
	d.EnableConsoleOutput = tfconfig.BoolVariable(enableConsoleOutput)
	return d
}

func (d *DatabaseModel) WithExternalVolume(externalVolume string) *DatabaseModel {
	d.ExternalVolume = tfconfig.StringVariable(externalVolume)
	return d
}

func (d *DatabaseModel) WithFullyQualifiedName(fullyQualifiedName string) *DatabaseModel {
	d.FullyQualifiedName = tfconfig.StringVariable(fullyQualifiedName)
	return d
}

func (d *DatabaseModel) WithIsTransient(isTransient bool) *DatabaseModel {
	d.IsTransient = tfconfig.BoolVariable(isTransient)
	return d
}

func (d *DatabaseModel) WithLogLevel(logLevel string) *DatabaseModel {
	d.LogLevel = tfconfig.StringVariable(logLevel)
	return d
}

func (d *DatabaseModel) WithMaxDataExtensionTimeInDays(maxDataExtensionTimeInDays int) *DatabaseModel {
	d.MaxDataExtensionTimeInDays = tfconfig.IntegerVariable(maxDataExtensionTimeInDays)
	return d
}

func (d *DatabaseModel) WithName(name string) *DatabaseModel {
	d.Name = tfconfig.StringVariable(name)
	return d
}

func (d *DatabaseModel) WithQuotedIdentifiersIgnoreCase(quotedIdentifiersIgnoreCase bool) *DatabaseModel {
	d.QuotedIdentifiersIgnoreCase = tfconfig.BoolVariable(quotedIdentifiersIgnoreCase)
	return d
}

func (d *DatabaseModel) WithReplaceInvalidCharacters(replaceInvalidCharacters bool) *DatabaseModel {
	d.ReplaceInvalidCharacters = tfconfig.BoolVariable(replaceInvalidCharacters)
	return d
}

// replication attribute type is not yet supported, so WithReplication can't be generated

func (d *DatabaseModel) WithStorageSerializationPolicy(storageSerializationPolicy string) *DatabaseModel {
	d.StorageSerializationPolicy = tfconfig.StringVariable(storageSerializationPolicy)
	return d
}

func (d *DatabaseModel) WithSuspendTaskAfterNumFailures(suspendTaskAfterNumFailures int) *DatabaseModel {
	d.SuspendTaskAfterNumFailures = tfconfig.IntegerVariable(suspendTaskAfterNumFailures)
	return d
}

func (d *DatabaseModel) WithTaskAutoRetryAttempts(taskAutoRetryAttempts int) *DatabaseModel {
	d.TaskAutoRetryAttempts = tfconfig.IntegerVariable(taskAutoRetryAttempts)
	return d
}

func (d *DatabaseModel) WithTraceLevel(traceLevel string) *DatabaseModel {
	d.TraceLevel = tfconfig.StringVariable(traceLevel)
	return d
}

func (d *DatabaseModel) WithUserTaskManagedInitialWarehouseSize(userTaskManagedInitialWarehouseSize string) *DatabaseModel {
	d.UserTaskManagedInitialWarehouseSize = tfconfig.StringVariable(userTaskManagedInitialWarehouseSize)
	return d
}

func (d *DatabaseModel) WithUserTaskMinimumTriggerIntervalInSeconds(userTaskMinimumTriggerIntervalInSeconds int) *DatabaseModel {
	d.UserTaskMinimumTriggerIntervalInSeconds = tfconfig.IntegerVariable(userTaskMinimumTriggerIntervalInSeconds)
	return d
}

func (d *DatabaseModel) WithUserTaskTimeoutMs(userTaskTimeoutMs int) *DatabaseModel {
	d.UserTaskTimeoutMs = tfconfig.IntegerVariable(userTaskTimeoutMs)
	return d
}

//////////////////////////////////////////
// below it's possible to set any value //
//////////////////////////////////////////

func (d *DatabaseModel) WithCatalogValue(value tfconfig.Variable) *DatabaseModel {
	d.Catalog = value
	return d
}

func (d *DatabaseModel) WithCommentValue(value tfconfig.Variable) *DatabaseModel {
	d.Comment = value
	return d
}

func (d *DatabaseModel) WithDataRetentionTimeInDaysValue(value tfconfig.Variable) *DatabaseModel {
	d.DataRetentionTimeInDays = value
	return d
}

func (d *DatabaseModel) WithDefaultDdlCollationValue(value tfconfig.Variable) *DatabaseModel {
	d.DefaultDdlCollation = value
	return d
}

func (d *DatabaseModel) WithDropPublicSchemaOnCreationValue(value tfconfig.Variable) *DatabaseModel {
	d.DropPublicSchemaOnCreation = value
	return d
}

func (d *DatabaseModel) WithEnableConsoleOutputValue(value tfconfig.Variable) *DatabaseModel {
	d.EnableConsoleOutput = value
	return d
}

func (d *DatabaseModel) WithExternalVolumeValue(value tfconfig.Variable) *DatabaseModel {
	d.ExternalVolume = value
	return d
}

func (d *DatabaseModel) WithFullyQualifiedNameValue(value tfconfig.Variable) *DatabaseModel {
	d.FullyQualifiedName = value
	return d
}

func (d *DatabaseModel) WithIsTransientValue(value tfconfig.Variable) *DatabaseModel {
	d.IsTransient = value
	return d
}

func (d *DatabaseModel) WithLogLevelValue(value tfconfig.Variable) *DatabaseModel {
	d.LogLevel = value
	return d
}

func (d *DatabaseModel) WithMaxDataExtensionTimeInDaysValue(value tfconfig.Variable) *DatabaseModel {
	d.MaxDataExtensionTimeInDays = value
	return d
}

func (d *DatabaseModel) WithNameValue(value tfconfig.Variable) *DatabaseModel {
	d.Name = value
	return d
}

func (d *DatabaseModel) WithQuotedIdentifiersIgnoreCaseValue(value tfconfig.Variable) *DatabaseModel {
	d.QuotedIdentifiersIgnoreCase = value
	return d
}

func (d *DatabaseModel) WithReplaceInvalidCharactersValue(value tfconfig.Variable) *DatabaseModel {
	d.ReplaceInvalidCharacters = value
	return d
}

func (d *DatabaseModel) WithReplicationValue(value tfconfig.Variable) *DatabaseModel {
	d.Replication = value
	return d
}

func (d *DatabaseModel) WithStorageSerializationPolicyValue(value tfconfig.Variable) *DatabaseModel {
	d.StorageSerializationPolicy = value
	return d
}

func (d *DatabaseModel) WithSuspendTaskAfterNumFailuresValue(value tfconfig.Variable) *DatabaseModel {
	d.SuspendTaskAfterNumFailures = value
	return d
}

func (d *DatabaseModel) WithTaskAutoRetryAttemptsValue(value tfconfig.Variable) *DatabaseModel {
	d.TaskAutoRetryAttempts = value
	return d
}

func (d *DatabaseModel) WithTraceLevelValue(value tfconfig.Variable) *DatabaseModel {
	d.TraceLevel = value
	return d
}

func (d *DatabaseModel) WithUserTaskManagedInitialWarehouseSizeValue(value tfconfig.Variable) *DatabaseModel {
	d.UserTaskManagedInitialWarehouseSize = value
	return d
}

func (d *DatabaseModel) WithUserTaskMinimumTriggerIntervalInSecondsValue(value tfconfig.Variable) *DatabaseModel {
	d.UserTaskMinimumTriggerIntervalInSeconds = value
	return d
}

func (d *DatabaseModel) WithUserTaskTimeoutMsValue(value tfconfig.Variable) *DatabaseModel {
	d.UserTaskTimeoutMs = value
	return d
}