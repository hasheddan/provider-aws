/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type ApplicationCodeConfigurationObservation struct {
}

type ApplicationCodeConfigurationParameters struct {

	// The location and type of the application code.
	// +kubebuilder:validation:Optional
	CodeContent []CodeContentParameters `json:"codeContent,omitempty" tf:"code_content,omitempty"`

	// Specifies whether the code content is in text or zip format. Valid values: PLAINTEXT, ZIPFILE.
	// +kubebuilder:validation:Required
	CodeContentType *string `json:"codeContentType" tf:"code_content_type,omitempty"`
}

type ApplicationConfigurationObservation struct {

	// The configuration of a SQL-based application.
	// +kubebuilder:validation:Optional
	SQLApplicationConfiguration []SQLApplicationConfigurationObservation `json:"sqlApplicationConfiguration,omitempty" tf:"sql_application_configuration,omitempty"`

	// The VPC configuration of a Flink-based application.
	// +kubebuilder:validation:Optional
	VPCConfiguration []VPCConfigurationObservation `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type ApplicationConfigurationParameters struct {

	// The code location and type parameters for the application.
	// +kubebuilder:validation:Required
	ApplicationCodeConfiguration []ApplicationCodeConfigurationParameters `json:"applicationCodeConfiguration" tf:"application_code_configuration,omitempty"`

	// Describes whether snapshots are enabled for a Flink-based application.
	// +kubebuilder:validation:Optional
	ApplicationSnapshotConfiguration []ApplicationSnapshotConfigurationParameters `json:"applicationSnapshotConfiguration,omitempty" tf:"application_snapshot_configuration,omitempty"`

	// Describes execution properties for a Flink-based application.
	// +kubebuilder:validation:Optional
	EnvironmentProperties []EnvironmentPropertiesParameters `json:"environmentProperties,omitempty" tf:"environment_properties,omitempty"`

	// The configuration of a Flink-based application.
	// +kubebuilder:validation:Optional
	FlinkApplicationConfiguration []FlinkApplicationConfigurationParameters `json:"flinkApplicationConfiguration,omitempty" tf:"flink_application_configuration,omitempty"`

	// Describes the starting properties for a Flink-based application.
	// +kubebuilder:validation:Optional
	RunConfiguration []RunConfigurationParameters `json:"runConfiguration,omitempty" tf:"run_configuration,omitempty"`

	// The configuration of a SQL-based application.
	// +kubebuilder:validation:Optional
	SQLApplicationConfiguration []SQLApplicationConfigurationParameters `json:"sqlApplicationConfiguration,omitempty" tf:"sql_application_configuration,omitempty"`

	// The VPC configuration of a Flink-based application.
	// +kubebuilder:validation:Optional
	VPCConfiguration []VPCConfigurationParameters `json:"vpcConfiguration,omitempty" tf:"vpc_configuration,omitempty"`
}

type ApplicationObservation struct {

	// The application's configuration
	// +kubebuilder:validation:Optional
	ApplicationConfiguration []ApplicationConfigurationObservation `json:"applicationConfiguration,omitempty" tf:"application_configuration,omitempty"`

	// The ARN of the application.
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	// A CloudWatch log stream to monitor application configuration errors.
	// +kubebuilder:validation:Optional
	CloudwatchLoggingOptions []CloudwatchLoggingOptionsObservation `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`

	// The current timestamp when the application was created.
	CreateTimestamp *string `json:"createTimestamp,omitempty" tf:"create_timestamp,omitempty"`

	// The application identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// The current timestamp when the application was last updated.
	LastUpdateTimestamp *string `json:"lastUpdateTimestamp,omitempty" tf:"last_update_timestamp,omitempty"`

	// The status of the application.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// A map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// The current application version. Kinesis Data Analytics updates the version_id each time the application is updated.
	VersionID *float64 `json:"versionId,omitempty" tf:"version_id,omitempty"`
}

type ApplicationParameters struct {

	// The application's configuration
	// +kubebuilder:validation:Optional
	ApplicationConfiguration []ApplicationConfigurationParameters `json:"applicationConfiguration,omitempty" tf:"application_configuration,omitempty"`

	// A CloudWatch log stream to monitor application configuration errors.
	// +kubebuilder:validation:Optional
	CloudwatchLoggingOptions []CloudwatchLoggingOptionsParameters `json:"cloudwatchLoggingOptions,omitempty" tf:"cloudwatch_logging_options,omitempty"`

	// A summary description of the application.
	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// Whether to force stop an unresponsive Flink-based application.
	// +kubebuilder:validation:Optional
	ForceStop *bool `json:"forceStop,omitempty" tf:"force_stop,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// The runtime environment for the application. Valid values: SQL-1_0, FLINK-1_6, FLINK-1_8, FLINK-1_11, FLINK-1_13.
	// +kubebuilder:validation:Required
	RuntimeEnvironment *string `json:"runtimeEnvironment" tf:"runtime_environment,omitempty"`

	// The ARN of the IAM role used by the application to access Kinesis data streams, Kinesis Data Firehose delivery streams, Amazon S3 objects, and other external resources.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ServiceExecutionRole *string `json:"serviceExecutionRole,omitempty" tf:"service_execution_role,omitempty"`

	// Reference to a Role in iam to populate serviceExecutionRole.
	// +kubebuilder:validation:Optional
	ServiceExecutionRoleRef *v1.Reference `json:"serviceExecutionRoleRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate serviceExecutionRole.
	// +kubebuilder:validation:Optional
	ServiceExecutionRoleSelector *v1.Selector `json:"serviceExecutionRoleSelector,omitempty" tf:"-"`

	// Whether to start or stop the application.
	// +kubebuilder:validation:Optional
	StartApplication *bool `json:"startApplication,omitempty" tf:"start_application,omitempty"`

	// A map of tags to assign to the application. If configured with a provider default_tags configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type ApplicationRestoreConfigurationObservation struct {
}

type ApplicationRestoreConfigurationParameters struct {

	// Specifies how the application should be restored. Valid values: RESTORE_FROM_CUSTOM_SNAPSHOT, RESTORE_FROM_LATEST_SNAPSHOT, SKIP_RESTORE_FROM_SNAPSHOT.
	// +kubebuilder:validation:Optional
	ApplicationRestoreType *string `json:"applicationRestoreType,omitempty" tf:"application_restore_type,omitempty"`

	// The identifier of an existing snapshot of application state to use to restart an application. The application uses this value if RESTORE_FROM_CUSTOM_SNAPSHOT is specified for application_restore_type.
	// +kubebuilder:validation:Optional
	SnapshotName *string `json:"snapshotName,omitempty" tf:"snapshot_name,omitempty"`
}

type ApplicationSnapshotConfigurationObservation struct {
}

type ApplicationSnapshotConfigurationParameters struct {

	// Describes whether snapshots are enabled for a Flink-based Kinesis Data Analytics application.
	// +kubebuilder:validation:Required
	SnapshotsEnabled *bool `json:"snapshotsEnabled" tf:"snapshots_enabled,omitempty"`
}

type CheckpointConfigurationObservation struct {
}

type CheckpointConfigurationParameters struct {

	// Describes the interval in milliseconds between checkpoint operations.
	// +kubebuilder:validation:Optional
	CheckpointInterval *float64 `json:"checkpointInterval,omitempty" tf:"checkpoint_interval,omitempty"`

	// Describes whether checkpointing is enabled for a Flink-based Kinesis Data Analytics application.
	// +kubebuilder:validation:Optional
	CheckpointingEnabled *bool `json:"checkpointingEnabled,omitempty" tf:"checkpointing_enabled,omitempty"`

	// Describes whether the application uses Kinesis Data Analytics' default checkpointing behavior. Valid values: CUSTOM, DEFAULT. Set this attribute to CUSTOM in order for any specified checkpointing_enabled, checkpoint_interval, or min_pause_between_checkpoints attribute values to be effective. If this attribute is set to DEFAULT, the application will always use the following values:
	// +kubebuilder:validation:Required
	ConfigurationType *string `json:"configurationType" tf:"configuration_type,omitempty"`

	// Describes the minimum time in milliseconds after a checkpoint operation completes that a new checkpoint operation can start.
	// +kubebuilder:validation:Optional
	MinPauseBetweenCheckpoints *float64 `json:"minPauseBetweenCheckpoints,omitempty" tf:"min_pause_between_checkpoints,omitempty"`
}

type CloudwatchLoggingOptionsObservation struct {

	// The application identifier.
	CloudwatchLoggingOptionID *string `json:"cloudwatchLoggingOptionId,omitempty" tf:"cloudwatch_logging_option_id,omitempty"`
}

type CloudwatchLoggingOptionsParameters struct {

	// The ARN of the CloudWatch log stream to receive application messages.
	// +kubebuilder:validation:Required
	LogStreamArn *string `json:"logStreamArn" tf:"log_stream_arn,omitempty"`
}

type CodeContentObservation struct {
}

type CodeContentParameters struct {

	// Information about the Amazon S3 bucket containing the application code.
	// +kubebuilder:validation:Optional
	S3ContentLocation []S3ContentLocationParameters `json:"s3ContentLocation,omitempty" tf:"s3_content_location,omitempty"`

	// The text-format code for the application.
	// +kubebuilder:validation:Optional
	TextContent *string `json:"textContent,omitempty" tf:"text_content,omitempty"`
}

type CsvMappingParametersObservation struct {
}

type CsvMappingParametersParameters struct {

	// The column delimiter. For example, in a CSV format, a comma (,) is the typical column delimiter.
	// +kubebuilder:validation:Required
	RecordColumnDelimiter *string `json:"recordColumnDelimiter" tf:"record_column_delimiter,omitempty"`

	// The row delimiter. For example, in a CSV format, \n is the typical row delimiter.
	// +kubebuilder:validation:Required
	RecordRowDelimiter *string `json:"recordRowDelimiter" tf:"record_row_delimiter,omitempty"`
}

type DestinationSchemaObservation struct {
}

type DestinationSchemaParameters struct {

	// The type of record format. Valid values: CSV, JSON.
	// +kubebuilder:validation:Required
	RecordFormatType *string `json:"recordFormatType" tf:"record_format_type,omitempty"`
}

type EnvironmentPropertiesObservation struct {
}

type EnvironmentPropertiesParameters struct {

	// Describes the execution property groups.
	// +kubebuilder:validation:Required
	PropertyGroup []PropertyGroupParameters `json:"propertyGroup" tf:"property_group,omitempty"`
}

type FlinkApplicationConfigurationObservation struct {
}

type FlinkApplicationConfigurationParameters struct {

	// Describes an application's checkpointing configuration.
	// +kubebuilder:validation:Optional
	CheckpointConfiguration []CheckpointConfigurationParameters `json:"checkpointConfiguration,omitempty" tf:"checkpoint_configuration,omitempty"`

	// Describes configuration parameters for CloudWatch logging for an application.
	// +kubebuilder:validation:Optional
	MonitoringConfiguration []MonitoringConfigurationParameters `json:"monitoringConfiguration,omitempty" tf:"monitoring_configuration,omitempty"`

	// Describes parameters for how an application executes multiple tasks simultaneously.
	// +kubebuilder:validation:Optional
	ParallelismConfiguration []ParallelismConfigurationParameters `json:"parallelismConfiguration,omitempty" tf:"parallelism_configuration,omitempty"`
}

type FlinkRunConfigurationObservation struct {
}

type FlinkRunConfigurationParameters struct {

	// When restoring from a snapshot, specifies whether the runtime is allowed to skip a state that cannot be mapped to the new program. Default is false.
	// +kubebuilder:validation:Optional
	AllowNonRestoredState *bool `json:"allowNonRestoredState,omitempty" tf:"allow_non_restored_state,omitempty"`
}

type InputLambdaProcessorObservation struct {
}

type InputLambdaProcessorParameters struct {

	// The ARN of the Lambda function that operates on records in the stream.
	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`
}

