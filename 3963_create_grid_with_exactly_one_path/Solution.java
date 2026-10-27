import java.util.*;

class Solution {
    public List<String> createGrid(int m, int n) {
        char[][] grid = new char[m][n];
        for (int i = 0; i < m; i++) {
            Arrays.fill(grid[i], '#');
        }
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
        List<String> result = new ArrayList<>();
        for (int i = 0; i < m; i++) {
            result.add(new String(grid[i]));
        }
        return result;
    }
}