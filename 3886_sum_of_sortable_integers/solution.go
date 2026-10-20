package main

func sortableIntegers(nums []int) int {
    n := len(nums)
    totalSum := 0

    canSortWithK := func(k int) bool {
        prevMax := -1
        for i := 0; i < n; i += k {
            drops := 0
            split := 0
            for j := 0; j < k; j++ {
                if nums[i+j] > nums[i+(j+1)%k] {
                    drops++
                    split = j + 1
                }
            }
            if drops > 1 {
                return false
            }

            minVal := nums[i+split]
            maxVal := nums[i+(split-1+k)%k]

            if prevMax != -1 && prevMax > minVal {
                return false
            }
            prevMax = maxVal
        }
        return true
    }

    for k := 1; k <= n; k++ {
        if n%k == 0 {
            if canSortWithK(k) {
                totalSum += k
            }
        }
    }

    return totalSum
}