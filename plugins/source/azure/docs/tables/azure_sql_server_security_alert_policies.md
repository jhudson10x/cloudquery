# Table: azure_sql_server_security_alert_policies


The primary key for this table is **id**.

## Relations
This table depends on [`azure_sql_servers`](azure_sql_servers.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|subscription_id|String|
|sql_server_id|UUID|
|state|String|
|disabled_alerts|StringArray|
|email_addresses|StringArray|
|email_account_admins|Bool|
|storage_endpoint|String|
|storage_account_access_key|String|
|retention_days|Int|
|creation_time|Timestamp|
|id (PK)|String|
|name|String|
|type|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|