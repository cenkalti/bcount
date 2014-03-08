package bcount

import "testing"

func Test1(t *testing.T) {
	c := New(100, 0.1)

	t.Logf("Cap: %d", c.Cap())

	c.Add([]byte("cenk"))
	if c.Count() != 1 {
		t.Fatalf("Invalid count: %d", c.Count())
	}

	c.Add([]byte("cenk"))
	if c.Count() != 1 {
		t.Fatalf("Invalid count: %d", c.Count())
	}

	c.Add([]byte("alti"))
	if c.Count() != 2 {
		t.Fatalf("Invalid count: %d", c.Count())
	}

	c.Reset()
	if c.Count() != 0 {
		t.Fatalf("Invalid count: %d", c.Count())
	}
}
