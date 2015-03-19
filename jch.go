// Package jch provides an implementation of the Jump Consistent Hash
// consistent hashing algorithm.
//
// Consistent hashing is designed to minimise hash changes when the number of
// buckets is changed, and is particularly useful for data sharding.  More
// information on consistent hashing is available at
// http://en.wikipedia.org/wiki/Consistent_hashing.
//
// Jump Consistent Hash was invented by John Lamping and Eric Veach, and is
// described in the paper "A Fast, Minimal Memory, Consistent Hash Algorithm"
// (2014) available at http://arxiv.org/abs/1406.2294v1.
package jch

// Hash generates a Jump Consistent Hash given a key and a number of buckets.
func Hash(key uint64, numBuckets int32) int32 {
	var b int64 = -1
	var j int64
	for j < int64(numBuckets) {
		b = j
		key = key*2862933555777941757 + 1
		j = int64(float64(b+1) * (float64(1<<31) / float64((key>>33)+1)))
	}
	return int32(b)
}
