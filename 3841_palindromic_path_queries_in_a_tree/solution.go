package main

import (
	"strconv"
	"strings"
)

type SegmentTree struct {
    n     int
    tree  []int
    lazy  []int
}

func NewSegmentTree(n int) *SegmentTree {
    return &SegmentTree{
        n:    n,
        tree: make([]int, 4*n),
        lazy: make([]int, 4*n),
    }
}

func (st *SegmentTree) apply(node, val int) {
    st.tree[node] ^= val
    st.lazy[node] ^= val
}

func (st *SegmentTree) push(node int) {
    if st.lazy[node] != 0 {
        st.apply(2*node, st.lazy[node])
        st.apply(2*node+1, st.lazy[node])
        st.lazy[node] = 0
    }
}

func (st *SegmentTree) Update(node, l, r, ql, qr, val int) {
    if ql <= l && r <= qr {
        st.apply(node, val)
        return
    }
    st.push(node)
    mid := (l + r) / 2
    if ql <= mid {
        st.Update(2*node, l, mid, ql, qr, val)
    }
    if qr > mid {
        st.Update(2*node+1, mid+1, r, ql, qr, val)
    }
    st.tree[node] = st.tree[2*node] ^ st.tree[2*node+1]
}

func (st *SegmentTree) Query(node, l, r, idx int) int {
    if l == r {
        return st.tree[node]
    }
    st.push(node)
    mid := (l + r) / 2
    if idx <= mid {
        return st.Query(2*node, l, mid, idx)
    }
    return st.Query(2*node+1, mid+1, r, idx)
}

func palindromePath(n int, edges [][]int, s string, queries []string) []bool {
    adj := make([][]int, n)
    for _, edge := range edges {
        u, v := edge[0], edge[1]
        adj[u] = append(adj[u], v)
        adj[v] = append(adj[v], u)
    }

    parent := make([]int, n)
    depth := make([]int, n)
    heavy := make([]int, n)
    for i := range heavy {
        heavy[i] = -1
    }
    head := make([]int, n)
    pos := make([]int, n)
    sz := make([]int, n)

    var dfs1 func(u, p, d int)
    dfs1 = func(u, p, d int) {
        parent[u] = p
        depth[u] = d
        sz[u] = 1
        maxSub := 0
        hChild := -1
        for _, v := range adj[u] {
            if v != p {
                dfs1(v, u, d+1)
                sz[u] += sz[v]
                if sz[v] > maxSub {
                    maxSub = sz[v]
                    hChild = v
                }
            }
        }
        heavy[u] = hChild
    }
    dfs1(0, -1, 0)

    curPos := 0
    var dfs2 func(u, h int)
    dfs2 = func(u, h int) {
        head[u] = h
        pos[u] = curPos
        curPos++
        if heavy[u] != -1 {
            dfs2(heavy[u], h)
        }
        for _, v := range adj[u] {
            if v != parent[u] && v != heavy[u] {
                dfs2(v, v)
            }
        }
    }
    dfs2(0, 0)

    seg := NewSegmentTree(n)
    charMasks := make([]int, n)
    for i := 0; i < n; i++ {
        charMasks[i] = 1 << (s[i] - 'a')
        seg.Update(1, 0, n-1, pos[i], pos[i], charMasks[i])
    }

    lca := func(u, v int) int {
        for head[u] != head[v] {
            if depth[head[u]] > depth[head[v]] {
                u = parent[head[u]]
            } else {
                v = parent[head[v]]
            }
        }
        if depth[u] < depth[v] {
            return u
        }
        return v
    }

    var ans []bool
    for _, q := range queries {
        parts := strings.Split(q, " ")
        if parts[0] == "update" {
            u, _ := strconv.Atoi(parts[1])
            c := parts[2][0]
            newMask := 1 << (c - 'a')
            oldMask := charMasks[u]
            if newMask != oldMask {
                diff := oldMask ^ newMask
                charMasks[u] = newMask
                seg.Update(1, 0, n-1, pos[u], pos[u]+sz[u]-1, diff)
            }
        } else {
            u, _ := strconv.Atoi(parts[1])
            v, _ := strconv.Atoi(parts[2])
            l := lca(u, v)
            res := seg.Query(1, 0, n-1, pos[u]) ^ seg.Query(1, 0, n-1, pos[v]) ^ charMasks[l]
            ans = append(ans, (res&(res-1)) == 0)
        }
    }
    return ans
}