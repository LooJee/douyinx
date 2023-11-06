package douyinx

import (
	"context"
	"github.com/loojee/douyinx/config"
	"github.com/loojee/douyinx/pkg/cache"
	"github.com/loojee/douyinx/pkg/traffic"
	"os"
	"testing"
)

func TestNewClientTokenRefresher(t *testing.T) {
	conf := config.NewConfig(os.Getenv("client_key"), os.Getenv("client_secret"))
	traffic.MustInit(conf)
	c := cache.NewDefaultCache()
	clientToken := NewClientTokenRefresher(conf, c)

	token, err := clientToken.GetToken(context.Background())
	if err != nil {
		t.Fatal(err)
	}

	t.Log(token)
}
