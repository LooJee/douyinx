package douyinx

import (
	"github.com/loojee/douyinx/config"
	"github.com/loojee/douyinx/traffic"
)

type App struct {
}

func NewApp(config config.Config) (*App, error) {
	traffic.MustInit()

	return &App{}, nil
}
