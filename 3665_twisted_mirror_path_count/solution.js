/**
 * @param {number[][]} grid
 * @return {number}
 */
var uniquePaths = function(grid) {
    const MOD = 1000000007;
    const m = grid.length;
    const n = grid[0].length;
    const dp = Array.from({ length: m }, () => Array(n).fill(0));
    dp[0][0] = 1;
    
    for (let i = 0; i < m; i++) {
        for (let j = 0; j < n; j++) {
            if (dp[i][j] === 0) continue;
            
            if (j + 1 < n) {
                let nr = i, nc = j + 1;
                let d = 'R';
                while (nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] === 1) {
                    if (d === 'R') {
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
                let nr = i + 1, nc = j;
                let d = 'D';
                while (nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] === 1) {
                    if (d === 'R') {
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
    return dp[m - 1][n - 1];
};