class Solution:
    def minCost(self, grid: list[list[int]]) -> int:
        m, n = len(grid), len(grid[0])
        dp = [set() for _ in range(n)]
        
        for r in range(m):
            new_dp = [set() for _ in range(n)]
            for c in range(n):
                if r == 0 and c == 0:
                    new_dp[c].add(grid[0][0])
                else:
                    if r > 0:
                        for val in dp[c]:
                            new_dp[c].add(val ^ grid[r][c])
                    if c > 0:
                        for val in new_dp[c - 1]:
                            new_dp[c].add(val ^ grid[r][c])
            dp = new_dp
            
        return min(dp[n - 1])