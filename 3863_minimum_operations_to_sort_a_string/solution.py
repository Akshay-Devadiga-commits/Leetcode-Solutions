from collections import deque

class Solution:
    def minOperations(self, s: str) -> int:
        if s == "".join(sorted(s)):
            return 0
            
        n = len(s)
        queue = deque([(s, 0)])
        visited = {s}
        
        while queue:
            curr, ops = queue.popleft()
            
            for i in range(n):
                for j in range(i + 2, n + 1):
                    if i == 0 and j == n:
                        continue
                        
                    sub = curr[i:j]
                    sorted_sub = "".join(sorted(sub))
                    next_s = curr[:i] + sorted_sub + curr[j:]
                    
                    if next_s == "".join(sorted(next_s)):
                        return ops + 1
                        
                    if next_s not in visited:
                        visited.add(next_s)
                        queue.append((next_s, ops + 1))
                        
        return -1