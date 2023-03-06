package fizzbuzz

import (
	"sync"
	"testing"
)

func TestRoundTrip(t *testing.T) {
	c := NewCounter()
	c.Increment("foo")
	c.Increment("foo")
	c.Increment("bar")
	c.Increment("foo")

	max := c.Max()

	if max.Key != "foo" {
		t.Errorf("expected foo, got %s", max.Key)
	}
	if max.Count != 3 {
		t.Errorf("expected 3, got %d", max.Count)
	}
}

func TestCounterConcurrent(t *testing.T) {
	times := 100_000

	c := NewCounter()
	var wg sync.WaitGroup
	for i := 0; i < times; i++ {
		wg.Add(1)
		go func() {
			c.Increment("foo")
			wg.Done()
		}()
	}
	wg.Wait()
	max := c.Max()
	if max.Key != "foo" {
		t.Errorf("expected foo, got %s", max.Key)
	}
	if max.Count != times {
		t.Errorf("expected %d, got %d", times, max.Count)
	}
}
