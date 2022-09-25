# Table: azure_cdn_security_policies


The primary key for this table is **id**.

## Relations
This table depends on [`azure_cdn_profiles`](azure_cdn_profiles.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|cdn_profile_id|UUID|
|provisioning_state|String|
|deployment_status|String|
|id (PK)|String|
|name|String|
|type|String|
|system_data|JSON|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|