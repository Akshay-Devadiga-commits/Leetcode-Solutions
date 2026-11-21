package main

import "fmt"

func findSubtreeSizes(parent []int, s string) []int {
    n := len(parent)
    
    // 1. Build the initial adjacency list (children for each node)
    adj := make([][]int, n)
    for i := 0; i < n; i++ {
        adj[i] = []int{}
    }
    for i := 1; i < n; i++ {
        adj[parent[i]] = append(adj[parent[i]], i)
    }
    
    // 2. Determine the new parent for each node x (from 1 to n-1)
    // newParentMap[x] will store the new parent of x, or -1 if no change
    newParentMap := make(map[int]int)
    
    for x := 1; x < n; x++ {
        // Traverse upwards to find the closest ancestor y with s[x] == s[y]
        ancestor := parent[x]
        for ancestor != -1 {
            if s[x] == s[ancestor] {
                newParentMap[x] = ancestor
                break
            }
            ancestor = parent[ancestor]
        }
    }
    
    // 3. Build the final adjacency list based on new parent assignments
    finalAdj := make([][]int, n)
    for i := 0; i < n; i++ {
        finalAdj[i] = []int{}
    }
    for i := 0; i < n; i++ {
        // If node i's parent changed, add it to the new parent's children list
        if newParent, ok := newParentMap[i]; ok {
            finalAdj[newParent] = append(finalAdj[newParent], i)
        }
        // Otherwise, it remains a child of its original parent
        else if i != 0 { // Node 0 is the root, has no parent to add to
            finalAdj[parent[i]] = append(finalAdj[parent[i]], i)
        }
    }
            
    // 4. Calculate subtree sizes using DFS (post-order traversal)
    answer := make([]int, n)
    
    var dfsSize func(node int) int
    dfsSize = func(node int) int {
        size := 1 // Count the node itself
        for _, child := range finalAdj[node] {
            size += dfsSize(child)
        }
        answer[node] = size
        return size
    }
        
dfsSize(0) // Start DFS from the root
    
    return answer
}
