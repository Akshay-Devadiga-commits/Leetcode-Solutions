class Solution {
    /**
     * @param {number} m
     * @param {number} n
     * @return {string[]}
     */
    createGrid(m, n) {
        const grid = Array.from({ length: m }, () => Array(n).fill('#'));
        let r = 0, c = 0;
        grid[r][c] = '.';
        while (r < m - 1 || c < n - 1) {
            if (r % 2 === 0) {
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
        return grid.map(row => row.join(''));
    }
}