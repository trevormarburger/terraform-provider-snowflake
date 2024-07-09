---
page_title: "snowflake_oauth_integration_for_custom_clients Resource - terraform-provider-snowflake"
subcategory: ""
description: |-
  
---

!> **V1 release candidate** This resource was reworked and is a release candidate for the V1. We do not expect significant changes in it before the V1. We will welcome any feedback and adjust the resource if needed. Any errors reported will be resolved with a higher priority. We encourage checking this resource out before the V1 release. Please follow the [migration guide](https://github.com/Snowflake-Labs/terraform-provider-snowflake/blob/main/MIGRATION_GUIDE.md#v0920--v0930) to use it.

# snowflake_oauth_integration_for_custom_clients (Resource)



## Example Usage

```terraform
# basic resource
resource "snowflake_oauth_integration_for_custom_clients" "basic" {
  name               = "saml_integration"
  oauth_client_type  = "CONFIDENTIAL"
  oauth_redirect_uri = "https://example.com"
  blocked_roles_list = ["ACCOUNTADMIN", "SECURITYADMIN"]
}

# resource with all fields set
resource "snowflake_oauth_integration_for_custom_clients" "complete" {
  name                             = "saml_integration"
  oauth_client_type                = "CONFIDENTIAL"
  oauth_redirect_uri               = "https://example.com"
  enabled                          = "true"
  oauth_allow_non_tls_redirect_uri = "true"
  oauth_enforce_pkce               = "true"
  oauth_use_secondary_roles        = "NONE"
  pre_authorized_roles_list        = ["role_id1", "role_id2"]
  blocked_roles_list               = ["ACCOUNTADMIN", "SECURITYADMIN", "role_id1", "role_id2"]
  oauth_issue_refresh_tokens       = "true"
  oauth_refresh_token_validity     = 87600
  network_policy                   = "network_policy_id"
  oauth_client_rsa_public_key      = file("rsa.pub")
  oauth_client_rsa_public_key_2    = file("rsa2.pub")
  comment                          = "my oauth integration"
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `blocked_roles_list` (Set of String) A set of Snowflake roles that a user cannot explicitly consent to using after authenticating.
- `name` (String) Specifies the name of the OAuth integration. This name follows the rules for Object Identifiers. The name should be unique among security integrations in your account.
- `oauth_client_type` (String) Specifies the type of client being registered. Snowflake supports both confidential and public clients. Valid options are: [PUBLIC CONFIDENTIAL]
- `oauth_redirect_uri` (String) Specifies the client URI. After a user is authenticated, the web browser is redirected to this URI.

### Optional

- `comment` (String) Specifies a comment for the OAuth integration.
- `enabled` (String) Specifies whether this OAuth integration is enabled or disabled. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
- `network_policy` (String) Specifies an existing network policy. This network policy controls network traffic that is attempting to exchange an authorization code for an access or refresh token or to use a refresh token to obtain a new access token.
- `oauth_allow_non_tls_redirect_uri` (String) If true, allows setting oauth_redirect_uri to a URI not protected by TLS. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
- `oauth_client_rsa_public_key` (String) Specifies a Base64-encoded RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource using `terraform taint`.
- `oauth_client_rsa_public_key_2` (String) Specifies a Base64-encoded RSA public key, without the -----BEGIN PUBLIC KEY----- and -----END PUBLIC KEY----- headers. External changes for this field won't be detected. In case you want to apply external changes, you can re-create the resource using `terraform taint`.
- `oauth_enforce_pkce` (String) Boolean that specifies whether Proof Key for Code Exchange (PKCE) should be required for the integration. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
- `oauth_issue_refresh_tokens` (String) Specifies whether to allow the client to exchange a refresh token for an access token when the current access token has expired. Available options are: "true" or "false". When the value is not set in the configuration the provider will put "default" there which means to use the Snowflake default for this value.
- `oauth_refresh_token_validity` (Number) Specifies how long refresh tokens should be valid (in seconds). OAUTH_ISSUE_REFRESH_TOKENS must be set to TRUE.
- `oauth_use_secondary_roles` (String) Specifies whether default secondary roles set in the user properties are activated by default in the session being opened. Valid options are: [IMPLICIT NONE]
- `pre_authorized_roles_list` (Set of String) A set of Snowflake roles that a user does not need to explicitly consent to using after authenticating.

### Read-Only

- `describe_output` (List of Object) Outputs the result of `DESCRIBE SECURITY INTEGRATION` for the given integration. (see [below for nested schema](#nestedatt--describe_output))
- `id` (String) The ID of this resource.
- `show_output` (List of Object) Outputs the result of `SHOW SECURITY INTEGRATION` for the given integration. (see [below for nested schema](#nestedatt--show_output))

<a id="nestedatt--describe_output"></a>
### Nested Schema for `describe_output`

Read-Only:

- `blocked_roles_list` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--blocked_roles_list))
- `comment` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--comment))
- `enabled` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--enabled))
- `network_policy` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--network_policy))
- `oauth_allow_non_tls_redirect_uri` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_allow_non_tls_redirect_uri))
- `oauth_allowed_authorization_endpoints` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_allowed_authorization_endpoints))
- `oauth_allowed_token_endpoints` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_allowed_token_endpoints))
- `oauth_authorization_endpoint` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_authorization_endpoint))
- `oauth_client_id` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_client_id))
- `oauth_client_rsa_public_key_2_fp` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_client_rsa_public_key_2_fp))
- `oauth_client_rsa_public_key_fp` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_client_rsa_public_key_fp))
- `oauth_client_type` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_client_type))
- `oauth_enforce_pkce` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_enforce_pkce))
- `oauth_issue_refresh_tokens` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_issue_refresh_tokens))
- `oauth_redirect_uri` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_redirect_uri))
- `oauth_refresh_token_validity` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_refresh_token_validity))
- `oauth_token_endpoint` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_token_endpoint))
- `oauth_use_secondary_roles` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--oauth_use_secondary_roles))
- `pre_authorized_roles_list` (List of Object) (see [below for nested schema](#nestedobjatt--describe_output--pre_authorized_roles_list))

<a id="nestedobjatt--describe_output--blocked_roles_list"></a>
### Nested Schema for `describe_output.blocked_roles_list`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--comment"></a>
### Nested Schema for `describe_output.comment`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--enabled"></a>
### Nested Schema for `describe_output.enabled`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--network_policy"></a>
### Nested Schema for `describe_output.network_policy`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_allow_non_tls_redirect_uri"></a>
### Nested Schema for `describe_output.oauth_allow_non_tls_redirect_uri`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_allowed_authorization_endpoints"></a>
### Nested Schema for `describe_output.oauth_allowed_authorization_endpoints`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_allowed_token_endpoints"></a>
### Nested Schema for `describe_output.oauth_allowed_token_endpoints`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_authorization_endpoint"></a>
### Nested Schema for `describe_output.oauth_authorization_endpoint`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_client_id"></a>
### Nested Schema for `describe_output.oauth_client_id`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_client_rsa_public_key_2_fp"></a>
### Nested Schema for `describe_output.oauth_client_rsa_public_key_2_fp`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_client_rsa_public_key_fp"></a>
### Nested Schema for `describe_output.oauth_client_rsa_public_key_fp`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_client_type"></a>
### Nested Schema for `describe_output.oauth_client_type`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_enforce_pkce"></a>
### Nested Schema for `describe_output.oauth_enforce_pkce`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_issue_refresh_tokens"></a>
### Nested Schema for `describe_output.oauth_issue_refresh_tokens`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_redirect_uri"></a>
### Nested Schema for `describe_output.oauth_redirect_uri`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_refresh_token_validity"></a>
### Nested Schema for `describe_output.oauth_refresh_token_validity`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_token_endpoint"></a>
### Nested Schema for `describe_output.oauth_token_endpoint`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--oauth_use_secondary_roles"></a>
### Nested Schema for `describe_output.oauth_use_secondary_roles`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)


<a id="nestedobjatt--describe_output--pre_authorized_roles_list"></a>
### Nested Schema for `describe_output.pre_authorized_roles_list`

Read-Only:

- `default` (String)
- `name` (String)
- `type` (String)
- `value` (String)



<a id="nestedatt--show_output"></a>
### Nested Schema for `show_output`

Read-Only:

- `category` (String)
- `comment` (String)
- `created_on` (String)
- `enabled` (Boolean)
- `integration_type` (String)
- `name` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import snowflake_oauth_integration_for_custom_clients.example "name"
```