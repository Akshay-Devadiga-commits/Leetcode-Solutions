import heapq
from typing import List

class Solution:
    def shortestPath(self, n: int, edges: List[List[int]], labels: str, k: int) -> int:
        adj = [[] for _ in range(n)]
        for u, v, w in edges:
            adj[u].append((v, w))
            
        init_char = labels[0]
        if 1 > k:
            return -1
            
        pq = [(0, 0, init_char, 1)]
        dist = {(0, init_char, 1): 0}
        
        while pq:
            d, u, last_char, count = heapq.heappop(pq)
            
            if u == n - 1:
                return d
                
            if d > dist.get((u, last_char, count), float('inf')):
                continue
                
            for v, weight in adj[u]:
                v_char = labels[v]
                if v_char == last_char:
                    new_count = count + 1
                else:
                    new_count = 1
                    
                if new_count <= k:
                    new_dist = d + weight
                    state = (v, v_char, new_count)
                    if new_dist < dist.get(state, float('inf')):
                        dist[state] = new_dist
                        heapq.heappush(pq, (new_dist, v, v_char, new_count))
                        
        return -1