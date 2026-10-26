class Solution {
    /**
     * @param {number[][]} grid
     * @param {number} k
     * @return {number}
     */
    countIslands(grid, k) {
        if (!grid || grid.length === 0 || grid[0].length === 0) {
            return 0;
        }

        const m = grid.length;
        const n = grid[0].length;
        const visited = Array.from({ length: m }, () => Array(n).fill(false));
        let count = 0;

        const directions = [[-1, 0], [1, 0], [0, -1], [0, 1]];

        for (let r = 0; r < m; r++) {
            for (let c = 0; c < n; c++) {
                if (grid[r][c] > 0 && !visited[r][c]) {
                    let islandSum = 0;
                    const queue = [];
                    queue.push([r, c]);
                    visited[r][c] = true;
                    let head = 0;

                    while (head < queue.length) {
                        const [currR, currC] = queue[head++];
                        islandSum += grid[currR][currC];

                        for (const [dr, dc] of directions) {
                            const nr = currR + dr;
                            const nc = currC + dc;
                            if (nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] && grid[nr][nc] > 0) {
                                visited[nr][nc] = true;
                                queue.push([nr, nc]);
                            }
                        }
                    }

                    if (islandSum % k === 0) {
                        count++;
                    }
                }
            }
        }

        return count;
    }
}