package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"
)

func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func calculateChunkSize(datasize int, numOfChunks int) int {
	chunkSize := math.Ceil(float64(datasize) / float64(numOfChunks))

	return int(chunkSize)
}

func concurrentSum(nums []int) int {
	numOfChunks := runtime.NumCPU()

	if numOfChunks > len(nums) {
		numOfChunks = len(nums)
	}

	sums := make([]int, numOfChunks)

	var wg sync.WaitGroup
	wg.Add(numOfChunks)

	chunkSize := calculateChunkSize(len(nums), numOfChunks)

	for i := 0; i < numOfChunks; i++ {
		go func(i int) {
			defer wg.Done()
			start := i * chunkSize
			end := (i + 1) * chunkSize

			if start >= len(nums) {
				return
			}

			if end > len(nums) {
				end = len(nums)
			}

			sums[i] = sum(nums[start:end])
		}(i)
	}

	wg.Wait()

	return sum(sums)
}

func main() {
	slice := []int{}

	for i := 1; i <= 100000000; i++ {
		slice = append(slice, i)
	}

	start := time.Now()

	total := concurrentSum(slice)
	fmt.Println("The Sum of the elements of slice is:", total)

	fmt.Println("Time taken:", time.Since(start))
}