type InputObservation struct {
	InAppStreamNames []*string `json:"inAppStreamNames,omitempty" tf:"in_app_stream_names,omitempty"`

	// The application identifier.
	InputID *string `json:"inputId,omitempty" tf:"input_id,omitempty"`
}

type InputParallelismObservation struct {
}

type InputParallelismParameters struct {

	// The number of in-application streams to create.
	// +kubebuilder:validation:Optional
	Count *float64 `json:"count,omitempty" tf:"count,omitempty"`
}

type InputParameters struct {

	// Describes the number of in-application streams to create.
	// +kubebuilder:validation:Optional
	InputParallelism []InputParallelismParameters `json:"inputParallelism,omitempty" tf:"input_parallelism,omitempty"`

	// The input processing configuration for the input.
	// An input processor transforms records as they are received from the stream, before the application's SQL code executes.
	// +kubebuilder:validation:Optional
	InputProcessingConfiguration []InputProcessingConfigurationParameters `json:"inputProcessingConfiguration,omitempty" tf:"input_processing_configuration,omitempty"`

	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns in the in-application stream that is being created.
	// +kubebuilder:validation:Required
	InputSchema []InputSchemaParameters `json:"inputSchema" tf:"input_schema,omitempty"`

	// The point at which the application starts processing records from the streaming source.
	// +kubebuilder:validation:Optional
	InputStartingPositionConfiguration []InputStartingPositionConfigurationParameters `json:"inputStartingPositionConfiguration,omitempty" tf:"input_starting_position_configuration,omitempty"`

	// If the streaming source is a Kinesis Data Firehose delivery stream, identifies the delivery stream's ARN.
	// +kubebuilder:validation:Optional
	KinesisFirehoseInput []KinesisFirehoseInputParameters `json:"kinesisFirehoseInput,omitempty" tf:"kinesis_firehose_input,omitempty"`

	// If the streaming source is a Kinesis data stream, identifies the stream's Amazon Resource Name (ARN).
	// +kubebuilder:validation:Optional
	KinesisStreamsInput []KinesisStreamsInputParameters `json:"kinesisStreamsInput,omitempty" tf:"kinesis_streams_input,omitempty"`

	// The name prefix to use when creating an in-application stream.
	// +kubebuilder:validation:Required
	NamePrefix *string `json:"namePrefix" tf:"name_prefix,omitempty"`
}

