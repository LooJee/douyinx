package auth

import (
	"context"
	"fmt"
	"github.com/loojee/douyinx/cache"
	"github.com/loojee/douyinx/config"
	"github.com/loojee/douyinx/constants"
	"github.com/loojee/douyinx/traffic"
	"github.com/loojee/douyinx/types"
	"time"
)

type ClientToken struct {
	config *config.Config
	cache  cache.Cache
	key    string
}

func NewClientTokenRefresher(config *config.Config, c cache.Cache) *ClientToken {
	token := &ClientToken{
		config: config,
		cache:  c,
		key:    config.CachePrefix + ":client_token",
	}

	token.mustRefresh(context.Background())

	return token
}

func (r *ClientToken) mustRefresh(ctx context.Context) {
	for {
		if err := r.refresh(ctx); err != nil {
			fmt.Println("refresh client token error: ", err)
			time.Sleep(time.Second)
			continue
		}

		break
	}
}

func (r *ClientToken) refresh(ctx context.Context) (err error) {
	defer func() {
		if e := recover(); e != nil {
			err = fmt.Errorf("refresh client token panic")
			return
		}
	}()

	rsp := types.GetClientTokenResp{}
	if err := traffic.PostJSON(ctx, constants.UriClientToken, types.GetClientTokenReq{
		ClientKey:    r.config.ClientKey,
		ClientSecret: r.config.ClientSecret,
		GrantType:    constants.GrantTypeClientCredential,
	}, &rsp); err != nil {
		return err
	}

	if rsp.Data.ErrorCode != 0 {
		return &rsp.Data.CommonData
	}

	if err := r.cache.Set(ctx, r.key, rsp.Data.AccessToken, time.Duration(rsp.Data.ExpiresIn)*time.Second, r.mustRefresh); err != nil {
		return err
	}

	return nil
}

// GetToken 获取 client_token
func (r *ClientToken) GetToken(ctx context.Context) (string, error) {
	value, ok, err := r.cache.Get(ctx, r.key)
	if err != nil {
		return "", err
	}

	if !ok {
		r.mustRefresh(ctx)
	}

	return value.(string), nil
}
