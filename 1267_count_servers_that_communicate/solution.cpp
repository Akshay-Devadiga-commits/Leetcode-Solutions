#include <vector>

class Solution {
public:
    int countServers(std::vector<std::vector<int>>& grid) {
        int m = grid.size();
        int n = grid[0].size();
        int ans = 0;
        std::vector<int> rows(m, 0);
        std::vector<int> cols(n, 0);

        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] == 1) {
                    rows[i]++;
                    cols[j]++;
                }
            }
        }

        for (int i = 0; i < m; ++i) {
            for (int j = 0; j < n; ++j) {
                if (grid[i][j] == 1 && (rows[i] > 1 || cols[j] > 1)) {
                    ans++;
                }
            }
        }

        return ans;
    }
};