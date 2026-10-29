class SegmentTree {
    constructor(n) {
        this.n = n;
        this.tree = new Int32Array(4 * n);
        this.lazy = new Int32Array(4 * n);
    }
    _apply(node, val) {
        this.tree[node] ^= val;
        this.lazy[node] ^= val;
    }
    _push(node) {
        if (this.lazy[node]) {
            this._apply(2 * node, this.lazy[node]);
            this._apply(2 * node + 1, this.lazy[node]);
            this.lazy[node] = 0;
        }
    }
    update(node, l, r, ql, qr, val) {
        if (ql <= l && r <= qr) {
            this._apply(node, val);
            return;
        }
        this._push(node);
        const mid = (l + r) >> 1;
        if (ql <= mid) this.update(2 * node, l, mid, ql, qr, val);
        if (qr > mid) this.update(2 * node + 1, mid + 1, r, ql, qr, val);
        this.tree[node] = this.tree[2 * node] ^ this.tree[2 * node + 1];
    }
    query(node, l, r, idx) {
        if (l === r) return this.tree[node];
        this._push(node);
        const mid = (l + r) >> 1;
        if (idx <= mid) return this.query(2 * node, l, mid, idx);
        else return this.query(2 * node + 1, mid + 1, r, idx);
    }
}

class Solution {
    palindromePath(n, edges, s, queries) {
        const adj = Array.from({ length: n }, () => []);
        for (const [u, v] of edges) {
            adj[u].push(v);
            adj[v].push(u);
        }

        const parent = new Int32Array(n);
        const depth = new Int32Array(n);
        const heavy = new Int32Array(n).fill(-1);
        const head = new Int32Array(n);
        const pos = new Int32Array(n);
        const sz = new Int32Array(n);

        const dfs1 = (u, p, d) => {
            parent[u] = p;
            depth[u] = d;
            sz[u] = 1;
            let max_sub = 0;
            let h_child = -1;
            for (const v of adj[u]) {
                if (v !== p) {
                    dfs1(v, u, d + 1);
                    sz[u] += sz[v];
                    if (sz[v] > max_sub) {
                        max_sub = sz[v];
                        h_child = v;
                    }
                }
            }
            heavy[u] = h_child;
        };
        dfs1(0, -1, 0);

        let cur_pos = 0;
        const dfs2 = (u, h) => {
            head[u] = h;
            pos[u] = cur_pos++;
            if (heavy[u] !== -1) {
                dfs2(heavy[u], h);
            }
            for (const v of adj[u]) {
                if (v !== parent[u] && v !== heavy[u]) {
                    dfs2(v, v);
                }
            }
        };
        dfs2(0, 0);

        const seg = new SegmentTree(n);
        const char_masks = new Int32Array(n);
        for (let i = 0; i < n; i++) {
            char_masks[i] = 1 << (s.charCodeAt(i) - 97);
            seg.update(1, 0, n - 1, pos[i], pos[i], char_masks[i]);
        }

        const update_subtree = (u, val) => {
            seg.update(1, 0, n - 1, pos[u], pos[u] + sz[u] - 1, val);
        };

        const get_path_xor = (u) => {
            return seg.query(1, 0, n - 1, pos[u]);
        };

        const lca = (u, v) => {
            while (head[u] !== head[v]) {
                if (depth[head[u]] > depth[head[v]]) {
                    u = parent[head[u]];
                } else {
                    v = parent[head[v]];
                }
            }
            return depth[u] < depth[v] ? u : v;
        };

        const ans = [];
        for (const q of queries) {
            const parts = q.split(' ');
            if (parts[0] === 'update') {
                const u = parseInt(parts[1]);
                const c = parts[2];
                const new_mask = 1 << (c.charCodeAt(0) - 97);
                const old_mask = char_masks[u];
                if (new_mask !== old_mask) {
                    const diff = old_mask ^ new_mask;
                    char_masks[u] = new_mask;
                    update_subtree(u, diff);
                }
            } else {
                const u = parseInt(parts[1]);
                const v = parseInt(parts[2]);
                const l = lca(u, v);
                const res = get_path_xor(u) ^ get_path_xor(v) ^ char_masks[l];
                ans.push((res & (res - 1)) === 0);
            }
        }
        return ans;
    }
}