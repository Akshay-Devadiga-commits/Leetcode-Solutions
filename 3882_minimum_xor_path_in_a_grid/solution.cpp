#include <vector>
#include <unordered_set>
#include <algorithm>

class Solution {
public:
    int minCost(std::vector<std::vector<int>>& grid) {
        int m = grid.size();
        int n = grid[0].size();
        std::vector<std::unordered_set<int>> dp(n);
        
        for (int r = 0; r < m; ++r) {
            std::vector<std::unordered_set<int>> new_dp(n);
            for (int c = 0; c < n; ++c) {
                if (r == 0 && c == 0) {
                    new_dp[c].insert(grid[0][0]);
                } else {
                    if (r > 0) {
                        for (int val : dp[c]) {
                            new_dp[c].insert(val ^ grid[r][c]);
                        }
                    }
                    if (c > 0) {
                        for (int val : new_dp[c - 1]) {
                            new_dp[c].insert(val ^ grid[r][c]);
                        }
                    }
                }
            }
            dp = std::move(new_dp);
        }
        
        int min_val = 2e9;
        for (int val : dp[n - 1]) {
            min_val = std::min(min_val, val);
        }
        return min_val;
    }
};