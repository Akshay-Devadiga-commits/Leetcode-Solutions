class Solution:
    def createGrid(self, m: int, n: int) -> list[str]:
        grid = [['#' for _ in range(n)] for _ in range(m)]
        r, c = 0, 0
        grid[r][c] = '.'
        while r < m - 1 or c < n - 1:
            if r % 2 == 0:
                if c < n - 1:
                    c += 1
                else:
                    r += 1
            else:
                if c > 0:
                    c -= 1
                else:
                    r += 1
            grid[r][c] = '.'
        return ["".join(row) for row in grid]