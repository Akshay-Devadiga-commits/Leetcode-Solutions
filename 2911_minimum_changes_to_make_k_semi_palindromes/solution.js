class Solution {
    minimumChanges(s, k) {
        const n = s.length;

        // Precompute the cost to make a substring s[i:j+1] a semi-palindrome
        // cost[i][j] stores the minimum changes for s[i...j]
        const cost = Array(n).fill(0).map(() => Array(n).fill(0));

        for (let i = 0; i < n; i++) {
            for (let j = i + 1; j < n; j++) {
                const subLen = j - i + 1;
                let minSubCost = Infinity;

                // Iterate through all possible divisors d (1 <= d < subLen)
                for (let d = 1; d < subLen; d++) {
                    if (subLen % d === 0) {
                        let currentDCost = 0;
                        // Calculate cost for this divisor d
                        for (let groupIdx = 0; groupIdx < d; groupIdx++) {
                            const groupChars = [];
                            // Collect characters for this group
                            for (let charIdx = groupIdx; charIdx < subLen; charIdx += d) {
                                groupChars.push(s[i + charIdx]);
                            }
                            
                            // Calculate palindrome cost for the group
                            const groupLen = groupChars.length;
                            for (let l = 0; l < Math.floor(groupLen / 2); l++) {
                                if (groupChars[l] !== groupChars[groupLen - 1 - l]) {
                                    currentDCost++;
                                }
                            }
                        }
                        minSubCost = Math.min(minSubCost, currentDCost);
                    }
                }
                cost[i][j] = minSubCost === Infinity ? 0 : minSubCost;
            }
        }

        // DP state: dp[i][p] = min changes to partition s[0...i-1] into p semi-palindromes
        // dp table size: (n+1) x (k+1)
        const dp = Array(n + 1).fill(0).map(() => Array(k + 1).fill(Infinity));

        // Base case: 0 changes for 0 length string and 0 partitions
        dp[0][0] = 0;

        // Fill the DP table
        for (let p = 1; p <= k; p++) {
            for (let i = 1; i <= n; i++) {
                // To partition s[0...i-1] into p parts, the last part is s[j...i-1]
                // where s[0...j-1] is partitioned into p-1 parts.
                for (let j = 0; j < i; j++) {
                    // The last substring is s[j...i-1]
                    // Its cost is cost[j][i-1]
                    if (dp[j][p - 1] !== Infinity) {
                        dp[i][p] = Math.min(dp[i][p], dp[j][p - 1] + cost[j][i - 1]);
                    }
                }
            }
        }

        return dp[n][k];
    }
}