package constants

const (
	// 用户授权相关 URI

	UriClientToken       = "/oauth/client_token/"        // 获取 client_token
	UriFetchAccessToken  = "/oauth/access_token/"        // 获取 access_token
	UriRenewRefreshToken = "/oauth/renew_refresh_token/" // 续期 refresh_token
	UriRenewAccessToken  = "/oauth/refresh_token/"       // 刷新 access_token

	UriUserInfo = "/oauth/userinfo/" // 获取用户公开信息
)
