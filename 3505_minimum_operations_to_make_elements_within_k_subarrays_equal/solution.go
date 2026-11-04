import (
    "math"
    "sort"
)

func minOperations(nums []int, x int, k int) int {
    n := len(nums)
    
    // Precompute the minimum cost to make each subarray of size x have equal elements.
    // subarrayCosts[i] will store the min cost for nums[i : i+x]
    subarrayCosts := make([]int, n-x+1)
    for i := 0; i <= n-x; i++ {
        subArray := make([]int, x)
        copy(subArray, nums[i:i+x])
        sort.Ints(subArray)
        median := subArray[x/2]
        cost := 0
        for _, val := range subArray {
            cost += abs(val - median)
        }
        subarrayCosts[i] = cost
    }
    
    // dp[i][j] = minimum operations to form exactly j non-overlapping subarrays of size x
    // using the first i elements of nums (i.e., nums[0...i-1]).
    // Dimensions: (n + 1) x (k + 1)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, k+1)
        for j := range dp[i] {
            dp[i][j] = math.MaxInt32
        }
    }
    
    // Base case: 0 operations are needed to form 0 subarrays.
    for i := 0; i <= n; i++ {
        dp[i][0] = 0
    }
    
    // Fill the DP table
    for j := 1; j <= k; j++ {
        for i := 1; i <= n; i++ {
            // Option 1: Do not form the j-th subarray ending at index i-1.
            // In this case, the minimum cost is the same as forming j subarrays using the first i-1 elements.
            dp[i][j] = dp[i-1][j]
            
            // Option 2: Form the j-th subarray ending at index i-1.
            // This subarray spans indices [i-x, i-1]. This is only possible if i >= x.
            // The cost for this subarray is subarrayCosts[i-x].
            // The previous j-1 subarrays must have been formed using the first i-x elements.
            if (i >= x) {
                prevCost := dp[i-x][j-1]
                currentSubarrayCost := subarrayCosts[i-x]
                
                if (prevCost != math.MaxInt32) {
                    dp[i][j] = min(dp[i][j], prevCost + currentSubarrayCost)
                }
            }
        }
    }
    
    // The result is the minimum cost to form k subarrays using all n elements.
    // Since the problem asks for 'at least k', and costs are non-negative, the minimum for 'at least k'
    // will be achieved by exactly k subarrays. If a solution with k+m subarrays has cost C,
    // then a solution with k subarrays (by removing m subarrays) will have cost <= C.
    return dp[n][k]
}

func abs(a int) int {
    if a < 0 {
        return -a
    }
    return a
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
