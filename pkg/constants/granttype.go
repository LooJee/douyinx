package constants

type GrantType string

const (
	GrantTypeClientCredential  = "client_credential"  // 生成 client_token
	GrantTypeAuthorizationCode = "authorization_code" // 获取 access_token
	GrantTypeRefreshToken      = "refresh_token"      // 刷新 access_token
)
