import (
    "math"
)

func minimumChanges(s string, k int) int {
    n := len(s)

    // Precompute the cost to make a substring s[i:j+1] a semi-palindrome
    // cost[i][j] stores the minimum changes for s[i...j]
    cost := make([][]int, n)
    for i := range cost {
        cost[i] = make([]int, n)
    }

    for i := 0; i < n; i++ {
        for j := i + 1; j < n; j++ {
            subLen := j - i + 1
            minSubCost := math.MaxInt32

            // Iterate through all possible divisors d (1 <= d < subLen)
            for d := 1; d < subLen; d++ {
                if subLen%d == 0 {
                    currentDCost := 0
                    // Calculate cost for this divisor d
                    for groupIdx := 0; groupIdx < d; groupIdx++ {
                        groupChars := []byte{}
                        // Collect characters for this group
                        for charIdx := groupIdx; charIdx < subLen; charIdx += d {
                            groupChars = append(groupChars, s[i+charIdx])
                        }
                        
                        // Calculate palindrome cost for the group
                        groupLen := len(groupChars)
                        for l := 0; l < groupLen/2; l++ {
                            if groupChars[l] != groupChars[groupLen-1-l] {
                                currentDCost++
                            }
                        }
                    }
                    if currentDCost < minSubCost {
                        minSubCost = currentDCost
                    }
                }
            }
            if minSubCost == math.MaxInt32 {
                cost[i][j] = 0
            } else {
                cost[i][j] = minSubCost
            }
        }
    }

    // DP state: dp[i][p] = min changes to partition s[0...i-1] into p semi-palindromes
    // dp table size: (n+1) x (k+1)
    dp := make([][]int, n+1)
    for i := range dp {
        dp[i] = make([]int, k+1)
        for j := range dp[i] {
            dp[i][j] = math.MaxInt32
        }
    }

    // Base case: 0 changes for 0 length string and 0 partitions
    dp[0][0] = 0

    // Fill the DP table
    for p := 1; p <= k; p++ {
        for i := 1; i <= n; i++ {
            // To partition s[0...i-1] into p parts, the last part is s[j...i-1]
            // where s[0...j-1] is partitioned into p-1 parts.
            for j := 0; j < i; j++ {
                // The last substring is s[j...i-1]
                // Its cost is cost[j][i-1]
                if dp[j][p-1] != math.MaxInt32 {
                    if dp[j][p-1] + cost[j][i-1] < dp[i][p] {
                        dp[i][p] = dp[j][p-1] + cost[j][i-1]
                    }
                }
            }
        }
    }

    return dp[n][k]
}
