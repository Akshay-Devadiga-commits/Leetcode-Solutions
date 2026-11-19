import (
    "sort"
)

func minimizeSum(nums []int) int {
    sort.Ints(nums)
    // Can always change the number to any other number in `nums`, so `low` becomes 0.
    // Thus, rephrase the problem as finding the minimum `high`.
    highOfChangingTwoMins := nums[len(nums)-1] - nums[2]
    highOfChangingTwoMaxs := nums[len(nums)-3] - nums[0]
    highOfChangingMinAndMax := nums[len(nums)-2] - nums[1]

    minVal := highOfChangingTwoMins
    if highOfChangingTwoMaxs < minVal {
        minVal = highOfChangingTwoMaxs
    }
    if highOfChangingMinAndMax < minVal {
        minVal = highOfChangingMinAndMax
    }
    return minVal
}