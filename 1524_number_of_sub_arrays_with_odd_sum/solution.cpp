#include <vector>

class Solution {
public:
    int numOfSubarrays(std::vector<int>& arr) {
        const int MOD = 1e9 + 7; // 10^9 + 7
        int oddCount = 0;
        int evenCount = 1; // Represents the prefix sum of 0 before the array starts
        int currentSumParity = 0;
        int totalOddSubarrays = 0;

        for (int num : arr) {
            // Update current sum parity
            // If num is odd, parity flips. If num is even, parity stays the same.
            currentSumParity = (currentSumParity + num) % 2;

            if (currentSumParity == 1) { // Current prefix sum is odd
                // We need previous prefix sums that were even to make the subarray sum odd
                totalOddSubarrays = (totalOddSubarrays + evenCount) % MOD;
                oddCount++;
            } else { // Current prefix sum is even
                // We need previous prefix sums that were odd to make the subarray sum odd
                totalOddSubarrays = (totalOddSubarrays + oddCount) % MOD;
                evenCount++;
            }
        }

        return totalOddSubarrays;
    }
};