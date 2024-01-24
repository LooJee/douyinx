package types

type GetOauthQrcodeReq struct {
	ClientKey            string `mapstructure:"client_key"`
	Scope                string `mapstructure:"scope"`
	Next                 string `mapstructure:"next"`
	JumpType             string `mapstructure:"jump_type"`
	OptionalScopeCheck   string `mapstructure:"optional_scope_check"`
	OptionalScopeUncheck string `mapstructure:"optional_scope_uncheck"`
	CustomizeParams      string `mapstructure:"customize_params"`
	State                string `mapstructure:"state"`
}

type GetOauthQrcodeResp struct {
	Data    OauthQrcodeData `json:"data"`
	Message string          `json:"message"`
}

type OauthQrcodeData struct {
	CommonData
	IsFrontier     bool   `json:"is_frontier"`
	Qrcode         string `json:"qrcode"`
	QrcodeIndexUrl string `json:"qrcode_index_url"`
	Token          string `json:"token"`
}

type OauthQrcodeState string

const (
	OauthQrcodeStateNew       OauthQrcodeState = "new"
	OauthQrcodeStateConfirmed OauthQrcodeState = "confirmed"
	OauthQrcodeStateExpired   OauthQrcodeState = "expired"
)

type CheckOauthQrcodeReq struct {
	ClientKey            string           `mapstructure:"client_key"`
	Scope                string           `mapstructure:"scope"`
	Next                 string           `mapstructure:"next"`
	JumpType             string           `mapstructure:"jump_type"`
	OptionalScopeCheck   string           `mapstructure:"optional_scope_check"`
	OptionalScopeUncheck string           `mapstructure:"optional_scope_uncheck"`
	CustomizeParams      string           `mapstructure:"customize_params"`
	State                OauthQrcodeState `mapstructure:"state"`
	Token                string           `mapstructure:"token"`
	Timestamp            string           `mapstructure:"timestamp"`
}

type CheckOauthQrcodeResp struct {
	Data    CheckOauthQrcodeData `json:"data"`
	Message string               `json:"message"`
}

type CheckOauthQrcodeData struct {
	OauthQrcodeData
	Status string `json:"status"`
}
