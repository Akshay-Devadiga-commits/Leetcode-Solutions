#include <vector>
#include <unordered_map>
using namespace std;

class Solution {
public:
    int smallestUniqueSubarray(vector<int>& nums) {
        int n = nums.size();
        
        for (int length = 1; length <= n; length++) {
            unordered_map<vector<int>, int, VectorHash> subarrayCounts;
            
            for (int i = 0; i <= n - length; i++) {
                vector<int> subarray(nums.begin() + i, nums.begin() + i + length);
                subarrayCounts[subarray]++;
            }
            
            for (const auto& pair : subarrayCounts) {
                if (pair.second == 1) {
                    return length;
                }
            }
        }
        
        return n;
    }
};

struct VectorHash {
    size_t operator()(const vector<int>& v) const {
        size_t seed = v.size();
        for (int x : v) {
            seed ^= hash<int>()(x) + 0x9e3779b9 + (seed << 6) + (seed >> 2);
        }
        return seed;
    }
};