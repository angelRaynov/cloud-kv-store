package main

import (
	"errors"
	"sync"
)

var store = struct {
	sync.RWMutex
	m map[string]string
}{m: make(map[string]string)}

var ErrNoSuchKey = errors.New("no such key")

func Put(k, v string) error {
	store.Lock()
	defer store.Unlock()
	store.m[k] = v
	return nil
}

func Get(key string) (string, error) {
	store.RLock()
	defer store.RUnlock()
	val, ok := store.m[key]
	if !ok {
		return "", ErrNoSuchKey
	}
	return val, nil
}

func Delete(key string) error {
	store.Lock()
	defer store.Unlock()
	delete(store.m, key)
	return nil
}
