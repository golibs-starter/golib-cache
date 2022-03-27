# Golib Cache

## Installation

```shell
go get gitlab.com/golibs-starter/golib-cache
```

## Configuration 
```yaml
app.cache:
  driver: "memory" #support memory, redis
  # If use memory
  memory:
    defaultExpiration: "30s" # 30s, 30m, 30h 
    cleanupInterval: "1s" # 1s, 1m, 1h
  # If use redis
  redis:
    host: localhost
    port: 6379
    database: 0
    user: username
    password: secret
    enableTLS: true #default: false
```

## Usage

Register to fx container

```go
package bootstrap

import (
	"gitlab.com/golibs-starter/golib-cache"
	"go.uber.org/fx"
)

func Register() fx.Option {
	return golibcache.EnableCache()
}
```

Remember function will get value in the cache if exists, if not exists, it will set to cache

```go
package app

import (
	"gitlab.com/golibs-starter/golib-cache"
	"time"
)

type NeedCache struct {
	cache *golibcache.Cache
}

func (nc *NeedCache) UseRemember() {
	// String
	str, err := nc.cache.Remember("key", 30*time.Second, func() (interface{}, error) {
		return "value", nil
	})
	// Number
	num, err := nc.cache.Remember("key", 30*time.Second, func() (interface{}, error) {
		return 100, nil
	})
	// Struct
	value, err := nc.cache.Remember("key", 30*time.Second, func() (interface{}, error) {
		return &Example{Data: "data"}, nil
	})
	data := value.(*Example).Data
}
```