#include <vector>

class Solution {
public:
    int sortableIntegers(std::vector<int>& nums) {
        int n = nums.size();
        int total_sum = 0;

        auto is_sorted_rotation = [](const std::vector<int>& sub) {
            int m = sub.size();
            if (m <= 1) return true;
            int drops = 0;
            for (int i = 0; i < m; ++i) {
                if (sub[i] > sub[(i + 1) % m]) {
                    drops++;
                }
            }
            return drops <= 1;
        };

        for (int k = 1; k <= n; ++k) {
            if (n % k != 0) continue;

            bool valid = true;
            std::vector<std::vector<int>> sorted_subarrays;

            for (int i = 0; i < n; i += k) {
                std::vector<int> sub(nums.begin() + i, nums.begin() + i + k);
                if (!is_sorted_rotation(sub)) {
                    valid = false;
                    break;
                }

                int m = sub.size();
                int start_idx = 0;
                for (int j = 0; j < m; ++j) {
                    if (sub[j] > sub[(j + 1) % m]) {
                        start_idx = (j + 1) % m;
                        break;
                    }
                }

                std::vector<int> rotated;
                rotated.reserve(m);
                for (int j = 0; j < m; ++j) {
                    rotated.push_back(sub[(start_idx + j) % m]);
                }
                sorted_subarrays.push_back(rotated);
            }

            if (!valid) continue;

            std::vector<int> flat;
            flat.reserve(n);
            for (const auto& sub : sorted_subarrays) {
                flat.insert(flat.end(), sub.begin(), sub.end());
            }

            bool is_non_decreasing = true;
            for (int i = 0; i < n - 1; ++i) {
                if (flat[i] > flat[i + 1]) {
                    is_non_decreasing = false;
                    break;
                }
            }

            if (is_non_decreasing) {
                total_sum += k;
            }
        }

        return total_sum;
    }
};