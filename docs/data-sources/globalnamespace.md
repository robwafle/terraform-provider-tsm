---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "tsm_globalnamespace Data Source - terraform-provider-tsm"
subcategory: ""
description: |-
  
---

# tsm_globalnamespace (Data Source)





<!-- schema generated by tfplugindocs -->
## Schema

### Read-Only

- `api_discovery_enabled` (Boolean)
- `ca` (String)
- `ca_type` (String)
- `color` (String)
- `description` (String)
- `display_name` (String)
- `domain_name` (String)
- `id` (String) The ID of this resource.
- `last_updated` (String)
- `match_conditions` (Set of Object) (see [below for nested schema](#nestedatt--match_conditions))
- `mtls_enforced` (Boolean)
- `name` (String)
- `use_shared_gateway` (Boolean)
- `version` (String)

<a id="nestedatt--match_conditions"></a>
### Nested Schema for `match_conditions`

Read-Only:

- `cluster_match` (String)
- `cluster_type` (String)
- `namespace_match` (String)
- `namespace_type` (String)


