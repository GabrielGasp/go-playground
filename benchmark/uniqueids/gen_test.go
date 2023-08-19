package uniqueids_test

import (
	"bench/uniqueids"
	"testing"
)

func BenchmarkUUIDV4Google(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniqueids.UUIDV4Google()
	}
}

func BenchmarkUUIDV4Gofrs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniqueids.UUIDV4Gofrs()
	}
}

func BenchmarkUUIDV7(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniqueids.UUIDV7()
	}
}

func BenchmarkULID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniqueids.ULID()
	}
}

func BenchmarkKSUID(b *testing.B) {
	for i := 0; i < b.N; i++ {
		uniqueids.KSUID()
	}
}
