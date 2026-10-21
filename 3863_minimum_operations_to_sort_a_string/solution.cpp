#include <string>
#include <vector>
#include <queue>
#include <unordered_set>
#include <algorithm>

using namespace std;

class Solution {
public:
    int minOperations(string s) {
        auto isSorted = [](const string& str) {
            for (size_t i = 1; i < str.length(); ++i) {
                if (str[i] < str[i - 1]) return false;
            }
            return true;
        };
        
        if (isSorted(s)) return 0;
        
        int n = s.length();
        unordered_set<string> visited;
        queue<pair<string, int>> q;
        
        visited.insert(s);
        q.push({s, 0});
        
        while (!q.empty()) {
            auto [curr, ops] = q.front();
            q.pop();
            
            for (int i = 0; i < n; ++i) {
                for (int j = i + 1; j <= n; ++j) {
                    if (i == 0 && j == n) continue;
                    
                    string nextStr = curr;
                    string sub = curr.substr(i, j - i);
                    sort(sub.begin(), sub.end());
                    nextStr.replace(i, j - i, sub);
                    
                    if (isSorted(nextStr)) return ops + 1;
                    
                    if (visited.find(nextStr) == visited.end()) {
                        visited.insert(nextStr);
                        q.push({nextStr, ops + 1});
                    }
                }
            }
        }
        
        return -1;
    }
};