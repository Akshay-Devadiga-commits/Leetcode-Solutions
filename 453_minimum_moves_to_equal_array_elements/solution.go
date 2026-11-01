import "math"

func minMoves(nums []int) int {
    mn := nums[0]
    for _, num := range nums {
        if num < mn {
            mn = num
        }
    }
    sum := 0
    for _, num := range nums {
        sum += num - mn
    }
    return sum
}