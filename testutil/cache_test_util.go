package golibCacheTestUtil

import (
	golibcache "gitlab.com/golibs-starter/golib-cache"
)

var cache *golibcache.Cache

// Cache return cache instance
func Cache() *golibcache.Cache {
	return cache
}
