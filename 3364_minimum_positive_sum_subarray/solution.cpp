#include <vector>
#include <limits>
#include <algorithm>

class Solution {
public:
    /**
     * @param nums
     * @param l
     * @param r
     * @return
     */
    int minimumSumSubarray(std::vector<int>& nums, int l, int r) {
        int n = nums.size();
        int minPositiveSum = std::numeric_limits<int>::max();
        bool found = false;

        // Iterate through all possible start indices of the subarray
        for (int i = 0; i < n; ++i) {
            int currentSum = 0;
            // Iterate through all possible end indices of the subarray starting from i
            for (int j = i; j < n; ++j) {
                currentSum += nums[j];
                int subarrayLength = j - i + 1;

                // Check if the subarray length is within the given range [l, r]
                if (l <= subarrayLength && subarrayLength <= r) {
                    // Check if the sum is positive
                    if (currentSum > 0) {
                        minPositiveSum = std::min(minPositiveSum, currentSum);
                        found = true;
                    }
                }
            }
        }

        // If no such subarray was found, return -1, otherwise return the minimum positive sum
        return found ? minPositiveSum : -1;
    }
};