#include <vector>
#include <algorithm>

class Solution {
public:
    int climbStairs(int n, std::vector<int>& costs) {
        // step_costs[k] will be the cost of step k.
        // We need costs for steps 0 to n. The input costs array has length n,
        // where costs[i] is the cost of step i+1.
        // We assume the cost of step 0 is 0.
        std::vector<int> step_costs(n + 1, 0);
        for (int i = 0; i < n; ++i) {
            step_costs[i + 1] = costs[i];
        }

        // dp[i] will store the minimum cost to reach step i.
        std::vector<int> dp(n + 1, INT_MAX);

        // Base case: starting at step 0 costs 0.
        dp[0] = 0;

        // Iterate through each step from 1 to n.
        for (int i = 1; i <= n; ++i) {
            // From step i-1 to i (jump of 1)
            if (i - 1 >= 0) {
                int jump_cost = 1 * 1; // (i - (i-1))^2 = 1^2
                if (dp[i - 1] != INT_MAX) {
                    dp[i] = std::min(dp[i], dp[i - 1] + step_costs[i] + jump_cost);
                }
            }

            // From step i-2 to i (jump of 2)
            if (i - 2 >= 0) {
                int jump_cost = 2 * 2; // (i - (i-2))^2 = 2^2
                if (dp[i - 2] != INT_MAX) {
                    dp[i] = std::min(dp[i], dp[i - 2] + step_costs[i] + jump_cost);
                }
            }

            // From step i-3 to i (jump of 3)
            if (i - 3 >= 0) {
                int jump_cost = 3 * 3; // (i - (i-3))^2 = 3^2
                if (dp[i - 3] != INT_MAX) {
                    dp[i] = std::min(dp[i], dp[i - 3] + step_costs[i] + jump_cost);
                }
            }
        }

        return dp[n];
    }
};
