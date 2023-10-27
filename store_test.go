package main

import (
	"errors"
	"testing"
)

func TestDelete(t *testing.T) {
	putKey := "putKey"
	putVal := "putVal"
	// use put to add key to the store
	_ = Put(putKey, putVal)

	defer delete(store,putKey)

	_ = Delete(putKey)

	// try to get the deleted value
	_, err := Get(putKey)
	if !errors.Is(err, ErrNoSuchKey) {
		t.Errorf("get err:%v", err)
	}

}

func TestGet(t *testing.T) {
	putKey := "putKey"
	putVal := "putVal"
	// use put to add key to the store
	_ = Put(putKey, putVal)

	defer delete(store,putKey)

	expectedVal := "putVal"
	// try to get the actual value
	actualVal, err := Get(putKey)
	if err != nil {
		t.Errorf("get err:%v", err)
	}

	if expectedVal != actualVal {
		t.Errorf("Get: expected %s, got %s", expectedVal, actualVal)
	}
}

func TestPut(t *testing.T) {
	initKey := "putKey"
	initVal := "putVal"

	defer delete(store,initKey)

	// use put to add key to the store
	_ = Put(initKey, initVal)

	changedVal := "changedVal"
	// use put to change the value in the store
	_ = Put(initKey, changedVal)

	// get the actual value
	actualVal, err := Get(initKey)
	if err != nil {
		t.Errorf("get err:%v", err)
	}

	if changedVal != actualVal {
		t.Errorf("Put: expected %s, got %s", changedVal, actualVal)
	}
}
