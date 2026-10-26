import "container/list"

func countIslands(grid [][]int, k int) int {
    if len(grid) == 0 || len(grid[0]) == 0 {
        return 0
    }

    m := len(grid)
    n := len(grid[0])
    visited := make([][]bool, m)
    for i := range visited {
        visited[i] = make([]bool, n)
    }

    count := 0
    directions := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

    for r := 0; r < m; r++ {
        for c := 0; c < n; c++ {
            if grid[r][c] > 0 && !visited[r][c] {
                islandSum := int64(0)
                queue := list.New()
                queue.PushBack([2]int{r, c})
                visited[r][c] = true

                for queue.Len() > 0 {
                    elem := queue.Front()
                    queue.Remove(elem)
                    curr := elem.Value.([2]int)
                    currR, currC := curr[0], curr[1]
                    islandSum += int64(grid[currR][currC])

                    for _, dir := range directions {
                        nr, nc := currR + dir[0], currC + dir[1]
                        if nr >= 0 && nr < m && nc >= 0 && nc < n && !visited[nr][nc] && grid[nr][nc] > 0 {
                            visited[nr][nc] = true
                            queue.PushBack([2]int{nr, nc})
                        }
                    }
                }

                if islandSum%int64(k) == 0 {
                    count++
                }
            }
        }
    }

    return count
}