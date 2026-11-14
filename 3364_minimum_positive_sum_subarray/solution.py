from typing import List

class Solution:
    def minimumSumSubarray(self, nums: List[int], l: int, r: int) -> int:
        n = len(nums)
        min_positive_sum = float('inf')
        found = False

        # Iterate through all possible start indices of the subarray
        for i in range(n):
            current_sum = 0
            # Iterate through all possible end indices of the subarray starting from i
            for j in range(i, n):
                current_sum += nums[j]
                subarray_length = j - i + 1

                # Check if the subarray length is within the given range [l, r]
                if l <= subarray_length <= r:
                    # Check if the sum is positive
                    if current_sum > 0:
                        min_positive_sum = min(min_positive_sum, current_sum)
                        found = True

        # If no such subarray was found, return -1, otherwise return the minimum positive sum
        return min_positive_sum if found else -1