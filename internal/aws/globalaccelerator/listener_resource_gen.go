// Code generated by generators/resource/main.go; DO NOT EDIT.

package globalaccelerator

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/tfsdk"
	"github.com/hashicorp/terraform-plugin-framework/types"
	. "github.com/hashicorp/terraform-provider-awscc/internal/generic"
	"github.com/hashicorp/terraform-provider-awscc/internal/registry"
	"github.com/hashicorp/terraform-provider-awscc/internal/validate"
)

func init() {
	registry.AddResourceTypeFactory("awscc_globalaccelerator_listener", listenerResourceType)
}

// listenerResourceType returns the Terraform awscc_globalaccelerator_listener resource type.
// This Terraform resource type corresponds to the CloudFormation AWS::GlobalAccelerator::Listener resource type.
func listenerResourceType(ctx context.Context) (tfsdk.ResourceType, error) {
	attributes := map[string]tfsdk.Attribute{
		"accelerator_arn": {
			// Property: AcceleratorArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the accelerator.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the accelerator.",
			Type:        types.StringType,
			Required:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.RequiresReplace(),
			},
		},
		"client_affinity": {
			// Property: ClientAffinity
			// CloudFormation resource type schema:
			// {
			//   "default": "NONE",
			//   "description": "Client affinity lets you direct all requests from a user to the same endpoint.",
			//   "enum": [
			//     "NONE",
			//     "SOURCE_IP"
			//   ],
			//   "type": "string"
			// }
			Description: "Client affinity lets you direct all requests from a user to the same endpoint.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"NONE",
					"SOURCE_IP",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.String{Value: "NONE"}),
				tfsdk.UseStateForUnknown(),
			},
		},
		"listener_arn": {
			// Property: ListenerArn
			// CloudFormation resource type schema:
			// {
			//   "description": "The Amazon Resource Name (ARN) of the listener.",
			//   "type": "string"
			// }
			Description: "The Amazon Resource Name (ARN) of the listener.",
			Type:        types.StringType,
			Computed:    true,
			PlanModifiers: []tfsdk.AttributePlanModifier{
				tfsdk.UseStateForUnknown(),
			},
		},
		"port_ranges": {
			// Property: PortRanges
			// CloudFormation resource type schema:
			// {
			//   "items": {
			//     "additionalProperties": false,
			//     "description": "A port range to support for connections from  clients to your accelerator.",
			//     "properties": {
			//       "FromPort": {
			//         "description": "A network port number",
			//         "maximum": 65535,
			//         "minimum": 0,
			//         "type": "integer"
			//       },
			//       "ToPort": {
			//         "description": "A network port number",
			//         "maximum": 65535,
			//         "minimum": 0,
			//         "type": "integer"
			//       }
			//     },
			//     "required": [
			//       "FromPort",
			//       "ToPort"
			//     ],
			//     "type": "object"
			//   },
			//   "type": "array"
			// }
			Attributes: tfsdk.ListNestedAttributes(
				map[string]tfsdk.Attribute{
					"from_port": {
						// Property: FromPort
						Description: "A network port number",
						Type:        types.Int64Type,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 65535),
						},
					},
					"to_port": {
						// Property: ToPort
						Description: "A network port number",
						Type:        types.Int64Type,
						Required:    true,
						Validators: []tfsdk.AttributeValidator{
							validate.IntBetween(0, 65535),
						},
					},
				},
				tfsdk.ListNestedAttributesOptions{},
			),
			Required: true,
		},
		"protocol": {
			// Property: Protocol
			// CloudFormation resource type schema:
			// {
			//   "default": "TCP",
			//   "description": "The protocol for the listener.",
			//   "enum": [
			//     "TCP",
			//     "UDP"
			//   ],
			//   "type": "string"
			// }
			Description: "The protocol for the listener.",
			Type:        types.StringType,
			Optional:    true,
			Computed:    true,
			Validators: []tfsdk.AttributeValidator{
				validate.StringInSlice([]string{
					"TCP",
					"UDP",
				}),
			},
			PlanModifiers: []tfsdk.AttributePlanModifier{
				DefaultValue(types.String{Value: "TCP"}),
				tfsdk.UseStateForUnknown(),
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
		Description: "Resource Type definition for AWS::GlobalAccelerator::Listener",
		Version:     1,
		Attributes:  attributes,
	}

	var opts ResourceTypeOptions

	opts = opts.WithCloudFormationTypeName("AWS::GlobalAccelerator::Listener").WithTerraformTypeName("awscc_globalaccelerator_listener")
	opts = opts.WithTerraformSchema(schema)
	opts = opts.WithSyntheticIDAttribute(true)
	opts = opts.WithAttributeNameMap(map[string]string{
		"accelerator_arn": "AcceleratorArn",
		"client_affinity": "ClientAffinity",
		"from_port":       "FromPort",
		"listener_arn":    "ListenerArn",
		"port_ranges":     "PortRanges",
		"protocol":        "Protocol",
		"to_port":         "ToPort",
	})

	opts = opts.WithCreateTimeoutInMinutes(0).WithDeleteTimeoutInMinutes(0)

	opts = opts.WithUpdateTimeoutInMinutes(0)

	resourceType, err := NewResourceType(ctx, opts...)

	if err != nil {
		return nil, err
	}

	return resourceType, nil
}
