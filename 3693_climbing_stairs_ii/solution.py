from typing import List

class Solution:
    def climbStairs(self, n: int, costs: List[int]) -> int:
        # dp[i] will store the minimum cost to reach step i.
        # We need n+1 steps, indexed from 0 to n.
        # The costs array is 1-indexed for steps 1 to n, but we can adjust.
        # Let's consider costs[i] as the cost associated with step i.
        # The problem states costs is of length n, and costs[i] is the cost of step i.
        # This implies costs[0] is for step 1, costs[1] for step 2, ..., costs[n-1] for step n.
        # However, the examples show costs[i] is for step i+1. Let's re-read.
        # "costs[i] is the cost of step i." This is ambiguous with 1-indexed array.
        # "costs of length n, where costs[i] is the cost of step i."
        # Example 1: n=4, costs=[1,2,3,4]. costs[0]=1 (step 0?), costs[1]=2 (step 1?), costs[2]=3 (step 2?), costs[3]=4 (step 3?).
        # The problem states "n + 1 steps, numbered from 0 to n".
        # "costs of length n, where costs[i] is the cost of step i."
        # This implies costs[0] is for step 0, costs[1] for step 1, ..., costs[n-1] for step n-1.
        # But the jump cost is costs[j] + (j-i)^2. If j is the destination step, then costs[j] is the cost of step j.
        # The examples clarify: costs is 0-indexed for steps 1 to n. So costs[0] is for step 1, costs[1] for step 2, ..., costs[n-1] for step n.
        # Let's adjust the costs array to be 0-indexed for steps 0 to n.
        # We can prepend a dummy cost for step 0, or handle it in logic.
        # The problem states "You start from step 0 with cost = 0."
        # The cost of jumping FROM step i TO step j is costs[j] + (j - i)^2.
        # This means costs[j] is the cost associated with the DESTINATION step j.
        # The input `costs` is of length `n`. `costs[i]` is the cost of step `i`.
        # This implies `costs[0]` is for step 0, `costs[1]` for step 1, ..., `costs[n-1]` for step n-1.
        # But the problem says "n + 1 steps, numbered from 0 to n".
        # And `costs` has length `n`. This means `costs` does NOT contain cost for step `n`.
        # The examples are key: n=4, costs=[1,2,3,4]. This means costs for steps 1, 2, 3, 4 are 1, 2, 3, 4 respectively.
        # So, `costs[i]` in the input corresponds to the cost of step `i+1`.
        # Let's create a new costs array that includes step 0 and is 0-indexed for steps 0 to n.
        # We can assume cost of step 0 is 0, as we start there with cost 0.
        
        # Let's re-interpret based on examples:
        # n = 4, costs = [1,2,3,4]
        # This means cost of step 1 is 1, step 2 is 2, step 3 is 3, step 4 is 4.
        # The input `costs` has length `n`. `costs[i]` is the cost of step `i+1`.
        # So, we need costs for steps 0 to n. We can assume cost of step 0 is 0.
        
        # Let's create a `step_costs` array where `step_costs[k]` is the cost of step `k`.
        # `step_costs` will have size `n+1`.
        step_costs = [0] * (n + 1)
        for i in range(n):
            step_costs[i+1] = costs[i]
        
        # dp[i] = minimum cost to reach step i.
        dp = [float('inf')] * (n + 1)
        
        # Base case: starting at step 0 costs 0.
        dp[0] = 0
        
        # Iterate through each step from 1 to n.
        for i in range(1, n + 1):
            # From step i-1 to i
            if i - 1 >= 0:
                jump_cost = (1)**2 # (i - (i-1))^2 = 1^2
                dp[i] = min(dp[i], dp[i-1] + step_costs[i] + jump_cost)
            
            # From step i-2 to i
            if i - 2 >= 0:
                jump_cost = (2)**2 # (i - (i-2))^2 = 2^2
                dp[i] = min(dp[i], dp[i-2] + step_costs[i] + jump_cost)
            
            # From step i-3 to i
            if i - 3 >= 0:
                jump_cost = (3)**2 # (i - (i-3))^2 = 3^2
                dp[i] = min(dp[i], dp[i-3] + step_costs[i] + jump_cost)
        
        return dp[n]
