import java.util.Queue;
import java.util.LinkedList;

class Solution {
    public int countIslands(int[][] grid, int k) {
        if (grid == null || grid.length == 0 || grid[0].length == 0) {
            return 0;
        }

        int m = grid.length;
        int n = grid[0].length;
        boolean[][] visited = new boolean[m][n];
        int count = 0;

        int[][] directions = {{-1, 0}, {1, 0}, {0, -1}, {0, 1}};

        for (int r = 0; r < m; r++) {
            for (int c = 0; c < n; c++) {
                if (grid[r][c] > 0 && !visited[r][c]) {
                    long islandSum = 0;
                    Queue<int[]> queue = new LinkedList<>();
                    queue.offer(new int[]{r, c});
                    visited[r][c] = true;

                    while (!queue.isEmpty()) {
                        int[] curr = queue.poll();
                        int currR = curr[0];
                        int currC = curr[1];
                        islandSum += grid[currR][currC];

                        for (int[] dir : directions) {
                            int nr = currR + dir[0];
                            int nc = currC + dir[1];
                            if (nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] && grid[nr][nc] > 0) {
                                visited[nr][nc] = true;
                                queue.offer(new int[]{nr, nc});
                            }
                        }
                    }

                    if (islandSum % k == 0) {
                        count++;
                    }
                }
            }
        }

        return count;
    }
}