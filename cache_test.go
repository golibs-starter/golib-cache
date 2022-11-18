package golibcache

import (
	"context"
	"github.com/eko/gocache/v2/store"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type Data struct {
	Value string
}

func TestCache_Exist(t *testing.T) {
	cache, err := NewCache(&CacheProperties{
		Driver: "memory",
		Memory: MemoryCacheProperties{
			DefaultExpiration: "1m",
			CleanupInterval:   "30s",
		},
	})
	assert.Nil(t, err)
	t.Run("Not Exist", func(t *testing.T) {
		exist := cache.Exist("not_exist_key")
		assert.False(t, exist)
	})
	t.Run("Exist", func(t *testing.T) {
		err := cache.cache.Set(context.Background(), "exist_key", "value", &store.Options{Expiration: time.Minute})
		assert.Nil(t, err)
		exist := cache.Exist("exist_key")
		assert.True(t, exist)
	})
}

func TestCache_Remember(t *testing.T) {
	cache, err := NewCache(&CacheProperties{
		Driver: "memory",
		Memory: MemoryCacheProperties{
			DefaultExpiration: "1m",
			CleanupInterval:   "30s",
		},
	})
	assert.Nil(t, err)
	t.Run("Not Exist With Integer", func(t *testing.T) {
		err := cache.cache.Delete(context.Background(), "not_exist_key")
		assert.Nil(t, err)
		value, err := cache.Remember("not_exist_key", time.Minute, func() (interface{}, error) {
			return 10, nil
		})
		assert.Nil(t, err)
		assert.Equal(t, 10, value)
		time.Sleep(200 * time.Millisecond)
		v, err := cache.cache.Get(context.Background(), "not_exist_key")
		assert.Nil(t, err)
		assert.Equal(t, 10, v)
	})
	t.Run("Not Exist With String", func(t *testing.T) {
		err := cache.cache.Delete(context.Background(), "not_exist_key")
		assert.Nil(t, err)
		value, err := cache.Remember("not_exist_key", time.Minute, func() (interface{}, error) {
			return "string", nil
		})
		assert.Nil(t, err)
		assert.Equal(t, "string", value)
		time.Sleep(200 * time.Millisecond)
		v, err := cache.cache.Get(context.Background(), "not_exist_key")
		assert.Nil(t, err)
		assert.Equal(t, "string", v)
	})
	t.Run("Not Exist With Struct", func(t *testing.T) {
		err := cache.cache.Delete(context.Background(), "not_exist_key")
		assert.Nil(t, err)
		value, err := cache.Remember("not_exist_key", time.Minute, func() (interface{}, error) {
			return &Data{Value: "value"}, nil
		})
		assert.Nil(t, err)
		assert.Equal(t, "value", value.(*Data).Value)
		time.Sleep(200 * time.Millisecond)
		v, err := cache.cache.Get(context.Background(), "not_exist_key")
		assert.Nil(t, err)
		assert.Equal(t, "value", v.(*Data).Value)
	})
	t.Run("Exist With Integer", func(t *testing.T) {
		err := cache.cache.Set(context.Background(), "exist_key", 10, &store.Options{Expiration: time.Minute})
		assert.Nil(t, err)
		value, err := cache.Remember("exist_key", time.Minute, func() (interface{}, error) {
			return 10, nil
		})
		assert.Nil(t, err)
		assert.Equal(t, 10, value)
	})
	t.Run("Exist With String", func(t *testing.T) {
		err := cache.cache.Set(context.Background(), "exist_key", "string", &store.Options{Expiration: time.Minute})
		assert.Nil(t, err)
		value, err := cache.Remember("exist_key", time.Minute, func() (interface{}, error) {
			return "string", nil
		})
		assert.Nil(t, err)
		assert.Equal(t, "string", value)
	})
	t.Run("Exist With Struct", func(t *testing.T) {
		err := cache.cache.Set(context.Background(), "exist_key", &Data{Value: "value"}, &store.Options{Expiration: time.Minute})
		assert.Nil(t, err)
		value, err := cache.Remember("exist_key", time.Minute, func() (interface{}, error) {
			return &Data{Value: "value"}, nil
		})
		assert.Nil(t, err)
		assert.Equal(t, "value", value.(*Data).Value)
	})
}

func TestCache_Delete(t *testing.T) {
	cache, err := NewCache(&CacheProperties{
		Driver: "memory",
		Memory: MemoryCacheProperties{
			DefaultExpiration: "1m",
			CleanupInterval:   "30s",
		},
	})
	assert.Nil(t, err)
	t.Run("Delete", func(t *testing.T) {
		setErr := cache.cache.Set(context.Background(), "delete_key", "value", &store.Options{Expiration: time.Minute})
		assert.Nil(t, setErr)
		beforeDelete := cache.Exist("delete_key")
		assert.True(t, beforeDelete)
		err := cache.Delete("delete_key")
		assert.Nil(t, err)
		afterDelete := cache.Exist("delete_key")
		assert.False(t, afterDelete)
	})
}
