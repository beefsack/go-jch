package jch

import (
	"fmt"
	"testing"
)

func TestHash(t *testing.T) {
	hashes := map[uint64]int32{}
	for key := uint64(0); key < 100000; key++ {
		for numBuckets := int32(1); numBuckets < 100; numBuckets++ {
			if numBuckets == 1 {
				// Initialise hash map.
				hashes[key] = Hash(key, numBuckets)
			} else {
				// Ensure the new bucket is equal to the last one or equal to
				// numBuckets-1.
				newHash := Hash(key, numBuckets)
				oldHash := hashes[key]
				if newHash != oldHash && newHash != numBuckets-1 {
					t.Errorf(
						"key: %d numBuckets: %d oldHash: %d newHash: %d",
						key,
						numBuckets,
						oldHash,
						newHash,
					)
				}
				hashes[key] = newHash
			}
		}
	}
}

func ExampleHash() {
	hash := Hash(28, 5)
	fmt.Print(hash)
	// Output: 2
}

func BenchmarkHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Hash(5123451233512351256, 512456124)
	}
}
