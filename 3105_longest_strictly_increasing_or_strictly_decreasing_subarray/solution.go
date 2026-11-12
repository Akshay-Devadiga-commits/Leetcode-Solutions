func longestMonotonicSubarray(nums []int) int {
    ans := 1
    increasing := 1
    decreasing := 1

    for i := 1; i < len(nums); i++ {
        if nums[i] > nums[i-1] {
            increasing++
            decreasing = 1
        } else if nums[i] < nums[i-1] {
            decreasing++
            increasing = 1
        } else {
            increasing = 1
            decreasing = 1
        }
        if increasing > ans {
            ans = increasing
        }
        if decreasing > ans {
            ans = decreasing
        }
    }

    return ans
}