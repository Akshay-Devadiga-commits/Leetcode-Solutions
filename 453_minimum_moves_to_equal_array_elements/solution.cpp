#include <vector>
#include <numeric>
#include <algorithm>

class Solution {
public:
    int minMoves(std::vector<int>& nums) {
        int mn = nums[0];
        for (int num : nums) {
            mn = std::min(mn, num);
        }
        int sum = 0;
        for (int num : nums) {
            sum += num - mn;
        }
        return sum;
    }
};