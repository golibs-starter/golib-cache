package golibCacheTestUtil

import (
	"gitlab.com/golibs-starter/golib-cache"
	"go.uber.org/fx"
)

var cache *golibcache.Cache

func EnableCacheTestUtil() fx.Option {
	return fx.Invoke(func(c *golibcache.Cache) {
		cache = c
	})
}

// Cache return cache instance
func Cache() *golibcache.Cache {
	return cache
}
