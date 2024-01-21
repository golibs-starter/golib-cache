package golibcache

import (
	"github.com/golibs-starter/golib/config"
	"time"
)

// CacheProperties represents ...
type CacheProperties struct {
	Driver string
	Memory MemoryCacheProperties
	Redis  RedisCacheProperties
}

// MemoryCacheProperties represent memory cache properties
type MemoryCacheProperties struct {
	DefaultExpiration time.Duration `default:"30s"`
	CleanupInterval   time.Duration `default:"30s"`
}

// RedisCacheProperties represents redis cache properties
type RedisCacheProperties struct {
	Host      string
	Port      int
	Database  int
	User      string
	Password  string
	EnableTLS bool
}

// NewCacheProperties return a new CacheProperties instance
func NewCacheProperties(loader config.Loader) (*CacheProperties, error) {
	props := CacheProperties{}
	err := loader.Bind(&props)
	return &props, err
}

// Prefix return config prefix
func (t *CacheProperties) Prefix() string {
	return "app.cache"
}
