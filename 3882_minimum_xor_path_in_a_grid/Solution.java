import java.util.HashSet;
import java.util.Set;
import java.util.Collections;
import java.util.List;
import java.util.ArrayList;

class Solution {
    public int minCost(int[][] grid) {
        int m = grid.length;
        int n = grid[0].length;
        List<Set<Integer>> dp = new ArrayList<>(n);
        for (int i = 0; i < n; i++) {
            dp.add(new HashSet<>());
        }
        
        for (int r = 0; r < m; r++) {
            List<Set<Integer>> newDp = new ArrayList<>(n);
            for (int i = 0; i < n; i++) {
                newDp.add(new HashSet<>());
            }
            for (int c = 0; c < n; c++) {
                if (r == 0 && c == 0) {
                    newDp.get(c).add(grid[0][0]);
                } else {
                    if (r > 0) {
                        for (int val : dp.get(c)) {
                            newDp.get(c).add(val ^ grid[r][c]);
                        }
                    }
                    if (c > 0) {
                        for (int val : newDp.get(c - 1)) {
                            newDp.get(c).add(val ^ grid[r][c]);
                        }
                    }
                }
            }
            dp = newDp;
        }
        
        return Collections.min(dp.get(n - 1));
    }
}