func countServers(grid [][]int) int {
    m := len(grid)
    n := len(grid[0])
    ans := 0
    rows := make([]int, m)
    cols := make([]int, n)

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 {
                rows[i]++
                cols[j]++
            }
        }
    }

    for i := 0; i < m; i++ {
        for j := 0; j < n; j++ {
            if grid[i][j] == 1 && (rows[i] > 1 || cols[j] > 1) {
                ans++
            }
        }
    }

    return ans
}