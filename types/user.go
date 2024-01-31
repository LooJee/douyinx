package types

type GetUserInfoReq struct {
	OpenId      string
	AccessToken string
}

func (r GetUserInfoReq) IntoForm() map[string]string {
	return map[string]string{
		"access_token": r.AccessToken,
		"open_id":      r.OpenId,
	}
}

type GetUserInfoItem struct {
	Avatar       string `json:"avatar"`
	AvatarLarger string `json:"avatar_larger"`
	EAccountRole string `json:"e_account_role"`
	Nickname     string `json:"nickname"`
	OpenId       string `json:"open_id"`
	UnionId      string `json:"union_id"`
	ClientKey    string `json:"client_key"`
}

type GetUserInfoResp struct {
	Data struct {
		GetUserInfoItem
		CommonData
	} `json:"data"`
	Extra   ExtraData `json:"extra"`
	Message string    `json:"message"`
}

type RoleLabel string

const (
	RoleLabelCompanyBand RoleLabel = "COMPANY_BAND"
	RoleLabelAuthCompany RoleLabel = "AUTH_COMPANY"
	RoleLabelStaff       RoleLabel = "STAFF"
	RoleLabelOpenBrand   RoleLabel = "OPEN_BRAND"
	RoleLabelOpenStaff   RoleLabel = "OPEN_STAFF"
	RoleLabelOpenPartner RoleLabel = "OPEN_PARTNER"
)

type RoleCheckReq struct {
	OpenId     string      `json:"open_id"`
	RoleLabels []RoleLabel `json:"role_labels"`
}

type RoleCheckItem struct {
	MatchResult bool               `json:"match_result"`
	FilterRole  map[RoleLabel]bool `json:"filter_role"`
}

type RoleCheckResp struct {
	Data   RoleCheckItem `json:"data"`
	ErrNo  int64         `json:"err_no"`
	ErrMsg string        `json:"err_msg"`
	LogId  string        `json:"log_id"`
}
