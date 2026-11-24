import math
from typing import List

class Solution:
    def maxLength(self, nums: List[int]) -> int:
        n = len(nums)
        max_len = 0
        for i in range(n):
            curr_prod = 1
            curr_gcd = 0
            curr_lcm = 1
            for j in range(i, n):
                val = nums[j]
                curr_prod *= val
                curr_gcd = math.gcd(curr_gcd, val)
                curr_lcm = (curr_lcm * val) // math.gcd(curr_lcm, val)
                if curr_prod == curr_lcm * curr_gcd:
                    max_len = max(max_len, j - i + 1)
        return max_len