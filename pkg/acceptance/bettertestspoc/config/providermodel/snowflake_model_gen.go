// Code generated by config model builder generator; DO NOT EDIT.

package providermodel

import (
	tfconfig "github.com/hashicorp/terraform-plugin-testing/config"

	"github.com/Snowflake-Labs/terraform-provider-snowflake/pkg/acceptance/bettertestspoc/config"
)

type SnowflakeModel struct {
	Account                        tfconfig.Variable `json:"account,omitempty"`
	AccountName                    tfconfig.Variable `json:"account_name,omitempty"`
	Authenticator                  tfconfig.Variable `json:"authenticator,omitempty"`
	BrowserAuth                    tfconfig.Variable `json:"browser_auth,omitempty"`
	ClientIp                       tfconfig.Variable `json:"client_ip,omitempty"`
	ClientRequestMfaToken          tfconfig.Variable `json:"client_request_mfa_token,omitempty"`
	ClientStoreTemporaryCredential tfconfig.Variable `json:"client_store_temporary_credential,omitempty"`
	ClientTimeout                  tfconfig.Variable `json:"client_timeout,omitempty"`
	DisableConsoleLogin            tfconfig.Variable `json:"disable_console_login,omitempty"`
	DisableQueryContextCache       tfconfig.Variable `json:"disable_query_context_cache,omitempty"`
	DisableTelemetry               tfconfig.Variable `json:"disable_telemetry,omitempty"`
	DriverTracing                  tfconfig.Variable `json:"driver_tracing,omitempty"`
	ExternalBrowserTimeout         tfconfig.Variable `json:"external_browser_timeout,omitempty"`
	Host                           tfconfig.Variable `json:"host,omitempty"`
	IncludeRetryReason             tfconfig.Variable `json:"include_retry_reason,omitempty"`
	InsecureMode                   tfconfig.Variable `json:"insecure_mode,omitempty"`
	JwtClientTimeout               tfconfig.Variable `json:"jwt_client_timeout,omitempty"`
	JwtExpireTimeout               tfconfig.Variable `json:"jwt_expire_timeout,omitempty"`
	KeepSessionAlive               tfconfig.Variable `json:"keep_session_alive,omitempty"`
	LoginTimeout                   tfconfig.Variable `json:"login_timeout,omitempty"`
	MaxRetryCount                  tfconfig.Variable `json:"max_retry_count,omitempty"`
	OauthAccessToken               tfconfig.Variable `json:"oauth_access_token,omitempty"`
	OauthClientId                  tfconfig.Variable `json:"oauth_client_id,omitempty"`
	OauthClientSecret              tfconfig.Variable `json:"oauth_client_secret,omitempty"`
	OauthEndpoint                  tfconfig.Variable `json:"oauth_endpoint,omitempty"`
	OauthRedirectUrl               tfconfig.Variable `json:"oauth_redirect_url,omitempty"`
	OauthRefreshToken              tfconfig.Variable `json:"oauth_refresh_token,omitempty"`
	OcspFailOpen                   tfconfig.Variable `json:"ocsp_fail_open,omitempty"`
	OktaUrl                        tfconfig.Variable `json:"okta_url,omitempty"`
	OrganizationName               tfconfig.Variable `json:"organization_name,omitempty"`
	Params                         tfconfig.Variable `json:"params,omitempty"`
	Passcode                       tfconfig.Variable `json:"passcode,omitempty"`
	PasscodeInPassword             tfconfig.Variable `json:"passcode_in_password,omitempty"`
	Password                       tfconfig.Variable `json:"password,omitempty"`
	Port                           tfconfig.Variable `json:"port,omitempty"`
	PrivateKey                     tfconfig.Variable `json:"private_key,omitempty"`
	PrivateKeyPassphrase           tfconfig.Variable `json:"private_key_passphrase,omitempty"`
	PrivateKeyPath                 tfconfig.Variable `json:"private_key_path,omitempty"`
	Profile                        tfconfig.Variable `json:"profile,omitempty"`
	Protocol                       tfconfig.Variable `json:"protocol,omitempty"`
	Region                         tfconfig.Variable `json:"region,omitempty"`
	RequestTimeout                 tfconfig.Variable `json:"request_timeout,omitempty"`
	Role                           tfconfig.Variable `json:"role,omitempty"`
	SessionParams                  tfconfig.Variable `json:"session_params,omitempty"`
	TmpDirectoryPath               tfconfig.Variable `json:"tmp_directory_path,omitempty"`
	Token                          tfconfig.Variable `json:"token,omitempty"`
	TokenAccessor                  tfconfig.Variable `json:"token_accessor,omitempty"`
	User                           tfconfig.Variable `json:"user,omitempty"`
	Username                       tfconfig.Variable `json:"username,omitempty"`
	ValidateDefaultParameters      tfconfig.Variable `json:"validate_default_parameters,omitempty"`
	Warehouse                      tfconfig.Variable `json:"warehouse,omitempty"`

	*config.ProviderModelMeta
}