type InputProcessingConfigurationObservation struct {
}

type InputProcessingConfigurationParameters struct {

	// Describes the Lambda function that is used to preprocess the records in the stream before being processed by your application code.
	// +kubebuilder:validation:Required
	InputLambdaProcessor []InputLambdaProcessorParameters `json:"inputLambdaProcessor" tf:"input_lambda_processor,omitempty"`
}

type InputSchemaObservation struct {
}

type InputSchemaParameters struct {

	// Describes the mapping of each data element in the streaming source to the corresponding column in the in-application stream.
	// +kubebuilder:validation:Required
	RecordColumn []RecordColumnParameters `json:"recordColumn" tf:"record_column,omitempty"`

	// Specifies the encoding of the records in the streaming source. For example, UTF-8.
	// +kubebuilder:validation:Optional
	RecordEncoding *string `json:"recordEncoding,omitempty" tf:"record_encoding,omitempty"`

	// Specifies the format of the records on the streaming source.
	// +kubebuilder:validation:Required
	RecordFormat []RecordFormatParameters `json:"recordFormat" tf:"record_format,omitempty"`
}

type InputStartingPositionConfigurationObservation struct {
}

type InputStartingPositionConfigurationParameters struct {

	// The starting position on the stream. Valid values: LAST_STOPPED_POINT, NOW, TRIM_HORIZON.
	// +kubebuilder:validation:Optional
	InputStartingPosition *string `json:"inputStartingPosition,omitempty" tf:"input_starting_position,omitempty"`
}

