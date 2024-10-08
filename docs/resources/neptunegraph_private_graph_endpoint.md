---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_neptunegraph_private_graph_endpoint Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::NeptuneGraph::PrivateGraphEndpoint resource creates an Amazon NeptuneGraph PrivateGraphEndpoint.
---

# awscc_neptunegraph_private_graph_endpoint (Resource)

The AWS::NeptuneGraph::PrivateGraphEndpoint resource creates an Amazon NeptuneGraph PrivateGraphEndpoint.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `graph_identifier` (String) The auto-generated Graph Id assigned by the service.
- `vpc_id` (String) The VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.

### Optional

- `security_group_ids` (List of String) The security group Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.
- `subnet_ids` (List of String) The subnet Ids associated with the VPC where you want the private graph endpoint to be created, ie, the graph will be reachable from within the VPC.

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `private_graph_endpoint_identifier` (String) PrivateGraphEndpoint resource identifier generated by concatenating the associated GraphIdentifier and VpcId with an underscore separator.

 For example, if GraphIdentifier is `g-12a3bcdef4` and VpcId is `vpc-0a12bc34567de8f90`, the generated PrivateGraphEndpointIdentifier will be `g-12a3bcdef4_vpc-0a12bc34567de8f90`
- `vpc_endpoint_id` (String) VPC endpoint that provides a private connection between the Graph and specified VPC.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_neptunegraph_private_graph_endpoint.example "private_graph_endpoint_identifier"
```
