from collections import deque

class Solution:
    def minOperations(self, s: str) -> int:
        if s == "".join(sorted(s)):
            return 0
            
        n = len(s)
        visited = {s}
        queue = deque([(s, 0)])
        
        while queue:
            curr, ops = queue.popleft()
            
            for i in range(n):
                for j in range(i + 1, n + 1):
                    if i == 0 and j == n:
                        continue
                    
                    sub = curr[i:j]
                    sorted_sub = "".join(sorted(sub))
                    
                    next_str = curr[:i] + sorted_sub + curr[j:]
                    
                    if next_str == "".join(sorted(next_str)):
                        return ops + 1
                        
                    if next_str not in visited:
                        visited.add(next_str)
                        queue.append((next_str, ops + 1))
                        
        return -1