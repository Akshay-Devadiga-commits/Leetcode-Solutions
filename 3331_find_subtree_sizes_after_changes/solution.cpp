#include <vector>
#include <string>
#include <unordered_map>
#include <functional>

class Solution {
public:
    std::vector<int> findSubtreeSizes(std::vector<int>& parent, std::string s) {
        int n = parent.size();
        
        // 1. Build the initial adjacency list (children for each node)
        std::vector<std::vector<int>> adj(n);
        for (int i = 1; i < n; ++i) {
            adj[parent[i]].push_back(i);
        }
        
        // 2. Determine the new parent for each node x (from 1 to n-1)
        // newParentMap[x] will store the new parent of x, or -1 if no change
        std::unordered_map<int, int> newParentMap;
        
        for (int x = 1; x < n; ++x) {
            int currentNode = x;
            // Traverse upwards to find the closest ancestor y with s[x] == s[y]
            int ancestor = parent[currentNode];
            while (ancestor != -1) {
                if (s[x] == s[ancestor]) {
                    newParentMap[x] = ancestor;
                    break;
                }
                ancestor = parent[ancestor];
            }
        }
        
        // 3. Build the final adjacency list based on new parent assignments
        std::vector<std::vector<int>> finalAdj(n);
        for (int i = 0; i < n; ++i) {
            // If node i's parent changed, add it to the new parent's children list
            if (newParentMap.count(i)) {
                finalAdj[newParentMap[i]].push_back(i);
            }
            // Otherwise, it remains a child of its original parent
            else if (i != 0) { // Node 0 is the root, has no parent to add to
                finalAdj[parent[i]].push_back(i);
            }
        }
                
        // 4. Calculate subtree sizes using DFS (post-order traversal)
        std::vector<int> answer(n);
        
        std::function<int(int)> dfsSize = 
            [&](int node) -> int {
            int size = 1; // Count the node itself
            for (int child : finalAdj[node]) {
                size += dfsSize(child);
            }
            answer[node] = size;
            return size;
        };
            
        dfsSize(0); // Start DFS from the root
        
        return answer;
    }
};
