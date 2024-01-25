// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// Code generated by generators/resource/main.go; DO NOT EDIT.

package ssmguiconnect

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/listvalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/objectplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
)

func init() {
	registry.AddResourceFactory("awscc_ssmguiconnect_preferences", preferencesResource)
}

// preferencesResource returns the Terraform awscc_ssmguiconnect_preferences resource.
// This Terraform resource corresponds to the CloudFormation AWS::SSMGuiConnect::Preferences resource.
func preferencesResource(ctx context.Context) (resource.Resource, error) {
	attributes := map[string]schema.Attribute{ /*START SCHEMA*/
		// Property: AccountId
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "The AWS Account Id that the preference is associated with, used as the unique identifier for this resource.",
		//	  "pattern": "\\d{12}",
		//	  "type": "string"
		//	}
		"account_id": schema.StringAttribute{ /*START ATTRIBUTE*/
			Description: "The AWS Account Id that the preference is associated with, used as the unique identifier for this resource.",
			Computed:    true,
			PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
				stringplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
		// Property: IdleConnection
		// CloudFormation resource type schema:
		//
		//	{
		//	  "description": "A map for Idle Connection Preferences",
		//	  "items": {
		//	    "additionalProperties": false,
		//	    "description": "Idle Connection Preferences",
		//	    "properties": {
		//	      "Alert": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Type": {
		//	            "enum": [
		//	              "MINUTES"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "default": 1,
		//	            "maximum": 60,
		//	            "minimum": 0,
		//	            "type": "integer"
		//	          }
		//	        },
		//	        "required": [
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      },
		//	      "Timeout": {
		//	        "additionalProperties": false,
		//	        "properties": {
		//	          "Type": {
		//	            "enum": [
		//	              "MINUTES"
		//	            ],
		//	            "type": "string"
		//	          },
		//	          "Value": {
		//	            "default": 10,
		//	            "maximum": 60,
		//	            "minimum": 1,
		//	            "type": "integer"
		//	          }
		//	        },
		//	        "required": [
		//	          "Value"
		//	        ],
		//	        "type": "object"
		//	      }
		//	    },
		//	    "type": "object"
		//	  },
		//	  "maxItems": 1,
		//	  "minItems": 1,
		//	  "type": "array",
		//	  "uniqueItems": true
		//	}
		"idle_connection": schema.ListNestedAttribute{ /*START ATTRIBUTE*/
			NestedObject: schema.NestedAttributeObject{ /*START NESTED OBJECT*/
				Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
					// Property: Alert
					"alert": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Type
							"type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"MINUTES",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.Int64{ /*START VALIDATORS*/
									int64validator.Between(0, 60),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									generic.Int64DefaultValue(1),
									int64planmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
						}, /*END SCHEMA*/
						Optional: true,
						Computed: true,
						PlanModifiers: []planmodifier.Object{ /*START PLAN MODIFIERS*/
							objectplanmodifier.UseStateForUnknown(),
						}, /*END PLAN MODIFIERS*/
					}, /*END ATTRIBUTE*/
					// Property: Timeout
					"timeout": schema.SingleNestedAttribute{ /*START ATTRIBUTE*/
						Attributes: map[string]schema.Attribute{ /*START SCHEMA*/
							// Property: Type
							"type": schema.StringAttribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.String{ /*START VALIDATORS*/
									stringvalidator.OneOf(
										"MINUTES",
									),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.String{ /*START PLAN MODIFIERS*/
									stringplanmodifier.UseStateForUnknown(),
								}, /*END PLAN MODIFIERS*/
							}, /*END ATTRIBUTE*/
							// Property: Value
							"value": schema.Int64Attribute{ /*START ATTRIBUTE*/
								Optional: true,
								Computed: true,
								Validators: []validator.Int64{ /*START VALIDATORS*/
									int64validator.Between(1, 60),
								}, /*END VALIDATORS*/
								PlanModifiers: []planmodifier.Int64{ /*START PLAN MODIFIERS*/
									generic.Int64DefaultValue(10),
									int64planmodifier.UseStateForUnknown(),
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
			}, /*END NESTED OBJECT*/
			Description: "A map for Idle Connection Preferences",
			Optional:    true,
			Computed:    true,
			Validators: []validator.List{ /*START VALIDATORS*/
				listvalidator.SizeBetween(1, 1),
				listvalidator.UniqueValues(),
			}, /*END VALIDATORS*/
			PlanModifiers: []planmodifier.List{ /*START PLAN MODIFIERS*/
				listplanmodifier.UseStateForUnknown(),
			}, /*END PLAN MODIFIERS*/
		}, /*END ATTRIBUTE*/
	} /*END SCHEMA*/

	attributes["id"] = schema.StringAttribute{
		Description: "Uniquely identifies the resource.",
		Computed:    true,
		PlanModifiers: []planmodifier.String{
			stringplanmodifier.UseStateForUnknown(),
		},
	}

	schema := schema.Schema{
		Description: "Definition of AWS::SSMGuiConnect::Preferences Resource Type",
		Version:     1,
		Attributes:  attributes,
	}

	var opts generic.ResourceOptions

	opts = opts.WithCloudFormationTypeName("AWS::SSMGuiConnect::Preferences").WithTerraformTypeName("awscc_ssmguiconnect_preferences")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"account_id":      "AccountId",
		"alert":           "Alert",
		"idle_connection": "IdleConnection",
		"timeout":         "Timeout",
		"type":            "Type",
		"value":           "Value",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	v, err := generic.NewResource(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return v, nil
}
