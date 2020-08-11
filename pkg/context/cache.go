package context

import (
	"time"

	"github.com/allegro/bigcache"
)

func (ctx *Application) Cache() *bigcache.BigCache {
	if ctx.cache == nil {
		var cachingTime = ctx.config.Get("cache", "timeout").Int(600000)
		cache, err := bigcache.NewBigCache(bigcache.DefaultConfig(time.Millisecond * time.Duration(cachingTime)))
		if err != nil {
			panic(err)
		}
		ctx.cache = cache
	}
	return ctx.cache
}
