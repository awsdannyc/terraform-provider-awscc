// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package cleanrooms

import (
	"context"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_cleanrooms_configured_table_association", configuredTableAssociationResource)
}

// configuredTableAssociationResource returns the Terraform awscc_cleanrooms_configured_table_association resource.
// This Terraform resource corresponds to the CloudFormation AWS::CleanRooms::ConfiguredTableAssociation resource.
func configuredTableAssociationResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: Arn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 100,
		//	  "type": "string"
		//	}
		"arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConfiguredTableAssociationAnalysisRules
		// CloudFormation resource type schema:
		//
		//	{
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Policy": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "V1": {
		//	            "properties": {
		//	              "Aggregation": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "AllowedAdditionalAnalyses": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "maxLength": 256,
		//	                      "pattern": "^arn:aws:cleanrooms:[\\w]{2}-[\\w]{4,9}-[\\d]:([\\d]{12}|\\*):membership\\/[\\*\\d\\w-]+\\/configuredaudiencemodelassociation\\/[\\*\\d\\w-]+$|^arn:aws[-a-z]*:cleanrooms-ml:[-a-z0-9]+:([0-9]{12}|\\*):configured-model-algorithm-association\\/([-a-zA-Z0-9_\\/.]+|\\*)$",
		//	                      "type": "string"
		//	                    },
		//	                    "maxItems": 25,
		//	                    "minItems": 0,
		//	                    "type": "array"
		//	                  },
		//	                  "AllowedResultReceivers": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "maxLength": 12,
		//	                      "minLength": 12,
		//	                      "pattern": "\\d+",
		//	                      "type": "string"
		//	                    },
		//	                    "minItems": 0,
		//	                    "type": "array"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              },
		//	              "Custom": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "AllowedAdditionalAnalyses": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "maxLength": 256,
		//	                      "pattern": "^arn:aws:cleanrooms:[\\w]{2}-[\\w]{4,9}-[\\d]:([\\d]{12}|\\*):membership\\/[\\*\\d\\w-]+\\/configuredaudiencemodelassociation\\/[\\*\\d\\w-]+$|^arn:aws[-a-z]*:cleanrooms-ml:[-a-z0-9]+:([0-9]{12}|\\*):configured-model-algorithm-association\\/([-a-zA-Z0-9_\\/.]+|\\*)$",
		//	                      "type": "string"
		//	                    },
		//	                    "maxItems": 25,
		//	                    "minItems": 0,
		//	                    "type": "array"
		//	                  },
		//	                  "AllowedResultReceivers": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "maxLength": 12,
		//	                      "minLength": 12,
		//	                      "pattern": "\\d+",
		//	                      "type": "string"
		//	                    },
		//	                    "minItems": 0,
		//	                    "type": "array"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              },
		//	              "List": {
		//	                "additionalProperties": false,
		//	                "properties": {
		//	                  "AllowedAdditionalAnalyses": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "maxLength": 256,
		//	                      "pattern": "^arn:aws:cleanrooms:[\\w]{2}-[\\w]{4,9}-[\\d]:([\\d]{12}|\\*):membership\\/[\\*\\d\\w-]+\\/configuredaudiencemodelassociation\\/[\\*\\d\\w-]+$|^arn:aws[-a-z]*:cleanrooms-ml:[-a-z0-9]+:([0-9]{12}|\\*):configured-model-algorithm-association\\/([-a-zA-Z0-9_\\/.]+|\\*)$",
		//	                      "type": "string"
		//	                    },
		//	                    "maxItems": 25,
		//	                    "minItems": 0,
		//	                    "type": "array"
		//	                  },
		//	                  "AllowedResultReceivers": {
		//	                    "insertionOrder": false,
		//	                    "items": {
		//	                      "maxLength": 12,
		//	                      "minLength": 12,
		//	                      "pattern": "\\d+",
		//	                      "type": "string"
		//	                    },
		//	                    "minItems": 0,
		//	                    "type": "array"
		//	                  }
		//	                },
		//	                "type": "object"
		//	              }
		//	            },
		//	            "type": "object"
		//	          }
		//	        },
		//	        "required": [
		//	          "V1"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Type": {
		//	        "enum": [
		//	          "AGGREGATION",
		//	          "LIST",
		//	          "CUSTOM"
		//	        ],
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Type",
		//	      "Policy"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1,
		//	  "minItems": 1,
		//	  "type": "array"
		//	}
		"configured_table_association_analysis_rules": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Policy
					"policy": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: V1
							"v1": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
								Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
									// Property: Aggregation
									"aggregation": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: AllowedAdditionalAnalyses
											"allowed_additional_analyses": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Optional:    true,
												Computed:    true,
												Validators: []validator.List{ /*START VALIDATORS*/
													listvalidator.SizeBetween(0, 25),
													listvalidator.ValueStringsAre(
														stringvalidator.LengthAtMost(256),
														stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws:cleanrooms:[\\w]{2}-[\\w]{4,9}-[\\d]:([\\d]{12}|\\*):membership\\/[\\*\\d\\w-]+\\/configuredaudiencemodelassociation\\/[\\*\\d\\w-]+$|^arn:aws[-a-z]*:cleanrooms-ml:[-a-z0-9]+:([0-9]{12}|\\*):configured-model-algorithm-association\\/([-a-zA-Z0-9_\\/.]+|\\*)$"), ""),
													),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
													generic.Multiset(),
													listplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: AllowedResultReceivers
											"allowed_result_receivers": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Optional:    true,
												Computed:    true,
												Validators: []validator.List{ /*START VALIDATORS*/
													listvalidator.SizeAtLeast(0),
													listvalidator.ValueStringsAre(
														stringvalidator.LengthBetween(12, 12),
														stringvalidator.RegexMatches(regexp.MustCompile("\\d+"), ""),
													),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
													generic.Multiset(),
													listplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
											objectplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: Custom
									"custom": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: AllowedAdditionalAnalyses
											"allowed_additional_analyses": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Optional:    true,
												Computed:    true,
												Validators: []validator.List{ /*START VALIDATORS*/
													listvalidator.SizeBetween(0, 25),
													listvalidator.ValueStringsAre(
														stringvalidator.LengthAtMost(256),
														stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws:cleanrooms:[\\w]{2}-[\\w]{4,9}-[\\d]:([\\d]{12}|\\*):membership\\/[\\*\\d\\w-]+\\/configuredaudiencemodelassociation\\/[\\*\\d\\w-]+$|^arn:aws[-a-z]*:cleanrooms-ml:[-a-z0-9]+:([0-9]{12}|\\*):configured-model-algorithm-association\\/([-a-zA-Z0-9_\\/.]+|\\*)$"), ""),
													),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
													generic.Multiset(),
													listplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: AllowedResultReceivers
											"allowed_result_receivers": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Optional:    true,
												Computed:    true,
												Validators: []validator.List{ /*START VALIDATORS*/
													listvalidator.SizeAtLeast(0),
													listvalidator.ValueStringsAre(
														stringvalidator.LengthBetween(12, 12),
														stringvalidator.RegexMatches(regexp.MustCompile("\\d+"), ""),
													),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
													generic.Multiset(),
													listplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
											objectplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
									// Property: List
									"list": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
										Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
											// Property: AllowedAdditionalAnalyses
											"allowed_additional_analyses": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Optional:    true,
												Computed:    true,
												Validators: []validator.List{ /*START VALIDATORS*/
													listvalidator.SizeBetween(0, 25),
													listvalidator.ValueStringsAre(
														stringvalidator.LengthAtMost(256),
														stringvalidator.RegexMatches(regexp.MustCompile("^arn:aws:cleanrooms:[\\w]{2}-[\\w]{4,9}-[\\d]:([\\d]{12}|\\*):membership\\/[\\*\\d\\w-]+\\/configuredaudiencemodelassociation\\/[\\*\\d\\w-]+$|^arn:aws[-a-z]*:cleanrooms-ml:[-a-z0-9]+:([0-9]{12}|\\*):configured-model-algorithm-association\\/([-a-zA-Z0-9_\\/.]+|\\*)$"), ""),
													),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
													generic.Multiset(),
													listplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
											// Property: AllowedResultReceivers
											"allowed_result_receivers": schema.ListAttribute{ /*START ATTRIBUTE*/
												ElementType: types.StringType,
												Optional:    true,
												Computed:    true,
												Validators: []validator.List{ /*START VALIDATORS*/
													listvalidator.SizeAtLeast(0),
													listvalidator.ValueStringsAre(
														stringvalidator.LengthBetween(12, 12),
														stringvalidator.RegexMatches(regexp.MustCompile("\\d+"), ""),
													),
												}, /*END VALIDATORS*/
												PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
													generic.Multiset(),
													listplanmodifier.UseStateForUnknown(),
												}, /*END PLAN MODIFIERS*/
											}, /*END ATTRIBUTE*/
										}, /*END SCHEMA*/
										Optional: true,
										Computed: true,
										PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
											objectplanmodifier.UseStateForUnknown(),
										}, /*END PLAN MODIFIERS*/
									}, /*END ATTRIBUTE*/
								}, /*END SCHEMA*/
								Required: true,
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Required: true,
					}, /*END ATTRIBUTE*/
					// Property: Type
					"type": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.OneOf(
								"AGGREGATION",
								"LIST",
								"CUSTOM",
							),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Optional: true,
			Computed: true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 1),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConfiguredTableAssociationIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}",
		//	  "type": "string"
		//	}
		"configured_table_association_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Computed: true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: ConfiguredTableIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}",
		//	  "type": "string"
		//	}
		"configured_table_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(36, 36),
				stringvalidator.RegexMatches(regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Description
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 255,
		//	  "pattern": "",
		//	  "type": "string"
		//	}
		"description": schema.StringAttribute{ /*START ATTRIBUTE*/
			Optional: true,
			Computed: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(255),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: MembershipIdentifier
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 36,
		//	  "minLength": 36,
		//	  "pattern": "[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}",
		//	  "type": "string"
		//	}
		"membership_identifier": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(36, 36),
				stringvalidator.RegexMatches(regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: Name
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 128,
		//	  "pattern": "^[a-zA-Z0-9_](([a-zA-Z0-9_ ]+-)*([a-zA-Z0-9_ ]+))?$",
		//	  "type": "string"
		//	}
		"name": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthAtMost(128),
				stringvalidator.RegexMatches(regexp.MustCompile("^[a-zA-Z0-9_](([a-zA-Z0-9_ ]+-)*([a-zA-Z0-9_ ]+))?$"), ""),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.RequiresReplace(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: RoleArn
		// CloudFormation resource type schema:
		//
		//	{
		//	  "maxLength": 512,
		//	  "minLength": 32,
		//	  "type": "string"
		//	}
		"role_arn": schema.StringAttribute{ /*START ATTRIBUTE*/
			Required: true,
			Validators: []validator.String{ /*START VALIDATORS*/
				stringvalidator.LengthBetween(32, 512),
			}, /*END VALIDATORS*/
		}, /*END ATTRIBUTE*/
		// Property: Tags
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.",
		//	  "insertionOrder": false,
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "properties": {
		//	      "Key": {
		//	        "maxLength": 128,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      },
		//	      "Value": {
		//	        "maxLength": 256,
		//	        "minLength": 1,
		//	        "type": "string"
		//	      }
		//	    },
		//	    "required": [
		//	      "Value",
		//	      "Key"
		//	    ],
		//	    "type": "object"
		//	  },
		//	  "type": "array"
		//	}
		"tags": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Key
					"key": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 128),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
					// Property: Value
					"value": schema.StringAttribute{ /*START ATTRIBUTE*/
						Required: true,
						Validators: []validator.String{ /*START VALIDATORS*/
							stringvalidator.LengthBetween(1, 256),
						}, /*END VALIDATORS*/
					}, /*END ATTRIBUTE*/
				}, /*END SCHEMA*/
			}, /*END NESTED OBJECT*/
			Description: "An arbitrary set of tags (key-value pairs) for this cleanrooms collaboration.",
			Optional:    true,
			Computed:    true,
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				generic.Multiset(),
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	// Corresponds to CloudFormation primaryIdentifier.
	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Represents a table that can be queried within a collaboration",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::CleanRooms::ConfiguredTableAssociation").WithTerraformTypeName("awscc_cleanrooms_configured_table_association")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithAttributeNameMap(map[string]string{
		"aggregation":                 "Aggregation",
		"allowed_additional_analyses": "AllowedAdditionalAnalyses",
		"allowed_result_receivers":    "AllowedResultReceivers",
		"arn":                         "Arn",
		"configured_table_association_analysis_rules": "ConfiguredTableAssociationAnalysisRules",
		"configured_table_association_identifier":     "ConfiguredTableAssociationIdentifier",
		"configured_table_identifier":                 "ConfiguredTableIdentifier",
		"custom":                                      "Custom",
		"description":                                 "Description",
		"key":                                         "Key",
		"list":                                        "List",
		"membership_identifier":                       "MembershipIdentifier",
		"name":                                        "Name",
		"policy":                                      "Policy",
		"role_arn":                                    "RoleArn",
		"tags":                                        "Tags",
		"type":                                        "Type",
		"v1":                                          "V1",
		"value":                                       "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
