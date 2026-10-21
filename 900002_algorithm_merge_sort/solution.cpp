#include <vector>

class Solution {
public:
    std::vector<int> solve(std::vector<int>& data) {
        if (data.size() <= 1) {
            return data;
        }
        
        int mid = data.size() / 2;
        std::vector<int> left_sub(data.begin(), data.begin() + mid);
        std::vector<int> right_sub(data.begin() + mid, data.end());
        
        std::vector<int> left_half = solve(left_sub);
        std::vector<int> right_half = solve(right_sub);
        
        return merge(left_half, right_half);
    }

private:
    std::vector<int> merge(const std::vector<int>& left, const std::vector<int>& right) {
        std::vector<int> sorted_list;
        sorted_list.reserve(left.size() + right.size());
        int i = 0, j = 0;
        
        while (i < left.size() && j < right.size()) {
            if (left[i] < right[j]) {
                sorted_list.push_back(left[i]);
                i++;
            } else {
                sorted_list.push_back(right[j]);
                j++;
            }
        }
        
        while (i < left.size()) {
            sorted_list.push_back(left[i]);
            i++;
        }
        
        while (j < right.size()) {
            sorted_list.push_back(right[j]);
            j++;
        }
        
        return sorted_list;
    }
};