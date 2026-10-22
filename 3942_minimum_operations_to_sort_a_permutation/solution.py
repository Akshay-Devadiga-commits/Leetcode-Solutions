from collections import deque
from typing import List

class Solution:
    def minOperations(self, nums: List[int]) -> int:
        target = tuple(range(len(nums)))
        start = tuple(nums)
        
        if start == target:
            return 0
            
        queue = deque([(start, 0)])
        visited = {start}
        
        while queue:
            curr, dist = queue.popleft()
            
            next_rev = tuple(reversed(curr))
            if next_rev == target:
                return dist + 1
            if next_rev not in visited:
                visited.add(next_rev)
                queue.append((next_rev, dist + 1))
                
            next_rot = curr[1:] + (curr[0],)
            if next_rot == target:
                return dist + 1
            if next_rot not in visited:
                visited.add(next_rot)
                queue.append((next_rot, dist + 1))
                
        return -1