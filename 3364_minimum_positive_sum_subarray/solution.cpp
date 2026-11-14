#include <vector>
#include <algorithm> // For std::min
#include <limits>    // For std::numeric_limits<int>::max()

class Solution {
public:
    int minimumSumSubarray(std::vector<int>& nums, int l, int r) {
        int n = nums.size();

        // Calculate prefix sums
        // prefixSum[k] stores the sum of nums[0...k-1]
        // prefixSum[0] = 0
        // Use long long for prefix sums to prevent potential overflow
        std::vector<long long> prefixSum(n + 1, 0);
        for (int i = 0; i < n; ++i) {
            prefixSum[i + 1] = prefixSum[i] + nums[i];
        }

        int minPositiveSum = std::numeric_limits<int>::max();

        // Iterate over all possible start indices i
        for (int i = 0; i < n; ++i) {
            // Iterate over all possible end indices j
            // j must be at least i to form a non-empty subarray
            for (int j = i; j < n; ++j) {
                // Calculate the length of the current subarray nums[i...j]
                int currentLength = j - i + 1;

                // Check if the length constraint is met
                if (currentLength >= l && currentLength <= r) {
                    // Calculate the sum of the subarray nums[i...j] using prefix sums
                    long long currentSumLong = prefixSum[j + 1] - prefixSum[i];
                    
                    // Check if the sum is positive and fits into an int
                    if (currentSumLong > 0 && currentSumLong <= std::numeric_limits<int>::max()) {
                        int currentSum = static_cast<int>(currentSumLong);
                        // Update the minimum positive sum found so far
                        minPositiveSum = std::min(minPositiveSum, currentSum);
                    }
                }
            }
        }

        // If minPositiveSum is still INT_MAX, it means no such subarray was found
        if (minPositiveSum == std::numeric_limits<int>::max()) {
            return -1;
        } else {
            return minPositiveSum;
        }
    }
};