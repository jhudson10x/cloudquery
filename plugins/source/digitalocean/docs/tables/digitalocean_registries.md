# Table: digitalocean_registries


The primary key for this table is **name**.

## Relations
The following tables depend on `digitalocean_registries`:
  - [`digitalocean_registry_repositories`](digitalocean_registry_repositories.md)

## Columns
| Name          | Type          |
| ------------- | ------------- |
|name (PK)|String|
|storage_usage_bytes|Int|
|storage_usage_bytes_updated_at|Timestamp|
|created_at|Timestamp|
|region|String|
|_cq_id|UUID|
|_cq_fetch_time|Timestamp|