/////////////////////////////////////////////////
// Basic builders (resource name and required) //
/////////////////////////////////////////////////

func SnowflakeProvider() *SnowflakeModel {
	s := &SnowflakeModel{ProviderModelMeta: config.DefaultProviderMeta("snowflake")}
	return s
}

func SnowflakeProviderAlias(
	alias string,
) *SnowflakeModel {
	s := &SnowflakeModel{ProviderModelMeta: config.ProviderMeta("snowflake", alias)}
	return s
}

/////////////////////////////////
// below all the proper values //
/////////////////////////////////

func (s *SnowflakeModel) WithAccount(account string) *SnowflakeModel {
	s.Account = tfconfig.StringVariable(account)
	return s
}

func (s *SnowflakeModel) WithAccountName(accountName string) *SnowflakeModel {
	s.AccountName = tfconfig.StringVariable(accountName)
	return s
}

func (s *SnowflakeModel) WithAuthenticator(authenticator string) *SnowflakeModel {
	s.Authenticator = tfconfig.StringVariable(authenticator)
	return s
}

func (s *SnowflakeModel) WithBrowserAuth(browserAuth bool) *SnowflakeModel {
	s.BrowserAuth = tfconfig.BoolVariable(browserAuth)
	return s
}

func (s *SnowflakeModel) WithClientIp(clientIp string) *SnowflakeModel {
	s.ClientIp = tfconfig.StringVariable(clientIp)
	return s
}

func (s *SnowflakeModel) WithClientRequestMfaToken(clientRequestMfaToken string) *SnowflakeModel {
	s.ClientRequestMfaToken = tfconfig.StringVariable(clientRequestMfaToken)
	return s
}

func (s *SnowflakeModel) WithClientStoreTemporaryCredential(clientStoreTemporaryCredential string) *SnowflakeModel {
	s.ClientStoreTemporaryCredential = tfconfig.StringVariable(clientStoreTemporaryCredential)
	return s
}

func (s *SnowflakeModel) WithClientTimeout(clientTimeout int) *SnowflakeModel {
	s.ClientTimeout = tfconfig.IntegerVariable(clientTimeout)
	return s
}

func (s *SnowflakeModel) WithDisableConsoleLogin(disableConsoleLogin string) *SnowflakeModel {
	s.DisableConsoleLogin = tfconfig.StringVariable(disableConsoleLogin)
	return s
}

func (s *SnowflakeModel) WithDisableQueryContextCache(disableQueryContextCache bool) *SnowflakeModel {
	s.DisableQueryContextCache = tfconfig.BoolVariable(disableQueryContextCache)
	return s
}

func (s *SnowflakeModel) WithDisableTelemetry(disableTelemetry bool) *SnowflakeModel {
	s.DisableTelemetry = tfconfig.BoolVariable(disableTelemetry)
	return s
}

func (s *SnowflakeModel) WithDriverTracing(driverTracing string) *SnowflakeModel {
	s.DriverTracing = tfconfig.StringVariable(driverTracing)
	return s
}

