class Solution {
    public int uniquePaths(int[][] grid) {
        int MOD = 1_000_000_007;
        int m = grid.length;
        int n = grid[0].length;
        long[][] dp = new long[m][n];
        dp[0][0] = 1;
        
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (dp[i][j] == 0) continue;
                
                if (j + 1 < n) {
                    int nr = i, nc = j + 1;
                    char d = 'R';
                    while (nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1) {
                        if (d == 'R') {
                            nr++;
                            d = 'D';
                        } else {
                            nc++;
                            d = 'R';
                        }
                    }
                    if (nr >= 0 && nr < m && nc >= 0 && nc < n) {
                        dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
                    }
                }
                
                if (i + 1 < m) {
                    int nr = i + 1, nc = j;
                    char d = 'D';
                    while (nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1) {
                        if (d == 'R') {
                            nr++;
                            d = 'D';
                        } else {
                            nc++;
                            d = 'R';
                        }
                    }
                    if (nr >= 0 && nr < m && nc >= 0 && nc < n) {
                        dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
                    }
                }
            }
        }
        return (int) dp[m - 1][n - 1];
    }
}