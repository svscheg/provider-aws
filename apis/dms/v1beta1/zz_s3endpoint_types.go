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

type S3EndpointObservation struct {

	// Whether to add column name information to the .csv output file. Default is false.
	AddColumnName *bool `json:"addColumnName,omitempty" tf:"add_column_name,omitempty"`

	// Whether to add padding. Default is false. (Ignored for source endpoints.)
	AddTrailingPaddingCharacter *bool `json:"addTrailingPaddingCharacter,omitempty" tf:"add_trailing_padding_character,omitempty"`

	// S3 object prefix.
	BucketFolder *string `json:"bucketFolder,omitempty" tf:"bucket_folder,omitempty"`

	// S3 bucket name.
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Predefined (canned) access control list for objects created in an S3 bucket. Valid values include none, private, public-read, public-read-write, authenticated-read, aws-exec-read, bucket-owner-read, and bucket-owner-full-control. Default is none.
	CannedACLForObjects *string `json:"cannedAclForObjects,omitempty" tf:"canned_acl_for_objects,omitempty"`

	// Whether to write insert and update operations to .csv or .parquet output files. Default is false.
	CdcInsertsAndUpdates *bool `json:"cdcInsertsAndUpdates,omitempty" tf:"cdc_inserts_and_updates,omitempty"`

	// Whether to write insert operations to .csv or .parquet output files. Default is false.
	CdcInsertsOnly *bool `json:"cdcInsertsOnly,omitempty" tf:"cdc_inserts_only,omitempty"`

	// Maximum length of the interval, defined in seconds, after which to output a file to Amazon S3. (AWS default is 60.)
	CdcMaxBatchInterval *float64 `json:"cdcMaxBatchInterval,omitempty" tf:"cdc_max_batch_interval,omitempty"`

	// Minimum file size condition as defined in kilobytes to output a file to Amazon S3. (AWS default is 32000 KB.)
	CdcMinFileSize *float64 `json:"cdcMinFileSize,omitempty" tf:"cdc_min_file_size,omitempty"`

	// Folder path of CDC files. If cdc_path is set, AWS DMS reads CDC files from this path and replicates the data changes to the target endpoint. Supported in AWS DMS versions 3.4.2 and later.
	CdcPath *string `json:"cdcPath,omitempty" tf:"cdc_path,omitempty"`

	// ARN for the certificate.
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Set to compress target files. Valid values are GZIP and NONE. Default is NONE. (Ignored for source endpoints.)
	CompressionType *string `json:"compressionType,omitempty" tf:"compression_type,omitempty"`

	// Delimiter used to separate columns in the source files. Default is ,.
	CsvDelimiter *string `json:"csvDelimiter,omitempty" tf:"csv_delimiter,omitempty"`

	// Only applies if output files for a CDC load are written in .csv format. If use_csv_no_sup_value is set to true, string to use for all columns not included in the supplemental log. If you do not specify a string value, DMS uses the null value for these columns regardless of use_csv_no_sup_value. (Ignored for source endpoints.)
	CsvNoSupValue *string `json:"csvNoSupValue,omitempty" tf:"csv_no_sup_value,omitempty"`

	// String to as null when writing to the target. (AWS default is NULL.)
	CsvNullValue *string `json:"csvNullValue,omitempty" tf:"csv_null_value,omitempty"`

	// Delimiter used to separate rows in the source files. Default is newline (i.e., \n).
	CsvRowDelimiter *string `json:"csvRowDelimiter,omitempty" tf:"csv_row_delimiter,omitempty"`

	// Output format for the files that AWS DMS uses to create S3 objects. Valid values are csv and parquet.  (Ignored for source endpoints -- only csv is valid.)
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`

	// Size of one data page in bytes. (AWS default is 1 MiB, i.e., 1048576.)
	DataPageSize *float64 `json:"dataPageSize,omitempty" tf:"data_page_size,omitempty"`

	// Date separating delimiter to use during folder partitioning. Valid values are SLASH, UNDERSCORE, DASH, and NONE. (AWS default is SLASH.) (Ignored for source endpoints.)
	DatePartitionDelimiter *string `json:"datePartitionDelimiter,omitempty" tf:"date_partition_delimiter,omitempty"`

	// Partition S3 bucket folders based on transaction commit dates. Default is false. (Ignored for source endpoints.)
	DatePartitionEnabled *bool `json:"datePartitionEnabled,omitempty" tf:"date_partition_enabled,omitempty"`

	// Date format to use during folder partitioning. Use this parameter when date_partition_enabled is set to true. Valid values are YYYYMMDD, YYYYMMDDHH, YYYYMM, MMYYYYDD, and DDMMYYYY. (AWS default is YYYYMMDD.) (Ignored for source endpoints.)
	DatePartitionSequence *string `json:"datePartitionSequence,omitempty" tf:"date_partition_sequence,omitempty"`

	// Convert the current UTC time to a timezone. The conversion occurs when a date partition folder is created and a CDC filename is generated. The timezone format is Area/Location (e.g., Europe/Paris). Use this when date_partition_enabled is true. (Ignored for source endpoints.)
	DatePartitionTimezone *string `json:"datePartitionTimezone,omitempty" tf:"date_partition_timezone,omitempty"`

	// Undocumented argument for use as directed by AWS Support.
	DetachTargetOnLobLookupFailureParquet *bool `json:"detachTargetOnLobLookupFailureParquet,omitempty" tf:"detach_target_on_lob_lookup_failure_parquet,omitempty"`

	// Maximum size in bytes of an encoded dictionary page of a column. (AWS default is 1 MiB, i.e., 1048576.)
	DictPageSizeLimit *float64 `json:"dictPageSizeLimit,omitempty" tf:"dict_page_size_limit,omitempty"`

	// Whether to enable statistics for Parquet pages and row groups. Default is true.
	EnableStatistics *bool `json:"enableStatistics,omitempty" tf:"enable_statistics,omitempty"`

	// Type of encoding to use. Value values are rle_dictionary, plain, and plain_dictionary. (AWS default is rle_dictionary.)
	EncodingType *string `json:"encodingType,omitempty" tf:"encoding_type,omitempty"`

	// Server-side encryption mode that you want to encrypt your .csv or .parquet object files copied to S3. Valid values are SSE_S3 and SSE_KMS. (AWS default is SSE_S3.) (Ignored for source endpoints -- only SSE_S3 is valid.)
	EncryptionMode *string `json:"encryptionMode,omitempty" tf:"encryption_mode,omitempty"`

	// ARN for the endpoint.
	EndpointArn *string `json:"endpointArn,omitempty" tf:"endpoint_arn,omitempty"`

	// Type of endpoint. Valid values are source, target.
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// Expanded name for the engine name.
	EngineDisplayName *string `json:"engineDisplayName,omitempty" tf:"engine_display_name,omitempty"`

	// Bucket owner to prevent sniping. Value is an AWS account ID.
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// Can be used for cross-account validation. Use it in another account with aws_dms_s3_endpoint to create the endpoint cross-account.
	ExternalID *string `json:"externalId,omitempty" tf:"external_id,omitempty"`

	// JSON document that describes how AWS DMS should interpret the data.
	ExternalTableDefinition *string `json:"externalTableDefinition,omitempty" tf:"external_table_definition,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// When this value is set to 1, DMS ignores the first row header in a .csv file. (AWS default is 0.)
	IgnoreHeaderRows *float64 `json:"ignoreHeaderRows,omitempty" tf:"ignore_header_rows,omitempty"`

	// Whether to enable a full load to write INSERT operations to the .csv output files only to indicate how the rows were added to the source database. Default is false.
	IncludeOpForFullLoad *bool `json:"includeOpForFullLoad,omitempty" tf:"include_op_for_full_load,omitempty"`

	// ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for kms_key_arn, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Maximum size (in KB) of any .csv file to be created while migrating to an S3 target during full load. Valid values are from 1 to 1048576. (AWS default is 1 GB, i.e., 1048576.)
	MaxFileSize *float64 `json:"maxFileSize,omitempty" tf:"max_file_size,omitempty"`

	// - Specifies the precision of any TIMESTAMP column values written to an S3 object file in .parquet format. Default is false. (Ignored for source endpoints.)
	ParquetTimestampInMillisecond *bool `json:"parquetTimestampInMillisecond,omitempty" tf:"parquet_timestamp_in_millisecond,omitempty"`

	// Version of the .parquet file format. Valid values are parquet-1-0 and parquet-2-0. (AWS default is parquet-1-0.) (Ignored for source endpoints.)
	ParquetVersion *string `json:"parquetVersion,omitempty" tf:"parquet_version,omitempty"`

	// Whether DMS saves the transaction order for a CDC load on the S3 target specified by cdc_path. Default is false. (Ignored for source endpoints.)
	PreserveTransactions *bool `json:"preserveTransactions,omitempty" tf:"preserve_transactions,omitempty"`

	// For an S3 source, whether each leading double quotation mark has to be followed by an ending double quotation mark. Default is true.
	Rfc4180 *bool `json:"rfc4180,omitempty" tf:"rfc_4180,omitempty"`

	// Number of rows in a row group. (AWS default is 10000.)
	RowGroupLength *float64 `json:"rowGroupLength,omitempty" tf:"row_group_length,omitempty"`

	// SSL mode to use for the connection. Valid values are none, require, verify-ca, verify-full. (AWS default is none.)
	SSLMode *string `json:"sslMode,omitempty" tf:"ssl_mode,omitempty"`

	// When encryption_mode is SSE_KMS, ARN for the AWS KMS key. (Ignored for source endpoints -- only SSE_S3 encryption_mode is valid.)
	ServerSideEncryptionKMSKeyID *string `json:"serverSideEncryptionKmsKeyId,omitempty" tf:"server_side_encryption_kms_key_id,omitempty"`

	// ARN of the IAM role with permissions to the S3 Bucket.
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn,omitempty" tf:"service_access_role_arn,omitempty"`

	// Status of the endpoint.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// Key-value map of resource tags.
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Map of tags assigned to the resource, including those inherited from the provider default_tags configuration block.
	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`

	// Column to add with timestamp information to the endpoint data for an Amazon S3 target.
	TimestampColumnName *string `json:"timestampColumnName,omitempty" tf:"timestamp_column_name,omitempty"`

	// Whether to use csv_no_sup_value for columns not included in the supplemental log. (Ignored for source endpoints.)
	UseCsvNoSupValue *bool `json:"useCsvNoSupValue,omitempty" tf:"use_csv_no_sup_value,omitempty"`

	// When set to true, uses the task start time as the timestamp column value instead of the time data is written to target. For full load, when set to true, each row of the timestamp column contains the task start time. For CDC loads, each row of the timestamp column contains the transaction commit time.When set to false, the full load timestamp in the timestamp column increments with the time data arrives at the target. Default is false.
	UseTaskStartTimeForFullLoadTimestamp *bool `json:"useTaskStartTimeForFullLoadTimestamp,omitempty" tf:"use_task_start_time_for_full_load_timestamp,omitempty"`
}

type S3EndpointParameters struct {

	// Whether to add column name information to the .csv output file. Default is false.
	// +kubebuilder:validation:Optional
	AddColumnName *bool `json:"addColumnName,omitempty" tf:"add_column_name,omitempty"`

	// Whether to add padding. Default is false. (Ignored for source endpoints.)
	// +kubebuilder:validation:Optional
	AddTrailingPaddingCharacter *bool `json:"addTrailingPaddingCharacter,omitempty" tf:"add_trailing_padding_character,omitempty"`

	// S3 object prefix.
	// +kubebuilder:validation:Optional
	BucketFolder *string `json:"bucketFolder,omitempty" tf:"bucket_folder,omitempty"`

	// S3 bucket name.
	// +kubebuilder:validation:Optional
	BucketName *string `json:"bucketName,omitempty" tf:"bucket_name,omitempty"`

	// Predefined (canned) access control list for objects created in an S3 bucket. Valid values include none, private, public-read, public-read-write, authenticated-read, aws-exec-read, bucket-owner-read, and bucket-owner-full-control. Default is none.
	// +kubebuilder:validation:Optional
	CannedACLForObjects *string `json:"cannedAclForObjects,omitempty" tf:"canned_acl_for_objects,omitempty"`

	// Whether to write insert and update operations to .csv or .parquet output files. Default is false.
	// +kubebuilder:validation:Optional
	CdcInsertsAndUpdates *bool `json:"cdcInsertsAndUpdates,omitempty" tf:"cdc_inserts_and_updates,omitempty"`

	// Whether to write insert operations to .csv or .parquet output files. Default is false.
	// +kubebuilder:validation:Optional
	CdcInsertsOnly *bool `json:"cdcInsertsOnly,omitempty" tf:"cdc_inserts_only,omitempty"`

	// Maximum length of the interval, defined in seconds, after which to output a file to Amazon S3. (AWS default is 60.)
	// +kubebuilder:validation:Optional
	CdcMaxBatchInterval *float64 `json:"cdcMaxBatchInterval,omitempty" tf:"cdc_max_batch_interval,omitempty"`

	// Minimum file size condition as defined in kilobytes to output a file to Amazon S3. (AWS default is 32000 KB.)
	// +kubebuilder:validation:Optional
	CdcMinFileSize *float64 `json:"cdcMinFileSize,omitempty" tf:"cdc_min_file_size,omitempty"`

	// Folder path of CDC files. If cdc_path is set, AWS DMS reads CDC files from this path and replicates the data changes to the target endpoint. Supported in AWS DMS versions 3.4.2 and later.
	// +kubebuilder:validation:Optional
	CdcPath *string `json:"cdcPath,omitempty" tf:"cdc_path,omitempty"`

	// ARN for the certificate.
	// +kubebuilder:validation:Optional
	CertificateArn *string `json:"certificateArn,omitempty" tf:"certificate_arn,omitempty"`

	// Set to compress target files. Valid values are GZIP and NONE. Default is NONE. (Ignored for source endpoints.)
	// +kubebuilder:validation:Optional
	CompressionType *string `json:"compressionType,omitempty" tf:"compression_type,omitempty"`

	// Delimiter used to separate columns in the source files. Default is ,.
	// +kubebuilder:validation:Optional
	CsvDelimiter *string `json:"csvDelimiter,omitempty" tf:"csv_delimiter,omitempty"`

	// Only applies if output files for a CDC load are written in .csv format. If use_csv_no_sup_value is set to true, string to use for all columns not included in the supplemental log. If you do not specify a string value, DMS uses the null value for these columns regardless of use_csv_no_sup_value. (Ignored for source endpoints.)
	// +kubebuilder:validation:Optional
	CsvNoSupValue *string `json:"csvNoSupValue,omitempty" tf:"csv_no_sup_value,omitempty"`

	// String to as null when writing to the target. (AWS default is NULL.)
	// +kubebuilder:validation:Optional
	CsvNullValue *string `json:"csvNullValue,omitempty" tf:"csv_null_value,omitempty"`

	// Delimiter used to separate rows in the source files. Default is newline (i.e., \n).
	// +kubebuilder:validation:Optional
	CsvRowDelimiter *string `json:"csvRowDelimiter,omitempty" tf:"csv_row_delimiter,omitempty"`

	// Output format for the files that AWS DMS uses to create S3 objects. Valid values are csv and parquet.  (Ignored for source endpoints -- only csv is valid.)
	// +kubebuilder:validation:Optional
	DataFormat *string `json:"dataFormat,omitempty" tf:"data_format,omitempty"`

	// Size of one data page in bytes. (AWS default is 1 MiB, i.e., 1048576.)
	// +kubebuilder:validation:Optional
	DataPageSize *float64 `json:"dataPageSize,omitempty" tf:"data_page_size,omitempty"`

	// Date separating delimiter to use during folder partitioning. Valid values are SLASH, UNDERSCORE, DASH, and NONE. (AWS default is SLASH.) (Ignored for source endpoints.)
	// +kubebuilder:validation:Optional
	DatePartitionDelimiter *string `json:"datePartitionDelimiter,omitempty" tf:"date_partition_delimiter,omitempty"`

	// Partition S3 bucket folders based on transaction commit dates. Default is false. (Ignored for source endpoints.)
	// +kubebuilder:validation:Optional
	DatePartitionEnabled *bool `json:"datePartitionEnabled,omitempty" tf:"date_partition_enabled,omitempty"`

	// Date format to use during folder partitioning. Use this parameter when date_partition_enabled is set to true. Valid values are YYYYMMDD, YYYYMMDDHH, YYYYMM, MMYYYYDD, and DDMMYYYY. (AWS default is YYYYMMDD.) (Ignored for source endpoints.)
	// +kubebuilder:validation:Optional
	DatePartitionSequence *string `json:"datePartitionSequence,omitempty" tf:"date_partition_sequence,omitempty"`

	// Convert the current UTC time to a timezone. The conversion occurs when a date partition folder is created and a CDC filename is generated. The timezone format is Area/Location (e.g., Europe/Paris). Use this when date_partition_enabled is true. (Ignored for source endpoints.)
	// +kubebuilder:validation:Optional
	DatePartitionTimezone *string `json:"datePartitionTimezone,omitempty" tf:"date_partition_timezone,omitempty"`

	// Undocumented argument for use as directed by AWS Support.
	// +kubebuilder:validation:Optional
	DetachTargetOnLobLookupFailureParquet *bool `json:"detachTargetOnLobLookupFailureParquet,omitempty" tf:"detach_target_on_lob_lookup_failure_parquet,omitempty"`

	// Maximum size in bytes of an encoded dictionary page of a column. (AWS default is 1 MiB, i.e., 1048576.)
	// +kubebuilder:validation:Optional
	DictPageSizeLimit *float64 `json:"dictPageSizeLimit,omitempty" tf:"dict_page_size_limit,omitempty"`

	// Whether to enable statistics for Parquet pages and row groups. Default is true.
	// +kubebuilder:validation:Optional
	EnableStatistics *bool `json:"enableStatistics,omitempty" tf:"enable_statistics,omitempty"`

	// Type of encoding to use. Value values are rle_dictionary, plain, and plain_dictionary. (AWS default is rle_dictionary.)
	// +kubebuilder:validation:Optional
	EncodingType *string `json:"encodingType,omitempty" tf:"encoding_type,omitempty"`

	// Server-side encryption mode that you want to encrypt your .csv or .parquet object files copied to S3. Valid values are SSE_S3 and SSE_KMS. (AWS default is SSE_S3.) (Ignored for source endpoints -- only SSE_S3 is valid.)
	// +kubebuilder:validation:Optional
	EncryptionMode *string `json:"encryptionMode,omitempty" tf:"encryption_mode,omitempty"`

	// Type of endpoint. Valid values are source, target.
	// +kubebuilder:validation:Optional
	EndpointType *string `json:"endpointType,omitempty" tf:"endpoint_type,omitempty"`

	// Bucket owner to prevent sniping. Value is an AWS account ID.
	// +kubebuilder:validation:Optional
	ExpectedBucketOwner *string `json:"expectedBucketOwner,omitempty" tf:"expected_bucket_owner,omitempty"`

	// JSON document that describes how AWS DMS should interpret the data.
	// +kubebuilder:validation:Optional
	ExternalTableDefinition *string `json:"externalTableDefinition,omitempty" tf:"external_table_definition,omitempty"`

	// When this value is set to 1, DMS ignores the first row header in a .csv file. (AWS default is 0.)
	// +kubebuilder:validation:Optional
	IgnoreHeaderRows *float64 `json:"ignoreHeaderRows,omitempty" tf:"ignore_header_rows,omitempty"`

	// Whether to enable a full load to write INSERT operations to the .csv output files only to indicate how the rows were added to the source database. Default is false.
	// +kubebuilder:validation:Optional
	IncludeOpForFullLoad *bool `json:"includeOpForFullLoad,omitempty" tf:"include_op_for_full_load,omitempty"`

	// ARN for the KMS key that will be used to encrypt the connection parameters. If you do not specify a value for kms_key_arn, then AWS DMS will use your default encryption key. AWS KMS creates the default encryption key for your AWS account. Your AWS account has a different default encryption key for each AWS region.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +kubebuilder:validation:Optional
	KMSKeyArn *string `json:"kmsKeyArn,omitempty" tf:"kms_key_arn,omitempty"`

	// Reference to a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnRef *v1.Reference `json:"kmsKeyArnRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate kmsKeyArn.
	// +kubebuilder:validation:Optional
	KMSKeyArnSelector *v1.Selector `json:"kmsKeyArnSelector,omitempty" tf:"-"`

	// Maximum size (in KB) of any .csv file to be created while migrating to an S3 target during full load. Valid values are from 1 to 1048576. (AWS default is 1 GB, i.e., 1048576.)
	// +kubebuilder:validation:Optional
	MaxFileSize *float64 `json:"maxFileSize,omitempty" tf:"max_file_size,omitempty"`

	// - Specifies the precision of any TIMESTAMP column values written to an S3 object file in .parquet format. Default is false. (Ignored for source endpoints.)
	// +kubebuilder:validation:Optional
	ParquetTimestampInMillisecond *bool `json:"parquetTimestampInMillisecond,omitempty" tf:"parquet_timestamp_in_millisecond,omitempty"`

	// Version of the .parquet file format. Valid values are parquet-1-0 and parquet-2-0. (AWS default is parquet-1-0.) (Ignored for source endpoints.)
	// +kubebuilder:validation:Optional
	ParquetVersion *string `json:"parquetVersion,omitempty" tf:"parquet_version,omitempty"`

	// Whether DMS saves the transaction order for a CDC load on the S3 target specified by cdc_path. Default is false. (Ignored for source endpoints.)
	// +kubebuilder:validation:Optional
	PreserveTransactions *bool `json:"preserveTransactions,omitempty" tf:"preserve_transactions,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +upjet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// For an S3 source, whether each leading double quotation mark has to be followed by an ending double quotation mark. Default is true.
	// +kubebuilder:validation:Optional
	Rfc4180 *bool `json:"rfc4180,omitempty" tf:"rfc_4180,omitempty"`

	// Number of rows in a row group. (AWS default is 10000.)
	// +kubebuilder:validation:Optional
	RowGroupLength *float64 `json:"rowGroupLength,omitempty" tf:"row_group_length,omitempty"`

	// SSL mode to use for the connection. Valid values are none, require, verify-ca, verify-full. (AWS default is none.)
	// +kubebuilder:validation:Optional
	SSLMode *string `json:"sslMode,omitempty" tf:"ssl_mode,omitempty"`

	// When encryption_mode is SSE_KMS, ARN for the AWS KMS key. (Ignored for source endpoints -- only SSE_S3 encryption_mode is valid.)
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/kms/v1beta1.Key
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractParamPath("arn",true)
	// +kubebuilder:validation:Optional
	ServerSideEncryptionKMSKeyID *string `json:"serverSideEncryptionKmsKeyId,omitempty" tf:"server_side_encryption_kms_key_id,omitempty"`

	// Reference to a Key in kms to populate serverSideEncryptionKmsKeyId.
	// +kubebuilder:validation:Optional
	ServerSideEncryptionKMSKeyIDRef *v1.Reference `json:"serverSideEncryptionKmsKeyIdRef,omitempty" tf:"-"`

	// Selector for a Key in kms to populate serverSideEncryptionKmsKeyId.
	// +kubebuilder:validation:Optional
	ServerSideEncryptionKMSKeyIDSelector *v1.Selector `json:"serverSideEncryptionKmsKeyIdSelector,omitempty" tf:"-"`

	// ARN of the IAM role with permissions to the S3 Bucket.
	// +crossplane:generate:reference:type=github.com/upbound/provider-aws/apis/iam/v1beta1.Role
	// +crossplane:generate:reference:extractor=github.com/upbound/provider-aws/config/common.ARNExtractor()
	// +kubebuilder:validation:Optional
	ServiceAccessRoleArn *string `json:"serviceAccessRoleArn,omitempty" tf:"service_access_role_arn,omitempty"`

	// Reference to a Role in iam to populate serviceAccessRoleArn.
	// +kubebuilder:validation:Optional
	ServiceAccessRoleArnRef *v1.Reference `json:"serviceAccessRoleArnRef,omitempty" tf:"-"`

	// Selector for a Role in iam to populate serviceAccessRoleArn.
	// +kubebuilder:validation:Optional
	ServiceAccessRoleArnSelector *v1.Selector `json:"serviceAccessRoleArnSelector,omitempty" tf:"-"`

	// Key-value map of resource tags.
	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// Column to add with timestamp information to the endpoint data for an Amazon S3 target.
	// +kubebuilder:validation:Optional
	TimestampColumnName *string `json:"timestampColumnName,omitempty" tf:"timestamp_column_name,omitempty"`

	// Whether to use csv_no_sup_value for columns not included in the supplemental log. (Ignored for source endpoints.)
	// +kubebuilder:validation:Optional
	UseCsvNoSupValue *bool `json:"useCsvNoSupValue,omitempty" tf:"use_csv_no_sup_value,omitempty"`

	// When set to true, uses the task start time as the timestamp column value instead of the time data is written to target. For full load, when set to true, each row of the timestamp column contains the task start time. For CDC loads, each row of the timestamp column contains the transaction commit time.When set to false, the full load timestamp in the timestamp column increments with the time data arrives at the target. Default is false.
	// +kubebuilder:validation:Optional
	UseTaskStartTimeForFullLoadTimestamp *bool `json:"useTaskStartTimeForFullLoadTimestamp,omitempty" tf:"use_task_start_time_for_full_load_timestamp,omitempty"`
}

// S3EndpointSpec defines the desired state of S3Endpoint
type S3EndpointSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     S3EndpointParameters `json:"forProvider"`
}

// S3EndpointStatus defines the observed state of S3Endpoint.
type S3EndpointStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        S3EndpointObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// S3Endpoint is the Schema for the S3Endpoints API. Provides a DMS (Data Migration Service) S3 endpoint resource.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,aws}
type S3Endpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.bucketName)",message="bucketName is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.endpointType)",message="endpointType is a required parameter"
	Spec   S3EndpointSpec   `json:"spec"`
	Status S3EndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// S3EndpointList contains a list of S3Endpoints
type S3EndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []S3Endpoint `json:"items"`
}

// Repository type metadata.
var (
	S3Endpoint_Kind             = "S3Endpoint"
	S3Endpoint_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: S3Endpoint_Kind}.String()
	S3Endpoint_KindAPIVersion   = S3Endpoint_Kind + "." + CRDGroupVersion.String()
	S3Endpoint_GroupVersionKind = CRDGroupVersion.WithKind(S3Endpoint_Kind)
)

func init() {
	SchemeBuilder.Register(&S3Endpoint{}, &S3EndpointList{})
}