type JSONMappingParametersObservation struct {
}

type JSONMappingParametersParameters struct {

	// The path to the top-level parent that contains the records.
	// +kubebuilder:validation:Required
	RecordRowPath *string `json:"recordRowPath" tf:"record_row_path,omitempty"`
}

type KinesisFirehoseInputObservation struct {
}

type KinesisFirehoseInputParameters struct {

	// The ARN of the Lambda function that operates on records in the stream.
	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`
}

type KinesisFirehoseOutputObservation struct {
}

type KinesisFirehoseOutputParameters struct {

	// The ARN of the Lambda function that operates on records in the stream.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/firehose/v1beta1.DeliveryStream
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",false)
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a DeliveryStream in firehose to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a DeliveryStream in firehose to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`
}

type KinesisStreamsInputObservation struct {
}

type KinesisStreamsInputParameters struct {

	// The ARN of the Lambda function that operates on records in the stream.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kinesis/v1beta1.Stream
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",false)
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a Stream in kinesis to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a Stream in kinesis to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`
}

type KinesisStreamsOutputObservation struct {
}

type KinesisStreamsOutputParameters struct {

	// The ARN of the Lambda function that operates on records in the stream.
	// +kubebuilder:validation:Required
	ResourceArn *string `json:"resourceArn" tf:"resource_arn,omitempty"`
}

type LambdaOutputObservation struct {
}

type LambdaOutputParameters struct {

	// The ARN of the Lambda function that operates on records in the stream.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/lambda/v1beta1.Function
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ResourceArn *string `json:"resourceArn,omitempty" tf:"resource_arn,omitempty"`

	// Reference to a Function in lambda to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnRef *v1.Reference `json:"resourceArnRef,omitempty" tf:"-"`

	// Selector for a Function in lambda to populate resourceArn.
	// +kubebuilder:validation:Optional
	ResourceArnSelector *v1.Selector `json:"resourceArnSelector,omitempty" tf:"-"`
}

type MappingParametersCsvMappingParametersObservation struct {
}

type MappingParametersCsvMappingParametersParameters struct {

	// The column delimiter. For example, in a CSV format, a comma (,) is the typical column delimiter.
	// +kubebuilder:validation:Required
	RecordColumnDelimiter *string `json:"recordColumnDelimiter" tf:"record_column_delimiter,omitempty"`

	// The row delimiter. For example, in a CSV format, \n is the typical row delimiter.
	// +kubebuilder:validation:Required
	RecordRowDelimiter *string `json:"recordRowDelimiter" tf:"record_row_delimiter,omitempty"`
}

type MappingParametersJSONMappingParametersObservation struct {
}

type MappingParametersJSONMappingParametersParameters struct {

	// The path to the top-level parent that contains the records.
	// +kubebuilder:validation:Required
	RecordRowPath *string `json:"recordRowPath" tf:"record_row_path,omitempty"`
}

type MappingParametersObservation struct {
}

type MappingParametersParameters struct {

	// Provides additional mapping information when the record format uses delimiters (for example, CSV).
	// +kubebuilder:validation:Optional
	CsvMappingParameters []CsvMappingParametersParameters `json:"csvMappingParameters,omitempty" tf:"csv_mapping_parameters,omitempty"`

	// Provides additional mapping information when JSON is the record format on the streaming source.
	// +kubebuilder:validation:Optional
	JSONMappingParameters []JSONMappingParametersParameters `json:"jsonMappingParameters,omitempty" tf:"json_mapping_parameters,omitempty"`
}

type MonitoringConfigurationObservation struct {
}

type MonitoringConfigurationParameters struct {

	// Describes whether the application uses Kinesis Data Analytics' default checkpointing behavior. Valid values: CUSTOM, DEFAULT. Set this attribute to CUSTOM in order for any specified checkpointing_enabled, checkpoint_interval, or min_pause_between_checkpoints attribute values to be effective. If this attribute is set to DEFAULT, the application will always use the following values:
	// +kubebuilder:validation:Required
	ConfigurationType *string `json:"configurationType" tf:"configuration_type,omitempty"`

	// Describes the verbosity of the CloudWatch Logs for an application. Valid values: DEBUG, ERROR, INFO, WARN.
	// +kubebuilder:validation:Optional
	LogLevel *string `json:"logLevel,omitempty" tf:"log_level,omitempty"`

	// Describes the granularity of the CloudWatch Logs for an application. Valid values: APPLICATION, OPERATOR, PARALLELISM, TASK.
	// +kubebuilder:validation:Optional
	MetricsLevel *string `json:"metricsLevel,omitempty" tf:"metrics_level,omitempty"`
}

type OutputObservation struct {

	// The application identifier.
	OutputID *string `json:"outputId,omitempty" tf:"output_id,omitempty"`
}

type OutputParameters struct {

	// Describes the data format when records are written to the destination.
	// +kubebuilder:validation:Required
	DestinationSchema []DestinationSchemaParameters `json:"destinationSchema" tf:"destination_schema,omitempty"`

	// Identifies a Kinesis Data Firehose delivery stream as the destination.
	// +kubebuilder:validation:Optional
	KinesisFirehoseOutput []KinesisFirehoseOutputParameters `json:"kinesisFirehoseOutput,omitempty" tf:"kinesis_firehose_output,omitempty"`

	// Identifies a Kinesis data stream as the destination.
	// +kubebuilder:validation:Optional
	KinesisStreamsOutput []KinesisStreamsOutputParameters `json:"kinesisStreamsOutput,omitempty" tf:"kinesis_streams_output,omitempty"`

	// Identifies a Lambda function as the destination.
	// +kubebuilder:validation:Optional
	LambdaOutput []LambdaOutputParameters `json:"lambdaOutput,omitempty" tf:"lambda_output,omitempty"`

	// The name of the application.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`
}

