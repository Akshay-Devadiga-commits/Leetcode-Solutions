#include <vector>

class Solution {
public:
    int sortableIntegers(std::vector<int>& nums) {
        int n = nums.size();
        int total_sum = 0;
        std::vector<int> divisors;
        for (int i = 1; i <= n; ++i) {
            if (n % i == 0) divisors.push_back(i);
        }

        for (int k : divisors) {
            bool possible = true;
            std::vector<std::vector<int>> subarrays;
            for (int i = 0; i < n; i += k) {
                std::vector<int> sub(nums.begin() + i, nums.begin() + i + k);
                int drops = 0;
                for (int j = 0; j < k; ++j) {
                    if (sub[j] > sub[(j + 1) % k]) {
                        drops++;
                    }
                }
                if (drops > 1) {
                    possible = false;
                    break;
                }
                std::vector<int> sorted_sub;
                if (drops == 1) {
                    int start_idx = 0;
                    for (int j = 0; j < k; ++j) {
                        if (sub[j] > sub[(j + 1) % k]) {
                            start_idx = (j + 1) % k;
                            break;
                        }
                    }
                    sorted_sub.reserve(k);
                    for (int j = 0; j < k; ++j) {
                        sorted_sub.push_back(sub[(start_idx + j) % k]);
                    }
                } else {
                    sorted_sub = sub;
                }
                subarrays.push_back(sorted_sub);
            }
            if (!possible) continue;
            std::vector<int> flat_array;
            flat_array.reserve(n);
            for (const auto& sub : subarrays) {
                flat_array.insert(flat_array.end(), sub.begin(), sub.end());
            }
            for (int i = 0; i < n - 1; ++i) {
                if (flat_array[i] > flat_array[i + 1]) {
                    possible = false;
                    break;
                }
            }
            if (possible) {
                total_sum += k;
            }
        }
        return total_sum;
    }
};