package golibcache

import (
	"crypto/tls"
	"fmt"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	gc "github.com/patrickmn/go-cache"
)

func NewStore(properties *CacheProperties) (store.StoreInterface, error) {
	switch properties.Driver {
	case "memory":
		return NewMemoryStore(properties.Memory)
	case "redis":
		return NewRedisStore(properties.Redis)
	default:
		return nil, fmt.Errorf("cache driver: %s is not supported", properties.Driver)
	}
}

func NewMemoryStore(properties MemoryCacheProperties) (store.StoreInterface, error) {
	client := gc.New(properties.DefaultExpiration, properties.CleanupInterval)
	return store.NewGoCache(client, nil), nil
}

func NewRedisStore(properties RedisCacheProperties) (store.StoreInterface, error) {
	options := &redis.Options{
		Addr:     fmt.Sprintf("%s:%d", properties.Host, properties.Port),
		Username: properties.User,
		Password: properties.Password,
		DB:       properties.Database,
	}
	if properties.EnableTLS {
		options.TLSConfig = &tls.Config{
			MinVersion: tls.VersionTLS12,
		}
	}
	return store.NewRedis(redis.NewClient(options), nil), nil
}
