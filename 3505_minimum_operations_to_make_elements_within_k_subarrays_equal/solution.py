import math
from typing import List

class Solution:
    def minOperations(self, nums: List[int], x: int, k: int) -> int:
        n = len(nums)
        
        # Precompute the minimum cost to make each subarray of size x have equal elements.
        # subarray_costs[i] will store the min cost for nums[i : i+x]
        subarray_costs = []
        for i in range(n - x + 1):
            sub_array = nums[i : i + x]
            sub_array.sort()
            median = sub_array[x // 2]
            cost = 0
            for val in nums[i : i + x]:
                cost += abs(val - median)
            subarray_costs.append(cost)
        
        # dp[i][j] = minimum operations to form exactly j non-overlapping subarrays of size x
        # using the first i elements of nums (i.e., nums[0...i-1]).
        # Dimensions: (n + 1) x (k + 1)
        dp = [[math.inf] * (k + 1) for _ in range(n + 1)]
        
        # Base case: 0 operations are needed to form 0 subarrays.
        for i in range(n + 1):
            dp[i][0] = 0
            
        # Fill the DP table
        for j in range(1, k + 1):
            for i in range(1, n + 1):
                # Option 1: Do not form the j-th subarray ending at index i-1.
                # In this case, the minimum cost is the same as forming j subarrays using the first i-1 elements.
                dp[i][j] = dp[i-1][j]
                
                # Option 2: Form the j-th subarray ending at index i-1.
                # This subarray spans indices [i-x, i-1]. This is only possible if i >= x.
                # The cost for this subarray is subarray_costs[i-x].
                # The previous j-1 subarrays must have been formed using the first i-x elements.
                if i >= x:
                    prev_cost = dp[i-x][j-1]
                    current_subarray_cost = subarray_costs[i-x]
                    
                    if prev_cost != math.inf:
                        dp[i][j] = min(dp[i][j], prev_cost + current_subarray_cost)
                        
        # The result is the minimum cost to form k subarrays using all n elements.
        # Since the problem asks for 'at least k', and costs are non-negative, the minimum for 'at least k'
        # will be achieved by exactly k subarrays. If a solution with k+m subarrays has cost C,
        # then a solution with k subarrays (by removing m subarrays) will have cost <= C.
        return dp[n][k]
