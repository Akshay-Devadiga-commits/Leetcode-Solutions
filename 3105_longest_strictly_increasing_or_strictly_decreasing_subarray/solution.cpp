#include <vector>
#include <algorithm>

class Solution {
public:
    int longestMonotonicSubarray(std::vector<int>& nums) {
        int ans = 1;
        int increasing = 1;
        int decreasing = 1;

        for (int i = 1; i < nums.size(); ++i) {
            if (nums[i] > nums[i - 1]) {
                increasing++;
                decreasing = 1;
            } else if (nums[i] < nums[i - 1]) {
                decreasing++;
                increasing = 1;
            } else {
                increasing = 1;
                decreasing = 1;
            }
            ans = std::max({ans, increasing, decreasing});
        }

        return ans;
    }
};