from typing import List
from collections import defaultdict

class Solution:
    def maxSelectedElements(self, nums: List[int]) -> int:
        nums.sort()
        dp = defaultdict(int)
        max_length = 0
        
        for x in nums:
            dp[x + 1] = dp[x] + 1
            dp[x] = dp[x - 1] + 1
            max_length = max(max_length, dp[x + 1], dp[x])
            
        return max_length