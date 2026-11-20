#include <vector>
#include <algorithm>
#include <unordered_map>

using namespace std;

class Solution {
public:
    int maxSelectedElements(vector<int>& nums) {
        sort(nums.begin(), nums.end());
        unordered_map<int, int> dp;
        int max_length = 0;
        
        for (int x : nums) {
            int dp_x = dp[x];
            int dp_x_minus_1 = dp[x - 1];
            
            dp[x + 1] = dp_x + 1;
            dp[x] = dp_x_minus_1 + 1;
            
            max_length = max({max_length, dp[x + 1], dp[x]});
        }
        
        return max_length;
    }
};