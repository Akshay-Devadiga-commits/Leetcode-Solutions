#include <vector>
#include <numeric>
#include <algorithm>

class Solution {
public:
    int maxLength(std::vector<int>& nums) {
        int n = nums.size();
        int maxLen = 0;
        for (int i = 0; i < n; ++i) {
            long long currProd = 1;
            long long currGcd = 0;
            long long currLcm = 1;
            for (int j = i; j < n; ++j) {
                long long val = nums[j];
                currProd *= val;
                currGcd = std::gcd(currGcd, val);
                currLcm = (currLcm / std::gcd(currLcm, val)) * val;
                if (currProd == currLcm * currGcd) {
                    maxLen = std::max(maxLen, j - i + 1);
                }
            }
        }
        return maxLen;
    }
};