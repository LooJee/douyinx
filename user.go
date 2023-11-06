package douyinx

import (
	"context"
	"github.com/loojee/douyinx/pkg/constants"
	"github.com/loojee/douyinx/pkg/errorx"
	"github.com/loojee/douyinx/pkg/traffic"
	"github.com/loojee/douyinx/types"
)

type User struct {
	accessToken *AccessToken
}

func NewUser(accessToken *AccessToken) *User {
	return &User{accessToken: accessToken}
}

func (u *User) GetUserInfo(ctx context.Context, openId string) (userInfo types.GetUserInfoItem, err error) {
	accessToken, err := u.accessToken.GetAccessToken(ctx, openId)
	if err != nil {
		return userInfo, err
	}
	req := types.GetUserInfoReq{
		OpenId:      openId,
		AccessToken: accessToken,
	}

	var resp types.GetUserInfoResp

	if err := traffic.PostUrlEncodeForm(ctx, constants.UriUserInfo, req, &resp); err != nil {
		return userInfo, err
	}

	if bizErr := errorx.CatchBizError(resp.Data.CommonData); bizErr != nil {
		return userInfo, bizErr
	}

	return resp.Data.GetUserInfoItem, nil
}
