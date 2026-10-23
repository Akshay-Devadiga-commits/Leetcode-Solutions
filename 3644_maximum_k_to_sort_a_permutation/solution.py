from typing import List

class Solution:
    def sortPermutation(self, nums: List[int]) -> int:
        n = len(nums)
        if nums == sorted(nums):
            return 0
            
        max_val = max(nums)
        
        for k in range(max_val, -1, -1):
            parent = list(range(n))
            
            def find(i):
                if parent[i] == i:
                    return i
                parent[i] = find(parent[i])
                return parent[i]
                
            def union(i, j):
                root_i = find(i)
                root_j = find(j)
                if root_i != root_j:
                    parent[root_i] = root_j
            
            for i in range(n):
                for j in range(i + 1, n):
                    if (nums[i] & nums[j]) == k:
                        union(i, j)
            
            components = {}
            for i in range(n):
                root = find(i)
                if root not in components:
                    components[root] = []
                components[root].append(i)
            
            possible = True
            for root, indices in components.items():
                current_values = sorted([nums[i] for i in indices])
                target_values = sorted(indices)
                if current_values != target_values:
                    possible = False
                    break
            
            if possible:
                return k
                
        return 0