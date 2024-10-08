---
page_title: "awscc_apigateway_model Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  The AWS::ApiGateway::Model resource defines the structure of a request or response payload for an API method.
---

# awscc_apigateway_model (Resource)

The ``AWS::ApiGateway::Model`` resource defines the structure of a request or response payload for an API method.

## Example Usage

```terraform
resource "awscc_apigateway_model" "example" {
  rest_api_id  = awscc_apigateway_rest_api.example.id
  name         = "example"
  description  = "example model"
  content_type = "application/json"

  schema = jsonencode({
    type = "object"
  })
}

resource "awscc_apigateway_rest_api" "example" {
  name        = "exampleAPI"
  description = "Example API"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `rest_api_id` (String) The string identifier of the associated RestApi.

### Optional

- `content_type` (String) The content-type for the model.
- `description` (String) The description of the model.
- `name` (String) A name for the model. If you don't specify a name, CFN generates a unique physical ID and uses that ID for the model name. For more information, see [Name Type](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-name.html).
  If you specify a name, you cannot perform updates that require replacement of this resource. You can perform updates that require no or some interruption. If you must replace the resource, specify a new name.
- `schema` (String) The schema for the model. For ``application/json`` models, this should be JSON schema draft 4 model. Do not include "\*/" characters in the description of any properties because such "\*/" characters may be interpreted as the closing marker for comments in some languages, such as Java or JavaScript, causing the installation of your API's SDK generated by API Gateway to fail.

### Read-Only

- `id` (String) Uniquely identifies the resource.

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_apigateway_model.example "rest_api_id|name"
```