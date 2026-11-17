#include <vector>
#include <string>
#include <algorithm>

class Solution {
public:
    int minimumChanges(std::string s, int k) {
        int n = s.length();

        // Precompute the cost to make a substring s[i:j+1] a semi-palindrome
        // cost[i][j] stores the minimum changes for s[i...j]
        std::vector<std::vector<int>> cost(n, std::vector<int>(n, 0));

        for (int i = 0; i < n; ++i) {
            for (int j = i + 1; j < n; ++j) {
                int subLen = j - i + 1;
                int minSubCost = 1e9; // Use a large number for infinity

                // Iterate through all possible divisors d (1 <= d < subLen)
                for (int d = 1; d < subLen; ++d) {
                    if (subLen % d == 0) {
                        int currentDCost = 0;
                        // Calculate cost for this divisor d
                        for (int groupIdx = 0; groupIdx < d; ++groupIdx) {
                            std::string groupChars;
                            // Collect characters for this group
                            for (int charIdx = groupIdx; charIdx < subLen; charIdx += d) {
                                groupChars += s[i + charIdx];
                            }
                            
                            // Calculate palindrome cost for the group
                            int groupLen = groupChars.length();
                            for (int l = 0; l < groupLen / 2; ++l) {
                                if (groupChars[l] != groupChars[groupLen - 1 - l]) {
                                    currentDCost++;
                                }
                            }
                        }
                        minSubCost = std::min(minSubCost, currentDCost);
                    }
                }
                cost[i][j] = (minSubCost == 1e9) ? 0 : minSubCost;
            }
        }

        // DP state: dp[i][p] = min changes to partition s[0...i-1] into p semi-palindromes
        // dp table size: (n+1) x (k+1)
        std::vector<std::vector<int>> dp(n + 1, std::vector<int>(k + 1, 1e9));

        // Base case: 0 changes for 0 length string and 0 partitions
        dp[0][0] = 0;

        // Fill the DP table
        for (int p = 1; p <= k; ++p) {
            for (int i = 1; i <= n; ++i) {
                // To partition s[0...i-1] into p parts, the last part is s[j...i-1]
                // where s[0...j-1] is partitioned into p-1 parts.
                for (int j = 0; j < i; ++j) {
                    // The last substring is s[j...i-1]
                    // Its cost is cost[j][i-1]
                    if (dp[j][p - 1] != 1e9) {
                        dp[i][p] = std::min(dp[i][p], dp[j][p - 1] + cost[j][i - 1]);
                    }
                }
            }
        }

        return dp[n][k];
    }
};