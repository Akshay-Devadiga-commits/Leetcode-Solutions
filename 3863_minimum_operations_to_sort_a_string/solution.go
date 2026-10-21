import (
    "sort"
)

type Node struct {
    str string
    ops int
}

func minOperations(s string) int {
    isSorted := func(str string) bool {
        for i := 1; i < len(str); i++ {
            if str[i] < str[i-1] {
                return false
            }
        }
        return true
    }

    if isSorted(s) {
        return 0
    }

    n := len(s)
    visited := make(map[string]bool)
    visited[s] = true
    
    queue := []Node{{s, 0}}
    
    for len(queue) > 0 {
        curr := queue[0]
        queue = queue[1:]
        
        for i := 0; i < n; i++ {
            for j := i + 1; j <= n; j++ {
                if i == 0 && j == n {
                    continue
                }
                
                subBytes := []byte(curr.str[i:j])
                sort.Slice(subBytes, func(a, b int) bool {
                    return subBytes[a] < subBytes[b]
                })
                
                nextStr := curr.str[:i] + string(subBytes) + curr.str[j:]
                
                if isSorted(nextStr) {
                    return curr.ops + 1
                }
                
                if !visited[nextStr] {
                    visited[nextStr] = true
                    queue = append(queue, Node{nextStr, curr.ops + 1})
                }
            }
        }
    }
    
    return -1
}