func (s *SnowflakeModel) WithExternalBrowserTimeout(externalBrowserTimeout int) *SnowflakeModel {
	s.ExternalBrowserTimeout = tfconfig.IntegerVariable(externalBrowserTimeout)
	return s
}

func (s *SnowflakeModel) WithHost(host string) *SnowflakeModel {
	s.Host = tfconfig.StringVariable(host)
	return s
}

func (s *SnowflakeModel) WithIncludeRetryReason(includeRetryReason string) *SnowflakeModel {
	s.IncludeRetryReason = tfconfig.StringVariable(includeRetryReason)
	return s
}

func (s *SnowflakeModel) WithInsecureMode(insecureMode bool) *SnowflakeModel {
	s.InsecureMode = tfconfig.BoolVariable(insecureMode)
	return s
}

func (s *SnowflakeModel) WithJwtClientTimeout(jwtClientTimeout int) *SnowflakeModel {
	s.JwtClientTimeout = tfconfig.IntegerVariable(jwtClientTimeout)
	return s
}

func (s *SnowflakeModel) WithJwtExpireTimeout(jwtExpireTimeout int) *SnowflakeModel {
	s.JwtExpireTimeout = tfconfig.IntegerVariable(jwtExpireTimeout)
	return s
}

func (s *SnowflakeModel) WithKeepSessionAlive(keepSessionAlive bool) *SnowflakeModel {
	s.KeepSessionAlive = tfconfig.BoolVariable(keepSessionAlive)
	return s
}

func (s *SnowflakeModel) WithLoginTimeout(loginTimeout int) *SnowflakeModel {
	s.LoginTimeout = tfconfig.IntegerVariable(loginTimeout)
	return s
}

func (s *SnowflakeModel) WithMaxRetryCount(maxRetryCount int) *SnowflakeModel {
	s.MaxRetryCount = tfconfig.IntegerVariable(maxRetryCount)
	return s
}

func (s *SnowflakeModel) WithOauthAccessToken(oauthAccessToken string) *SnowflakeModel {
	s.OauthAccessToken = tfconfig.StringVariable(oauthAccessToken)
	return s
}

func (s *SnowflakeModel) WithOauthClientId(oauthClientId string) *SnowflakeModel {
	s.OauthClientId = tfconfig.StringVariable(oauthClientId)
	return s
}

func (s *SnowflakeModel) WithOauthClientSecret(oauthClientSecret string) *SnowflakeModel {
	s.OauthClientSecret = tfconfig.StringVariable(oauthClientSecret)
	return s
}

func (s *SnowflakeModel) WithOauthEndpoint(oauthEndpoint string) *SnowflakeModel {
	s.OauthEndpoint = tfconfig.StringVariable(oauthEndpoint)
	return s
}

func (s *SnowflakeModel) WithOauthRedirectUrl(oauthRedirectUrl string) *SnowflakeModel {
	s.OauthRedirectUrl = tfconfig.StringVariable(oauthRedirectUrl)
	return s
}

func (s *SnowflakeModel) WithOauthRefreshToken(oauthRefreshToken string) *SnowflakeModel {
	s.OauthRefreshToken = tfconfig.StringVariable(oauthRefreshToken)
	return s
}

func (s *SnowflakeModel) WithOcspFailOpen(ocspFailOpen string) *SnowflakeModel {
	s.OcspFailOpen = tfconfig.StringVariable(ocspFailOpen)
	return s
}

func (s *SnowflakeModel) WithOktaUrl(oktaUrl string) *SnowflakeModel {
	s.OktaUrl = tfconfig.StringVariable(oktaUrl)
	return s
}

func (s *SnowflakeModel) WithOrganizationName(organizationName string) *SnowflakeModel {
	s.OrganizationName = tfconfig.StringVariable(organizationName)
	return s
}

// params attribute type is not yet supported, so WithParams can't be generated

