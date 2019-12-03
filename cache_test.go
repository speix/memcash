package main

import (
	"testing"
	"time"

	"github.com/speix/memcash/app/services"
)

var service = services.Engine{
	GRPCNetwork: nil,
	Cache:       services.NewCache(),
}

func TestAddItemToCache(t *testing.T) {

	key := "item"
	value := "data"
	expiration := time.Duration(20) * time.Second

	if err := service.Set(key, value, expiration); err != nil {
		t.Errorf("Fail: %v", err.Error())
	}

	key2 := "item2"
	value2 := "expired data"
	expiration2 := time.Duration(1) * time.Second

	if err := service.Set(key2, value2, expiration2); err != nil {
		t.Errorf("Fail: %v", err.Error())
	}

}

func TestGetItemFromCache(t *testing.T) {

	key := "item"
	expected := "data"

	item, err := service.Get(key)
	if err != nil {
		t.Errorf("Fail: %v", err.Error())
	}

	if item != nil {
		if item.Value != expected {
			t.Errorf("Fail: want %v got %v", expected, item.Value)
		}
	} else {
		t.Error("Item is nil")
	}
}

func TestItemExpired(t *testing.T) {

	key := "item2"
	expectedErr := "Item has expired and will be purged "

	time.Sleep(2 * time.Second)

	_, err := service.Get(key)
	if err != nil {
		if err.Error() != expectedErr {
			t.Errorf("Fail: want %v got %v", expectedErr, err.Error())
		}
	} else {
		t.Error("Item is not expired")
	}

}

func TestItemNotFound(t *testing.T) {

	key := "aasdfsa"
	expectedErr := "Item not found "

	_, err := service.Get(key)
	if err != nil {
		if err.Error() != expectedErr {
			t.Errorf("Fail: want %v got %v", expectedErr, err.Error())
		}
	} else {
		t.Error("Item is found")
	}

}

func TestItemIsAlreadySet(t *testing.T) {

	key := "item"
	value := "some other data"
	expiration := time.Duration(20) * time.Second
	expectedErr := "Item with the selected key is already set "

	if err := service.Set(key, value, expiration); err != nil {
		if err.Error() != expectedErr {
			t.Errorf("Fail: want %v got %v", expectedErr, err.Error())
		}
	}
}
