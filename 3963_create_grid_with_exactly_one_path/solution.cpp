#include <vector>
#include <string>

class Solution {
public:
    std::vector<std::string> createGrid(int m, int n) {
        std::vector<std::string> grid(m, std::string(n, '#'));
        int r = 0, c = 0;
        grid[r][c] = '.';
        while (r < m - 1 || c < n - 1) {
            if (r % 2 == 0) {
                if (c < n - 1) {
                    c++;
                } else {
                    r++;
                }
            } else {
                if (c > 0) {
                    c--;
                } else {
                    r++;
                }
            }
            grid[r][c] = '.';
        }
        return grid;
    }
};