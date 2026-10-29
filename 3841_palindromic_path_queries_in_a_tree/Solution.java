import java.util.*;

class SegmentTree {
    int n;
    int[] tree;
    int[] lazy;

    public SegmentTree(int n) {
        this.n = n;
        this.tree = new int[4 * n];
        this.lazy = new int[4 * n];
    }

    private void apply(int node, int val) {
        tree[node] ^= val;
        lazy[node] ^= val;
    }

    private void push(int node) {
        if (lazy[node] != 0) {
            apply(2 * node, lazy[node]);
            apply(2 * node + 1, lazy[node]);
            lazy[node] = 0;
        }
    }

    public void update(int node, int l, int r, int ql, int qr, int val) {
        if (ql <= l && r <= qr) {
            apply(node, val);
            return;
        }
        push(node);
        int mid = (l + r) / 2;
        if (ql <= mid) update(2 * node, l, mid, ql, qr, val);
        if (qr > mid) update(2 * node + 1, mid + 1, r, ql, qr, val);
        tree[node] = tree[2 * node] ^ tree[2 * node + 1];
    }

    public int query(int node, int l, int r, int idx) {
        if (l == r) return tree[node];
        push(node);
        int mid = (l + r) / 2;
        if (idx <= mid) return query(2 * node, l, mid, idx);
        else return query(2 * node + 1, mid + 1, r, idx);
    }
}

class Solution {
    List<List<Integer>> adj;
    int[] parent, depth, heavy, head, pos, sz, charMasks;
    int curPos = 0;

    void dfs1(int u, int p, int d) {
        parent[u] = p;
        depth[u] = d;
        sz[u] = 1;
        int maxSub = 0;
        int hChild = -1;
        for (int v : adj.get(u)) {
            if (v != p) {
                dfs1(v, u, d + 1);
                sz[u] += sz[v];
                if (sz[v] > maxSub) {
                    maxSub = sz[v];
                    hChild = v;
                }
            }
        }
        heavy[u] = hChild;
    }

    void dfs2(int u, int h) {
        head[u] = h;
        pos[u] = curPos++;
        if (heavy[u] != -1) {
            dfs2(heavy[u], h);
        }
        for (int v : adj.get(u)) {
            if (v != parent[u] && v != heavy[u]) {
                dfs2(v, v);
            }
        }
    }

    int lca(int u, int v) {
        while (head[u] != head[v]) {
            if (depth[head[u]] > depth[head[v]]) {
                u = parent[head[u]];
            } else {
                v = parent[head[v]];
            }
        }
        return depth[u] < depth[v] ? u : v;
    }

    public List<Boolean> palindromePath(int n, int[][] edges, String s, List<String> queries) {
        adj = new ArrayList<>(n);
        for (int i = 0; i < n; i++) adj.add(new ArrayList<>());
        for (int[] edge : edges) {
            adj.get(edge[0]).add(edge[1]);
            adj.get(edge[1]).add(edge[0]);
        }

        parent = new int[n];
        depth = new int[n];
        heavy = new int[n];
        Arrays.fill(heavy, -1);
        head = new int[n];
        pos = new int[n];
        sz = new int[n];
        curPos = 0;

        dfs1(0, -1, 0);
        dfs2(0, 0);

        SegmentTree seg = new SegmentTree(n);
        charMasks = new int[n];
        for (int i = 0; i < n; i++) {
            charMasks[i] = 1 << (s.charAt(i) - 'a');
            seg.update(1, 0, n - 1, pos[i], pos[i], charMasks[i]);
        }

        List<Boolean> ans = new ArrayList<>();
        for (String q : queries) {
            String[] parts = q.split(" ");
            if (parts[0].equals("update")) {
                int u = Integer.parseInt(parts[1]);
                int newMask = 1 << (parts[2].charAt(0) - 'a');
                int oldMask = charMasks[u];
                if (newMask != oldMask) {
                    int diff = oldMask ^ newMask;
                    charMasks[u] = newMask;
                    seg.update(1, 0, n - 1, pos[u], pos[u] + sz[u] - 1, diff);
                }
            } else {
                int u = Integer.parseInt(parts[1]);
                int v = Integer.parseInt(parts[2]);
                int l = lca(u, v);
                int res = seg.query(1, 0, n - 1, pos[u]) ^ seg.query(1, 0, n - 1, pos[v]) ^ charMasks[l];
                ans.add((res & (res - 1)) == 0);
            }
        }
        return ans;
    }
}