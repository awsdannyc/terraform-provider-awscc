---
page_title: "awscc_backup_backup_vault Resource - terraform-provider-awscc"
subcategory: ""
description: |-
  Resource Type definition for AWS::Backup::BackupVault
---

# awscc_backup_backup_vault (Resource)

Resource Type definition for AWS::Backup::BackupVault

## Example Usage

### First example
Create a backup vault:
```terraform
resource "awscc_backup_backup_vault" "example" {
  backup_vault_name = "example_backup_vault"
}
```

### Second example
Create a backup vault with a KMS key:
```terraform
resource "awscc_backup_backup_vault" "example" {
  backup_vault_name  = "example_backup_vault_kms"
  encryption_key_arn = awscc_kms_key.example.arn
}

resource "awscc_kms_key" "example" {
  description = "KMS Key for root"
  key_policy = jsonencode({
    "Version" : "2012-10-17",
    "Id" : "KMS-Key-Policy-For-Root",
    "Statement" : [
      {
        "Sid" : "Enable IAM User Permissions",
        "Effect" : "Allow",
        "Principal" : {
          "AWS" : "arn:aws:iam::111122223333:root"
        },
        "Action" : "kms:*",
        "Resource" : "*"
      },
    ],
    }
  )
}

resource "awscc_kms_alias" "example" {
  alias_name    = "alias/backup-kms-example"
  target_key_id = awscc_kms_key.example.key_id
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `backup_vault_name` (String)

### Optional

- `access_policy` (String)
- `backup_vault_tags` (Map of String)
- `encryption_key_arn` (String)
- `lock_configuration` (Attributes) (see [below for nested schema](#nestedatt--lock_configuration))
- `notifications` (Attributes) (see [below for nested schema](#nestedatt--notifications))

### Read-Only

- `backup_vault_arn` (String)
- `id` (String) Uniquely identifies the resource.

<a id="nestedatt--lock_configuration"></a>
### Nested Schema for `lock_configuration`

Required:

- `min_retention_days` (Number)

Optional:

- `changeable_for_days` (Number)
- `max_retention_days` (Number)


<a id="nestedatt--notifications"></a>
### Nested Schema for `notifications`

Required:

- `backup_vault_events` (List of String)
- `sns_topic_arn` (String)

## Import

Import is supported using the following syntax:

```shell
$ terraform import awscc_backup_backup_vault.example "backup_vault_name"
```