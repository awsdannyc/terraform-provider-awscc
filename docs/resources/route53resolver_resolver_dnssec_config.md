---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_route53resolver_resolver_dnssec_config Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.
---

# awscc_route53resolver_resolver_dnssec_config (Resource)

Resource schema for AWS::Route53Resolver::ResolverDNSSECConfig.



<!-- schema generated by tfplugindocs -->
## Schema

### Optional

- `resource_id` (String) ResourceId

### Read-Only

- `id` (String) Uniquely identifies the resource.
- `owner_id` (String) AccountId
- `resolver_dnssec_config_id` (String) Id
- `validation_status` (String) ResolverDNSSECValidationStatus, possible values are ENABLING, ENABLED, DISABLING AND DISABLED.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_route53resolver_resolver_dnssec_config.example "id"
```
