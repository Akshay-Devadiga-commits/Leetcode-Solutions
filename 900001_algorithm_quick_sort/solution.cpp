#include <vector>

class Solution {
public:
    std::vector<int> solve(std::vector<int>& data) {
        if (data.empty() || data.size() <= 1) {
            return data;
        }
        int pivot = data[data.size() / 2];
        std::vector<int> left, middle, right;
        for (int x : data) {
            if (x < pivot) left.push_back(x);
            else if (x == pivot) middle.push_back(x);
            else right.push_back(x);
        }
        std::vector<int> sorted_left = solve(left);
        std::vector<int> sorted_right = solve(right);
        
        sorted_left.insert(sorted_left.end(), middle.begin(), middle.end());
        sorted_left.insert(sorted_left.end(), sorted_right.begin(), sorted_right.end());
        return sorted_left;
    }
};