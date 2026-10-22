#include <vector>
#include <queue>
#include <string>
#include <numeric>
#include <algorithm>
#include <unordered_set>

using namespace std;

class Solution {
public:
    int minOperations(vector<int>& nums) {
        int n = nums.size();
        vector<int> target(n);
        iota(target.begin(), target.end(), 0);
        
        if (nums == target) return 0;
        
        queue<pair<vector<int>, int>> q;
        set<vector<int>> visited;
        
        q.push({nums, 0});
        visited.insert(nums);
        
        while (!q.empty()) {
            auto [curr, dist] = q.front();
            q.pop();
            
            vector<int> next_rev = curr;
            reverse(next_rev.begin(), next_rev.end());
            if (next_rev == target) return dist + 1;
            if (visited.find(next_rev) == visited.end()) {
                visited.insert(next_rev);
                q.push({next_rev, dist + 1});
            }
            
            vector<int> next_rot(n);
            for (int i = 0; i < n - 1; ++i) {
                next_rot[i] = curr[i + 1];
            }
            next_rot[n - 1] = curr[0];
            if (next_rot == target) return dist + 1;
            if (visited.find(next_rot) == visited.end()) {
                visited.insert(next_rot);
                q.push({next_rot, dist + 1});
            }
        }
        return -1;
    }
};