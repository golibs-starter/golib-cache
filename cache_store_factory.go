package golibcache

import (
	"crypto/tls"
	"fmt"
	"github.com/eko/gocache/v2/store"
	"github.com/go-redis/redis/v8"
	gc "github.com/patrickmn/go-cache"
	"time"
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
	defaultExpiration, err := time.ParseDuration(properties.DefaultExpiration)
	if err != nil {
		return nil, fmt.Errorf("memory cache: parse default expiration: %v: %v", properties.DefaultExpiration, err)
	}
	cleanupInterval, err := time.ParseDuration(properties.CleanupInterval)
	if err != nil {
		return nil, fmt.Errorf("memory cache: parse cleanup interval: %v: %v", properties.CleanupInterval, err)
	}
	client := gc.New(defaultExpiration, cleanupInterval)
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
