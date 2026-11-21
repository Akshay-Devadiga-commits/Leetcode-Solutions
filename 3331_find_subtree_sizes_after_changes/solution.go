func findSubtreeSizes(parent []int, s string) []int {
	n := len(parent)
	ans := make([]int, n)
	newParent := make([]int, n)
	copy(newParent, parent)
	tree := make([][]int, n)
	for i := range tree {
		tree[i] = []int{}
	}

	for i := 1; i < n; i++ {
		closest := findClosestAncestor(i, parent, s)
		if closest != -1 {
			newParent[i] = closest
		}
	}

	for i := 1; i < n; i++ {
		tree[newParent[i]] = append(tree[newParent[i]], i)
	}

	dfs(tree, 0, ans)
	return ans
}

func findClosestAncestor(u int, parent []int, s string) int {
	curr := parent[u]
	for curr != -1 {
		if s[curr] == s[u] {
			return curr
		}
		curr = parent[curr]
	}
	return -1
}

func dfs(tree [][]int, u int, ans []int) int {
	sz := 1
	for _, v := range tree[u] {
		sz += dfs(tree, v, ans)
	}
	ans[u] = sz
	return sz
}