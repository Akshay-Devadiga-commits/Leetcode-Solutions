class Solution:
    def numOfSubarrays(self, arr: list[int]) -> int:
        MOD = 10**9 + 7
        odd_count = 0
        even_count = 1  # Represents the prefix sum of 0 before the array starts
        current_sum_parity = 0
        total_odd_subarrays = 0

        for num in arr:
            # Update current sum parity
            # If num is odd, parity flips. If num is even, parity stays the same.
            current_sum_parity = (current_sum_parity + num) % 2

            if current_sum_parity == 1:  # Current prefix sum is odd
                # We need previous prefix sums that were even to make the subarray sum odd
                total_odd_subarrays = (total_odd_subarrays + even_count) % MOD
                odd_count += 1
            else:  # Current prefix sum is even
                # We need previous prefix sums that were odd to make the subarray sum odd
                total_odd_subarrays = (total_odd_subarrays + odd_count) % MOD
                even_count += 1
        
        return total_odd_subarrays