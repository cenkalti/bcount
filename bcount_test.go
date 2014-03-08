package bcount

import (
	"fmt"
	"testing"
)

func ExampleBCount() {
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

func BenchmarkBCountAdd100(b *testing.B)     { benchmarkBCountAdd(b, 100, 0.01) }
func BenchmarkBCountAdd1000(b *testing.B)    { benchmarkBCountAdd(b, 1000, 0.01) }
func BenchmarkBCountAdd10000(b *testing.B)   { benchmarkBCountAdd(b, 10000, 0.01) }
func BenchmarkBCountAdd100000(b *testing.B)  { benchmarkBCountAdd(b, 100000, 0.01) }
func BenchmarkBCountAdd1000000(b *testing.B) { benchmarkBCountAdd(b, 1000000, 0.01) }

func benchmarkBCountAdd(b *testing.B, n uint, fp float64) {
	c := New(n, fp)
	b.Logf("Cap: %d", c.Cap())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		c.Add([]byte("foo"))
	}
}

// Benchmark simple increment operation
func BenchmarkIncrement(b *testing.B) {
	c := &Counter{}
	for i := 0; i < b.N; i++ {
		c.Increment()
	}
}

type Counter struct {
	C int
}

func (c *Counter) Increment() {
	c.C += 1
}
