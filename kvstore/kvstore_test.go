package kvstore

import (
	"testing"

	"go.melnyk.org/concept"
)

func TestEmpty(t *testing.T) {
	var store = &Empty[string, int]{}
	var err error
	var ok bool

	// Load
	_, err = store.Load("foo")
	if err != concept.ErrKeyNotFound {
		t.Errorf("Load: expected ErrKeyNotFound, got %v", err)
	}

	// Exists
	ok, err = store.Exists("foo")
	if err != nil {
		t.Errorf("Exists: unexpected error %v", err)
	}
	if ok {
		t.Errorf("Exists: expected false, got true")
	}

	// Save
	err = store.Save("foo", 42)
	if err != concept.ErrNotImplemented {
		t.Errorf("Save: expected ErrNotImplemented, got %v", err)
	}

	// Delete
	err = store.Delete("foo")
	if err != concept.ErrKeyNotFound {
		t.Errorf("Delete: expected ErrKeyNotFound, got %v", err)
	}

	// Enumerate
	err = store.Enumerate(func(key string, value int) bool {
		t.Errorf("Enumerate: unexpected callback")
		return false
	})
	if err != nil {
		t.Errorf("Enumerate: unexpected error %v", err)
	}
}