type ParallelismConfigurationObservation struct {
}

type ParallelismConfigurationParameters struct {

	// Describes whether the Kinesis Data Analytics service can increase the parallelism of the application in response to increased throughput.
	// +kubebuilder:validation:Optional
	AutoScalingEnabled *bool `json:"autoScalingEnabled,omitempty" tf:"auto_scaling_enabled,omitempty"`

	// Describes whether the application uses Kinesis Data Analytics' default checkpointing behavior. Valid values: CUSTOM, DEFAULT. Set this attribute to CUSTOM in order for any specified checkpointing_enabled, checkpoint_interval, or min_pause_between_checkpoints attribute values to be effective. If this attribute is set to DEFAULT, the application will always use the following values:
	// +kubebuilder:validation:Required
	ConfigurationType *string `json:"configurationType" tf:"configuration_type,omitempty"`

	// Describes the initial number of parallel tasks that a Flink-based Kinesis Data Analytics application can perform.
	// +kubebuilder:validation:Optional
	Parallelism *float64 `json:"parallelism,omitempty" tf:"parallelism,omitempty"`

	// Describes the number of parallel tasks that a Flink-based Kinesis Data Analytics application can perform per Kinesis Processing Unit (KPU) used by the application.
	// +kubebuilder:validation:Optional
	ParallelismPerKpu *float64 `json:"parallelismPerKpu,omitempty" tf:"parallelism_per_kpu,omitempty"`
}

