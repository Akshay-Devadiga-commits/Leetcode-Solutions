func uniquePaths(grid [][]int) int {
    MOD := int64(1000000007)
    m := len(grid)
    n := len(grid[0])
    dp := make([][]int64, m)
    for i := range dp {
        dp[i] = make([]int64, n)
    }
    dp[0][0] = 1
    
    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if dp[i][j] == 0 {
                continue
            }
            
            if j + 1 < n {
                nr, nc := i, j + 1
                d := 'R'
                for nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
                    if d == 'R' {
                        nr++
                        d = 'D'
                    } else {
                        nc++
                        d = 'R'
                    }
                }
                if nr >= 0 && nr < m && nc >= 0 && nc < n {
                    dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD
                }
            }
            
            if i + 1 < m {
                nr, nc := i + 1, j
                d := 'D'
                for nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1 {
                    if d == 'R' {
                        nr++
                        d = 'D'
                    } else {
                        nc++
                        d = 'R'
                    }
                }
                if nr >= 0 && nr < m && nc >= 0 && nc < n {
                    dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD
                }
            }
        }
    }
    return int(dp[m - 1][n - 1])
}