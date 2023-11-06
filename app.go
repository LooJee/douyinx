package douyinx

import (
	"github.com/loojee/douyinx/config"
	"github.com/loojee/douyinx/pkg/cache"
	"github.com/loojee/douyinx/pkg/traffic"
)

type App struct {
	clientToken *ClientToken
	accessToken *AccessToken

	userClient *User
}

func NewApp(conf config.Config, c cache.Cache) (*App, error) {
	traffic.MustInit(&conf)

	app := App{
		clientToken: NewClientTokenRefresher(&conf, c),
		accessToken: NewAccessToken(&conf, c),
	}

	app.userClient = NewUser(app.accessToken)

	return &app, nil
}

func (a *App) User() *User {
	return a.userClient
}
