package context

import (
	"github.com/infinytum/go-website/internal/config"
	"github.com/infinytum/go-website/pkg/language"
)

func (ctx *Application) LanguageService() *language.Language {
	if ctx.language == nil {
		ctx.language = &language.Language{
			Config: ctx.config,
			Cache:  make(map[string]config.Config),
		}
	}
	return ctx.language
}
