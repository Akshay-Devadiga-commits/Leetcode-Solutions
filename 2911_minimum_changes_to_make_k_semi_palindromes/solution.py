import math

class Solution:
    def minimumChanges(self, s: str, k: int) -> int:
        n = len(s)

        # Precompute the cost to make a substring s[i:j+1] a semi-palindrome
        # cost[i][j] stores the minimum changes for s[i...j]
        cost = [[0] * n for _ in range(n)]

        for i in range(n):
            for j in range(i + 1, n):
                sub_len = j - i + 1
                min_sub_cost = float('inf')

                # Iterate through all possible divisors d (1 <= d < sub_len)
                for d in range(1, sub_len):
                    if sub_len % d == 0:
                        current_d_cost = 0
                        # Calculate cost for this divisor d
                        for group_idx in range(d):
                            group_chars = []
                            # Collect characters for this group
                            for char_idx in range(group_idx, sub_len, d):
                                group_chars.append(s[i + char_idx])
                            
                            # Calculate palindrome cost for the group
                            group_len = len(group_chars)
                            for l in range(group_len // 2):
                                if group_chars[l] != group_chars[group_len - 1 - l]:
                                    current_d_cost += 1
                        min_sub_cost = min(min_sub_cost, current_d_cost)
                
                # If no valid divisor (sub_len=1), cost is 0. But problem constraints ensure sub_len >= 2 for valid divisors.
                # If sub_len is prime, only d=1 is possible. If sub_len=1, no valid divisor.
                # For sub_len=1, the loop for d won't run. We need to handle this implicitly or explicitly.
                # The problem states d must be a *positive* divisor and d < string's length. For length 1, no such d exists.
                # For length 2, only d=1 is valid. For length 3, only d=1 is valid.
                # If min_sub_cost remains inf, it means no valid divisor was found (e.g., length 1), or all divisors lead to inf cost (not possible here).
                # For length 1, cost is 0. For length > 1, there's always d=1.
                if sub_len > 1 and min_sub_cost == float('inf'): # This case should ideally not happen for sub_len > 1
                    # If sub_len is prime, only d=1 is checked. If sub_len is composite, multiple d are checked.
                    # The only case min_sub_cost remains inf is if sub_len=1, which is handled by j > i.
                    # Let's ensure d=1 is always considered if sub_len > 1.
                    # The loop `for d in range(1, sub_len)` correctly includes d=1.
                    # If sub_len is prime, only d=1 is checked.
                    # If sub_len is 1, the loop `for j in range(i + 1, n)` ensures sub_len >= 2.
                    pass # Should not reach here if sub_len >= 2
                
                cost[i][j] = min_sub_cost if min_sub_cost != float('inf') else 0 # Default to 0 if no valid divisor (e.g., length 1, though not possible here)

        # DP state: dp[i][p] = min changes to partition s[0...i-1] into p semi-palindromes
        # dp table size: (n+1) x (k+1)
        dp = [[float('inf')] * (k + 1) for _ in range(n + 1)]

        # Base case: 0 changes for 0 length string and 0 partitions
        dp[0][0] = 0

        # Fill the DP table
        for p in range(1, k + 1):
            for i in range(1, n + 1):
                # To partition s[0...i-1] into p parts, the last part is s[j...i-1]
                # where s[0...j-1] is partitioned into p-1 parts.
                for j in range(i):
                    # The last substring is s[j...i-1]
                    # Its cost is cost[j][i-1]
                    if dp[j][p - 1] != float('inf'):
                        dp[i][p] = min(dp[i][p], dp[j][p - 1] + cost[j][i - 1])

        return dp[n][k]
