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
