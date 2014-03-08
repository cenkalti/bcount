package bcount

import (
	"fmt"
	"testing"
)

func Example() {
	// Create a new counter for 100 items with 1% false positive rate
	counter := New(100, 0.01)

	// Count some values
	counter.Add([]byte("foo"))
	counter.Add([]byte("bar"))
	counter.Add([]byte("baz"))

	// Get the counter value
	fmt.Println(counter.Count())
	// Output: 3
}

func TestBCount(t *testing.T) {
	// Create a new counter
	c := New(100, 0.01)

	// Prints capacity
	t.Logf("Cap: %d", c.Cap())

	// Count a value
	c.Add([]byte("cenk"))
	if c.Count() != 1 {
		t.Fatalf("Invalid count: %d", c.Count())
	}

	// Count existing value
	c.Add([]byte("cenk"))
	if c.Count() != 1 {
		t.Fatalf("Invalid count: %d", c.Count())
	}

	// Count new value
	c.Add([]byte("alti"))
	if c.Count() != 2 {
		t.Fatalf("Invalid count: %d", c.Count())
	}

	// Test reset
	c.Reset()
	if c.Count() != 0 {
		t.Fatalf("Invalid count: %d", c.Count())
	}
}
