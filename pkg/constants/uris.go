package constants

const (
	UriClientToken       = "/oauth/client_token/"        // 获取 client_token
	UriFetchAccessToken  = "/oauth/access_token/"        // 获取 access_token
	UriRenewRefreshToken = "/oauth/renew_refresh_token/" // 续期 refresh_token
	UriRenewAccessToken  = "/oauth/refresh_token/"       // 刷新 access_token

	UriUserInfo  = "/oauth/userinfo/"           // 获取用户公开信息
	UriRoleCheck = "/api/douyin/v1/role/check/" // 用户经营身份管理

	UriDirectMessage = "/im/send/msg/"               // 发送私信消息
	UriImageUpload   = "/tool/imagex/client_upload/" // 上传图片
	UriRecallMessage = "/im/recall/msg/"             // 撤回消息
	UriJsbTicket     = "/js/getticket/"

	UriOauthQrcode      = "/oauth/get_qrcode/"   // 获取授权二维码
	UriCheckOauthQrcode = "/oauth/check_qrcode/" // 检查授权二维码
)
