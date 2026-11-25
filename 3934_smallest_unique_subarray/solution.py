from typing import List

class Solution:
    def smallestUniqueSubarray(self, nums: List[int]) -> int:
        n = len(nums)
        
        # Iterate over all possible subarray lengths from 1 to n
        for length in range(1, n + 1):
            # Use a dictionary to count occurrences of each subarray of this length
            subarray_counts = {}
            
            # Generate all subarrays of the current length
            for i in range(n - length + 1):
                subarray = tuple(nums[i:i + length])
                subarray_counts[subarray] = subarray_counts.get(subarray, 0) + 1
            
            # Check if any subarray of this length appears exactly once
            for count in subarray_counts.values():
                if count == 1:
                    return length
        
        # This case is theoretically unreachable because the entire array is always unique
        return n