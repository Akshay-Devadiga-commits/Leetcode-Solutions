/**
 * @param {number[][]} grid
 * @return {number}
 */
var closedIsland = function(grid) {
    if (!grid || grid.length === 0 || grid[0].length === 0) return 0;
    
    const rows = grid.length;
    const cols = grid[0].length;
    
    function dfs(r, c) {
        if (r < 0 || r >= rows || c < 0 || c >= cols || grid[r][c] === 1) return;
        grid[r][c] = 1;
        dfs(r + 1, c);
        dfs(r - 1, c);
        dfs(r, c + 1);
        dfs(r, c - 1);
    }
    
    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if ((r === 0 || r === rows - 1 || c === 0 || c === cols - 1) && grid[r][c] === 0) {
                dfs(r, c);
            }
        }
    }
    
    let closedIslands = 0;
    for (let r = 0; r < rows; r++) {
        for (let c = 0; c < cols; c++) {
            if (grid[r][c] === 0) {
                closedIslands++;
                dfs(r, c);
            }
        }
    }
    
    return closedIslands;
};