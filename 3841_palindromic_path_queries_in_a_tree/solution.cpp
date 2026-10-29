#include <vector>
#include <string>
#include <sstream>
#include <algorithm>

using namespace std;

class SegmentTree {
    int n;
    vector<int> tree;
    vector<int> lazy;

    void apply(int node, int val) {
        tree[node] ^= val;
        lazy[node] ^= val;
    }

    void push(int node) {
        if (lazy[node]) {
            apply(2 * node, lazy[node]);
            apply(2 * node + 1, lazy[node]);
            lazy[node] = 0;
        }
    }

public:
    SegmentTree(int n) : n(n), tree(4 * n, 0), lazy(4 * n, 0) {}

    void update(int node, int l, int r, int ql, int qr, int val) {
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

    int query(int node, int l, int r, int idx) {
        if (l == r) return tree[node];
        push(node);
        int mid = (l + r) / 2;
        if (idx <= mid) return query(2 * node, l, mid, idx);
        else return query(2 * node + 1, mid + 1, r, idx);
    }
};

class Solution {
    vector<vector<int>> adj;
    vector<int> parent, depth, heavy, head, pos, sz, char_masks;
    int cur_pos = 0;

    void dfs1(int u, int p, int d) {
        parent[u] = p;
        depth[u] = d;
        sz[u] = 1;
        int max_sub = 0;
        int h_child = -1;
        for (int v : adj[u]) {
            if (v != p) {
                dfs1(v, u, d + 1);
                sz[u] += sz[v];
                if (sz[v] > max_sub) {
                    max_sub = sz[v];
                    h_child = v;
                }
            }
        }
        heavy[u] = h_child;
    }

    void dfs2(int u, int h) {
        head[u] = h;
        pos[u] = cur_pos++;
        if (heavy[u] != -1) {
            dfs2(heavy[u], h);
        }
        for (int v : adj[u]) {
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

public:
    vector<bool> palindromePath(int n, vector<vector<int>>& edges, string s, vector<string>& queries) {
        adj.assign(n, vector<int>());
        for (const auto& edge : edges) {
            adj[edge[0]].push_back(edge[1]);
            adj[edge[1]].push_back(edge[0]);
        }

        parent.assign(n, 0);
        depth.assign(n, 0);
        heavy.assign(n, -1);
        head.assign(n, 0);
        pos.assign(n, 0);
        sz.assign(n, 0);
        cur_pos = 0;

        dfs1(0, -1, 0);
        dfs2(0, 0);

        SegmentTree seg(n);
        char_masks.assign(n, 0);
        for (int i = 0; i < n; ++i) {
            char_masks[i] = 1 << (s[i] - 'a');
            seg.update(1, 0, n - 1, pos[i], pos[i], char_masks[i]);
        }

        vector<bool> ans;
        for (const string& q : queries) {
            stringstream ss(q);
            string type;
            ss >> type;
            if (type == "update") {
                int u;
                char c;
                ss >> u >> c;
                int new_mask = 1 << (c - 'a');
                int old_mask = char_masks[u];
                if (new_mask != old_mask) {
                    int diff = old_mask ^ new_mask;
                    char_masks[u] = new_mask;
                    seg.update(1, 0, n - 1, pos[u], pos[u] + sz[u] - 1, diff);
                }
            } else {
                int u, v;
                ss >> u >> v;
                int l = lca(u, v);
                int res = seg.query(1, 0, n - 1, pos[u]) ^ seg.query(1, 0, n - 1, pos[v]) ^ char_masks[l];
                ans.push_back((res & (res - 1)) == 0);
            }
        }
        return ans;
    }
};