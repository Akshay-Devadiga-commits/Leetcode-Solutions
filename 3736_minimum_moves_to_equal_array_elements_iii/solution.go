import "math"

func minMoves(nums []int) int {
    if len(nums) == 0 {
        return 0
    }
    
    maxVal := nums[0]
    totalSum := 0
    
    for _, num := range nums {
        if num > maxVal {
            maxVal = num
        }
        totalSum += num
    }
    
    // The total number of moves is the sum of (max_val - num) for each num in nums.
    // This can be rewritten as (n * max_val) - sum(nums).
    // Use int64 for intermediate calculation to prevent potential overflow.
    return int(int64(len(nums)) * int64(maxVal) - int64(totalSum))
}