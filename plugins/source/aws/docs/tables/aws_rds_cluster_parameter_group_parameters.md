# Table: aws_rds_cluster_parameter_group_parameters


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_rds_cluster_parameter_groups`](aws_rds_cluster_parameter_groups.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|cluster_parameter_group_arn|String|
|allowed_values|String|
|apply_method|String|
|apply_type|String|
|data_type|String|
|description|String|
|is_modifiable|Bool|
|minimum_engine_version|String|
|parameter_name|String|
|parameter_value|String|
|source|String|
|supported_engine_modes|StringArray|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|