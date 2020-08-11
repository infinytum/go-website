package ui

import (
	"net/http"

	"github.com/infinytum/go-website/pkg/context"
)

var ApplicationContext *context.Application

func NewPageContext(title string) *context.PageContext {
	return &context.PageContext{
		Title: title,
	}
}

// ListenAndServe always returns a non-nil error.
func ListenAndServe(ctx *context.Application) error {
	ApplicationContext = ctx
	http.HandleFunc("/service/maps", maptoken)
	http.HandleFunc("/assets/", assets)
	return http.ListenAndServe(":8090", nil)
}
