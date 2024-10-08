---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "awscc_appflow_connector_profile Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::AppFlow::ConnectorProfile
---

# awscc_appflow_connector_profile (Resource)

Resource Type definition for AWS::AppFlow::ConnectorProfile



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `connection_mode` (String) Mode in which data transfer should be enabled. Private connection mode is currently enabled for Salesforce, Snowflake, Trendmicro and Singular
- `connector_profile_name` (String) The maximum number of items to retrieve in a single batch.
- `connector_type` (String) List of Saas providers that need connector profile to be created

### Optional

- `connector_label` (String) The label of the connector. The label is unique for each ConnectorRegistration in your AWS account. Only needed if calling for CUSTOMCONNECTOR connector type/.
- `connector_profile_config` (Attributes) Connector specific configurations needed to create connector profile (see [below for nested schema](#nestedatt--connector_profile_config))
- `kms_arn` (String) The ARN of the AWS Key Management Service (AWS KMS) key that's used to encrypt your function's environment variables. If it's not provided, AWS Lambda uses a default service key.

### Read-Only

- `connector_profile_arn` (String) Unique identifier for connector profile resources
- `credentials_arn` (String) A unique Arn for Connector-Profile resource
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--connector_profile_config"></a>
### Nested Schema for `connector_profile_config`

Optional:

- `connector_profile_credentials` (Attributes) Connector specific configuration needed to create connector profile based on Authentication mechanism (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials))
- `connector_profile_properties` (Attributes) Connector specific properties needed to create connector profile - currently not needed for Amplitude, Trendmicro, Googleanalytics and Singular (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties))

<a id="nestedatt--connector_profile_config--connector_profile_credentials"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials`

Optional:

- `amplitude` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--amplitude))
- `custom_connector` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--custom_connector))
- `datadog` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--datadog))
- `dynatrace` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--dynatrace))
- `google_analytics` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--google_analytics))
- `infor_nexus` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--infor_nexus))
- `marketo` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--marketo))
- `pardot` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--pardot))
- `redshift` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--redshift))
- `salesforce` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--salesforce))
- `sapo_data` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--sapo_data))
- `service_now` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--service_now))
- `singular` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--singular))
- `slack` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--slack))
- `snowflake` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--snowflake))
- `trendmicro` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--trendmicro))
- `veeva` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--veeva))
- `zendesk` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--zendesk))

<a id="nestedatt--connector_profile_config--connector_profile_credentials--amplitude"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.amplitude`

Required:

- `api_key` (String) A unique alphanumeric identi?er used to authenticate a user, developer, or calling program to your API.
- `secret_key` (String)


<a id="nestedatt--connector_profile_config--connector_profile_credentials--custom_connector"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.custom_connector`

Required:

- `authentication_type` (String)

Optional:

- `api_key` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--custom_connector--api_key))
- `basic` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--custom_connector--basic))
- `custom` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--custom_connector--custom))
- `oauth_2` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--custom_connector--oauth_2))

<a id="nestedatt--connector_profile_config--connector_profile_credentials--custom_connector--api_key"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.custom_connector.api_key`

Required:

- `api_key` (String)

Optional:

- `api_secret_key` (String)


<a id="nestedatt--connector_profile_config--connector_profile_credentials--custom_connector--basic"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.custom_connector.basic`

Required:

- `password` (String)
- `username` (String)


<a id="nestedatt--connector_profile_config--connector_profile_credentials--custom_connector--custom"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.custom_connector.custom`

Required:

- `custom_authentication_type` (String)

Optional:

- `credentials_map` (Map of String) A map for properties for custom authentication.


<a id="nestedatt--connector_profile_config--connector_profile_credentials--custom_connector--oauth_2"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.custom_connector.oauth_2`

Optional:

- `access_token` (String)
- `client_id` (String)
- `client_secret` (String)
- `o_auth_request` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--custom_connector--oauth_2--o_auth_request))
- `refresh_token` (String)

<a id="nestedatt--connector_profile_config--connector_profile_credentials--custom_connector--oauth_2--o_auth_request"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.custom_connector.oauth_2.o_auth_request`

Optional:

- `auth_code` (String) The code provided by the connector when it has been authenticated via the connected app.
- `redirect_uri` (String) The URL to which the authentication server redirects the browser after authorization has been
granted.




<a id="nestedatt--connector_profile_config--connector_profile_credentials--datadog"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.datadog`

Required:

- `api_key` (String) A unique alphanumeric identi?er used to authenticate a user, developer, or calling program to your API.
- `application_key` (String) Application keys, in conjunction with your API key, give you full access to Datadog?s programmatic API. Application keys are associated with the user account that created them. The application key is used to log all requests made to the API.


<a id="nestedatt--connector_profile_config--connector_profile_credentials--dynatrace"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.dynatrace`

