package concept

import "time"

// Cache is a generic cache interface.
type Cache[K comparable, V any] interface {
	// Get returns the value stored in the cache for a key, or nil if no value is present.
	Get(key K) (V, bool)

	// Set stores a value for a key.
	Set(key K, v V, ttl time.Duration)

	// Delete removes a key from the cache.
	Delete(key ...K)

	// Reset clears the cache.
	Reset() error
}
