package main

type Solution struct {}

func (s *Solution) FindSubtreeSizes(parent []int, str string) []int {
	n := len(parent)
	ans := make([]int, n)
	newParent := make([]int, n)
	copy(newParent, parent)
	tree := make([][]int, n)
	for i := range tree {
		tree[i] = []int{}
	}

	for i := 1; i < n; i++ {
		closest := s.findClosestAncestor(i, parent, str)
		if closest != -1 {
			newParent[i] = closest
		}
	}

	for i := 1; i < n; i++ {
		tree[newParent[i]] = append(tree[newParent[i]], i)
	}

	s.dfs(tree, 0, ans)
	return ans
}

// Returns the closest ancestor of node `u` that has the same value as `u`.
func (s *Solution) findClosestAncestor(u int, parent []int, str string) int {
	curr := parent[u]
	for curr != -1 {
		if str[curr] == str[u] {
			return curr
		}
		curr = parent[curr]
	}
	return -1
}

func (s *Solution) dfs(tree [][]int, u int, ans []int) int {
	sz := 1
	for _, v := range tree[u] {
		sz += s.dfs(tree, v, ans)
	}
	ans[u] = sz
	return sz
}