type PropertyGroupObservation struct {
}

type PropertyGroupParameters struct {

	// The key of the application execution property key-value map.
	// +kubebuilder:validation:Required
	PropertyGroupID *string `json:"propertyGroupId" tf:"property_group_id,omitempty"`

	// Application execution property key-value map.
	// +kubebuilder:validation:Required
	PropertyMap map[string]*string `json:"propertyMap" tf:"property_map,omitempty"`
}

type RecordColumnObservation struct {
}

type RecordColumnParameters struct {

	// A reference to the data element in the streaming input or the reference data source.
	// +kubebuilder:validation:Optional
	Mapping *string `json:"mapping,omitempty" tf:"mapping,omitempty"`

	// The name of the application.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The type of column created in the in-application input stream or reference table.
	// +kubebuilder:validation:Required
	SQLType *string `json:"sqlType" tf:"sql_type,omitempty"`
}

type RecordFormatMappingParametersObservation struct {
}

type RecordFormatMappingParametersParameters struct {

	// Provides additional mapping information when the record format uses delimiters (for example, CSV).
	// +kubebuilder:validation:Optional
	CsvMappingParameters []MappingParametersCsvMappingParametersParameters `json:"csvMappingParameters,omitempty" tf:"csv_mapping_parameters,omitempty"`

	// Provides additional mapping information when JSON is the record format on the streaming source.
	// +kubebuilder:validation:Optional
	JSONMappingParameters []MappingParametersJSONMappingParametersParameters `json:"jsonMappingParameters,omitempty" tf:"json_mapping_parameters,omitempty"`
}

type RecordFormatObservation struct {
}

type RecordFormatParameters struct {

	// Provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
	// +kubebuilder:validation:Required
	MappingParameters []MappingParametersParameters `json:"mappingParameters" tf:"mapping_parameters,omitempty"`

	// The type of record format. Valid values: CSV, JSON.
	// +kubebuilder:validation:Required
	RecordFormatType *string `json:"recordFormatType" tf:"record_format_type,omitempty"`
}

