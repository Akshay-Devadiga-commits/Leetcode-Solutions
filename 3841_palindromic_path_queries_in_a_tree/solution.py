class SegmentTree:
    def __init__(self, n):
        self.n = n
        self.tree = [0] * (4 * n)
        self.lazy = [0] * (4 * n)

    def _apply(self, node, val):
        self.tree[node] ^= val
        self.lazy[node] ^= val

    def _push(self, node):
        if self.lazy[node]:
            self._apply(2 * node, self.lazy[node])
            self._apply(2 * node + 1, self.lazy[node])
            self.lazy[node] = 0

    def update(self, node, l, r, ql, qr, val):
        if ql <= l and r <= qr:
            self._apply(node, val)
            return
        self._push(node)
        mid = (l + r) // 2
        if ql <= mid:
            self.update(2 * node, l, mid, ql, qr, val)
        if qr > mid:
            self.update(2 * node + 1, mid + 1, r, ql, qr, val)
        self.tree[node] = self.tree[2 * node] ^ self.tree[2 * node + 1]

    def query(self, node, l, r, idx):
        if l == r:
            return self.tree[node]
        self._push(node)
        mid = (l + r) // 2
        if idx <= mid:
            return self.query(2 * node, l, mid, idx)
        else:
            return self.query(2 * node + 1, mid + 1, r, idx)

class Solution:
    def palindromePath(self, n: int, edges: list[list[int]], s: str, queries: list[str]) -> list[bool]:
        import sys
        sys.setrecursionlimit(200000)
        
        adj = [[] for _ in range(n)]
        for u, v in edges:
            adj[u].append(v)
            adj[v].append(u)
            
        parent = [0] * n
        depth = [0] * n
        heavy = [-1] * n
        head = [0] * n
        pos = [0] * n
        sz = [0] * n
        
        def dfs1(u, p, d):
            parent[u] = p
            depth[u] = d
            sz[u] = 1
            max_sub = 0
            h_child = -1
            for v in adj[u]:
                if v != p:
                    dfs1(v, u, d + 1)
                    sz[u] += sz[v]
                    if sz[v] > max_sub:
                        max_sub = sz[v]
                        h_child = v
            heavy[u] = h_child
            
        dfs1(0, -1, 0)
        
        cur_pos = 0
        def dfs2(u, h):
            nonlocal cur_pos
            head[u] = h
            pos[u] = cur_pos
            cur_pos += 1
            if heavy[u] != -1:
                dfs2(heavy[u], h)
            for v in adj[u]:
                if v != parent[u] and v != heavy[u]:
                    dfs2(v, v)
                    
        dfs2(0, 0)
        
        seg = SegmentTree(n)
        char_masks = [1 << (ord(c) - ord('a')) for c in s]
        
        for i in range(n):
            seg.update(1, 0, n - 1, pos[i], pos[i], char_masks[i])
            
        def update_subtree(u, val):
            seg.update(1, 0, n - 1, pos[u], pos[u] + sz[u] - 1, val)
            
        def get_path_xor(u):
            return seg.query(1, 0, n - 1, pos[u])
            
        def lca(u, v):
            while head[u] != head[v]:
                if depth[head[u]] > depth[head[v]]:
                    u = parent[head[u]]
                else:
                    v = parent[head[v]]
            return u if depth[u] < depth[v] else v
            
        ans = []
        for q in queries:
            parts = q.split()
            if parts[0] == 'update':
                u = int(parts[1])
                c = parts[2]
                new_mask = 1 << (ord(c) - ord('a'))
                old_mask = char_masks[u]
                if new_mask != old_mask:
                    diff = old_mask ^ new_mask
                    char_masks[u] = new_mask
                    update_subtree(u, diff)
            else:
                u = int(parts[1])
                v = int(parts[2])
                l = lca(u, v)
                
                res = get_path_xor(u) ^ get_path_xor(v) ^ char_masks[l]
                ans.append((res & (res - 1)) == 0)
                
        return ans