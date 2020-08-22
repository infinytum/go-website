package context

import "github.com/infinytum/go-website/pkg/language"

func (ctx *Application) LanguageService() *language.Language {
	if ctx.language == nil {
		ctx.language = &language.Language{}
	}
	return ctx.language
}
