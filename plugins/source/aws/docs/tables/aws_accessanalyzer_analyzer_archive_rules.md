# Table: aws_accessanalyzer_analyzer_archive_rules


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`aws_accessanalyzer_analyzers`](aws_accessanalyzer_analyzers.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|account_id|String|
|region|String|
|analyzer_arn|String|
|created_at|Timestamp|
|filter|JSON|
|rule_name|String|
|updated_at|Timestamp|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|