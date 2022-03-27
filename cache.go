package golibcache

import (
	"context"
	"github.com/eko/gocache/v2/cache"
	"github.com/eko/gocache/v2/store"
	"gitlab.com/golibs-starter/golib/log"
	"time"
)

type Cache struct {
	properties *CacheProperties
	cache      *cache.Cache
}

func NewCache(properties *CacheProperties) (*Cache, error) {
	cacheStore, err := NewStore(properties)
	if err != nil {
		return nil, err
	}
	c := cache.New(cacheStore)
	return &Cache{
		properties: properties,
		cache:      c,
	}, nil
}

func (c *Cache) Exist(key string) bool {
	_, err := c.Get(key)
	return err == nil
}

func (c *Cache) Remember(key string, duration time.Duration, fn func() (interface{}, error)) (interface{}, error) {
	value, err := c.Get(key)
	if err == nil {
		return value, nil
	}
	v, err := fn()
	if err == nil {
		c.AsyncSet(key, v, duration)
	}
	return v, err
}

func (c *Cache) AsyncSet(key string, value interface{}, duration time.Duration) {
	go func() {
		err := c.cache.Set(context.Background(), key, value, &store.Options{
			Expiration: duration,
		})
		if err != nil {
			log.Warnf("cache: async set: %v", err)
		}
	}()
}

func (c *Cache) Get(key string) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel()
	value, err := c.cache.Get(ctx, key)
	if err != nil {
		log.Warnf("cache: get by key: %v: %v", key, err)
		return nil, err
	}
	return value, nil
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) error {
	return c.cache.Set(context.Background(), key, value, &store.Options{
		Expiration: duration,
	})
}
