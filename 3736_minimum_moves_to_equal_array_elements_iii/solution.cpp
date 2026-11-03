#include <vector>
#include <numeric>
#include <algorithm>

class Solution {
public:
    /**
     * Calculates the minimum number of moves to make all elements in an array equal by only increasing elements.
     *
     * @param nums A vector of integers.
     * @return The minimum total number of moves required.
     */
    int minMoves(std::vector<int>& nums) {
        if (nums.empty()) {
            return 0;
        }
        
        int maxVal = nums[0];
        long long totalSum = 0; // Use long long to prevent potential overflow for sum
        
        for (int num : nums) {
            if (num > maxVal) {
                maxVal = num;
            }
            totalSum += num;
        }
        
        // The total number of moves is the sum of (max_val - num) for each num in nums.
        // This can be rewritten as (n * max_val) - sum(nums).
        // Cast nums.size() to long long before multiplication to prevent overflow if maxVal is large.
        return static_cast<int>((static_cast<long long>(nums.size()) * maxVal) - totalSum);
    }
};