#include <vector>
#include <queue>
#include <tuple>
#include <algorithm>

using namespace std;

class Solution {
public:
    vector<int> assignTasks(vector<int>& servers, vector<int>& tasks) {
        int n = servers.size();
        int m = tasks.size();
        vector<int> ans(m);

        // Min-priority queue for available servers: (weight, index)
        priority_queue<pair<int, int>, vector<pair<int, int>>, greater<pair<int, int>>> available;
        for (int i = 0; i < n; ++i) {
            available.push({servers[i], i});
        }

        // Min-priority queue for busy servers: (free_time, weight, index)
        priority_queue<tuple<long long, int, int>, vector<tuple<long long, int, int>>, greater<tuple<long long, int, int>>> busy;

        long long current_time = 0;
        int task_idx = 0;

        while (task_idx < m) {
            // 1. Free up servers
            while (!busy.empty() && get<0>(busy.top()) <= current_time) {
                auto [free_time, weight, index] = busy.top();
                busy.pop();
                available.push({weight, index});
            }

            // 2. Assign task
            if (!available.empty()) {
                auto [weight, index] = available.top();
                available.pop();
                ans[task_idx] = index;

                long long free_time = current_time + tasks[task_idx];
                busy.push({free_time, weight, index});
                task_idx++;
            } else {
                // 3. Jump to next available server's completion time
                current_time = get<0>(busy.top());
            }

            // 4. Advance time to next task arrival if needed
            if (task_idx < m) {
                current_time = max(current_time, (long long)task_idx);
            }
        }

        return ans;
    }
};