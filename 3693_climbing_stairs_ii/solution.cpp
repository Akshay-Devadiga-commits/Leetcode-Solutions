#include <vector>
#include <algorithm>

class Solution {
public:
    int climbStairs(int n, std::vector<int>& costs) {
        std::vector<int> dp(n + 1, 0);
        for (int i = 1; i <= n; ++i) {
            int minPrevCost = dp[i - 1] + 1;
            if (i >= 2) {
                minPrevCost = std::min(minPrevCost, dp[i - 2] + 4);
            }
            if (i >= 3) {
                minPrevCost = std::min(minPrevCost, dp[i - 3] + 9);
            }
            dp[i] = costs[i - 1] + minPrevCost;
        }
        return dp[n];
    }
};