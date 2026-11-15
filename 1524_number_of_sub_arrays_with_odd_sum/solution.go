package main

type Solution struct{}

func (s *Solution) NumOfSubarrays(arr []int) int {
    const MOD int = 1e9 + 7 // 10^9 + 7
    oddCount := 0
    evenCount := 1 // Represents the prefix sum of 0 before the array starts
    currentSumParity := 0
    totalOddSubarrays := 0

    for _, num := range arr {
        // Update current sum parity
        // If num is odd, parity flips. If num is even, parity stays the same.
        currentSumParity = (currentSumParity + num) % 2

        if currentSumParity == 1 { // Current prefix sum is odd
            // We need previous prefix sums that were even to make the subarray sum odd
            totalOddSubarrays = (totalOddSubarrays + evenCount) % MOD
            oddCount++
        } else { // Current prefix sum is even
            // We need previous prefix sums that were odd to make the subarray sum odd
            totalOddSubarrays = (totalOddSubarrays + oddCount) % MOD
            evenCount++
        }
    }

    return totalOddSubarrays
}