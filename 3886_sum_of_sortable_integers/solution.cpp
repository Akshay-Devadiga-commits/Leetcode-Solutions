#include <vector>
#include <numeric>
#include <algorithm>
#include <string>

class Solution {
public:
    int sortableIntegers(std::vector<int>& nums) {
        int n = nums.size();
        std::vector<int> target = nums;
        std::sort(target.begin(), target.end());

        auto is_cyclic_shift = [&](const std::vector<int>& sub, const std::vector<int>& orig) {
            if (sub.size() != orig.size()) {
                return false;
            }
            std::vector<int> double_orig = orig;
            double_orig.insert(double_orig.end(), orig.begin(), orig.end());
            
            std::string sub_str;
            for(int x : sub) sub_str += std::to_string(x) + ",";
            if (!sub_str.empty()) sub_str.pop_back();

            std::string double_orig_str;
            for(int x : double_orig) double_orig_str += std::to_string(x) + ",";
            if (!double_orig_str.empty()) double_orig_str.pop_back();

            return double_orig_str.find(sub_str) != std::string::npos;
        };

        int ans = 0;
        for (int k = 1; k <= n; ++k) {
            if (n % k == 0) {
                bool possible = true;
                for (int i = 0; i < n; i += k) {
                    std::vector<int> sub(nums.begin() + i, nums.begin() + i + k);
                    std::vector<int> orig(target.begin() + i, target.begin() + i + k);
                    if (!is_cyclic_shift(sub, orig)) {
                        possible = false;
                        break;
                    }
                }
                if (possible) {
                    ans += k;
                }
            }
        }
        return ans;
    }
};