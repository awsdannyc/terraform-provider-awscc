---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_groundstation_dataflow_endpoint_group Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  AWS Ground Station DataflowEndpointGroup schema for CloudFormation
---

# awscc_groundstation_dataflow_endpoint_group (Resource)

AWS Ground Station DataflowEndpointGroup schema for CloudFormation



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `endpoint_details` (Attributes List) (see [below for nested schema](#nestedatt--endpoint_details))

### Optional

- `contact_post_pass_duration_seconds` (Number) Amount of time, in seconds, after a contact ends that the Ground Station Dataflow Endpoint Group will be in a POSTPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the POSTPASS state.
- `contact_pre_pass_duration_seconds` (Number) Amount of time, in seconds, before a contact starts that the Ground Station Dataflow Endpoint Group will be in a PREPASS state. A Ground Station Dataflow Endpoint Group State Change event will be emitted when the Dataflow Endpoint Group enters and exits the PREPASS state.
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `arn` (String)
- `dataflow_endpoint_group_id` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--endpoint_details"></a>
### Nested Schema for `endpoint_details`

Optional:

- `aws_ground_station_agent_endpoint` (Attributes) Information about AwsGroundStationAgentEndpoint. (see [below for nested schema](#nestedatt--endpoint_details--aws_ground_station_agent_endpoint))
- `endpoint` (Attributes) (see [below for nested schema](#nestedatt--endpoint_details--endpoint))
- `security_details` (Attributes) (see [below for nested schema](#nestedatt--endpoint_details--security_details))

<a id="nestedatt--endpoint_details--aws_ground_station_agent_endpoint"></a>
### Nested Schema for `endpoint_details.aws_ground_station_agent_endpoint`

Optional:

- `agent_status` (String) The status of AgentEndpoint.
- `audit_results` (String) The results of the audit.
- `egress_address` (Attributes) Egress address of AgentEndpoint with an optional mtu. (see [below for nested schema](#nestedatt--endpoint_details--aws_ground_station_agent_endpoint--egress_address))
- `ingress_address` (Attributes) Ingress address of AgentEndpoint with a port range and an optional mtu. (see [below for nested schema](#nestedatt--endpoint_details--aws_ground_station_agent_endpoint--ingress_address))
- `name` (String)

<a id="nestedatt--endpoint_details--aws_ground_station_agent_endpoint--egress_address"></a>
### Nested Schema for `endpoint_details.aws_ground_station_agent_endpoint.egress_address`

Optional:

- `mtu` (Number) Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
- `socket_address` (Attributes) (see [below for nested schema](#nestedatt--endpoint_details--aws_ground_station_agent_endpoint--egress_address--socket_address))

<a id="nestedatt--endpoint_details--aws_ground_station_agent_endpoint--egress_address--socket_address"></a>
### Nested Schema for `endpoint_details.aws_ground_station_agent_endpoint.egress_address.socket_address`

Optional:

- `name` (String)
- `port` (Number)



<a id="nestedatt--endpoint_details--aws_ground_station_agent_endpoint--ingress_address"></a>
### Nested Schema for `endpoint_details.aws_ground_station_agent_endpoint.ingress_address`

Optional:

- `mtu` (Number) Maximum transmission unit (MTU) size in bytes of a dataflow endpoint.
- `socket_address` (Attributes) A socket address with a port range. (see [below for nested schema](#nestedatt--endpoint_details--aws_ground_station_agent_endpoint--ingress_address--socket_address))

<a id="nestedatt--endpoint_details--aws_ground_station_agent_endpoint--ingress_address--socket_address"></a>
### Nested Schema for `endpoint_details.aws_ground_station_agent_endpoint.ingress_address.socket_address`

Optional:

- `name` (String) IPv4 socket address.
- `port_range` (Attributes) Port range of a socket address. (see [below for nested schema](#nestedatt--endpoint_details--aws_ground_station_agent_endpoint--ingress_address--socket_address--port_range))

<a id="nestedatt--endpoint_details--aws_ground_station_agent_endpoint--ingress_address--socket_address--port_range"></a>
### Nested Schema for `endpoint_details.aws_ground_station_agent_endpoint.ingress_address.socket_address.port_range`

Optional:

- `maximum` (Number) A maximum value.
- `minimum` (Number) A minimum value.





<a id="nestedatt--endpoint_details--endpoint"></a>
### Nested Schema for `endpoint_details.endpoint`

Optional:

- `address` (Attributes) (see [below for nested schema](#nestedatt--endpoint_details--endpoint--address))
- `mtu` (Number)
- `name` (String)

<a id="nestedatt--endpoint_details--endpoint--address"></a>
### Nested Schema for `endpoint_details.endpoint.address`

Optional:

- `name` (String)
- `port` (Number)



<a id="nestedatt--endpoint_details--security_details"></a>
### Nested Schema for `endpoint_details.security_details`

Optional:

- `role_arn` (String)
- `security_group_ids` (List of String)
- `subnet_ids` (List of String)



<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_groundstation_dataflow_endpoint_group.example "id"
```