type ReferenceDataSourceObservation struct {

	// The application identifier.
	ReferenceID *string `json:"referenceId,omitempty" tf:"reference_id,omitempty"`
}

type ReferenceDataSourceParameters struct {

	// Describes the format of the data in the streaming source, and how each data element maps to corresponding columns created in the in-application stream.
	// +kubebuilder:validation:Required
	ReferenceSchema []ReferenceSchemaParameters `json:"referenceSchema" tf:"reference_schema,omitempty"`

	// Identifies the S3 bucket and object that contains the reference data.
	// +kubebuilder:validation:Required
	S3ReferenceDataSource []S3ReferenceDataSourceParameters `json:"s3ReferenceDataSource" tf:"s3_reference_data_source,omitempty"`

	// The name of the in-application table to create.
	// +kubebuilder:validation:Required
	TableName *string `json:"tableName" tf:"table_name,omitempty"`
}

type ReferenceSchemaObservation struct {
}

type ReferenceSchemaParameters struct {

	// Describes the mapping of each data element in the streaming source to the corresponding column in the in-application stream.
	// +kubebuilder:validation:Required
	RecordColumn []ReferenceSchemaRecordColumnParameters `json:"recordColumn" tf:"record_column,omitempty"`

	// Specifies the encoding of the records in the streaming source. For example, UTF-8.
	// +kubebuilder:validation:Optional
	RecordEncoding *string `json:"recordEncoding,omitempty" tf:"record_encoding,omitempty"`

	// Specifies the format of the records on the streaming source.
	// +kubebuilder:validation:Required
	RecordFormat []ReferenceSchemaRecordFormatParameters `json:"recordFormat" tf:"record_format,omitempty"`
}

type ReferenceSchemaRecordColumnObservation struct {
}

type ReferenceSchemaRecordColumnParameters struct {

	// A reference to the data element in the streaming input or the reference data source.
	// +kubebuilder:validation:Optional
	Mapping *string `json:"mapping,omitempty" tf:"mapping,omitempty"`

	// The name of the application.
	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// The type of column created in the in-application input stream or reference table.
	// +kubebuilder:validation:Required
	SQLType *string `json:"sqlType" tf:"sql_type,omitempty"`
}

type ReferenceSchemaRecordFormatObservation struct {
}

type ReferenceSchemaRecordFormatParameters struct {

	// Provides additional mapping information specific to the record format (such as JSON, CSV, or record fields delimited by some delimiter) on the streaming source.
	// +kubebuilder:validation:Required
	MappingParameters []RecordFormatMappingParametersParameters `json:"mappingParameters" tf:"mapping_parameters,omitempty"`

	// The type of record format. Valid values: CSV, JSON.
	// +kubebuilder:validation:Required
	RecordFormatType *string `json:"recordFormatType" tf:"record_format_type,omitempty"`
}

type RunConfigurationObservation struct {
}

type RunConfigurationParameters struct {

	// The restore behavior of a restarting application.
	// +kubebuilder:validation:Optional
	ApplicationRestoreConfiguration []ApplicationRestoreConfigurationParameters `json:"applicationRestoreConfiguration,omitempty" tf:"application_restore_configuration,omitempty"`

	// The starting parameters for a Flink-based Kinesis Data Analytics application.
	// +kubebuilder:validation:Optional
	FlinkRunConfiguration []FlinkRunConfigurationParameters `json:"flinkRunConfiguration,omitempty" tf:"flink_run_configuration,omitempty"`
}

type S3ContentLocationObservation struct {
}

