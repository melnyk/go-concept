package concept

// KVStore is a generic key-value storage interface.
type KVStore[K comparable, V any] interface {
    // Load returns the value associated with the given key.
    // If the key is not present, it returns the zero value and ErrKeyNotFound.
    Load(key K) (V, error)

    // Exists returns true if the key is present in the storage.
    // Returns (false, nil) if the key is absent; returns a non-nil error only for storage failures.
    Exists(key K) (bool, error)

    // Save associates the given value with the given key.
    Save(key K, value V) error

	// Delete removes the key from the storage.
	Delete(key K) error

	// Enumerate all entries in the storage.
	// The callback function is called for each entry.
	// If the callback returns false, the enumeration is stopped.
	Enumerate(callback func(key K, value V) bool) error
}
