package stringconcat_test

import (
	"bench/stringconcat"
	"testing"
)

func BenchmarkWithSprintf(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringconcat.WithSprintf()
	}
}

func BenchmarkWithPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringconcat.WithPlus()
	}
}
