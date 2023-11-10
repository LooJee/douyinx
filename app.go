package douyinx

import (
	"github.com/loojee/douyinx/config"
	"github.com/loojee/douyinx/pkg/cache"
	"github.com/loojee/douyinx/pkg/traffic"
)

type App struct {
	c cache.Cache

	clientToken *ClientToken
	AccessToken *AccessToken

	UserClient    *User
	Webhook       *Webhook
	DirectMessage *DirectMessage
	Tool          *Tool
}

type Option func(*App)

func WithCache(c cache.Cache) func(*App) {
	return func(app *App) {
		app.c = c
	}
}

func NewApp(conf *config.Config, options ...Option) (*App, error) {
	traffic.MustInit(conf)

	app := &App{}

	for _, opt := range options {
		opt(app)
	}

	if app.c == nil {
		app.c = cache.NewDefaultCache()
	}

	app.clientToken = NewClientTokenRefresher(conf, app.c)
	app.AccessToken = NewAccessToken(conf, app.c)

	app.UserClient = NewUser(app.AccessToken)
	app.Webhook = NewWebhook()
	app.DirectMessage = NewDirectMessage(app.AccessToken)
	app.Tool = NewTool(app.clientToken)

	return app, nil
}
