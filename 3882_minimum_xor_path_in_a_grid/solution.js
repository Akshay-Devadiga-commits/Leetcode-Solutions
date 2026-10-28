/**
 * @param {number[][]} grid
 * @return {number}
 */
var minCost = function(grid) {
    const m = grid.length;
    const n = grid[0].length;
    let dp = Array.from({ length: n }, () => new Set());
    
    for (let r = 0; r < m; r++) {
        const newDp = Array.from({ length: n }, () => new Set());
        for (let c = 0; c < n; c++) {
            if (r === 0 && c === 0) {
                newDp[c].add(grid[0][0]);
            } else {
                if (r > 0) {
                    for (const val of dp[c]) {
                        newDp[c].add(val ^ grid[r][c]);
                    }
                }
                if (c > 0) {
                    for (const val of newDp[c - 1]) {
                        newDp[c].add(val ^ grid[r][c]);
                    }
                }
            }
        }
        dp = newDp;
    }
    
    let minVal = Infinity;
    for (const val of dp[n - 1]) {
        if (val < minVal) minVal = val;
    }
    return minVal;
};