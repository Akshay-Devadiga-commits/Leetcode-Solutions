type Solution struct {
}

func (this *Solution) FindSubtreeSizes(parent []int, s string) []int {
    n := len(parent)
    ans := make([]int, n)
    newParent := make([]int, n)
    copy(newParent, parent) // Copy parent slice
    tree := make([][]int, n) // Adjacency list

    for i := 0; i < n; i++ {
        tree[i] = make([]int, 0)
    }

    for i := 1; i < n; i++ {
        closest := this.findClosestAncestor(i, parent, s)
        if closest != -1 {
            newParent[i] = closest
        }
    }

    for i := 1; i < n; i++ {
        tree[newParent[i]] = append(tree[newParent[i]], i)
    }

    this.dfs(tree, 0, ans)
    return ans
}

// Returns the closest ancestor of node `u` that has the same value as `u`.
func (this *Solution) findClosestAncestor(u int, parent []int, s string) int {
    curr := parent[u]
    for curr != -1 {
        if s[curr] == s[u] {
            return curr
        }
        curr = parent[curr]
    }
    return -1
}

func (this *Solution) dfs(tree [][]int, u int, ans []int) int {
    sz := 1
    for _, v := range tree[u] {
        sz += this.dfs(tree, v, ans)
    }
    ans[u] = sz
    return sz
}