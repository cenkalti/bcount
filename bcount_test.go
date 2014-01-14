package bcount

import "testing"

func Test1(t *testing.T) {
	c := New(100, 0.01)

	c.Add([]byte("cenk"))
	if c.Count() != 1 {
		t.Errorf("Invalid count: %d", c.Count())
		return
	}

	c.Add([]byte("cenk"))
	if c.Count() != 1 {
		t.Errorf("Invalid count: %d", c.Count())
		return
	}

	c.Add([]byte("alti"))
	if c.Count() != 2 {
		t.Errorf("Invalid count: %d", c.Count())
		return
	}
}
