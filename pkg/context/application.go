package context

import (
	"github.com/allegro/bigcache"
	"github.com/infinytum/go-website/internal/config"
	"github.com/infinytum/go-website/pkg/language"
	"github.com/shurcooL/graphql"
)

type Application struct {
	config   config.Config
	cache    *bigcache.BigCache
	language *language.Language
	squidex  *graphql.Client
}

func (ctx *Application) Config() config.Config {
	return ctx.config
}

func NewApplication(config config.Config) *Application {
	return &Application{
		config: config,
	}
}
