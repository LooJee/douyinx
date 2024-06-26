package douyinx

import (
	"context"
	"fmt"
	"github.com/loojee/douyinx/config"
	"github.com/loojee/douyinx/pkg/cache"
	"github.com/loojee/douyinx/pkg/constants"
	"github.com/loojee/douyinx/pkg/errorx"
	"github.com/loojee/douyinx/pkg/traffic"
	"github.com/loojee/douyinx/types"
	"time"
)

type AccessToken struct {
	conf            *config.Config
	c               cache.Cache
	accessTokenKey  string
	refreshTokenKey string
}

func NewAccessToken(conf *config.Config, c cache.Cache) *AccessToken {
	token := &AccessToken{
		conf:            conf,
		c:               c,
		accessTokenKey:  conf.CachePrefix + ":access_token",
		refreshTokenKey: conf.CachePrefix + ":refresh_token",
	}

	token.c.SetExpireHook(token.accessTokenKey, token.AccessTokenExpireHook)
	token.c.SetExpireHook(token.refreshTokenKey, token.RefreshTokenExpireHook)

	return token
}

// FetchToken 获取 access_token
func (a *AccessToken) FetchToken(ctx context.Context, code string) (types.GetAccessTokenItem, error) {
	resp := types.GetAccessTokenResp{}
	err := traffic.PostJSON(ctx, constants.UriFetchAccessToken, types.GetAccessTokenReq{
		GrantType:    constants.GrantTypeAuthorizationCode,
		ClientKey:    a.conf.ClientKey,
		ClientSecret: a.conf.ClientSecret,
		Code:         code,
	}, &resp)

	if err != nil {
		return types.GetAccessTokenItem{}, err
	}

	if bizErr := errorx.CatchBizError(resp.Data.CommonData); bizErr != nil {
		return types.GetAccessTokenItem{}, bizErr
	}

	accessTokenItem := resp.Data.GetAccessTokenItem

	if err := a.c.Set(ctx, a.accessTokenKey, accessTokenItem.OpenId, accessTokenItem.AccessToken, time.Duration(accessTokenItem.ExpiresIn)*time.Second); err != nil {
		return types.GetAccessTokenItem{}, err
	}

	if err := a.c.Set(ctx, a.refreshTokenKey, accessTokenItem.OpenId, accessTokenItem.RefreshToken, time.Duration(accessTokenItem.RefreshExpiresIn)*time.Second); err != nil {
		return types.GetAccessTokenItem{}, err
	}

	return accessTokenItem, err
}

// RenewAccessToken 刷新 access_token
func (a *AccessToken) RenewAccessToken(ctx context.Context, refreshToken string) (types.GetAccessTokenItem, error) {
	resp := types.RenewAccessTokenResp{}

	err := traffic.PostUrlEncodeForm(ctx, constants.UriRenewAccessToken, &types.RenewAccessTokenReq{
		ClientKey:    a.conf.ClientKey,
		GrantType:    constants.GrantTypeRefreshToken,
		RefreshToken: refreshToken,
	}, &resp)

	if err != nil {
		return types.GetAccessTokenItem{}, err
	}

	if bizErr := errorx.CatchBizError(resp.Data.CommonData); bizErr != nil {
		return types.GetAccessTokenItem{}, bizErr
	}

	accessTokenItem := resp.Data.GetAccessTokenItem

	if err := a.c.Set(ctx, a.accessTokenKey, accessTokenItem.OpenId, accessTokenItem.AccessToken, time.Duration(accessTokenItem.ExpiresIn)*time.Second); err != nil {
		return types.GetAccessTokenItem{}, err
	}

	if err := a.c.Set(ctx, a.refreshTokenKey, accessTokenItem.OpenId, accessTokenItem.RefreshToken, time.Duration(accessTokenItem.RefreshExpiresIn)*time.Second); err != nil {
		return types.GetAccessTokenItem{}, err
	}

	return accessTokenItem, err
}

// RenewRefreshToken 刷新 refresh_token
func (a *AccessToken) RenewRefreshToken(ctx context.Context, openId, refreshToken string) (types.RenewRefreshTokenItem, error) {
	resp := types.RenewRefreshTokenResp{}

	err := traffic.PostUrlEncodeForm(ctx, constants.UriRenewRefreshToken, &types.RenewRefreshTokenReq{
		ClientKey:    a.conf.ClientKey,
		RefreshToken: refreshToken,
	}, &resp)

	if err != nil {
		return types.RenewRefreshTokenItem{}, err
	}

	if bizErr := errorx.CatchBizError(resp.Data.CommonData); bizErr != nil {
		return types.RenewRefreshTokenItem{}, bizErr
	}

	refreshItem := resp.Data.RenewRefreshTokenItem

	return refreshItem, a.c.Set(ctx, a.refreshTokenKey, openId, refreshItem.RefreshToken, time.Duration(refreshItem.ExpiresIn)*time.Second)
}

func (a *AccessToken) RefreshTokenExpireHook(ctx context.Context, openId string) {
	refreshToken, err := a.GetRefreshToken(ctx, openId)
	if err != nil {
		fmt.Printf("get refresh token for openId %s failed: %v\n", openId, err)
		return
	}

	_, err = a.RenewRefreshToken(ctx, openId, refreshToken)
	if err != nil {
		fmt.Printf("renew refresh token for openId %s failed: %v\n", openId, err)
		return
	}
}

func (a *AccessToken) AccessTokenExpireHook(ctx context.Context, openId string) {
	refreshToken, err := a.GetRefreshToken(ctx, openId)
	if err != nil {
		fmt.Printf("get refresh token for openId %s failed: %v\n", openId, err)
		return
	}
	_, err = a.RenewAccessToken(ctx, refreshToken)
	if err != nil {
		fmt.Printf("renew access token for openId %s failed: %v\n", openId, err)
		return
	}
}

func (a *AccessToken) GetAccessToken(ctx context.Context, openId string) (string, error) {
	value, ok, err := a.c.Get(ctx, a.accessTokenKey, openId)
	if err != nil {
		return "", err
	}

	if !ok {
		return "", &errorx.BizError{
			Code:    -1,
			Message: "需要授权",
		}
	}

	return value.(string), nil
}

func (a *AccessToken) GetRefreshToken(ctx context.Context, openId string) (string, error) {
	value, ok, err := a.c.Get(ctx, a.refreshTokenKey, openId)
	if err != nil {
		return "", err
	}

	if !ok {
		return "", &errorx.BizError{
			Code:    -1,
			Message: "需要授权",
		}
	}

	return value.(string), nil
}

func (a *AccessToken) ResetToken(ctx context.Context, openId string) error {
	if err := a.c.Del(ctx, a.accessTokenKey, openId); err != nil {
		return err
	}

	return a.c.Del(ctx, a.refreshTokenKey, openId)
}
