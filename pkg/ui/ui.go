package ui

import (
	"net/http"

	"github.com/infinytum/go-website/pkg/context"
)

var ApplicationContext *context.Application

func NewPageContext(title string, req http.Request) *context.PageContext {
	lang := req.Header.Get("Accept-Language")
	return &context.PageContext{
		Title:    title,
		Language: ApplicationContext.LanguageService().GetOrLoad(lang).Map(),
	}
}

// ListenAndServe always returns a non-nil error.
func ListenAndServe(ctx *context.Application) error {
	ApplicationContext = ctx
	http.HandleFunc("/service/maps", maptoken)
	http.HandleFunc("/assets/", assets)
	return http.ListenAndServe(":8090", nil)
}
