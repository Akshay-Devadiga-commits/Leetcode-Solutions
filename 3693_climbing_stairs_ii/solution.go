import "math"

func climbStairs(n int, costs []int) int {
    // step_costs[k] will be the cost of step k.
    // We need costs for steps 0 to n. The input costs array has length n,
    // where costs[i] is the cost of step i+1.
    // We assume the cost of step 0 is 0.
    step_costs := make([]int, n+1)
    for i := 0; i < n; i++ {
        step_costs[i+1] = costs[i]
    }

    // dp[i] will store the minimum cost to reach step i.
    dp := make([]int, n+1)
    for i := range dp {
        dp[i] = math.MaxInt32 // Use a large integer for infinity
    }

    // Base case: starting at step 0 costs 0.
    dp[0] = 0

    // Iterate through each step from 1 to n.
    for i := 1; i <= n; i++ {
        // From step i-1 to i (jump of 1)
        if i-1 >= 0 {
            jump_cost := 1 * 1 // (i - (i-1))^2 = 1^2
            if dp[i-1] != math.MaxInt32 {
                dp[i] = min(dp[i], dp[i-1] + step_costs[i] + jump_cost)
            }
        }

        // From step i-2 to i (jump of 2)
        if i-2 >= 0 {
            jump_cost := 2 * 2 // (i - (i-2))^2 = 2^2
            if dp[i-2] != math.MaxInt32 {
                dp[i] = min(dp[i], dp[i-2] + step_costs[i] + jump_cost)
            }
        }

        // From step i-3 to i (jump of 3)
        if i-3 >= 0 {
            jump_cost := 3 * 3 // (i - (i-3))^2 = 3^2
            if dp[i-3] != math.MaxInt32 {
                dp[i] = min(dp[i], dp[i-3] + step_costs[i] + jump_cost)
            }
        }
    }

    return dp[n]
}

// Helper function for min, as Go doesn't have a built-in min for integers.
func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}
