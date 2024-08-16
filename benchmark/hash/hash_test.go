package hash_test

import (
	"bench/hash"
	"testing"
)

var testString = "55252f79-8874-473b-b0de-c6816a0ef06e.431091077177.rollout-tenants.0"

func BenchmarkSum256(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hash.Sum256(testString)
	}
}

func BenchmarkSumMD5(b *testing.B) {
	for i := 0; i < b.N; i++ {
		hash.SumMD5(testString)
	}
}