func (s *SnowflakeModel) WithPasscode(passcode string) *SnowflakeModel {
	s.Passcode = tfconfig.StringVariable(passcode)
	return s
}

func (s *SnowflakeModel) WithPasscodeInPassword(passcodeInPassword bool) *SnowflakeModel {
	s.PasscodeInPassword = tfconfig.BoolVariable(passcodeInPassword)
	return s
}

func (s *SnowflakeModel) WithPassword(password string) *SnowflakeModel {
	s.Password = tfconfig.StringVariable(password)
	return s
}

func (s *SnowflakeModel) WithPort(port int) *SnowflakeModel {
	s.Port = tfconfig.IntegerVariable(port)
	return s
}

func (s *SnowflakeModel) WithPrivateKey(privateKey string) *SnowflakeModel {
	s.PrivateKey = tfconfig.StringVariable(privateKey)
	return s
}

func (s *SnowflakeModel) WithPrivateKeyPassphrase(privateKeyPassphrase string) *SnowflakeModel {
	s.PrivateKeyPassphrase = tfconfig.StringVariable(privateKeyPassphrase)
	return s
}

func (s *SnowflakeModel) WithPrivateKeyPath(privateKeyPath string) *SnowflakeModel {
	s.PrivateKeyPath = tfconfig.StringVariable(privateKeyPath)
	return s
}

func (s *SnowflakeModel) WithProfile(profile string) *SnowflakeModel {
	s.Profile = tfconfig.StringVariable(profile)
	return s
}

func (s *SnowflakeModel) WithProtocol(protocol string) *SnowflakeModel {
	s.Protocol = tfconfig.StringVariable(protocol)
	return s
}

func (s *SnowflakeModel) WithRegion(region string) *SnowflakeModel {
	s.Region = tfconfig.StringVariable(region)
	return s
}

func (s *SnowflakeModel) WithRequestTimeout(requestTimeout int) *SnowflakeModel {
	s.RequestTimeout = tfconfig.IntegerVariable(requestTimeout)
	return s
}

func (s *SnowflakeModel) WithRole(role string) *SnowflakeModel {
	s.Role = tfconfig.StringVariable(role)
	return s
}

// session_params attribute type is not yet supported, so WithSessionParams can't be generated

func (s *SnowflakeModel) WithTmpDirectoryPath(tmpDirectoryPath string) *SnowflakeModel {
	s.TmpDirectoryPath = tfconfig.StringVariable(tmpDirectoryPath)
	return s
}

func (s *SnowflakeModel) WithToken(token string) *SnowflakeModel {
	s.Token = tfconfig.StringVariable(token)
	return s
}

// token_accessor attribute type is not yet supported, so WithTokenAccessor can't be generated

func (s *SnowflakeModel) WithUser(user string) *SnowflakeModel {
	s.User = tfconfig.StringVariable(user)
	return s
}

func (s *SnowflakeModel) WithUsername(username string) *SnowflakeModel {
	s.Username = tfconfig.StringVariable(username)
	return s
}

func (s *SnowflakeModel) WithValidateDefaultParameters(validateDefaultParameters string) *SnowflakeModel {
	s.ValidateDefaultParameters = tfconfig.StringVariable(validateDefaultParameters)
	return s
}

func (s *SnowflakeModel) WithWarehouse(warehouse string) *SnowflakeModel {
	s.Warehouse = tfconfig.StringVariable(warehouse)
	return s
}

//////////////////////////////////////////
// below it's possible to set any value //
//////////////////////////////////////////

func (s *SnowflakeModel) WithAccountValue(value tfconfig.Variable) *SnowflakeModel {
	s.Account = value
	return s
}

func (s *SnowflakeModel) WithAccountNameValue(value tfconfig.Variable) *SnowflakeModel {
	s.AccountName = value
	return s
}

func (s *SnowflakeModel) WithAuthenticatorValue(value tfconfig.Variable) *SnowflakeModel {
	s.Authenticator = value
	return s
}

