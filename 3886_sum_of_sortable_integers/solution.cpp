#include <vector>

class Solution {
public:
    int sortableIntegers(std::vector<int>& nums) {
        int n = nums.size();
        int total_sum = 0;

        auto can_sort_with_k = [&](int k) {
            int prev_max = -1;
            for (int i = 0; i < n; i += k) {
                int drops = 0;
                int split = 0;
                for (int j = 0; j < k; ++j) {
                    if (nums[i + j] > nums[i + (j + 1) % k]) {
                        drops++;
                        split = j + 1;
                    }
                }
                if (drops > 1) return false;

                int min_val = nums[i + split];
                int max_val = nums[i + (split - 1 + k) % k];

                if (prev_max != -1 && prev_max > min_val) {
                    return false;
                }
                prev_max = max_val;
            }
            return true;
        };

        for (int k = 1; k <= n; ++k) {
            if (n % k == 0) {
                if (can_sort_with_k(k)) {
                    total_sum += k;
                }
            }
        }

        return total_sum;
    }
};