Required:

- `api_token` (String) The API tokens used by Dynatrace API to authenticate various API calls.


<a id="nestedatt--connector_profile_config--connector_profile_credentials--google_analytics"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.google_analytics`

Required:

- `client_id` (String) The identi?er for the desired client.
- `client_secret` (String) The client secret used by the oauth client to authenticate to the authorization server.

Optional:

- `access_token` (String) The credentials used to access protected resources.
- `connector_o_auth_request` (Attributes) The oauth needed to request security tokens from the connector endpoint. (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--google_analytics--connector_o_auth_request))
- `refresh_token` (String) The credentials used to acquire new access tokens.

<a id="nestedatt--connector_profile_config--connector_profile_credentials--google_analytics--connector_o_auth_request"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.google_analytics.connector_o_auth_request`

Optional:

- `auth_code` (String) The code provided by the connector when it has been authenticated via the connected app.
- `redirect_uri` (String) The URL to which the authentication server redirects the browser after authorization has been
granted.



<a id="nestedatt--connector_profile_config--connector_profile_credentials--infor_nexus"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.infor_nexus`

Required:

- `access_key_id` (String) The Access Key portion of the credentials.
- `datakey` (String) The encryption keys used to encrypt data.
- `secret_access_key` (String) The secret key used to sign requests.
- `user_id` (String) The identi?er for the user.


<a id="nestedatt--connector_profile_config--connector_profile_credentials--marketo"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.marketo`

Required:

- `client_id` (String) The identi?er for the desired client.
- `client_secret` (String) The client secret used by the oauth client to authenticate to the authorization server.

Optional:

- `access_token` (String) The credentials used to access protected resources.
- `connector_o_auth_request` (Attributes) The oauth needed to request security tokens from the connector endpoint. (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--marketo--connector_o_auth_request))

<a id="nestedatt--connector_profile_config--connector_profile_credentials--marketo--connector_o_auth_request"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.marketo.connector_o_auth_request`

Optional:

- `auth_code` (String) The code provided by the connector when it has been authenticated via the connected app.
- `redirect_uri` (String) The URL to which the authentication server redirects the browser after authorization has been
granted.



<a id="nestedatt--connector_profile_config--connector_profile_credentials--pardot"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.pardot`

Optional:

- `access_token` (String) The credentials used to access protected resources.
- `client_credentials_arn` (String) The client credentials to fetch access token and refresh token.
- `connector_o_auth_request` (Attributes) The oauth needed to request security tokens from the connector endpoint. (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--pardot--connector_o_auth_request))
- `refresh_token` (String) The credentials used to acquire new access tokens.

<a id="nestedatt--connector_profile_config--connector_profile_credentials--pardot--connector_o_auth_request"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.pardot.connector_o_auth_request`

Optional:

- `auth_code` (String) The code provided by the connector when it has been authenticated via the connected app.
- `redirect_uri` (String) The URL to which the authentication server redirects the browser after authorization has been
granted.



<a id="nestedatt--connector_profile_config--connector_profile_credentials--redshift"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.redshift`

Optional:

- `password` (String) The password that corresponds to the username.
- `username` (String) The name of the user.


<a id="nestedatt--connector_profile_config--connector_profile_credentials--salesforce"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.salesforce`

Optional:

- `access_token` (String) The credentials used to access protected resources.
- `client_credentials_arn` (String) The client credentials to fetch access token and refresh token.
- `connector_o_auth_request` (Attributes) The oauth needed to request security tokens from the connector endpoint. (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--salesforce--connector_o_auth_request))
- `refresh_token` (String) The credentials used to acquire new access tokens.

<a id="nestedatt--connector_profile_config--connector_profile_credentials--salesforce--connector_o_auth_request"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.salesforce.connector_o_auth_request`

Optional:

- `auth_code` (String) The code provided by the connector when it has been authenticated via the connected app.
- `redirect_uri` (String) The URL to which the authentication server redirects the browser after authorization has been
granted.



<a id="nestedatt--connector_profile_config--connector_profile_credentials--sapo_data"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.sapo_data`

Optional:

- `basic_auth_credentials` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--sapo_data--basic_auth_credentials))
- `o_auth_credentials` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--sapo_data--o_auth_credentials))

<a id="nestedatt--connector_profile_config--connector_profile_credentials--sapo_data--basic_auth_credentials"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.sapo_data.basic_auth_credentials`

Required:

- `password` (String)
- `username` (String)


