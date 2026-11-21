import collections

class Solution:
    def findSubtreeSizes(self, parent: list[int], s: str) -> list[int]:
        n = len(parent)
        
        # 1. Build the initial adjacency list (children for each node)
        adj = collections.defaultdict(list)
        for i in range(1, n):
            adj[parent[i]].append(i)
            
        # 2. Determine the new parent for each node x (from 1 to n-1)
        # new_parent_map[x] will store the new parent of x, or -1 if no change
        new_parent_map = {} 
        
        for x in range(1, n):
            current_node = x
            # Traverse upwards to find the closest ancestor y with s[x] == s[y]
            ancestor = parent[current_node]
            while ancestor != -1:
                if s[x] == s[ancestor]:
                    new_parent_map[x] = ancestor
                    break
                ancestor = parent[ancestor]
        
        # 3. Build the final adjacency list based on new parent assignments
        final_adj = collections.defaultdict(list)
        for i in range(n):
            # If node i's parent changed, add it to the new parent's children list
            if i in new_parent_map:
                final_adj[new_parent_map[i]].append(i)
            # Otherwise, it remains a child of its original parent
            elif i != 0: # Node 0 is the root, has no parent to add to
                final_adj[parent[i]].append(i)
                
        # 4. Calculate subtree sizes using DFS (post-order traversal)
        answer = [0] * n
        
        def dfs_size(node):
            size = 1 # Count the node itself
            for child in final_adj[node]:
                size += dfs_size(child)
            answer[node] = size
            return size
            
        dfs_size(0) # Start DFS from the root
        
        return answer
