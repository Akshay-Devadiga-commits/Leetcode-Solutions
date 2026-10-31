class Solution {
    public int uniquePaths(int[][] grid) {
        int MOD = 1000000007;
        int m = grid.length;
        int n = grid[0].length;
        int[][] dp_right = new int[m][n];
        int[][] dp_down = new int[m][n];
        if (n > 1) {
            int nr = 0, nc = 1, nd = 0;
            while (nc < n && grid[nr][nc] == 1) {
                nr++;
                nd = 0;
            }
            if (nr >= 0 && nr < m && nc >= 0 && nc < n) {
                if (nd == 0) {
                    dp_right[nr][nc] = (dp_right[nr][nc] + 1) % MOD;
                } else {
                    dp_down[nr][nc] = (dp_down[nr][nc] + 1) % MOD;
                }
            }
        }
        if (m > 1) {
            int nr = 1, nc = 0, nd = 1;
            while (nr < m && grid[nr][nc] == 1) {
                nc++;
                nd = 1;
            }
            if (nr >= 0 && nr < m && nc >= 0 && nc < n) {
                if (nd == 0) {
                    dp_right[nr][nc] = (dp_right[nr][nc] + 1) % MOD;
                } else {
                    dp_down[nr][nc] = (dp_down[nr][nc] + 1) % MOD;
                }
            }
        }
        for (int r = 0; r < m; r++) {
            for (int c = 0; c < n; c++) {
                int ways_right = dp_right[r][c];
                int ways_down = dp_down[r][c];
                int total_ways = (ways_right + ways_down) % MOD;
                if (total_ways == 0) continue;
                if (r == m - 1 && c == n - 1) continue;
                if (c + 1 < n) {
                    int curr_r = r, curr_c = c + 1, curr_d = 0;
                    while (curr_r >= 0 && curr_r < m && curr_c >= 0 && curr_c < n && grid[curr_r][curr_c] == 1) {
                        if (curr_d == 0) {
                            curr_r++;
                            curr_d = 0;
                        } else {
                            curr_c++;
                            curr_d = 1;
                        }
                    }
                    if (curr_r >= 0 && curr_r < m && curr_c >= 0 && curr_c < n) {
                        if (curr_d == 0) {
                            dp_right[curr_r][curr_c] = (dp_right[curr_r][curr_c] + total_ways) % MOD;
                        } else {
                            dp_down[curr_r][curr_c] = (dp_down[curr_r][curr_c] + total_ways) % MOD;
                        }
                    }
                }
                if (r + 1 < m) {
                    int curr_r = r + 1, curr_c = c, curr_d = 1;
                    while (curr_r >= 0 && curr_r < m && curr_c >= 0 && curr_c < n && grid[curr_r][curr_c] == 1) {
                        if (curr_d == 0) {
                            curr_r++;
                            curr_d = 0;
                        } else {
                            curr_c++;
                            curr_d = 1;
                        }
                    }
                    if (curr_r >= 0 && curr_r < m && curr_c >= 0 && curr_c < n) {
                        if (curr_d == 0) {
                            dp_right[curr_r][curr_c] = (dp_right[curr_r][curr_c] + total_ways) % MOD;
                        } else {
                            dp_down[curr_r][curr_c] = (dp_down[curr_r][curr_c] + total_ways) % MOD;
                        }
                    }
                }
            }
        }
        return (dp_right[m - 1][n - 1] + dp_down[m - 1][n - 1]) % MOD;
    }
}