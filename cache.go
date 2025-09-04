package concept

import "time"

// Cache is a generic cache interface.
type Cache[K comparable, V any] interface {
    // Get returns the value and a boolean indicating presence.
    // Returns (zero value, false) if the key is absent or expired.
    Get(key K) (V, bool)

    // Set stores a value for a key.
    // Zero TTL means no expiration.
    Set(key K, v V, ttl time.Duration)

    // Delete removes one or more keys from the cache.
    Delete(key ...K)

    // Reset clears the cache.
    Reset() error
}
