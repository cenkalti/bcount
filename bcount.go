// Package bcount implements a counter for counting distinct values.
package bcount

import "github.com/willf/bloom"

// BCount is a distict value counter that uses a bloom filter internally.
// It uses constant space in memory regardless of number of items counted.
// Mehtods of this struct are not goroutine-safe.
type BCount struct {
	filter *bloom.BloomFilter
	count  uint64
}

// New creates a new counter for about n items with fp false positive rate.
func New(n uint, fp float64) *BCount {
	return &BCount{
		filter: bloom.NewWithEstimates(n, fp),
	}
}

// Add adds data to to counter. The count will only be incremented if the data
// is not added before.
func (c *BCount) Add(data []byte) {
	if !c.filter.TestAndAdd(data) {
		c.count++
	}
}

// Count returns the count of distinct items.
func (c *BCount) Count() uint64 {
	return c.count
}

// Cap returns the capacity of a Bloom filter.
func (c *BCount) Cap() uint {
	return c.filter.Cap()
}

// Clear clears all the data in a Bloom filter, removing all keys.
func (c *BCount) Reset() {
	c.count = 0
	c.filter.ClearAll()
}
