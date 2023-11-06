package douyinx

import (
	"context"
	"errors"
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
	return &AccessToken{
		conf:            conf,
		c:               c,
		accessTokenKey:  conf.CachePrefix + ":access_token",
		refreshTokenKey: conf.CachePrefix + ":refresh_token",
	}
}

func (a *AccessToken) genAccessTokenKey(openId string) string {
	return a.accessTokenKey + ":" + openId
}

func (a *AccessToken) genRefreshTokenKey(openId string) string {
	return a.refreshTokenKey + ":" + openId
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

	if err := a.setAccessToken(ctx, accessTokenItem.OpenId, accessTokenItem.AccessToken, time.Duration(accessTokenItem.ExpiresIn)*time.Second); err != nil {
		return types.GetAccessTokenItem{}, err
	}

	if err := a.setRefreshToken(ctx, accessTokenItem.OpenId, accessTokenItem.RefreshToken, time.Duration(accessTokenItem.RefreshExpiresIn)*time.Second); err != nil {
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

	if err := a.setAccessToken(ctx, accessTokenItem.OpenId, accessTokenItem.AccessToken, time.Duration(accessTokenItem.ExpiresIn)*time.Second); err != nil {
		return types.GetAccessTokenItem{}, err
	}

	if err := a.setRefreshToken(ctx, accessTokenItem.OpenId, accessTokenItem.RefreshToken, time.Duration(accessTokenItem.RefreshExpiresIn)*time.Second); err != nil {
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

	return refreshItem, a.setRefreshToken(ctx, openId, refreshToken, time.Duration(refreshItem.ExpiresIn)*time.Second)
}

func (a *AccessToken) setAccessToken(ctx context.Context, openId, token string, expiration time.Duration) error {
	if err := a.c.Set(ctx, a.genAccessTokenKey(openId), token, expiration, func(ctx context.Context) {
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
	}); err != nil {
		return err
	}

	return nil
}

func (a *AccessToken) setRefreshToken(ctx context.Context, openId, token string, expiration time.Duration) error {
	if err := a.c.Set(ctx, a.genRefreshTokenKey(openId), token, expiration, func(ctx context.Context) {
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
	}); err != nil {
		return err
	}

	return nil
}

func (a *AccessToken) GetAccessToken(ctx context.Context, openId string) (string, error) {
	value, ok, err := a.c.Get(ctx, a.genAccessTokenKey(openId))
	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("需要授权")
	}

	return value.(string), nil
}

func (a *AccessToken) GetRefreshToken(ctx context.Context, openId string) (string, error) {
	value, ok, err := a.c.Get(ctx, a.genRefreshTokenKey(openId))
	if err != nil {
		return "", err
	}

	if !ok {
		return "", errors.New("需要授权")
	}

	return value.(string), nil
}
