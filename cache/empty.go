package cache

import (
	"time"

	"go.melnyk.org/concept"
)

// Empty is an empty cache implementation.
type Empty[K comparable, V any] struct {
}

func (*Empty[K, V]) Get(key K) (V, bool) {
	var v V
	return v, false
}

func (*Empty[K, V]) Set(key K, val V, ttl time.Duration) {
}

func (*Empty[K, V]) Delete(key ...K) {
}

func (*Empty[K, V]) Reset() error {
	return concept.ErrNotImplemented
}

var (
	_ concept.Cache[any, any] = &Empty[any, any]{}
)
