// Code generated by generators/resource/main.go; DO NOT EDIT.

package cassandra

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_cassandra_table", tableResourceType)
}

// tableResourceType returns the Terraform awscc_cassandra_table resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Cassandra::Table resource type.
func tableResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"billing_mode": {
			// Property: BillingMode
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "Mode": {
			//       "default": "ON_DEMAND",
			//       "description": "Capacity mode for the specified table",
			//       "enum": [
			//         "PROVISIONED",
			//         "ON_DEMAND"
			//       ],
			//       "type": "string"
			//     },
			//     "ProvisionedThroughput": {
			//       "additionalProperties": false,
			//       "description": "Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits",
			//       "properties": {
			//         "ReadCapacityUnits": {
			//           "minimum": 1,
			//           "type": "integer"
			//         },
			//         "WriteCapacityUnits": {
			//           "minimum": 1,
			//           "type": "integer"
			//         }
			//       },
			//       "required": [
			//         "ReadCapacityUnits",
			//         "WriteCapacityUnits"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "required": [
			//     "Mode"
			//   ],
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"mode": {
						// Property: Mode
						Description: "Capacity mode for the specified table",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"PROVISIONED",
								"ON_DEMAND",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							DefaultValue(types.String{Value: "ON_DEMAND"}),
							tfsdk.UseStateForUnknown(),
						},
					},
					"provisioned_throughput": {
						// Property: ProvisionedThroughput
						Description: "Throughput for the specified table, which consists of values for ReadCapacityUnits and WriteCapacityUnits",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"read_capacity_units": {
									// Property: ReadCapacityUnits
									Type:     types.Int64Type,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntAtLeast(1),
									},
								},
								"write_capacity_units": {
									// Property: WriteCapacityUnits
									Type:     types.Int64Type,
									Required: true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntAtLeast(1),
									},
								},
							},
						),
						Optional: true,
					},
				},
			),
			Optional: true,
		},
		"clustering_key_columns": {
			// Property: ClusteringKeyColumns
			// CloudFormation resource type schema:
			// {
			//   "description": "Clustering key columns of the table",
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Column": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "ColumnName": {
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "ColumnType": {
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "ColumnName",
			//           "ColumnType"
			//         ],
			//         "type": "object"
			//       },
			//       "OrderBy": {
			//         "default": "ASC",
			//         "enum": [
			//           "ASC",
			//           "DESC"
			//         ],
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Column"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Clustering key columns of the table",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"column": {
						// Property: Column
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"column_name": {
									// Property: ColumnName
									Type:     types.StringType,
									Required: true,
								},
								"column_type": {
									// Property: ColumnType
									Type:     types.StringType,
									Required: true,
								},
							},
						),
						Required: true,
					},
					"order_by": {
						// Property: OrderBy
						Type:     types.StringType,
						Optional: true,
						Computed: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"ASC",
								"DESC",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							DefaultValue(types.String{Value: "ASC"}),
							tfsdk.UseStateForUnknown(),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.UniqueItems(),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"default_time_to_live": {
			// Property: DefaultTimeToLive
			// CloudFormation resource type schema:
			// {
			//   "description": "Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.",
			//   "minimum": 0,
			//   "type": "integer"
			// }
			Description: "Default TTL (Time To Live) in seconds, where zero is disabled. If the value is greater than zero, TTL is enabled for the entire table and an expiration timestamp is added to each column.",
			Type:        types.Int64Type,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.IntAtLeast(0),
			},
		},
		"encryption_specification": {
			// Property: EncryptionSpecification
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Represents the settings used to enable server-side encryption",
			//   "properties": {
			//     "EncryptionType": {
			//       "default": "AWS_OWNED_KMS_KEY",
			//       "description": "Server-side encryption type",
			//       "enum": [
			//         "AWS_OWNED_KMS_KEY",
			//         "CUSTOMER_MANAGED_KMS_KEY"
			//       ],
			//       "type": "string"
			//     },
			//     "KmsKeyIdentifier": {
			//       "description": "The AWS KMS customer master key (CMK) that should be used for the AWS KMS encryption. To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. ",
			//       "type": "string"
			//     }
			//   },
			//   "required": [
			//     "EncryptionType"
			//   ],
			//   "type": "object"
			// }
			Description: "Represents the settings used to enable server-side encryption",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"encryption_type": {
						// Property: EncryptionType
						Description: "Server-side encryption type",
						Type:        types.StringType,
						Optional:    true,
						Computed:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringInSlice([]string{
								"AWS_OWNED_KMS_KEY",
								"CUSTOMER_MANAGED_KMS_KEY",
							}),
						},
						PlanModifiers: []tfsdk.AttributePlanModifier{
							DefaultValue(types.String{Value: "AWS_OWNED_KMS_KEY"}),
							tfsdk.UseStateForUnknown(),
						},
					},
					"kms_key_identifier": {
						// Property: KmsKeyIdentifier
						Description: "The AWS KMS customer master key (CMK) that should be used for the AWS KMS encryption. To specify a CMK, use its key ID, Amazon Resource Name (ARN), alias name, or alias ARN. ",
						Type:        types.StringType,
						Optional:    true,
					},
				},
			),
			Optional: true,
		},
		"keyspace_name": {
			// Property: KeyspaceName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name for Cassandra keyspace",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name for Cassandra keyspace",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"partition_key_columns": {
			// Property: PartitionKeyColumns
			// CloudFormation resource type schema:
			// {
			//   "description": "Partition key columns of the table",
			//   "insertionOrder": true,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ColumnName": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "ColumnType": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "ColumnName",
			//       "ColumnType"
			//     ],
			//     "type": "object"
			//   },
			//   "minItems": 1,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Partition key columns of the table",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"column_name": {
						// Property: ColumnName
						Type:     types.StringType,
						Required: true,
					},
					"column_type": {
						// Property: ColumnType
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Required: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenAtLeast(1),
				validate.UniqueItems(),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"point_in_time_recovery_enabled": {
			// Property: PointInTimeRecoveryEnabled
			// CloudFormation resource type schema:
			// {
			//   "description": "Indicates whether point in time recovery is enabled (true) or disabled (false) on the table",
			//   "type": "boolean"
			// }
			Description: "Indicates whether point in time recovery is enabled (true) or disabled (false) on the table",
			Type:        types.BoolType,
			Optional:    true,
		},
		"regular_columns": {
			// Property: RegularColumns
			// CloudFormation resource type schema:
			// {
			//   "description": "Non-key columns of the table",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "ColumnName": {
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "ColumnType": {
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "ColumnName",
			//       "ColumnType"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Non-key columns of the table",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"column_name": {
						// Property: ColumnName
						Type:     types.StringType,
						Required: true,
					},
					"column_type": {
						// Property: ColumnType
						Type:     types.StringType,
						Required: true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Optional: true,
		},
		"table_name": {
			// Property: TableName
			// CloudFormation resource type schema:
			// {
			//   "description": "Name for Cassandra table",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Name for Cassandra table",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "An array of key-value pairs to apply to this resource",
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to apply to the resource",
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 256,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Value",
			//       "Key"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "An array of key-value pairs to apply to this resource",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Type:     types.StringType,
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 50),
				validate.UniqueItems(),
			},
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Computed:    true,
		PlanModifiers: []tfsdk.AttributePlanModifier{
			tfsdk.UseStateForUnknown(),
		},
	}

	schema := tfsdk.Schema{
		Description: "Resource schema for AWS::Cassandra::Table",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Cassandra::Table").WithTerraformTypeName("awscc_cassandra_table")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"billing_mode":                   "BillingMode",
		"clustering_key_columns":         "ClusteringKeyColumns",
		"column":                         "Column",
		"column_name":                    "ColumnName",
		"column_type":                    "ColumnType",
		"default_time_to_live":           "DefaultTimeToLive",
		"encryption_specification":       "EncryptionSpecification",
		"encryption_type":                "EncryptionType",
		"key":                            "Key",
		"keyspace_name":                  "KeyspaceName",
		"kms_key_identifier":             "KmsKeyIdentifier",
		"mode":                           "Mode",
		"order_by":                       "OrderBy",
		"partition_key_columns":          "PartitionKeyColumns",
		"point_in_time_recovery_enabled": "PointInTimeRecoveryEnabled",
		"provisioned_throughput":         "ProvisionedThroughput",
		"read_capacity_units":            "ReadCapacityUnits",
		"regular_columns":                "RegularColumns",
		"table_name":                     "TableName",
		"tags":                           "Tags",
		"value":                          "Value",
		"write_capacity_units":           "WriteCapacityUnits",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
