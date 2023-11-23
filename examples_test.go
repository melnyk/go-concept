//go:build ignore

package main

import (
	"go.melnyk.org/concept/cache"
	"go.melnyk.org/concept/kvstore"
)

func ExampleEmptyCache() {
	var c Cache[any, any] = &cache.Empty[any, any]{}
	c.Set("key", "value", 0)
	_, ok := c.Get("key")
	c.Delete("key")
	c.Reset()
	// Output:
}

func ExampleEmptyKVStore() {
	var kv KVStore[any, any] = &kvstore.Empty[any, any]{}
	kv.Save("key", "value")
	_, ok := kv.Load("key")
	_, err := kv.Exists("key")
	kv.Enumerate(func(key any, value any) bool {
		return true
	})
	kv.Delete("key")
	// Output:
}