type S3ContentLocationParameters struct {

	// The ARN for the S3 bucket containing the application code.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	BucketArn *string `json:"bucketArn,omitempty" tf:"bucket_arn,omitempty"`

	// Reference to a Bucket in s3 to populate bucketArn.
	// +kubebuilder:validation:Optional
	BucketArnRef *v1.Reference `json:"bucketArnRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucketArn.
	// +kubebuilder:validation:Optional
	BucketArnSelector *v1.Selector `json:"bucketArnSelector,omitempty" tf:"-"`

	// The file key for the object containing the application code.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Object
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("key",false)
	// +kubebuilder:validation:Optional
	FileKey *string `json:"fileKey,omitempty" tf:"file_key,omitempty"`

	// Reference to a Object in s3 to populate fileKey.
	// +kubebuilder:validation:Optional
	FileKeyRef *v1.Reference `json:"fileKeyRef,omitempty" tf:"-"`

	// Selector for a Object in s3 to populate fileKey.
	// +kubebuilder:validation:Optional
	FileKeySelector *v1.Selector `json:"fileKeySelector,omitempty" tf:"-"`

	// The version of the object containing the application code.
	// +kubebuilder:validation:Optional
	ObjectVersion *string `json:"objectVersion,omitempty" tf:"object_version,omitempty"`
}

type S3ReferenceDataSourceObservation struct {
}

type S3ReferenceDataSourceParameters struct {

	// The ARN for the S3 bucket containing the application code.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/s3/v1beta1.Bucket
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	BucketArn *string `json:"bucketArn,omitempty" tf:"bucket_arn,omitempty"`

	// Reference to a Bucket in s3 to populate bucketArn.
	// +kubebuilder:validation:Optional
	BucketArnRef *v1.Reference `json:"bucketArnRef,omitempty" tf:"-"`

	// Selector for a Bucket in s3 to populate bucketArn.
	// +kubebuilder:validation:Optional
	BucketArnSelector *v1.Selector `json:"bucketArnSelector,omitempty" tf:"-"`

	// The file key for the object containing the application code.
	// +kubebuilder:validation:Required
	FileKey *string `json:"fileKey" tf:"file_key,omitempty"`
}

type SQLApplicationConfigurationObservation struct {

	// The input stream used by the application.
	// +kubebuilder:validation:Optional
	Input []InputObservation `json:"input,omitempty" tf:"input,omitempty"`

	// The destination streams used by the application.
	// +kubebuilder:validation:Optional
	Output []OutputObservation `json:"output,omitempty" tf:"output,omitempty"`

	// The reference data source used by the application.
	// +kubebuilder:validation:Optional
	ReferenceDataSource []ReferenceDataSourceObservation `json:"referenceDataSource,omitempty" tf:"reference_data_source,omitempty"`
}

type SQLApplicationConfigurationParameters struct {

	// The input stream used by the application.
	// +kubebuilder:validation:Optional
	Input []InputParameters `json:"input,omitempty" tf:"input,omitempty"`

	// The destination streams used by the application.
	// +kubebuilder:validation:Optional
	Output []OutputParameters `json:"output,omitempty" tf:"output,omitempty"`

	// The reference data source used by the application.
	// +kubebuilder:validation:Optional
	ReferenceDataSource []ReferenceDataSourceParameters `json:"referenceDataSource,omitempty" tf:"reference_data_source,omitempty"`
}

type VPCConfigurationObservation struct {

	// The application identifier.
	VPCConfigurationID *string `json:"vpcConfigurationId,omitempty" tf:"vpc_configuration_id,omitempty"`

	// The application identifier.
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`
}

type VPCConfigurationParameters struct {

	// The Security Group IDs used by the VPC configuration.
	// +kubebuilder:validation:Required
	SecurityGroupIds []*string `json:"securityGroupIds" tf:"security_group_ids,omitempty"`

	// The Subnet IDs used by the VPC configuration.
	// +kubebuilder:validation:Required
	SubnetIds []*string `json:"subnetIds" tf:"subnet_ids,omitempty"`
}

// ApplicationSpec defines the desired state of Application
type ApplicationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     ApplicationParameters `json:"forProvider"`
}

// ApplicationStatus defines the observed state of Application.
type ApplicationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        ApplicationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Application is the Schema for the Applications API. Manages a Kinesis Analytics v2 Application.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type Application struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              ApplicationSpec   `json:"spec"`
	Status            ApplicationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ApplicationList contains a list of Applications
type ApplicationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Application `json:"items"`
}

// Repository type metadata.
var (
	Application_Kind             = "Application"
	Application_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Application_Kind}.String()
	Application_KindAPIVersion   = Application_Kind + "." + CRDGroupVersion.String()
	Application_GroupVersionKind = CRDGroupVersion.WithKind(Application_Kind)
)

func init() {
	SchemeBuilder.Register(&Application{}, &ApplicationList{})
}
