package main

import (
	"math"
)

// Define a helper min function for int, as math.Min works on float64
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type Solution struct{}

func (s *Solution) minimumSumSubarray(nums []int, l int, r int) int {
	n := len(nums)

	// Calculate prefix sums
	// prefixSum[k] stores the sum of nums[0...k-1]
	// prefixSum[0] = 0
	// Use int64 for prefix sums to prevent potential overflow
	prefixSum := make([]int64, n+1)
	for i := 0; i < n; i++ {
		prefixSum[i+1] = prefixSum[i] + int64(nums[i])
	}

	minPositiveSum := math.MaxInt32 // Initialize with max int value

	// Iterate over all possible start indices i
	for i := 0; i < n; i++ {
		// Iterate over all possible end indices j
		// j must be at least i to form a non-empty subarray
		for j := i; j < n; j++ {
			// Calculate the length of the current subarray nums[i...j]
			currentLength := j - i + 1

			// Check if the length constraint is met
			if currentLength >= l && currentLength <= r {
				// Calculate the sum of the subarray nums[i...j] using prefix sums
				currentSumLong := prefixSum[j+1] - prefixSum[i]

				// Check if the sum is positive and fits into an int
				if currentSumLong > 0 && currentSumLong <= int64(math.MaxInt32) {
					currentSum := int(currentSumLong)
					// Update the minimum positive sum found so far
					minPositiveSum = min(minPositiveSum, currentSum)
				}
			}
		}
	}

	// If minPositiveSum is still math.MaxInt32, it means no such subarray was found
	if minPositiveSum == math.MaxInt32 {
		return -1;
	} else {
		return minPositiveSum;
	}
}