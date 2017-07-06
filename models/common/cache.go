package common

import (
	"time"

	"github.com/iteny/hmgo/go-cache"
)

var Cache *cache.Cache

func init() {
	Cache = cache.New(1*time.Minute, 1*time.Minute)
}
func CacheSetConfineTime(key string, val interface{}) {
	Cache.Set(key, val, cache.DefaultExpiration)
}
func CacheSetAlwaysTime(key string, val interface{}) {
	Cache.Set(key, val, cache.NoExpiration)
}
func CacheGet(key string) (interface{}, bool) {
	val, found := Cache.Get(key)
	return val, found
}
func CacheDel(key string) {
	Cache.Delete(key)
}
