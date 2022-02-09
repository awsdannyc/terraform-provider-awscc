// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package greengrassv2

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_greengrassv2_component_version", componentVersionDataSourceType)
}

// componentVersionDataSourceType returns the Terraform awscc_greengrassv2_component_version data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::GreengrassV2::ComponentVersion resource type.
func componentVersionDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"arn": {
			// Property: Arn
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"component_name": {
			// Property: ComponentName
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"component_version": {
			// Property: ComponentVersion
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"inline_recipe": {
			// Property: InlineRecipe
			// CloudFormation resource type schema:
			// {
			//   "type": "string"
			// }
			Type:     types.StringType,
			Computed: true,
		},
		"lambda_function": {
			// Property: LambdaFunction
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "properties": {
			//     "ComponentDependencies": {
			//       "additionalProperties": false,
			//       "patternProperties": {
			//         "": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "DependencyType": {
			//               "enum": [
			//                 "SOFT",
			//                 "HARD"
			//               ],
			//               "type": "string"
			//             },
			//             "VersionRequirement": {
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "ComponentLambdaParameters": {
			//       "additionalProperties": false,
			//       "properties": {
			//         "EnvironmentVariables": {
			//           "additionalProperties": false,
			//           "patternProperties": {
			//             "": {
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "EventSources": {
			//           "insertionOrder": false,
			//           "items": {
			//             "additionalProperties": false,
			//             "properties": {
			//               "Topic": {
			//                 "type": "string"
			//               },
			//               "Type": {
			//                 "enum": [
			//                   "PUB_SUB",
			//                   "IOT_CORE"
			//                 ],
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "type": "array"
			//         },
			//         "ExecArgs": {
			//           "insertionOrder": true,
			//           "items": {
			//             "type": "string"
			//           },
			//           "type": "array"
			//         },
			//         "InputPayloadEncodingType": {
			//           "enum": [
			//             "json",
			//             "binary"
			//           ],
			//           "type": "string"
			//         },
			//         "LinuxProcessParams": {
			//           "additionalProperties": false,
			//           "properties": {
			//             "ContainerParams": {
			//               "additionalProperties": false,
			//               "properties": {
			//                 "Devices": {
			//                   "insertionOrder": false,
			//                   "items": {
			//                     "additionalProperties": false,
			//                     "properties": {
			//                       "AddGroupOwner": {
			//                         "type": "boolean"
			//                       },
			//                       "Path": {
			//                         "type": "string"
			//                       },
			//                       "Permission": {
			//                         "enum": [
			//                           "ro",
			//                           "rw"
			//                         ],
			//                         "type": "string"
			//                       }
			//                     },
			//                     "type": "object"
			//                   },
			//                   "type": "array"
			//                 },
			//                 "MemorySizeInKB": {
			//                   "type": "integer"
			//                 },
			//                 "MountROSysfs": {
			//                   "type": "boolean"
			//                 },
			//                 "Volumes": {
			//                   "insertionOrder": false,
			//                   "items": {
			//                     "additionalProperties": false,
			//                     "properties": {
			//                       "AddGroupOwner": {
			//                         "type": "boolean"
			//                       },
			//                       "DestinationPath": {
			//                         "type": "string"
			//                       },
			//                       "Permission": {
			//                         "enum": [
			//                           "ro",
			//                           "rw"
			//                         ],
			//                         "type": "string"
			//                       },
			//                       "SourcePath": {
			//                         "type": "string"
			//                       }
			//                     },
			//                     "type": "object"
			//                   },
			//                   "type": "array"
			//                 }
			//               },
			//               "type": "object"
			//             },
			//             "IsolationMode": {
			//               "enum": [
			//                 "GreengrassContainer",
			//                 "NoContainer"
			//               ],
			//               "type": "string"
			//             }
			//           },
			//           "type": "object"
			//         },
			//         "MaxIdleTimeInSeconds": {
			//           "type": "integer"
			//         },
			//         "MaxInstancesCount": {
			//           "type": "integer"
			//         },
			//         "MaxQueueSize": {
			//           "type": "integer"
			//         },
			//         "Pinned": {
			//           "type": "boolean"
			//         },
			//         "StatusTimeoutInSeconds": {
			//           "type": "integer"
			//         },
			//         "TimeoutInSeconds": {
			//           "type": "integer"
			//         }
			//       },
			//       "type": "object"
			//     },
			//     "ComponentName": {
			//       "type": "string"
			//     },
			//     "ComponentPlatforms": {
			//       "insertionOrder": false,
			//       "items": {
			//         "additionalProperties": false,
			//         "properties": {
			//           "Attributes": {
			//             "additionalProperties": false,
			//             "patternProperties": {
			//               "": {
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "Name": {
			//             "type": "string"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "type": "array"
			//     },
			//     "ComponentVersion": {
			//       "type": "string"
			//     },
			//     "LambdaArn": {
			//       "pattern": "",
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			Attributes: tfsdk.SingleNestedAttributes(
				map[string]tfsdk.Attribute{
					"component_dependencies": {
						// Property: ComponentDependencies
						// Pattern: ""
						Attributes: tfsdk.MapNestedAttributes(
							map[string]tfsdk.Attribute{
								"dependency_type": {
									// Property: DependencyType
									Type:     types.StringType,
									Computed: true,
								},
								"version_requirement": {
									// Property: VersionRequirement
									Type:     types.StringType,
									Computed: true,
								},
							},
							tfsdk.MapNestedAttributesOptions{},
						),
						Computed: true,
					},
					"component_lambda_parameters": {
						// Property: ComponentLambdaParameters
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"environment_variables": {
									// Property: EnvironmentVariables
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Computed: true,
								},
								"event_sources": {
									// Property: EventSources
									Attributes: tfsdk.ListNestedAttributes(
										map[string]tfsdk.Attribute{
											"topic": {
												// Property: Topic
												Type:     types.StringType,
												Computed: true,
											},
											"type": {
												// Property: Type
												Type:     types.StringType,
												Computed: true,
											},
										},
										tfsdk.ListNestedAttributesOptions{},
									),
									Computed: true,
								},
								"exec_args": {
									// Property: ExecArgs
									Type:     types.ListType{ElemType: types.StringType},
									Computed: true,
								},
								"input_payload_encoding_type": {
									// Property: InputPayloadEncodingType
									Type:     types.StringType,
									Computed: true,
								},
								"linux_process_params": {
									// Property: LinuxProcessParams
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"container_params": {
												// Property: ContainerParams
												Attributes: tfsdk.SingleNestedAttributes(
													map[string]tfsdk.Attribute{
														"devices": {
															// Property: Devices
															Attributes: tfsdk.ListNestedAttributes(
																map[string]tfsdk.Attribute{
																	"add_group_owner": {
																		// Property: AddGroupOwner
																		Type:     types.BoolType,
																		Computed: true,
																	},
																	"path": {
																		// Property: Path
																		Type:     types.StringType,
																		Computed: true,
																	},
																	"permission": {
																		// Property: Permission
																		Type:     types.StringType,
																		Computed: true,
																	},
																},
																tfsdk.ListNestedAttributesOptions{},
															),
															Computed: true,
														},
														"memory_size_in_kb": {
															// Property: MemorySizeInKB
															Type:     types.Int64Type,
															Computed: true,
														},
														"mount_ro_sysfs": {
															// Property: MountROSysfs
															Type:     types.BoolType,
															Computed: true,
														},
														"volumes": {
															// Property: Volumes
															Attributes: tfsdk.ListNestedAttributes(
																map[string]tfsdk.Attribute{
																	"add_group_owner": {
																		// Property: AddGroupOwner
																		Type:     types.BoolType,
																		Computed: true,
																	},
																	"destination_path": {
																		// Property: DestinationPath
																		Type:     types.StringType,
																		Computed: true,
																	},
																	"permission": {
																		// Property: Permission
																		Type:     types.StringType,
																		Computed: true,
																	},
																	"source_path": {
																		// Property: SourcePath
																		Type:     types.StringType,
																		Computed: true,
																	},
																},
																tfsdk.ListNestedAttributesOptions{},
															),
															Computed: true,
														},
													},
												),
												Computed: true,
											},
											"isolation_mode": {
												// Property: IsolationMode
												Type:     types.StringType,
												Computed: true,
											},
										},
									),
									Computed: true,
								},
								"max_idle_time_in_seconds": {
									// Property: MaxIdleTimeInSeconds
									Type:     types.Int64Type,
									Computed: true,
								},
								"max_instances_count": {
									// Property: MaxInstancesCount
									Type:     types.Int64Type,
									Computed: true,
								},
								"max_queue_size": {
									// Property: MaxQueueSize
									Type:     types.Int64Type,
									Computed: true,
								},
								"pinned": {
									// Property: Pinned
									Type:     types.BoolType,
									Computed: true,
								},
								"status_timeout_in_seconds": {
									// Property: StatusTimeoutInSeconds
									Type:     types.Int64Type,
									Computed: true,
								},
								"timeout_in_seconds": {
									// Property: TimeoutInSeconds
									Type:     types.Int64Type,
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"component_name": {
						// Property: ComponentName
						Type:     types.StringType,
						Computed: true,
					},
					"component_platforms": {
						// Property: ComponentPlatforms
						Attributes: tfsdk.ListNestedAttributes(
							map[string]tfsdk.Attribute{
								"attributes": {
									// Property: Attributes
									// Pattern: ""
									Type:     types.MapType{ElemType: types.StringType},
									Computed: true,
								},
								"name": {
									// Property: Name
									Type:     types.StringType,
									Computed: true,
								},
							},
							tfsdk.ListNestedAttributesOptions{},
						),
						Computed: true,
					},
					"component_version": {
						// Property: ComponentVersion
						Type:     types.StringType,
						Computed: true,
					},
					"lambda_arn": {
						// Property: LambdaArn
						Type:     types.StringType,
						Computed: true,
					},
				},
			),
			Computed: true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "patternProperties": {
			//     "": {
			//       "maxLength": 256,
			//       "type": "string"
			//     }
			//   },
			//   "type": "object"
			// }
			// Pattern: ""
			Type:     types.MapType{ElemType: types.StringType},
			Computed: true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::GreengrassV2::ComponentVersion",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GreengrassV2::ComponentVersion").WithTerraformTypeName("awscc_greengrassv2_component_version")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"add_group_owner":             "AddGroupOwner",
		"arn":                         "Arn",
		"attributes":                  "Attributes",
		"component_dependencies":      "ComponentDependencies",
		"component_lambda_parameters": "ComponentLambdaParameters",
		"component_name":              "ComponentName",
		"component_platforms":         "ComponentPlatforms",
		"component_version":           "ComponentVersion",
		"container_params":            "ContainerParams",
		"dependency_type":             "DependencyType",
		"destination_path":            "DestinationPath",
		"devices":                     "Devices",
		"environment_variables":       "EnvironmentVariables",
		"event_sources":               "EventSources",
		"exec_args":                   "ExecArgs",
		"inline_recipe":               "InlineRecipe",
		"input_payload_encoding_type": "InputPayloadEncodingType",
		"isolation_mode":              "IsolationMode",
		"lambda_arn":                  "LambdaArn",
		"lambda_function":             "LambdaFunction",
		"linux_process_params":        "LinuxProcessParams",
		"max_idle_time_in_seconds":    "MaxIdleTimeInSeconds",
		"max_instances_count":         "MaxInstancesCount",
		"max_queue_size":              "MaxQueueSize",
		"memory_size_in_kb":           "MemorySizeInKB",
		"mount_ro_sysfs":              "MountROSysfs",
		"name":                        "Name",
		"path":                        "Path",
		"permission":                  "Permission",
		"pinned":                      "Pinned",
		"source_path":                 "SourcePath",
		"status_timeout_in_seconds":   "StatusTimeoutInSeconds",
		"tags":                        "Tags",
		"timeout_in_seconds":          "TimeoutInSeconds",
		"topic":                       "Topic",
		"type":                        "Type",
		"version_requirement":         "VersionRequirement",
		"volumes":                     "Volumes",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
