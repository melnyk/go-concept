package kvstore

import (
	"go.melnyk.org/concept"
)

// Empty is an empty key-value store implementation.
type Empty[K comparable, V any] struct {
}

func (*Empty[K, V]) Load(key K) (V, error) {
	var v V
	return v, concept.ErrKeyNotFound
}

func (*Empty[K, V]) Exists(key K) (bool, error) {
	return false, nil
}

func (*Empty[K, V]) Save(key K, value V) error {
	return concept.ErrNotImplemented
}

func (*Empty[K, V]) Delete(key K) error {
	return concept.ErrKeyNotFound
}

func (*Empty[K, V]) Enumerate(callback func(key K, value V) bool) error {
	return nil
}

var (
	_ concept.KVStore[any, any] = &Empty[any, any]{}
)
