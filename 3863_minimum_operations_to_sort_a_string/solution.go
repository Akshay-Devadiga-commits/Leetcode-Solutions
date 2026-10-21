import (
    "sort"
)

type Item struct {
    s   string
    ops int
}

func minOperations(s string) int {
    runes := []rune(s)
    sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
    target := string(runes)
    if s == target {
        return 0
    }

    n := len(s)
    queue := []Item{{s, 0}}
    visited := map[string]bool{s: true}

    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]

        currRunes := []rune(curr.s)
        for i := 0; i < n; i++ {
            for j := i + 2; j <= n; j++ {
                if i == 0 && j == n {
                    continue
                }

                sub := make([]rune, j-i)
                copy(sub, currRunes[i:j])
                sort.Slice(sub, func(x, y int) bool { return sub[x] < sub[y] })

                nextRunes := append([]rune{}, currRunes[:i]...)
                nextRunes = append(nextRunes, sub...)
                nextRunes = append(nextRunes, currRunes[j:]...)
                nextS := string(nextRunes)

                nextSortedRunes := append([]rune{}, nextRunes...)
                sort.Slice(nextSortedRunes, func(x, y int) bool { return nextSortedRunes[x] < nextSortedRunes[y] })
                if nextS == string(nextSortedRunes) {
                    return curr.ops + 1
                }

                if !visited[nextS] {
                    visited[nextS] = true
                    queue = append(queue, Item{nextS, curr.ops + 1})
                }
            }
        }
    }

    return -1
}