func uniquePaths(grid [][]int) int {
    MOD := 1000000007
    m := len(grid)
    n := len(grid[0])
    dp_right := make([][]int, m)
    dp_down := make([][]int, m)
    for i := 0; i < m; i++ {
        dp_right[i] = make([]int, n)
        dp_down[i] = make([]int, n)
    }
    if n > 1 {
        nr, nc, nd := 0, 1, 0
        for nc < n && grid[nr][nc] == 1 {
            nr++
            nd = 0
        }
        if nr >= 0 && nr < m && nc >= 0 && nc < n {
            if nd == 0 {
                dp_right[nr][nc] = (dp_right[nr][nc] + 1) % MOD
            } else {
                dp_down[nr][nc] = (dp_down[nr][nc] + 1) % MOD
            }
        }
    }
    if m > 1 {
        nr, nc, nd := 1, 0, 1
        for nr < m && grid[nr][nc] == 1 {
            nc++
            nd = 1
        }
        if nr >= 0 && nr < m && nc >= 0 && nc < n {
            if nd == 0 {
                dp_right[nr][nc] = (dp_right[nr][nc] + 1) % MOD
            } else {
                dp_down[nr][nc] = (dp_down[nr][nc] + 1) % MOD
            }
        }
    }
    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            ways_right := dp_right[r][c]
            ways_down := dp_down[r][c]
            total_ways := (ways_right + ways_down) % MOD
            if total_ways == 0 {
                continue
            }
            if r == m - 1 && c == n - 1 {
                continue
            }
            if c + 1 < n {
                curr_r, curr_c, curr_d := r, c + 1, 0
                for curr_r >= 0 && curr_r < m && curr_c >= 0 && curr_c < n && grid[curr_r][curr_c] == 1 {
                    if curr_d == 0 {
                        curr_r++
                        curr_d = 0
                    } else {
                        curr_c++
                        curr_d = 1
                    }
                }
                if curr_r >= 0 && curr_r < m && curr_c >= 0 && curr_c < n {
                    if curr_d == 0 {
                        dp_right[curr_r][curr_c] = (dp_right[curr_r][curr_c] + total_ways) % MOD
                    } else {
                        dp_down[curr_r][curr_c] = (dp_down[curr_r][curr_c] + total_ways) % MOD
                    }
                }
            }
            if r + 1 < m {
                curr_r, curr_c, curr_d := r + 1, c, 1
                for curr_r >= 0 && curr_r < m && curr_c >= 0 && curr_c < n && grid[curr_r][curr_c] == 1 {
                    if curr_d == 0 {
                        curr_r++
                        curr_d = 0
                    } else {
                        curr_c++
                        curr_d = 1
                    }
                }
                if curr_r >= 0 && curr_r < m && curr_c >= 0 && curr_c < n {
                    if curr_d == 0 {
                        dp_right[curr_r][curr_c] = (dp_right[curr_r][curr_c] + total_ways) % MOD
                    } else {
                        dp_down[curr_r][curr_c] = (dp_down[curr_r][curr_c] + total_ways) % MOD
                    }
                }
            }
        }
    }
    return (dp_right[m - 1][n - 1] + dp_down[m - 1][n - 1]) % MOD
}