package main

import (
	"fmt"
	"sync"
)

func sum(nums []int) int {
	total := 0

	for _, num := range nums {
		total += num
	}

	return total
}

func concurrentSum(nums []int) int {
	numOfChunks := 4
	sums := make([]int, numOfChunks)
	var wg sync.WaitGroup
	wg.Add(numOfChunks)

	chunkSize := (len(nums) + numOfChunks - 1) / numOfChunks

	fmt.Println("Chunk Size:", chunkSize)

	for i := 0; i < numOfChunks; i++ {
		go func(i int) {
			defer wg.Done()
			start := i * chunkSize
			end := (i + 1) * chunkSize

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
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	fmt.Println("Numbers:", slice)

	total := concurrentSum(slice)
	fmt.Println("The Sum of the elements of slice is:", total)
}
