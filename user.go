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
	clientToken *ClientToken
}

func NewUser(accessToken *AccessToken, clientToken *ClientToken) *User {
	return &User{accessToken: accessToken, clientToken: clientToken}
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

func (u *User) RoleCheck(ctx context.Context, openId string, roles []types.RoleLabel) (types.RoleCheckItem, error) {
	req := types.RoleCheckReq{
		OpenId:     openId,
		RoleLabels: roles,
	}

	clientToken, err := u.clientToken.GetToken(ctx)
	if err != nil {
		return types.RoleCheckItem{}, err
	}

	var resp types.RoleCheckResp

	if err := traffic.PostJSON(ctx, constants.UriRoleCheck, req, &resp, traffic.WithAccessTokenHeader(clientToken)); err != nil {
		return types.RoleCheckItem{}, err
	}

	if resp.ErrNo != 0 {
		return types.RoleCheckItem{}, &errorx.BizError{
			Code:    int(resp.ErrNo),
			Message: resp.ErrMsg,
		}
	}

	return resp.Data, nil
}