func (s *SnowflakeModel) WithBrowserAuthValue(value tfconfig.Variable) *SnowflakeModel {
	s.BrowserAuth = value
	return s
}

func (s *SnowflakeModel) WithClientIpValue(value tfconfig.Variable) *SnowflakeModel {
	s.ClientIp = value
	return s
}

func (s *SnowflakeModel) WithClientRequestMfaTokenValue(value tfconfig.Variable) *SnowflakeModel {
	s.ClientRequestMfaToken = value
	return s
}

func (s *SnowflakeModel) WithClientStoreTemporaryCredentialValue(value tfconfig.Variable) *SnowflakeModel {
	s.ClientStoreTemporaryCredential = value
	return s
}

func (s *SnowflakeModel) WithClientTimeoutValue(value tfconfig.Variable) *SnowflakeModel {
	s.ClientTimeout = value
	return s
}

func (s *SnowflakeModel) WithDisableConsoleLoginValue(value tfconfig.Variable) *SnowflakeModel {
	s.DisableConsoleLogin = value
	return s
}

func (s *SnowflakeModel) WithDisableQueryContextCacheValue(value tfconfig.Variable) *SnowflakeModel {
	s.DisableQueryContextCache = value
	return s
}

func (s *SnowflakeModel) WithDisableTelemetryValue(value tfconfig.Variable) *SnowflakeModel {
	s.DisableTelemetry = value
	return s
}

func (s *SnowflakeModel) WithDriverTracingValue(value tfconfig.Variable) *SnowflakeModel {
	s.DriverTracing = value
	return s
}

func (s *SnowflakeModel) WithExternalBrowserTimeoutValue(value tfconfig.Variable) *SnowflakeModel {
	s.ExternalBrowserTimeout = value
	return s
}

func (s *SnowflakeModel) WithHostValue(value tfconfig.Variable) *SnowflakeModel {
	s.Host = value
	return s
}

func (s *SnowflakeModel) WithIncludeRetryReasonValue(value tfconfig.Variable) *SnowflakeModel {
	s.IncludeRetryReason = value
	return s
}

func (s *SnowflakeModel) WithInsecureModeValue(value tfconfig.Variable) *SnowflakeModel {
	s.InsecureMode = value
	return s
}

func (s *SnowflakeModel) WithJwtClientTimeoutValue(value tfconfig.Variable) *SnowflakeModel {
	s.JwtClientTimeout = value
	return s
}

func (s *SnowflakeModel) WithJwtExpireTimeoutValue(value tfconfig.Variable) *SnowflakeModel {
	s.JwtExpireTimeout = value
	return s
}

func (s *SnowflakeModel) WithKeepSessionAliveValue(value tfconfig.Variable) *SnowflakeModel {
	s.KeepSessionAlive = value
	return s
}

func (s *SnowflakeModel) WithLoginTimeoutValue(value tfconfig.Variable) *SnowflakeModel {
	s.LoginTimeout = value
	return s
}

func (s *SnowflakeModel) WithMaxRetryCountValue(value tfconfig.Variable) *SnowflakeModel {
	s.MaxRetryCount = value
	return s
}

func (s *SnowflakeModel) WithOauthAccessTokenValue(value tfconfig.Variable) *SnowflakeModel {
	s.OauthAccessToken = value
	return s
}

func (s *SnowflakeModel) WithOauthClientIdValue(value tfconfig.Variable) *SnowflakeModel {
	s.OauthClientId = value
	return s
}

func (s *SnowflakeModel) WithOauthClientSecretValue(value tfconfig.Variable) *SnowflakeModel {
	s.OauthClientSecret = value
	return s
}

func (s *SnowflakeModel) WithOauthEndpointValue(value tfconfig.Variable) *SnowflakeModel {
	s.OauthEndpoint = value
	return s
}

func (s *SnowflakeModel) WithOauthRedirectUrlValue(value tfconfig.Variable) *SnowflakeModel {
	s.OauthRedirectUrl = value
	return s
}

