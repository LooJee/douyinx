package types

import "github.com/loojee/douyinx/pkg/constants"

type GetClientTokenReq struct {
	ClientKey    string              `json:"client_key"`
	ClientSecret string              `json:"client_secret"`
	GrantType    constants.GrantType `json:"grant_type"`
}

type ClientToken struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`
}

type GetClientTokenResp struct {
	Data struct {
		ClientToken
		CommonData
	} `json:"data"`
	Extra   ExtraData `json:"extra"`
	Message string    `json:"message"`
}

type GetAccessTokenReq struct {
	GrantType    constants.GrantType `json:"grant_type"`
	ClientKey    string              `json:"client_key"`
	ClientSecret string              `json:"client_secret"`
	Code         string              `json:"code"`
}

type GetAccessTokenItem struct {
	AccessToken      string `json:"access_token"`
	ExpiresIn        int    `json:"expires_in"`
	OpenId           string `json:"open_id"`
	RefreshExpiresIn int    `json:"refresh_expires_in"`
	RefreshToken     string `json:"refresh_token"`
	Scope            string `json:"scope"`
}

type GetAccessTokenResp struct {
	Data struct {
		GetAccessTokenItem
		CommonData
	} `json:"data"`
	Extra   ExtraData `json:"extra"`
	Message string    `json:"message"`
}

type RenewAccessTokenReq struct {
	ClientKey    string
	GrantType    constants.GrantType
	RefreshToken string
}

func (req *RenewAccessTokenReq) IntoForm() map[string]string {
	return map[string]string{
		"client_key":    req.ClientKey,
		"grant_type":    string(req.GrantType),
		"refresh_token": req.RefreshToken,
	}
}

type RenewAccessTokenResp struct {
	Data struct {
		GetAccessTokenItem
		CommonData
	} `json:"data"`
	Extra   ExtraData `json:"extra"`
	Message string    `json:"message"`
}

type RenewRefreshTokenReq struct {
	ClientKey    string
	RefreshToken string
}

func (req *RenewRefreshTokenReq) IntoForm() map[string]string {
	return map[string]string{
		"client_key":    req.ClientKey,
		"refresh_token": req.RefreshToken,
	}
}

type RenewRefreshTokenItem struct {
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type RenewRefreshTokenResp struct {
	Data struct {
		RenewRefreshTokenItem
		CommonData
	} `json:"data"`
	Extra   ExtraData `json:"extra"`
	Message string    `json:"message"`
}
