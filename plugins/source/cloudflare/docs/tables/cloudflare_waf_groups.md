# Table: cloudflare_waf_groups


The primary key for this table is **_cq_id**.

## Relations
This table depends on [`cloudflare_waf_packages`](cloudflare_waf_packages.md).

## Columns
| Name          | Type          |
| ------------- | ------------- |
|waf_package_cq_id|UUID|
|id|String|
|name|String|
|description|String|
|rules_count|Int|
|modified_rules_count|Int|
|package_id|String|
|mode|String|
|allowed_modes|StringArray|
|_cq_id (PK)|UUID|
|_cq_fetch_time|Timestamp|