/**
 * @param {number[]} parent
 * @param {string} s
 * @return {number[]}
 */
var findSubtreeSizes = function(parent, s) {
    const n = parent.length;
    
    // 1. Build the initial adjacency list (children for each node)
    const adj = Array.from({ length: n }, () => []);
    for (let i = 1; i < n; i++) {
        adj[parent[i]].push(i);
    }
    
    // 2. Determine the new parent for each node x (from 1 to n-1)
    // newParentMap[x] will store the new parent of x, or -1 if no change
    const newParentMap = new Map(); 
    
    for (let x = 1; x < n; x++) {
        let currentNode = x;
        // Traverse upwards to find the closest ancestor y with s[x] == s[y]
        let ancestor = parent[currentNode];
        while (ancestor !== -1) {
            if (s[x] === s[ancestor]) {
                newParentMap.set(x, ancestor);
                break;
            }
            ancestor = parent[ancestor];
        }
    }
    
    // 3. Build the final adjacency list based on new parent assignments
    const finalAdj = Array.from({ length: n }, () => []);
    for (let i = 0; i < n; i++) {
        // If node i's parent changed, add it to the new parent's children list
        if (newParentMap.has(i)) {
            finalAdj[newParentMap.get(i)].push(i);
        }
        // Otherwise, it remains a child of its original parent
        else if (i !== 0) { // Node 0 is the root, has no parent to add to
            finalAdj[parent[i]].push(i);
        }
    }
            
    // 4. Calculate subtree sizes using DFS (post-order traversal)
    const answer = new Array(n).fill(0);
    
    function dfsSize(node) {
        let size = 1; // Count the node itself
        for (const child of finalAdj[node]) {
            size += dfsSize(child);
        }
        answer[node] = size;
        return size;
    }
        
dfsSize(0); // Start DFS from the root
    
    return answer;
};
