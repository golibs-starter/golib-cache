package golibCacheTestUtil

import (
	golibcache "gitlab.com/golibs-starter/golib-cache"
	"go.uber.org/fx"
)

func EnableCacheTestUtil() fx.Option {
	return fx.Invoke(func(c *golibcache.Cache) {
		cache = c
	})
}
