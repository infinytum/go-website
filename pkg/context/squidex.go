package context

import (
	"context"
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/shurcooL/graphql"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
)

func (ctx *Application) Squidex() *graphql.Client {
	if ctx.squidex == nil {
		data := url.Values{
			"grant_type":    {"client_credentials"},
			"client_id":     {ctx.config.Get("squidex", "clientId").String("")},
			"client_secret": {ctx.config.Get("squidex", "clientSecret").String("")},
			"scope":         {"squidex-api"},
		}

		src := oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: "nope"},
		)
		httpClient := oauth2.NewClient(context.Background(), src)

		resp, err := http.PostForm(ctx.config.Get("squidex", "url").String(""), data)
		if err != nil {
			logrus.Error("Invalid CMS credentials")
			return graphql.NewClient(ctx.config.Get("graphql", "url").String(""), httpClient)
		}
		var res map[string]interface{}
		if json.NewDecoder(resp.Body).Decode(&res) != nil {
			logrus.Error("Invalid CMS credentials")
			return graphql.NewClient(ctx.config.Get("graphql", "url").String(""), httpClient)
		}

		token, ok := res["access_token"]
		if !ok {
			logrus.Error("Invalid CMS credentials")
			return graphql.NewClient(ctx.config.Get("graphql", "url").String(""), httpClient)
		}

		src = oauth2.StaticTokenSource(
			&oauth2.Token{AccessToken: token.(string)},
		)
		httpClient = oauth2.NewClient(context.Background(), src)

		ctx.squidex = graphql.NewClient(ctx.config.Get("graphql", "url").String(""), httpClient)
	}
	return ctx.squidex
}
