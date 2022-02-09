// Code generated by generators/resource/main.go; DO NOT EDIT.

package sagemaker

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_sagemaker_app_image_config", appImageConfigResourceType)
}

// appImageConfigResourceType returns the Terraform awscc_sagemaker_app_image_config resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::SageMaker::AppImageConfig resource type.
func appImageConfigResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"app_image_config_arn": {
			// Property: AppImageConfigArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the AppImageConfig.",
			//   "maxLength": 256,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the AppImageConfig.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"app_image_config_name": {
			// Property: AppImageConfigName
			// CloudFormation resource type schema:
			// {
			//   "description": "The Name of the AppImageConfig.",
			//   "maxLength": 63,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "The Name of the AppImageConfig.",
			Type:        types.StringType,
			Required:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringLenBetween(1, 63),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"kernel_gateway_image_config": {
			// Property: KernelGatewayImageConfig
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "The KernelGatewayImageConfig.",
			//   "properties": {
			//     "FileSystemConfig": {
			//       "additionalProperties": false,
			//       "description": "The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.",
			//       "properties": {
			//         "DefaultGid": {
			//           "description": "The default POSIX group ID (GID). If not specified, defaults to 100.",
			//           "maximum": 65535,
			//           "minimum": 0,
			//           "type": "integer"
			//         },
			//         "DefaultUid": {
			//           "description": "The default POSIX user ID (UID). If not specified, defaults to 1000.",
			//           "maximum": 65535,
			//           "minimum": 0,
			//           "type": "integer"
			//         },
			//         "MountPath": {
			//           "description": "The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.",
			//           "maxLength": 1024,
			//           "minLength": 1,
			//           "pattern": "",
			//           "type": "string"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "KernelSpecs": {
			//       "description": "The specification of the Jupyter kernels in the image.",
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "DisplayName": {
			//             "description": "The display name of the kernel.",
			//             "maxLength": 1024,
			//             "minLength": 1,
			//             "type": "string"
			//           },
			//           "Name": {
			//             "description": "The name of the kernel.",
			//             "maxLength": 1024,
			//             "minLength": 1,
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "Name"
			//         ],
			//         "type": "object"
			//       },
			//       "maxItems": 1,
			//       "minItems": 1,
			//       "type": "array"
			//     }
			//   },
			//   "required": [
			//     "KernelSpecs"
			//   ],
			//   "type": "object"
			// }
			Description: "The KernelGatewayImageConfig.",
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"file_system_config": {
						// Property: FileSystemConfig
						Description: "The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"default_gid": {
									// Property: DefaultGid
									Description: "The default POSIX group ID (GID). If not specified, defaults to 100.",
									Type:        types.Int64Type,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(0, 65535),
									},
								},
								"default_uid": {
									// Property: DefaultUid
									Description: "The default POSIX user ID (UID). If not specified, defaults to 1000.",
									Type:        types.Int64Type,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.IntBetween(0, 65535),
									},
								},
								"mount_path": {
									// Property: MountPath
									Description: "The path within the image to mount the user's EFS home directory. The directory should be empty. If not specified, defaults to /home/sagemaker-user.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 1024),
									},
								},
							},
						),
						Optional: true,
					},
					"kernel_specs": {
						// Property: KernelSpecs
						Description: "The specification of the Jupyter kernels in the image.",
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"display_name": {
									// Property: DisplayName
									Description: "The display name of the kernel.",
									Type:        types.StringType,
									Optional:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 1024),
									},
								},
								"name": {
									// Property: Name
									Description: "The name of the kernel.",
									Type:        types.StringType,
									Required:    true,
									Validators: []tfsdk.AttributeValidator{
										validate.StringLenBetween(1, 1024),
									},
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Required: true,
						Validators: []tfsdk.AttributeValidator{
							validate.ArrayLenBetween(1, 1),
						},
					},
				},
			),
			Optional: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of tags to apply to the AppImageConfig.",
			//   "items": {
			//     "additionalProperties": false,
			//     "properties": {
			//       "Key": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       }
			//     },
			//     "required": [
			//       "Key",
			//       "Value"
			//     ],
			//     "type": "object"
			//   },
			//   "maxItems": 50,
			//   "minItems": 0,
			//   "type": "array",
			//   "uniqueItems": false
			// }
			Description: "A list of tags to apply to the AppImageConfig.",
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
							validate.StringLenBetween(1, 128),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Optional: true,
			Computed: true,
			Validators: []tfsdk.AttributeValidator{
				validate.ArrayLenBetween(0, 50),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
				tfsdk.RequiresReplace(),
			},
			// Tags is a write-only property.
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
		Description: "Resource Type definition for AWS::SageMaker::AppImageConfig",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::SageMaker::AppImageConfig").WithTerraformTypeName("awscc_sagemaker_app_image_config")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"app_image_config_arn":        "AppImageConfigArn",
		"app_image_config_name":       "AppImageConfigName",
		"default_gid":                 "DefaultGid",
		"default_uid":                 "DefaultUid",
		"display_name":                "DisplayName",
		"file_system_config":          "FileSystemConfig",
		"kernel_gateway_image_config": "KernelGatewayImageConfig",
		"kernel_specs":                "KernelSpecs",
		"key":                         "Key",
		"mount_path":                  "MountPath",
		"name":                        "Name",
		"tags":                        "Tags",
		"value":                       "Value",
	})

	opts = opts.WithWriteOnlyPropertyPaths([]string{
		"/properties/Tags",
	})
	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
