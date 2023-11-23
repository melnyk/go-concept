package cache

import (
	"testing"
	"time"

	"go.melnyk.org/concept"
)

func TestEmpty(t *testing.T) {
	var c Empty[int, int]

	if _, ok := c.Get(0); ok {
		t.Error("unexpected ok")
	}

	c.Set(0, 0, time.Second)

	c.Delete(0)

	if err := c.Reset(); err != concept.ErrNotImplemented {
		t.Error("unexpected error")
	}
}