func (s *SnowflakeModel) WithOauthRefreshTokenValue(value tfconfig.Variable) *SnowflakeModel {
	s.OauthRefreshToken = value
	return s
}

func (s *SnowflakeModel) WithOcspFailOpenValue(value tfconfig.Variable) *SnowflakeModel {
	s.OcspFailOpen = value
	return s
}

func (s *SnowflakeModel) WithOktaUrlValue(value tfconfig.Variable) *SnowflakeModel {
	s.OktaUrl = value
	return s
}

func (s *SnowflakeModel) WithOrganizationNameValue(value tfconfig.Variable) *SnowflakeModel {
	s.OrganizationName = value
	return s
}

func (s *SnowflakeModel) WithParamsValue(value tfconfig.Variable) *SnowflakeModel {
	s.Params = value
	return s
}

func (s *SnowflakeModel) WithPasscodeValue(value tfconfig.Variable) *SnowflakeModel {
	s.Passcode = value
	return s
}

func (s *SnowflakeModel) WithPasscodeInPasswordValue(value tfconfig.Variable) *SnowflakeModel {
	s.PasscodeInPassword = value
	return s
}

func (s *SnowflakeModel) WithPasswordValue(value tfconfig.Variable) *SnowflakeModel {
	s.Password = value
	return s
}

func (s *SnowflakeModel) WithPortValue(value tfconfig.Variable) *SnowflakeModel {
	s.Port = value
	return s
}

func (s *SnowflakeModel) WithPrivateKeyValue(value tfconfig.Variable) *SnowflakeModel {
	s.PrivateKey = value
	return s
}

func (s *SnowflakeModel) WithPrivateKeyPassphraseValue(value tfconfig.Variable) *SnowflakeModel {
	s.PrivateKeyPassphrase = value
	return s
}

func (s *SnowflakeModel) WithPrivateKeyPathValue(value tfconfig.Variable) *SnowflakeModel {
	s.PrivateKeyPath = value
	return s
}

func (s *SnowflakeModel) WithProfileValue(value tfconfig.Variable) *SnowflakeModel {
	s.Profile = value
	return s
}

func (s *SnowflakeModel) WithProtocolValue(value tfconfig.Variable) *SnowflakeModel {
	s.Protocol = value
	return s
}

func (s *SnowflakeModel) WithRegionValue(value tfconfig.Variable) *SnowflakeModel {
	s.Region = value
	return s
}

func (s *SnowflakeModel) WithRequestTimeoutValue(value tfconfig.Variable) *SnowflakeModel {
	s.RequestTimeout = value
	return s
}

func (s *SnowflakeModel) WithRoleValue(value tfconfig.Variable) *SnowflakeModel {
	s.Role = value
	return s
}

func (s *SnowflakeModel) WithSessionParamsValue(value tfconfig.Variable) *SnowflakeModel {
	s.SessionParams = value
	return s
}

func (s *SnowflakeModel) WithTmpDirectoryPathValue(value tfconfig.Variable) *SnowflakeModel {
	s.TmpDirectoryPath = value
	return s
}

func (s *SnowflakeModel) WithTokenValue(value tfconfig.Variable) *SnowflakeModel {
	s.Token = value
	return s
}

func (s *SnowflakeModel) WithTokenAccessorValue(value tfconfig.Variable) *SnowflakeModel {
	s.TokenAccessor = value
	return s
}

func (s *SnowflakeModel) WithUserValue(value tfconfig.Variable) *SnowflakeModel {
	s.User = value
	return s
}

func (s *SnowflakeModel) WithUsernameValue(value tfconfig.Variable) *SnowflakeModel {
	s.Username = value
	return s
}

func (s *SnowflakeModel) WithValidateDefaultParametersValue(value tfconfig.Variable) *SnowflakeModel {
	s.ValidateDefaultParameters = value
	return s
}

func (s *SnowflakeModel) WithWarehouseValue(value tfconfig.Variable) *SnowflakeModel {
	s.Warehouse = value
	return s
}
