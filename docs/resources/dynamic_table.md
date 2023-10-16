---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "snowflake_dynamic_table Resource - terraform-provider-snowflake"
subcategory: ""
description: |-
  
---

# snowflake_dynamic_table (Resource)



## Example Usage

```terraform
# https://docs.snowflake.com/en/sql-reference/sql/create-dynamic-table#examples
resource "snowflake_dynamic_table" "dt" {
  name     = "product"
  database = "mydb"
  schema   = "myschema"
  target_lag {
    maximum_duration = "20 minutes"
  }
  warehouse = "mywh"
  query     = "SELECT product_id, product_name FROM \"mydb\".\"myschema\".\"staging_table\""
  comment   = "example comment"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `database` (String) The database in which to create the dynamic table.
- `name` (String) Specifies the identifier (i.e. name) for the dynamic table; must be unique for the schema in which the dynamic table is created.
- `query` (String) Specifies the query to use to populate the dynamic table.
- `schema` (String) The schema in which to create the dynamic table.
- `target_lag` (Block List, Min: 1, Max: 1) Specifies the target lag time for the dynamic table. (see [below for nested schema](#nestedblock--target_lag))
- `warehouse` (String) The warehouse in which to create the dynamic table.

### Optional

- `comment` (String) Specifies a comment for the dynamic table.
- `or_replace` (Boolean) Specifies whether to replace the dynamic table if it already exists.

### Read-Only

- `automatic_clustering` (Boolean) Whether auto-clustering is enabled on the dynamic table. Not currently supported for dynamic tables.
- `bytes` (Number) Number of bytes that will be scanned if the entire dynamic table is scanned in a query.
- `cluster_by` (String) The clustering key for the dynamic table.
- `data_timestamp` (String) Timestamp of the data in the base object(s) that is included in the dynamic table.
- `id` (String) The ID of this resource.
- `is_clone` (Boolean) TRUE if the dynamic table has been cloned, else FALSE.
- `is_replica` (Boolean) TRUE if the dynamic table is a replica. else FALSE.
- `last_suspended_on` (String) Timestamp of last suspension.
- `owner` (String) Role that owns the dynamic table.
- `refresh_mode` (String) INCREMENTAL if the dynamic table will use incremental refreshes, or FULL if it will recompute the whole table on every refresh.
- `refresh_mode_reason` (String) Explanation for why FULL refresh mode was chosen. NULL if refresh mode is not FULL.
- `rows` (Number) Number of rows in the table.
- `scheduling_state` (String) Displays RUNNING for dynamic tables that are actively scheduling refreshes and SUSPENDED for suspended dynamic tables.

<a id="nestedblock--target_lag"></a>
### Nested Schema for `target_lag`

Optional:

- `downstream` (Boolean) Specifies whether the target lag time is downstream.
- `maximum_duration` (String) Specifies the maximum target lag time for the dynamic table.

## Import

Import is supported using the following syntax:

```shell
terraform import snowflake_dynamic_table.dt "mydb|myschema|product"
```