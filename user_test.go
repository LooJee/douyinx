package douyinx

import (
	"context"
	"github.com/loojee/douyinx/config"
	"github.com/loojee/douyinx/pkg/cache"
	"github.com/loojee/douyinx/pkg/traffic"
	"github.com/loojee/douyinx/types"
	"os"
	"testing"
)

func TestUser_GetUserInfo(t *testing.T) {
	conf := config.NewConfig(os.Getenv("client_key"), os.Getenv("client_secret"))
	traffic.MustInit(conf)
	c := cache.NewDefaultCache()
	accessToken := NewAccessToken(conf, c)
	clientToken := NewClientTokenRefresher(conf, c)

	token, err := accessToken.FetchToken(context.Background(), "a4316d37fc2442faRLSZhRjmLsQQi08I0f7h")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("token: ", token)

	user := NewUser(accessToken, clientToken)

	userInfo, err := user.GetUserInfo(context.Background(), token.OpenId)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("user: ", userInfo)
}

func TestUser_RoleCheck(t *testing.T) {
	conf := config.NewConfig(os.Getenv("client_key"), os.Getenv("client_secret"))
	traffic.MustInit(conf)
	c := cache.NewDefaultCache()
	accessToken := NewAccessToken(conf, c)
	clientToken := NewClientTokenRefresher(conf, c)

	user := NewUser(accessToken, clientToken)

	userInfo, err := user.RoleCheck(context.Background(), os.Getenv("open_id"), []types.RoleLabel{types.RoleLabelCompanyBand, types.RoleLabelAuthCompany, types.RoleLabelStaff})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("user: ", userInfo)
}
