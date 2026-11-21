import java.util.*;

class Solution {
    public int[] findSubtreeSizes(int[] parent, String s) {
        int n = parent.length;
        
        // 1. Build the initial adjacency list (children for each node)
        List<List<Integer>> adj = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            adj.add(new ArrayList<>());
        }
        for (int i = 1; i < n; i++) {
            adj.get(parent[i]).add(i);
        }
        
        // 2. Determine the new parent for each node x (from 1 to n-1)
        // newParentMap[x] will store the new parent of x, or -1 if no change
        Map<Integer, Integer> newParentMap = new HashMap<>(); 
        
        for (int x = 1; x < n; x++) {
            int currentNode = x;
            // Traverse upwards to find the closest ancestor y with s[x] == s[y]
            int ancestor = parent[currentNode];
            while (ancestor != -1) {
                if (s.charAt(x) == s.charAt(ancestor)) {
                    newParentMap.put(x, ancestor);
                    break;
                }
                ancestor = parent[ancestor];
            }
        }
        
        // 3. Build the final adjacency list based on new parent assignments
        List<List<Integer>> finalAdj = new ArrayList<>();
        for (int i = 0; i < n; i++) {
            finalAdj.add(new ArrayList<>());
        }
        for (int i = 0; i < n; i++) {
            // If node i's parent changed, add it to the new parent's children list
            if (newParentMap.containsKey(i)) {
                finalAdj.get(newParentMap.get(i)).add(i);
            }
            // Otherwise, it remains a child of its original parent
            else if (i != 0) { // Node 0 is the root, has no parent to add to
                finalAdj.get(parent[i]).add(i);
            }
        }
                
        // 4. Calculate subtree sizes using DFS (post-order traversal)
        int[] answer = new int[n];
        
        dfsSize(0, finalAdj, answer);
        
        return answer;
    }

    private int dfsSize(int node, List<List<Integer>> finalAdj, int[] answer) {
        int size = 1; // Count the node itself
        for (int child : finalAdj.get(node)) {
            size += dfsSize(child, finalAdj, answer);
        }
        answer[node] = size;
        return size;
    }
}
