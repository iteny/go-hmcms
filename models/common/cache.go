package common

import (
	"time"

	"github.com/iteny/hmgo/go-cache"
)

var Cache *cache.Cache

func init() {
	Cache = cache.New(1*time.Minute, 1*time.Minute)
}
