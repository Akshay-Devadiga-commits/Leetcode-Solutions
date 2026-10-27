func createGrid(m int, n int) []string {
    grid := make([][]byte, m)
    for i := 0; i < m; i++ {
        grid[i] = make([]byte, n)
        for j := 0; j < n; j++ {
            grid[i][j] = '#'
        }
    }
    r, c := 0, 0
    grid[r][c] = '.'
    for r < m-1 || c < n-1 {
        if r%2 == 0 {
            if c < n-1 {
                c++
            } else {
                r++
            }
        } else {
            if c > 0 {
                c--
            } else {
                r++
            }
        }
        grid[r][c] = '.'
    }
    result := make([]string, m)
    for i := 0; i < m; i++ {
        result[i] = string(grid[i])
    }
    return result
}