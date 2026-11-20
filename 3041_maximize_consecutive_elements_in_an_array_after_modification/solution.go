package main

import (
    "sort"
)

func maxSelectedElements(nums []int) int {
    sort.Ints(nums)
    dp := make(map[int]int)
    maxLength := 0
    
    for _, x := range nums {
        dpX := dp[x]
        dpXMinus1 := dp[x-1]
        
        dp[x+1] = dpX + 1
        dp[x] = dpXMinus1 + 1
        
        if dp[x+1] > maxLength {
            maxLength = dp[x+1]
        }
        if dp[x] > maxLength {
            maxLength = dp[x]
        }
    }
    
    return maxLength
}