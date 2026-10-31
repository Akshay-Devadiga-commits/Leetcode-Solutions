class Solution:
    def uniquePaths(self, grid: List[List[int]]) -> int:
        MOD = 10**9 + 7
        m = len(grid)
        n = len(grid[0])
        dp_right = [[0] * n for _ in range(m)]
        dp_down = [[0] * n for _ in range(m)]
        if n > 1:
            nr, nc, nd = 0, 1, 0
            while nc < n and grid[nr][nc] == 1:
                nr += 1
                nd = 0
            if 0 <= nr < m and 0 <= nc < n:
                if nd == 0:
                    dp_right[nr][nc] = (dp_right[nr][nc] + 1) % MOD
                else:
                    dp_down[nr][nc] = (dp_down[nr][nc] + 1) % MOD
        if m > 1:
            nr, nc, nd = 1, 0, 1
            while nr < m and grid[nr][nc] == 1:
                nc += 1
                nd = 1
            if 0 <= nr < m and 0 <= nc < n:
                if nd == 0:
                    dp_right[nr][nc] = (dp_right[nr][nc] + 1) % MOD
                else:
                    dp_down[nr][nc] = (dp_down[nr][nc] + 1) % MOD
        for r in range(m):
            for c in range(n):
                ways_right = dp_right[r][c]
                ways_down = dp_down[r][c]
                total_ways = (ways_right + ways_down) % MOD
                if total_ways == 0:
                    continue
                if r == m - 1 and c == n - 1:
                    continue
                if c + 1 < n:
                    curr_r, curr_c, curr_d = r, c + 1, 0
                    while 0 <= curr_r < m and 0 <= curr_c < n and grid[curr_r][curr_c] == 1:
                        if curr_d == 0:
                            curr_r += 1
                            curr_d = 0
                        else:
                            curr_c += 1
                            curr_d = 1
                    if 0 <= curr_r < m and 0 <= curr_c < n:
                        if curr_d == 0:
                            dp_right[curr_r][curr_c] = (dp_right[curr_r][curr_c] + total_ways) % MOD
                        else:
                            dp_down[curr_r][curr_c] = (dp_down[curr_r][curr_c] + total_ways) % MOD
                if r + 1 < m:
                    curr_r, curr_c, curr_d = r + 1, c, 1
                    while 0 <= curr_r < m and 0 <= curr_c < n and grid[curr_r][curr_c] == 1:
                        if curr_d == 0:
                            curr_r += 1
                            curr_d = 0
                        else:
                            curr_c += 1
                            curr_d = 1
                    if 0 <= curr_r < m and 0 <= curr_c < n:
                        if curr_d == 0:
                            dp_right[curr_r][curr_c] = (dp_right[curr_r][curr_c] + total_ways) % MOD
                        else:
                            dp_down[curr_r][curr_c] = (dp_down[curr_r][curr_c] + total_ways) % MOD
        return (dp_right[m - 1][n - 1] + dp_down[m - 1][n - 1]) % MOD