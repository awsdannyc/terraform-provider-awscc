---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_databrew_project Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource schema for AWS::DataBrew::Project.
---

# awscc_databrew_project (Resource)

Resource schema for AWS::DataBrew::Project.



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `dataset_name` (String) Dataset name
- `name` (String) Project name
- `recipe_name` (String) Recipe name
- `role_arn` (String) Role arn

### Optional

- `sample` (Attributes) Sample (see [below for nested schema](#nestedatt--sample))
- `tags` (Attributes List) (see [below for nested schema](#nestedatt--tags))

### Read-Only

- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--sample"></a>
### Nested Schema for `sample`

Required:

- `type` (String) Sample type

Optional:

- `size` (Number) Sample size


<a id="nestedatt--tags"></a>
### Nested Schema for `tags`

Required:

- `key` (String)
- `value` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_databrew_project.example "name"
```
