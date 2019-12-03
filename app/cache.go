package app

import (
	"sync"
	"time"
)

type Cache struct {
	Supervisor *Supervisor
	Items      map[string]Item
	Mu         sync.Mutex
}

type CacheService interface {
	Get(key string) (*Item, error)
	Set(key, value string, duration time.Duration) error
	Unset(key string)
	Purge()
}
