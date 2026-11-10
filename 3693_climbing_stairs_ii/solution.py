class Solution:
    def climbStairs(self, n: int, costs: List[int]) -> int:
        dp = [0] * (n + 1)
        for i in range(1, n + 1):
            min_prev_cost = dp[i - 1] + 1
            if i >= 2:
                min_prev_cost = min(min_prev_cost, dp[i - 2] + 4)
            if i >= 3:
                min_prev_cost = min(min_prev_cost, dp[i - 3] + 9)
            dp[i] = costs[i - 1] + min_prev_cost
        return dp[n]