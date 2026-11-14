import "math"

func minimumSumSubarray(nums []int, l int, r int) int {
    n := len(nums)
    minPositiveSum := math.MaxInt32
    found := false

    // Iterate through all possible start indices of the subarray
    for i := 0; i < n; i++ {
        currentSum := 0
        // Iterate through all possible end indices of the subarray starting from i
        for j := i; j < n; j++ {
            currentSum += nums[j]
            subarrayLength := j - i + 1

            // Check if the subarray length is within the given range [l, r]
            if l <= subarrayLength && subarrayLength <= r {
                // Check if the sum is positive
                if currentSum > 0 {
                    if currentSum < minPositiveSum {
                        minPositiveSum = currentSum
                    }
                    found = true
                }
            }
        }
    }

    // If no such subarray was found, return -1, otherwise return the minimum positive sum
    if !found {
        return -1
    }
    return minPositiveSum
}