// Code generated by generators/resource/main.go; DO NOT EDIT.

package glue

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_glue_schema", schemaResourceType)
}

// schemaResourceType returns the Terraform awscc_glue_schema resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::Glue::Schema resource type.
func schemaResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "description": "Amazon Resource Name for the Schema.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Amazon Resource Name for the Schema.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"checkpoint_version": {
			// Property: CheckpointVersion
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Specify checkpoint version for update. This is only required to update the Compatibility.",
			//   "properties": {
			//     "IsLatest": {
			//       "description": "Indicates if the latest version needs to be updated.",
			//       "type": "boolean"
			//     },
			//     "VersionNumber": {
			//       "description": "Indicates the version number in the schema to update.",
			//       "maximum": 100000,
			//       "minimum": 1,
			//       "type": "integer"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Specify checkpoint version for update. This is only required to update the Compatibility.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"is_latest": {
						// Property: IsLatest
						Description: "Indicates if the latest version needs to be updated.",
						Type:        types.BoolType,
						Optional:    true,
					},
					"version_number": {
						// Property: VersionNumber
						Description: "Indicates the version number in the schema to update.",
						Type:        types.Int64Type,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(1, 100000),
						},
					},
				},
			),
			Optional: true,
		},
		"compatibility": {
			// Property: Compatibility
			// CloudFormation resource type schema:
			// {
			//   "description": "Compatibility setting for the schema.",
			//   "enum": [
			//     "NONE",
			//     "DISABLED",
			//     "BACKWARD",
			//     "BACKWARD_ALL",
			//     "FORWARD",
			//     "FORWARD_ALL",
			//     "FULL",
			//     "FULL_ALL"
			//   ],
			//   "type": "string"
			// }
			Description: "Compatibility setting for the schema.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"NONE",
					"DISABLED",
					"BACKWARD",
					"BACKWARD_ALL",
					"FORWARD",
					"FORWARD_ALL",
					"FULL",
					"FULL_ALL",
				}),
			},
		},
		"data_format": {
			// Property: DataFormat
			// CloudFormation resource type schema:
			// {
			//   "description": "Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF'",
			//   "enum": [
			//     "AVRO",
			//     "JSON",
			//     "PROTOBUF"
			//   ],
			//   "type": "string"
			// }
			Description: "Data format name to use for the schema. Accepted values: 'AVRO', 'JSON', 'PROTOBUF'",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"AVRO",
					"JSON",
					"PROTOBUF",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"description": {
			// Property: Description
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the schema. If description is not provided, there will not be any default value for this.",
			//   "maxLength": 1000,
			//   "minLength": 0,
			//   "type": "string"
			// }
			Description: "A description of the schema. If description is not provided, there will not be any default value for this.",
			Type:        types.StringType,
			Optional:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(0, 1000),
			},
		},
		"initial_schema_version_id": {
			// Property: InitialSchemaVersionId
			// CloudFormation resource type schema:
			// {
			//   "description": "Represents the version ID associated with the initial schema version.",
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "Represents the version ID associated with the initial schema version.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"name": {
			// Property: Name
			// CloudFormation resource type schema:
			// {
			//   "description": "Name of the schema.",
			//   "maxLength": 255,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Name of the schema.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 255),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"registry": {
			// Property: Registry
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Identifier for the registry which the schema is part of.",
			//   "properties": {
			//     "Arn": {
			//       "description": "Amazon Resource Name for the Registry.",
			//       "pattern": "",
			//       "type": "string"
			//     },
			//     "Name": {
			//       "description": "Name of the registry in which the schema will be created.",
			//       "maxLength": 255,
			//       "minLength": 1,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Identifier for the registry which the schema is part of.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"arn": {
						// Property: Arn
						Description: "Amazon Resource Name for the Registry.",
						Type:        types.StringType,
						Optional:    true,
					},
					"name": {
						// Property: Name
						Description: "Name of the registry in which the schema will be created.",
						Type:        types.StringType,
						Optional:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 255),
						},
					},
				},
			),
			Optional: true,
			Computed: true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
		},
		"schema_definition": {
			// Property: SchemaDefinition
			// CloudFormation resource type schema:
			// {
			//   "description": "Definition for the initial schema version in plain-text.",
			//   "maxLength": 170000,
			//   "minLength": 1,
			//   "type": "string"
			// }
			Description: "Definition for the initial schema version in plain-text.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 170000),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
			// SchemaDefinition is a write-only property.
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "List of tags to tag the schema",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "description": "A key to identify the tag.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "Corresponding tag value for the key.",
			//         "maxLength": 256,
			//         "minLength": 0,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 10,
			//   "minItems": 0,
			//   "type": "array"
			// }
			Description: "List of tags to tag the schema",
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "A key to identify the tag.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(1, 128),
						},
					},
					"value": {
						// Property: Value
						Description: "Corresponding tag value for the key.",
						Type:        types.StringType,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.StringLenBetween(0, 256),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 10),
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
		Description: "This resource represents a schema of Glue Schema Registry.",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::Glue::Schema").WithTerraformTypeName("awscc_glue_schema")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"arn":                       "Arn",
		"checkpoint_version":        "CheckpointVersion",
		"compatibility":             "Compatibility",
		"data_format":               "DataFormat",
		"description":               "Description",
		"initial_schema_version_id": "InitialSchemaVersionId",
		"is_latest":                 "IsLatest",
		"key":                       "Key",
		"name":                      "Name",
		"registry":                  "Registry",
		"schema_definition":         "SchemaDefinition",
		"tags":                      "Tags",
		"value":                     "Value",
		"version_number":            "VersionNumber",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/SchemaDefinition",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
