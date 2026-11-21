class Solution {
  /**
   * @param {number[]} parent
   * @param {string} s
   * @return {number[]}
   */
  findSubtreeSizes(parent, s) {
    const n = parent.length;
    const ans = new Array(n).fill(0);
    const newParent = [...parent]; // Create a shallow copy of parent array
    const tree = Array.from({ length: n }, () => []); // Adjacency list

    for (let i = 1; i < n; ++i) {
      const closest = this._findClosestAncestor(i, parent, s);
      if (closest !== -1) {
        newParent[i] = closest;
      }
    }

    for (let i = 1; i < n; ++i) {
      tree[newParent[i]].push(i);
    }

    this._dfs(tree, 0, ans);
    return ans;
  }

  /**
   * Returns the closest ancestor of node `u` that has the same value as `u`.
   * @param {number} u
   * @param {number[]} parent
   * @param {string} s
   * @return {number}
   */
  _findClosestAncestor(u, parent, s) {
    let curr = parent[u];
    while (curr !== -1) {
      if (s[curr] === s[u]) {
        return curr;
      }
      curr = parent[curr];
    }
    return -1;
  }

  /**
   * @param {number[][]} tree
   * @param {number} u
   * @param {number[]} ans
   * @return {number}
   */
  _dfs(tree, u, ans) {
    let sz = 1;
    for (const v of tree[u]) {
      sz += this._dfs(tree, v, ans);
    }
    ans[u] = sz;
    return sz;
  }
}