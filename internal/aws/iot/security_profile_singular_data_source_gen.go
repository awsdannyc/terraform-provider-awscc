// Code generated by generators/singular-data-source/main.go; DO NOT EDIT.

package iot

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddDataSourceTypeFactory("awscc_iot_security_profile", securityProfileDataSourceType)
}

// securityProfileDataSourceType returns the Terraform awscc_iot_security_profile data source type.
// This Terraform data source type corresponds to the CloudFormation AWS::IoT::SecurityProfile resource type.
func securityProfileDataSourceType(ctx context.Context) (tfsdk.DataSourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"additional_metrics_to_retain_v2": {
			// Property: AdditionalMetricsToRetainV2
			// CloudFormation resource type schema:
			// {
			//   "description": "A list of metrics whose data is retained (stored). By default, data is retained for any metric used in the profile's behaviors, but it is also retained for any metric specified here.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "The metric you want to retain. Dimensions are optional.",
			//     "properties": {
			//       "Metric": {
			//         "description": "What is measured by the behavior.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "MetricDimension": {
			//         "additionalProperties": false,
			//         "description": "The dimension of a metric.",
			//         "properties": {
			//           "DimensionName": {
			//             "description": "A unique identifier for the dimension.",
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Operator": {
			//             "description": "Defines how the dimensionValues of a dimension are interpreted.",
			//             "enum": [
			//               "IN",
			//               "NOT_IN"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "DimensionName"
			//         ],
			//         "type": "object"
			//       }
			//     },
			//     "required": [
			//       "Metric"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A list of metrics whose data is retained (stored). By default, data is retained for any metric used in the profile's behaviors, but it is also retained for any metric specified here.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"metric": {
						// Property: Metric
						Description: "What is measured by the behavior.",
						Type:        types.StringType,
						Computed:    true,
					},
					"metric_dimension": {
						// Property: MetricDimension
						Description: "The dimension of a metric.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"dimension_name": {
									// Property: DimensionName
									Description: "A unique identifier for the dimension.",
									Type:        types.StringType,
									Computed:    true,
								},
								"operator": {
									// Property: Operator
									Description: "Defines how the dimensionValues of a dimension are interpreted.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"alert_targets": {
			// Property: AlertTargets
			// CloudFormation resource type schema:
			// {
			//   "additionalProperties": false,
			//   "description": "Specifies the destinations to which alerts are sent.",
			//   "patternProperties": {
			//     "": {
			//       "additionalProperties": false,
			//       "description": "A structure containing the alert target ARN and the role ARN.",
			//       "properties": {
			//         "AlertTargetArn": {
			//           "description": "The ARN of the notification target to which alerts are sent.",
			//           "maxLength": 2048,
			//           "type": "string"
			//         },
			//         "RoleArn": {
			//           "description": "The ARN of the role that grants permission to send alerts to the notification target.",
			//           "maxLength": 2048,
			//           "minLength": 20,
			//           "type": "string"
			//         }
			//       },
			//       "required": [
			//         "AlertTargetArn",
			//         "RoleArn"
			//       ],
			//       "type": "object"
			//     }
			//   },
			//   "type": "object"
			// }
			Description: "Specifies the destinations to which alerts are sent.",
			// Pattern: ""
			Attributes: tfsdk.MapNestedAttributes(
				map[string]tfsdk.Attribute{
					"alert_target_arn": {
						// Property: AlertTargetArn
						Description: "The ARN of the notification target to which alerts are sent.",
						Type:        types.StringType,
						Computed:    true,
					},
					"role_arn": {
						// Property: RoleArn
						Description: "The ARN of the role that grants permission to send alerts to the notification target.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.MapNestedAttributesOptions{},
			),
			Computed: true,
		},
		"behaviors": {
			// Property: Behaviors
			// CloudFormation resource type schema:
			// {
			//   "description": "Specifies the behaviors that, when violated by a device (thing), cause an alert.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A security profile behavior.",
			//     "properties": {
			//       "Criteria": {
			//         "additionalProperties": false,
			//         "description": "The criteria by which the behavior is determined to be normal.",
			//         "properties": {
			//           "ComparisonOperator": {
			//             "description": "The operator that relates the thing measured (metric) to the criteria (containing a value or statisticalThreshold).",
			//             "enum": [
			//               "less-than",
			//               "less-than-equals",
			//               "greater-than",
			//               "greater-than-equals",
			//               "in-cidr-set",
			//               "not-in-cidr-set",
			//               "in-port-set",
			//               "not-in-port-set",
			//               "in-set",
			//               "not-in-set"
			//             ],
			//             "type": "string"
			//           },
			//           "ConsecutiveDatapointsToAlarm": {
			//             "description": "If a device is in violation of the behavior for the specified number of consecutive datapoints, an alarm occurs. If not specified, the default is 1.",
			//             "maximum": 10,
			//             "minimum": 1,
			//             "type": "integer"
			//           },
			//           "ConsecutiveDatapointsToClear": {
			//             "description": "If an alarm has occurred and the offending device is no longer in violation of the behavior for the specified number of consecutive datapoints, the alarm is cleared. If not specified, the default is 1.",
			//             "maximum": 10,
			//             "minimum": 1,
			//             "type": "integer"
			//           },
			//           "DurationSeconds": {
			//             "description": "Use this to specify the time duration over which the behavior is evaluated.",
			//             "type": "integer"
			//           },
			//           "MlDetectionConfig": {
			//             "additionalProperties": false,
			//             "description": "The configuration of an ML Detect Security Profile.",
			//             "properties": {
			//               "ConfidenceLevel": {
			//                 "description": "The sensitivity of anomalous behavior evaluation. Can be Low, Medium, or High.",
			//                 "enum": [
			//                   "LOW",
			//                   "MEDIUM",
			//                   "HIGH"
			//                 ],
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "StatisticalThreshold": {
			//             "additionalProperties": false,
			//             "description": "A statistical ranking (percentile) which indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior.",
			//             "properties": {
			//               "Statistic": {
			//                 "description": "The percentile which resolves to a threshold value by which compliance with a behavior is determined",
			//                 "enum": [
			//                   "Average",
			//                   "p0",
			//                   "p0.1",
			//                   "p0.01",
			//                   "p1",
			//                   "p10",
			//                   "p50",
			//                   "p90",
			//                   "p99",
			//                   "p99.9",
			//                   "p99.99",
			//                   "p100"
			//                 ],
			//                 "type": "string"
			//               }
			//             },
			//             "type": "object"
			//           },
			//           "Value": {
			//             "additionalProperties": false,
			//             "description": "The value to be compared with the metric.",
			//             "properties": {
			//               "Cidrs": {
			//                 "description": "If the ComparisonOperator calls for a set of CIDRs, use this to specify that set to be compared with the metric.",
			//                 "insertionOrder": false,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array",
			//                 "uniqueItems": true
			//               },
			//               "Count": {
			//                 "description": "If the ComparisonOperator calls for a numeric value, use this to specify that (integer) numeric value to be compared with the metric.",
			//                 "minimum": 0,
			//                 "type": "string"
			//               },
			//               "Number": {
			//                 "description": "The numeral value of a metric.",
			//                 "type": "number"
			//               },
			//               "Numbers": {
			//                 "description": "The numeral values of a metric.",
			//                 "insertionOrder": false,
			//                 "items": {
			//                   "type": "number"
			//                 },
			//                 "type": "array",
			//                 "uniqueItems": true
			//               },
			//               "Ports": {
			//                 "description": "If the ComparisonOperator calls for a set of ports, use this to specify that set to be compared with the metric.",
			//                 "insertionOrder": false,
			//                 "items": {
			//                   "maximum": 65535,
			//                   "minimum": 0,
			//                   "type": "integer"
			//                 },
			//                 "type": "array",
			//                 "uniqueItems": true
			//               },
			//               "Strings": {
			//                 "description": "The string values of a metric.",
			//                 "insertionOrder": false,
			//                 "items": {
			//                   "type": "string"
			//                 },
			//                 "type": "array",
			//                 "uniqueItems": true
			//               }
			//             },
			//             "type": "object"
			//           }
			//         },
			//         "type": "object"
			//       },
			//       "Metric": {
			//         "description": "What is measured by the behavior.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "MetricDimension": {
			//         "additionalProperties": false,
			//         "description": "The dimension of a metric.",
			//         "properties": {
			//           "DimensionName": {
			//             "description": "A unique identifier for the dimension.",
			//             "maxLength": 128,
			//             "minLength": 1,
			//             "pattern": "",
			//             "type": "string"
			//           },
			//           "Operator": {
			//             "description": "Defines how the dimensionValues of a dimension are interpreted.",
			//             "enum": [
			//               "IN",
			//               "NOT_IN"
			//             ],
			//             "type": "string"
			//           }
			//         },
			//         "required": [
			//           "DimensionName"
			//         ],
			//         "type": "object"
			//       },
			//       "Name": {
			//         "description": "The name for the behavior.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "pattern": "",
			//         "type": "string"
			//       },
			//       "SuppressAlerts": {
			//         "description": "Manage Detect alarm SNS notifications by setting behavior notification to on or suppressed. Detect will continue to performing device behavior evaluations. However, suppressed alarms wouldn't be forwarded for SNS notification.",
			//         "type": "boolean"
			//       }
			//     },
			//     "required": [
			//       "Name"
			//     ],
			//     "type": "object"
			//   },
			//   "maxLength": 100,
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Specifies the behaviors that, when violated by a device (thing), cause an alert.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"criteria": {
						// Property: Criteria
						Description: "The criteria by which the behavior is determined to be normal.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"comparison_operator": {
									// Property: ComparisonOperator
									Description: "The operator that relates the thing measured (metric) to the criteria (containing a value or statisticalThreshold).",
									Type:        types.StringType,
									Computed:    true,
								},
								"consecutive_datapoints_to_alarm": {
									// Property: ConsecutiveDatapointsToAlarm
									Description: "If a device is in violation of the behavior for the specified number of consecutive datapoints, an alarm occurs. If not specified, the default is 1.",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"consecutive_datapoints_to_clear": {
									// Property: ConsecutiveDatapointsToClear
									Description: "If an alarm has occurred and the offending device is no longer in violation of the behavior for the specified number of consecutive datapoints, the alarm is cleared. If not specified, the default is 1.",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"duration_seconds": {
									// Property: DurationSeconds
									Description: "Use this to specify the time duration over which the behavior is evaluated.",
									Type:        types.Int64Type,
									Computed:    true,
								},
								"ml_detection_config": {
									// Property: MlDetectionConfig
									Description: "The configuration of an ML Detect Security Profile.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"confidence_level": {
												// Property: ConfidenceLevel
												Description: "The sensitivity of anomalous behavior evaluation. Can be Low, Medium, or High.",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"statistical_threshold": {
									// Property: StatisticalThreshold
									Description: "A statistical ranking (percentile) which indicates a threshold value by which a behavior is determined to be in compliance or in violation of the behavior.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"statistic": {
												// Property: Statistic
												Description: "The percentile which resolves to a threshold value by which compliance with a behavior is determined",
												Type:        types.StringType,
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
								"value": {
									// Property: Value
									Description: "The value to be compared with the metric.",
									Attributes: tfsdk.SingleNestedAttributes(
										map[string]tfsdk.Attribute{
											"cidrs": {
												// Property: Cidrs
												Description: "If the ComparisonOperator calls for a set of CIDRs, use this to specify that set to be compared with the metric.",
												Type:        types.SetType{ElemType: types.StringType},
												Computed:    true,
											},
											"count": {
												// Property: Count
												Description: "If the ComparisonOperator calls for a numeric value, use this to specify that (integer) numeric value to be compared with the metric.",
												Type:        types.StringType,
												Computed:    true,
											},
											"number": {
												// Property: Number
												Description: "The numeral value of a metric.",
												Type:        types.Float64Type,
												Computed:    true,
											},
											"numbers": {
												// Property: Numbers
												Description: "The numeral values of a metric.",
												Type:        types.SetType{ElemType: types.Float64Type},
												Computed:    true,
											},
											"ports": {
												// Property: Ports
												Description: "If the ComparisonOperator calls for a set of ports, use this to specify that set to be compared with the metric.",
												Type:        types.SetType{ElemType: types.Int64Type},
												Computed:    true,
											},
											"strings": {
												// Property: Strings
												Description: "The string values of a metric.",
												Type:        types.SetType{ElemType: types.StringType},
												Computed:    true,
											},
										},
									),
									Computed: true,
								},
							},
						),
						Computed: true,
					},
					"metric": {
						// Property: Metric
						Description: "What is measured by the behavior.",
						Type:        types.StringType,
						Computed:    true,
					},
					"metric_dimension": {
						// Property: MetricDimension
						Description: "The dimension of a metric.",
						Attributes: tfsdk.SingleNestedAttributes(
							map[string]tfsdk.Attribute{
								"dimension_name": {
									// Property: DimensionName
									Description: "A unique identifier for the dimension.",
									Type:        types.StringType,
									Computed:    true,
								},
								"operator": {
									// Property: Operator
									Description: "Defines how the dimensionValues of a dimension are interpreted.",
									Type:        types.StringType,
									Computed:    true,
								},
							},
						),
						Computed: true,
					},
					"name": {
						// Property: Name
						Description: "The name for the behavior.",
						Type:        types.StringType,
						Computed:    true,
					},
					"suppress_alerts": {
						// Property: SuppressAlerts
						Description: "Manage Detect alarm SNS notifications by setting behavior notification to on or suppressed. Detect will continue to performing device behavior evaluations. However, suppressed alarms wouldn't be forwarded for SNS notification.",
						Type:        types.BoolType,
						Computed:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"security_profile_arn": {
			// Property: SecurityProfileArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The ARN (Amazon resource name) of the created security profile.",
			//   "type": "string"
			// }
			Description: "The ARN (Amazon resource name) of the created security profile.",
			Type:        types.StringType,
			Computed:    true,
		},
		"security_profile_description": {
			// Property: SecurityProfileDescription
			// CloudFormation resource type schema:
			// {
			//   "description": "A description of the security profile.",
			//   "maxLength": 1000,
			//   "type": "string"
			// }
			Description: "A description of the security profile.",
			Type:        types.StringType,
			Computed:    true,
		},
		"security_profile_name": {
			// Property: SecurityProfileName
			// CloudFormation resource type schema:
			// {
			//   "description": "A unique identifier for the security profile.",
			//   "maxLength": 128,
			//   "minLength": 1,
			//   "pattern": "",
			//   "type": "string"
			// }
			Description: "A unique identifier for the security profile.",
			Type:        types.StringType,
			Computed:    true,
		},
		"tags": {
			// Property: Tags
			// CloudFormation resource type schema:
			// {
			//   "description": "Metadata that can be used to manage the security profile.",
			//   "insertionOrder": false,
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A key-value pair to associate with a resource.",
			//     "properties": {
			//       "Key": {
			//         "description": "The tag's key.",
			//         "maxLength": 128,
			//         "minLength": 1,
			//         "type": "string"
			//       },
			//       "Value": {
			//         "description": "The tag's value.",
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
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "Metadata that can be used to manage the security profile.",
			Attributes: tfsdk.SetNestedAttributes(
				map[string]tfsdk.Attribute{
					"key": {
						// Property: Key
						Description: "The tag's key.",
						Type:        types.StringType,
						Computed:    true,
					},
					"value": {
						// Property: Value
						Description: "The tag's value.",
						Type:        types.StringType,
						Computed:    true,
					},
				},
				tfsdk.SetNestedAttributesOptions{},
			),
			Computed: true,
		},
		"target_arns": {
			// Property: TargetArns
			// CloudFormation resource type schema:
			// {
			//   "description": "A set of target ARNs that the security profile is attached to.",
			//   "insertionOrder": false,
			//   "items": {
			//     "description": "The ARN of the target to which the security profile is attached.",
			//     "maxLength": 2048,
			//     "type": "string"
			//   },
			//   "type": "array",
			//   "uniqueItems": true
			// }
			Description: "A set of target ARNs that the security profile is attached to.",
			Type:        types.SetType{ElemType: types.StringType},
			Computed:    true,
		},
	}

	attributes["id"] = tfsdk.Attribute{
		Description: "Uniquely identifies the resource.",
		Type:        types.StringType,
		Required:    true,
	}

	schema := tfsdk.Schema{
		Description: "Data Source schema for AWS::IoT::SecurityProfile",
		Version:     1,
		Attributes:  attributes,
	}

	var opts DataSourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::IoT::SecurityProfile").WithTerraformTypeName("awscc_iot_security_profile")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"additional_metrics_to_retain_v2": "AdditionalMetricsToRetainV2",
		"alert_target_arn":                "AlertTargetArn",
		"alert_targets":                   "AlertTargets",
		"behaviors":                       "Behaviors",
		"cidrs":                           "Cidrs",
		"comparison_operator":             "ComparisonOperator",
		"confidence_level":                "ConfidenceLevel",
		"consecutive_datapoints_to_alarm": "ConsecutiveDatapointsToAlarm",
		"consecutive_datapoints_to_clear": "ConsecutiveDatapointsToClear",
		"count":                           "Count",
		"criteria":                        "Criteria",
		"dimension_name":                  "DimensionName",
		"duration_seconds":                "DurationSeconds",
		"key":                             "Key",
		"metric":                          "Metric",
		"metric_dimension":                "MetricDimension",
		"ml_detection_config":             "MlDetectionConfig",
		"name":                            "Name",
		"number":                          "Number",
		"numbers":                         "Numbers",
		"operator":                        "Operator",
		"ports":                           "Ports",
		"role_arn":                        "RoleArn",
		"security_profile_arn":            "SecurityProfileArn",
		"security_profile_description":    "SecurityProfileDescription",
		"security_profile_name":           "SecurityProfileName",
		"statistic":                       "Statistic",
		"statistical_threshold":           "StatisticalThreshold",
		"strings":                         "Strings",
		"suppress_alerts":                 "SuppressAlerts",
		"tags":                            "Tags",
		"target_arns":                     "TargetArns",
		"value":                           "Value",
	})

	singularDataSourceType, err := NewSingularDataSourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return singularDataSourceType, nil
}
