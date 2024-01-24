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

func TestTool_GetJsbTicket(t *testing.T) {
	conf := config.NewConfig(os.Getenv("client_key"), os.Getenv("client_secret"))
	traffic.MustInit(conf)
	c := cache.NewDefaultCache()
	clientToken := NewClientTokenRefresher(conf, c)

	toolSrv := NewTool(clientToken)

	ticket, err := toolSrv.GetJsbTicket(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Log("ticket: ", ticket)
}

func TestTool_GetOauthQrcode(t *testing.T) {
	conf := config.NewConfig(os.Getenv("client_key"), os.Getenv("client_secret"))
	traffic.MustInit(conf)
	c := cache.NewDefaultCache()
	clientToken := NewClientTokenRefresher(conf, c)

	data, err := NewTool(clientToken).GetOauthQrcode(context.TODO(), types.GetOauthQrcodeReq{
		Scope: os.Getenv("scope"),
		Next:  os.Getenv("redirect_url"),
		State: "hello",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("data: ", data)
}

func TestTool_CheckOauthQrcode(t *testing.T) {
	conf := config.NewConfig(os.Getenv("client_key"), os.Getenv("client_secret"))
	traffic.MustInit(conf)
	c := cache.NewDefaultCache()
	clientToken := NewClientTokenRefresher(conf, c)

	data, err := NewTool(clientToken).CheckOauthQrcode(context.TODO(), types.CheckOauthQrcodeReq{
		Scope: os.Getenv("scope"),
		Next:  os.Getenv("redirect_url"),
		State: "hello",
		Token: "335f6a69a58b6b3b9b8c23cb6f85d7fa_lf",
	})
	if err != nil {
		t.Fatal(err)
	}

	t.Log("data: ", data)
}
