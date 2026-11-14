import math
from typing import List

class Solution:
    def minimumSumSubarray(self, nums: List[int], l: int, r: int) -> int:
        n = len(nums)
        
        # Calculate prefix sums
        # prefix_sum[k] stores the sum of nums[0...k-1]
        # prefix_sum[0] = 0
        prefix_sum = [0] * (n + 1)
        for i in range(n):
            prefix_sum[i+1] = prefix_sum[i] + nums[i]
        
        min_positive_sum = math.inf
        
        # Iterate over all possible start indices i
        for i in range(n):
            # Iterate over all possible end indices j
            # j must be at least i to form a non-empty subarray
            for j in range(i, n):
                # Calculate the length of the current subarray nums[i...j]
                current_length = j - i + 1
                
                # Check if the length constraint is met
                if l <= current_length <= r:
                    # Calculate the sum of the subarray nums[i...j] using prefix sums
                    current_sum = prefix_sum[j+1] - prefix_sum[i]
                    
                    # Check if the sum is positive
                    if current_sum > 0:
                        # Update the minimum positive sum found so far
                        min_positive_sum = min(min_positive_sum, current_sum)
        
        # If min_positive_sum is still infinity, it means no such subarray was found
        if min_positive_sum == math.inf:
            return -1
        else:
            return int(min_positive_sum)