<a id="nestedatt--connector_profile_config--connector_profile_credentials--sapo_data--o_auth_credentials"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.sapo_data.o_auth_credentials`

Optional:

- `access_token` (String)
- `client_id` (String)
- `client_secret` (String)
- `connector_o_auth_request` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--sapo_data--o_auth_credentials--connector_o_auth_request))
- `refresh_token` (String)

<a id="nestedatt--connector_profile_config--connector_profile_credentials--sapo_data--o_auth_credentials--connector_o_auth_request"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.sapo_data.o_auth_credentials.connector_o_auth_request`

Optional:

- `auth_code` (String) The code provided by the connector when it has been authenticated via the connected app.
- `redirect_uri` (String) The URL to which the authentication server redirects the browser after authorization has been
granted.




<a id="nestedatt--connector_profile_config--connector_profile_credentials--service_now"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.service_now`

Required:

- `password` (String) The password that corresponds to the username.
- `username` (String) The name of the user.


<a id="nestedatt--connector_profile_config--connector_profile_credentials--singular"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.singular`

Required:

- `api_key` (String) A unique alphanumeric identi?er used to authenticate a user, developer, or calling program to your API.


<a id="nestedatt--connector_profile_config--connector_profile_credentials--slack"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.slack`

Required:

- `client_id` (String) The identi?er for the desired client.
- `client_secret` (String) The client secret used by the oauth client to authenticate to the authorization server.

Optional:

- `access_token` (String) The credentials used to access protected resources.
- `connector_o_auth_request` (Attributes) The oauth needed to request security tokens from the connector endpoint. (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--slack--connector_o_auth_request))

<a id="nestedatt--connector_profile_config--connector_profile_credentials--slack--connector_o_auth_request"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.slack.connector_o_auth_request`

Optional:

- `auth_code` (String) The code provided by the connector when it has been authenticated via the connected app.
- `redirect_uri` (String) The URL to which the authentication server redirects the browser after authorization has been
granted.



<a id="nestedatt--connector_profile_config--connector_profile_credentials--snowflake"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.snowflake`

Required:

- `password` (String) The password that corresponds to the username.
- `username` (String) The name of the user.


<a id="nestedatt--connector_profile_config--connector_profile_credentials--trendmicro"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.trendmicro`

Required:

- `api_secret_key` (String) The Secret Access Key portion of the credentials.


<a id="nestedatt--connector_profile_config--connector_profile_credentials--veeva"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.veeva`

Required:

- `password` (String) The password that corresponds to the username.
- `username` (String) The name of the user.


<a id="nestedatt--connector_profile_config--connector_profile_credentials--zendesk"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.zendesk`

Required:

- `client_id` (String) The identi?er for the desired client.
- `client_secret` (String) The client secret used by the oauth client to authenticate to the authorization server.

Optional:

- `access_token` (String) The credentials used to access protected resources.
- `connector_o_auth_request` (Attributes) The oauth needed to request security tokens from the connector endpoint. (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_credentials--zendesk--connector_o_auth_request))

<a id="nestedatt--connector_profile_config--connector_profile_credentials--zendesk--connector_o_auth_request"></a>
### Nested Schema for `connector_profile_config.connector_profile_credentials.zendesk.connector_o_auth_request`

Optional:

- `auth_code` (String) The code provided by the connector when it has been authenticated via the connected app.
- `redirect_uri` (String) The URL to which the authentication server redirects the browser after authorization has been
granted.




<a id="nestedatt--connector_profile_config--connector_profile_properties"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties`

Optional:

- `custom_connector` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--custom_connector))
- `datadog` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--datadog))
- `dynatrace` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--dynatrace))
- `infor_nexus` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--infor_nexus))
- `marketo` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--marketo))
- `pardot` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--pardot))
- `redshift` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--redshift))
- `salesforce` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--salesforce))
- `sapo_data` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--sapo_data))
- `service_now` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--service_now))
- `slack` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--slack))
- `snowflake` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--snowflake))
- `veeva` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--veeva))
- `zendesk` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--zendesk))

<a id="nestedatt--connector_profile_config--connector_profile_properties--custom_connector"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.custom_connector`

Optional:

- `o_auth_2_properties` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--custom_connector--o_auth_2_properties))
- `profile_properties` (Map of String) A map for properties for custom connector.

<a id="nestedatt--connector_profile_config--connector_profile_properties--custom_connector--o_auth_2_properties"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.custom_connector.o_auth_2_properties`

Optional:

- `o_auth_2_grant_type` (String)
- `token_url` (String)
- `token_url_custom_properties` (Map of String) A map for properties for custom connector Token Url.



<a id="nestedatt--connector_profile_config--connector_profile_properties--datadog"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.datadog`

