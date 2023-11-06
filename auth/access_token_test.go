package auth

import (
	"context"
	"github.com/loojee/douyinx/cache"
	"github.com/loojee/douyinx/config"
	"github.com/loojee/douyinx/traffic"
	"os"
	"testing"
)

func TestAccessToken_AccessToken(t *testing.T) {
	conf := config.NewConfig(os.Getenv("client_key"), os.Getenv("client_secret"))
	traffic.MustInit(conf)
	c := cache.NewDefaultCache()

	accessToken := NewAccessToken(conf, c)

	token, err := accessToken.FetchToken(context.Background(), "a4316d37fc2442faBDxVLRz2O1imofZ4wYps")
	if err != nil {
		t.Fatal(err)
	}

	t.Log(token)

	token, err = accessToken.RenewAccessToken(context.Background(), token.RefreshToken)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(token)

	refreshToken, err := accessToken.RenewRefreshToken(context.Background(), token.RefreshToken)
	if err != nil {
		panic(err)
	}

	t.Log("refresh token: ", refreshToken)
}
