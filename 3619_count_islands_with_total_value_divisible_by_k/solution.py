from typing import List
from collections import deque

class Solution:
    def countIslands(self, grid: List[List[int]], k: int) -> int:
        if not grid or not grid[0]:
            return 0
            
        m, n = len(grid), len(grid[0])
        visited = [[False] * n for _ in range(m)]
        count = 0
        
        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]
        
        for r in range(m):
            for c in range(n):
                if grid[r][c] > 0 and not visited[r][c]:
                    island_sum = 0
                    queue = deque([(r, c)])
                    visited[r][c] = True
                    
                    while queue:
                        curr_r, curr_c = queue.popleft()
                        island_sum += grid[curr_r][curr_c]
                        
                        for dr, dc in directions:
                            nr, nc = curr_r + dr, curr_c + dc
                            if 0 <= nr < m and 0 <= nc < n and not visited[nr][nc] and grid[nr][nc] > 0:
                                visited[nr][nc] = True
                                queue.append((nr, nc))
                                
                    if island_sum % k == 0:
                        count += 1
                        
        return count