class Solution:
    def uniquePaths(self, grid: List[List[int]]) -> int:
        MOD = 10**9 + 7
        m = len(grid)
        n = len(grid[0])
        dp = [[0] * n for _ in range(m)]
        dp[0][0] = 1
        
        for i in range(m):
            for j in range(n):
                if dp[i][j] == 0:
                    continue
                
                if j + 1 < n:
                    nr, nc = i, j + 1
                    d = 'R'
                    while 0 <= nr < m and 0 <= nc < n and grid[nr][nc] == 1:
                        if d == 'R':
                            nr += 1
                            d = 'D'
                        else:
                            nc += 1
                            d = 'R'
                    if 0 <= nr < m and 0 <= nc < n:
                        dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD
                        
                if i + 1 < m:
                    nr, nc = i + 1, j
                    d = 'D'
                    while 0 <= nr < m and 0 <= nc < n and grid[nr][nc] == 1:
                        if d == 'R':
                            nr += 1
                            d = 'D'
                        else:
                            nc += 1
                            d = 'R'
                    if 0 <= nr < m and 0 <= nc < n:
                        dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD
                        
        return dp[m - 1][n - 1]