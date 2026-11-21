class Solution {
  findSubtreeSizes(parent, s) {
    const n = parent.length;
    const ans = new Array(n).fill(0);
    const newParent = [...parent];
    const tree = Array.from({ length: n }, () => []);

    for (let i = 1; i < n; ++i) {
      const closest = this.findClosestAncestor(i, parent, s);
      if (closest !== -1) {
        newParent[i] = closest;
      }
    }

    for (let i = 1; i < n; ++i) {
      tree[newParent[i]].push(i);
    }

    this.dfs(tree, 0, ans);
    return ans;
  }

  findClosestAncestor(u, parent, s) {
    let curr = parent[u];
    while (curr !== -1) {
      if (s[curr] === s[u]) {
        return curr;
      }
      curr = parent[curr];
    }
    return -1;
  }

  dfs(tree, u, ans) {
    let sz = 1;
    for (const v of tree[u]) {
      sz += this.dfs(tree, v, ans);
    }
    ans[u] = sz;
    return sz;
  }
}