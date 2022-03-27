package golibcache

import (
	"gitlab.com/golibs-starter/golib/config"
)

// CacheProperties represents ...
type CacheProperties struct {
	Driver string
	Memory MemoryCacheProperties
	Redis  RedisCacheProperties
}

// MemoryCacheProperties represent memory cache properties
type MemoryCacheProperties struct {
	DefaultExpiration string
	CleanupInterval   string
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
