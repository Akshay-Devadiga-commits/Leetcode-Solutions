class Solution {
    countServers(grid) {
        const m = grid.length;
        const n = grid[0].length;
        let ans = 0;
        const rows = new Array(m).fill(0);
        const cols = new Array(n).fill(0);

        for (let i = 0; i < m; i++) {
            for (let j = 0; j < n; j++) {
                if (grid[i][j] === 1) {
                    rows[i]++;
                    cols[j]++;
                }
            }
        }

        for (let i = 0; i < m; i++) {
            for (let j = 0; j < n; j++) {
                if (grid[i][j] === 1 && (rows[i] > 1 || cols[j] > 1)) {
                    ans++;
                }
            }
        }

        return ans;
    }
}