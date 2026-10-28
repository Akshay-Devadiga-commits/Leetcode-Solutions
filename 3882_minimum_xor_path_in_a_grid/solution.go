import "math"

func minCost(grid [][]int) int {
    m, n := len(grid), len(grid[0])
    dp := make([]map[int]bool, n)
    for i := range dp {
        dp[i] = make(map[int]bool)
    }
    
    for r := 0; r < m; r++ {
        newDp := make([]map[int]bool, n)
        for i := range newDp {
            newDp[i] = make(map[int]bool)
        }
        for c := 0; c < n; c++ {
            if r == 0 && c == 0 {
                newDp[c][grid[0][0]] = true
            } else {
                if r > 0 {
                    for val := range dp[c] {
                        newDp[c][val^grid[r][c]] = true
                    }
                }
                if c > 0 {
                    for val := range newDp[c-1] {
                        newDp[c][val^grid[r][c]] = true
                    }
                }
            }
        }
        dp = newDp
    }
    
    minVal := math.MaxInt32
    for val := range dp[n-1] {
        if val < minVal {
            minVal = val
        }
    }
    return minVal
}