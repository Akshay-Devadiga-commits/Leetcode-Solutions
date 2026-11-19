#include <vector>
#include <algorithm>
#include <limits>

class Solution {
public:
    int minimizeSum(std::vector<int>& nums) {
        std::sort(nums.begin(), nums.end());
        // Can always change the number to any other number in `nums`, so `low` becomes 0.
        // Thus, rephrase the problem as finding the minimum `high`.
        int highOfChangingTwoMins = nums[nums.size() - 1] - nums[2];
        int highOfChangingTwoMaxs = nums[nums.size() - 3] - nums[0];
        int highOfChangingMinAndMax = nums[nums.size() - 2] - nums[1];
        return std::min({highOfChangingTwoMins, highOfChangingTwoMaxs, highOfChangingMinAndMax});
    }
};