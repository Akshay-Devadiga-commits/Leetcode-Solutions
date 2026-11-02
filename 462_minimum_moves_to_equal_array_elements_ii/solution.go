import (
    "sort"
    "math"
)

func minMoves2(nums []int) int {
    sort.Ints(nums)
    median := nums[len(nums)/2]
    moves := 0
    for _, num := range nums {
        moves += int(math.Abs(float64(num - median)))
    }
    return moves
}