#include <vector>

using namespace std;

class Solution {
public:
    int uniquePaths(vector<vector<int>>& grid) {
        long long MOD = 1e9 + 7;
        int m = grid.size();
        int n = grid[0].size();
        vector<vector<long long>> dp(m, vector<long long>(n, 0));
        dp[0][0] = 1;
        
        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (dp[i][j] == 0) continue;
                
                if (j + 1 < n) {
                    int nr = i, nc = j + 1;
                    char d = 'R';
                    while (nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1) {
                        if (d == 'R') {
                            nr++;
                            d = 'D';
                        } else {
                            nc++;
                            d = 'R';
                        }
                    }
                    if (nr >= 0 && nr < m && nc >= 0 && nc < n) {
                        dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
                    }
                }
                
                if (i + 1 < m) {
                    int nr = i + 1, nc = j;
                    char d = 'D';
                    while (nr >= 0 && nr < m && nc >= 0 && nc < n && grid[nr][nc] == 1) {
                        if (d == 'R') {
                            nr++;
                            d = 'D';
                        } else {
                            nc++;
                            d = 'R';
                        }
                    }
                    if (nr >= 0 && nr < m && nc >= 0 && nc < n) {
                        dp[nr][nc] = (dp[nr][nc] + dp[i][j]) % MOD;
                    }
                }
            }
        }
        return dp[m - 1][n - 1];
    }
};