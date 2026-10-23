#include <vector>
#include <algorithm>
#include <numeric>
#include <map>

class Solution {
public:
    int sortPermutation(std::vector<int>& nums) {
        int n = nums.size();
        if (std::is_sorted(nums.begin(), nums.end())) {
            return 0;
        }

        int max_val = *std::max_element(nums.begin(), nums.end());

        for (int k = max_val; k >= 0; --k) {
            std::vector<int> parent(n);
            std::iota(parent.begin(), parent.end(), 0);

            auto find = [&](auto& self, int i) -> int {
                if (parent[i] == i) return i;
                return parent[i] = self(self, parent[i]);
            };

            auto union_sets = [&](int i, int j) {
                int root_i = find(find, i);
                int root_j = find(find, j);
                if (root_i != root_j) {
                    parent[root_i] = root_j;
                }
            };

            for (int i = 0; i < n; ++i) {
                for (int j = i + 1; j < n; ++j) {
                    if ((nums[i] & nums[j]) == k) {
                        union_sets(i, j);
                    }
                }
            }

            std::map<int, std::vector<int>> components;
            for (int i = 0; i < n; ++i) {
                int root = find(find, i);
                components[root].push_back(i);
            }

            bool possible = true;
            for (auto& [root, indices] : components) {
                std::vector<int> current_values;
                for (int idx : indices) {
                    current_values.push_back(nums[idx]);
                }
                std::sort(current_values.begin(), current_values.end());
                
                std::vector<int> target_values = indices;
                std::sort(target_values.begin(), target_values.end());

                if (current_values != target_values) {
                    possible = false;
                    break;
                }
            }

            if (possible) return k;
        }

        return 0;
    }
};