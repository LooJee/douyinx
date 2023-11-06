package douyinx

import (
	"github.com/loojee/douyinx/auth"
	"github.com/loojee/douyinx/cache"
	"github.com/loojee/douyinx/config"
	"github.com/loojee/douyinx/traffic"
)

type App struct {
	clientToken *auth.ClientToken
	accessToken *auth.AccessToken
}

func NewApp(conf config.Config, c cache.Cache) (*App, error) {
	traffic.MustInit(&conf)

	app := App{
		clientToken: auth.NewClientTokenRefresher(&conf, c),
		accessToken: auth.NewAccessToken(&conf, c),
	}

	return &app, nil
}
