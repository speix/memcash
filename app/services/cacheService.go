package services

import (
	"errors"
	"os"
	"strconv"
	"time"

	"github.com/speix/memcash/app"
)

func (e *Engine) Set(key, value string, duration time.Duration) error {

	e.Cache.Mu.Lock()
	defer e.Cache.Mu.Unlock()

	_, found := e.Cache.Items[key]
	if found {
		return errors.New("Item with the selected key is already set ")
	}

	if len(key) == 0 || len(value) == 0 || duration == 0 {
		return errors.New("Key, value and duration must be set correctly ")
	}

	e.Cache.Items[key] = app.Item{
		Value:      value,
		Expiration: time.Now().Add(duration).Unix(),
	}

	return nil
}

func (e *Engine) Get(key string) (*app.Item, error) {

	e.Cache.Mu.Lock()
	defer e.Cache.Mu.Unlock()

	item, found := e.Cache.Items[key]
	if !found {
		return nil, errors.New("Item not found ")
	}

	if time.Now().Unix() > item.Expiration {
		return nil, errors.New("Item has expired and will be purged ")
	}

	return &item, nil
}

func (e *Engine) Unset(key string) {
	e.Cache.Mu.Lock()
	defer e.Cache.Mu.Unlock()

	delete(e.Cache.Items, key)
}

func (e *Engine) Purge() {
	now := time.Now().Unix()

	for k, v := range e.Cache.Items {

		if now > v.Expiration {
			e.Unset(k)
		}

	}

}

func NewCache() *app.Cache {

	duration, err := strconv.Atoi(os.Getenv("MEMCASH_TICKER_INTERVAL"))
	if err != nil {
		duration = 300
	}

	return &app.Cache{
		Items: make(map[string]app.Item),
		Supervisor: &app.Supervisor{
			Interval: time.Duration(duration) * time.Second,
			Stop:     make(chan bool),
		},
	}
}
