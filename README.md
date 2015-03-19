[![GoDoc](https://godoc.org/github.com/beefsack/go-jch?status.png)](http://godoc.org/github.com/beefsack/go-jch) [![Build Status](https://travis-ci.org/beefsack/go-jch.svg?branch=master)](https://travis-ci.org/beefsack/go-jch)

Package jch provides an implementation of the Jump Consistent Hash
consistent hashing algorithm in Go.

Consistent hashing is designed to minimise hash changes when the number of
buckets is changed, and is particularly useful for data sharding.  More
information on consistent hashing is available at
http://en.wikipedia.org/wiki/Consistent_hashing.

Jump Consistent Hash was invented by John Lamping and Eric Veach, and is
described in the paper "A Fast, Minimal Memory, Consistent Hash Algorithm"
(2014) available at http://arxiv.org/abs/1406.2294v1.

```go
import "github.com/beefsack/go-jch"

func ExampleHash() {
	hash := jch.Hash(28, 5)
	fmt.Print(hash)
	// Output: 2
}
```
