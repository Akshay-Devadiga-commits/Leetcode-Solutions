#include <string>
#include <queue>
#include <unordered_set>
#include <algorithm>

class Solution {
public:
    int minOperations(std::string s) {
        std::string sorted_s = s;
        std::sort(sorted_s.begin(), sorted_s.end());
        if (s == sorted_s) return 0;

        int n = s.length();
        std::queue<std::pair<std::string, int>> q;
        std::unordered_set<std::string> visited;

        q.push({s, 0});
        visited.insert(s);

        while (!q.empty()) {
            auto [curr, ops] = q.front();
            q.pop();

            for (int i = 0; i < n; ++i) {
                for (int j = i + 2; j <= n; ++j) {
                    if (i == 0 && j == n) continue;

                    std::string sub = curr.substr(i, j - i);
                    std::sort(sub.begin(), sub.end());
                    std::string next_s = curr.substr(0, i) + sub + curr.substr(j);

                    std::string next_sorted = next_s;
                    std::sort(next_sorted.begin(), next_sorted.end());
                    if (next_s == next_sorted) {
                        return ops + 1;
                    }

                    if (visited.find(next_s) == visited.end()) {
                        visited.insert(next_s);
                        q.push({next_s, ops + 1});
                    }
                }
            }
        }

        return -1;
    }
};