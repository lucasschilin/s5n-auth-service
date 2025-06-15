package cache

import "time"

type Cache interface {
	Set(key string, val any, exp time.Duration) error
	Get(key string) (string, error)
}