Required:

- `instance_url` (String) The location of the Datadog resource


<a id="nestedatt--connector_profile_config--connector_profile_properties--dynatrace"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.dynatrace`

Required:

- `instance_url` (String) The location of the Dynatrace resource


<a id="nestedatt--connector_profile_config--connector_profile_properties--infor_nexus"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.infor_nexus`

Required:

- `instance_url` (String) The location of the InforNexus resource


<a id="nestedatt--connector_profile_config--connector_profile_properties--marketo"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.marketo`

Required:

- `instance_url` (String) The location of the Marketo resource


<a id="nestedatt--connector_profile_config--connector_profile_properties--pardot"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.pardot`

Required:

- `business_unit_id` (String) The Business unit id of Salesforce Pardot instance to be connected

Optional:

- `instance_url` (String) The location of the Salesforce Pardot resource
- `is_sandbox_environment` (Boolean) Indicates whether the connector profile applies to a demo or production environment


<a id="nestedatt--connector_profile_config--connector_profile_properties--redshift"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.redshift`

Required:

- `bucket_name` (String) The name of the Amazon S3 bucket associated with Redshift.
- `role_arn` (String) The Amazon Resource Name (ARN) of the IAM role.

Optional:

- `bucket_prefix` (String) The object key for the destination bucket in which Amazon AppFlow will place the ?les.
- `cluster_identifier` (String) The unique identifier of the Amazon Redshift cluster.
- `data_api_role_arn` (String) The Amazon Resource Name (ARN) of the IAM role that grants Amazon AppFlow access to the data through the Amazon Redshift Data API.
- `database_name` (String) The name of the Amazon Redshift database that will store the transferred data.
- `database_url` (String) The JDBC URL of the Amazon Redshift cluster.
- `is_redshift_serverless` (Boolean) If Amazon AppFlow will connect to Amazon Redshift Serverless or Amazon Redshift cluster.
- `workgroup_name` (String) The name of the Amazon Redshift serverless workgroup


<a id="nestedatt--connector_profile_config--connector_profile_properties--salesforce"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.salesforce`

Optional:

- `instance_url` (String) The location of the Salesforce resource
- `is_sandbox_environment` (Boolean) Indicates whether the connector profile applies to a sandbox or production environment


<a id="nestedatt--connector_profile_config--connector_profile_properties--sapo_data"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.sapo_data`

Optional:

- `application_host_url` (String)
- `application_service_path` (String)
- `client_number` (String)
- `logon_language` (String)
- `o_auth_properties` (Attributes) (see [below for nested schema](#nestedatt--connector_profile_config--connector_profile_properties--sapo_data--o_auth_properties))
- `port_number` (Number)
- `private_link_service_name` (String)

<a id="nestedatt--connector_profile_config--connector_profile_properties--sapo_data--o_auth_properties"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.sapo_data.o_auth_properties`

Optional:

- `auth_code_url` (String)
- `o_auth_scopes` (List of String)
- `token_url` (String)



<a id="nestedatt--connector_profile_config--connector_profile_properties--service_now"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.service_now`

Required:

- `instance_url` (String) The location of the ServiceNow resource


<a id="nestedatt--connector_profile_config--connector_profile_properties--slack"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.slack`

Required:

- `instance_url` (String) The location of the Slack resource


<a id="nestedatt--connector_profile_config--connector_profile_properties--snowflake"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.snowflake`

Required:

- `bucket_name` (String) The name of the Amazon S3 bucket associated with Snow?ake.
- `stage` (String) The name of the Amazon S3 stage that was created while setting up an Amazon S3 stage in the
Snow?ake account. This is written in the following format: < Database>< Schema><Stage Name>.
- `warehouse` (String) The name of the Snow?ake warehouse.

Optional:

- `account_name` (String) The name of the account.
- `bucket_prefix` (String) The bucket prefix that refers to the Amazon S3 bucket associated with Snow?ake.
- `private_link_service_name` (String) The Snow?ake Private Link service name to be used for private data transfers.
- `region` (String) The region of the Snow?ake account.


<a id="nestedatt--connector_profile_config--connector_profile_properties--veeva"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.veeva`

Required:

- `instance_url` (String) The location of the Veeva resource


<a id="nestedatt--connector_profile_config--connector_profile_properties--zendesk"></a>
### Nested Schema for `connector_profile_config.connector_profile_properties.zendesk`

Required:

- `instance_url` (String) The location of the Zendesk resource

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_appflow_connector_profile.example "connector_profile_name"
```
