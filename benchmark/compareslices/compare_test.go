package compareslices_test

import (
	"bench/compareslices"
	"strconv"
	"testing"
)

func MakeSlices() ([]string, []string) {
	aSliceSize := 100
	bSliceSize := 50

	sliceA := make([]string, 0, aSliceSize)
	sliceB := make([]string, 0, bSliceSize)

	for i := 0; i < aSliceSize; i++ {
		sliceA = append(sliceA, strconv.Itoa(i))
	}

	for i := 0; i < bSliceSize; i++ {
		sliceB = append(sliceB, strconv.Itoa(i))
	}

	return sliceA, sliceB
}

func BenchmarkCompareSlicesUsingMap(b *testing.B) {
	sliceA, sliceB := MakeSlices()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		compareslices.CompareSlicesUsingMap(sliceA, sliceB)
	}
}

func BenchmarkCompareSlicesUsingNestedLoop(b *testing.B) {
	sliceA, sliceB := MakeSlices()

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		compareslices.CompareSlicesUsingNestedLoop(sliceA, sliceB)
	}
}
