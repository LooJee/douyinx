package types

import "github.com/loojee/douyinx/constants"

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
	ExtraData ExtraData `json:"extra_data"`
	Message   string    `json:"message"`
}
