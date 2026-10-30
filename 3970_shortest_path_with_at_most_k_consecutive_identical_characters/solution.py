import heapq
from typing import List

class Solution:
    def shortestPath(self, n: int, edges: List[List[int]], labels: str, k: int) -> int:
        adj = [[] for _ in range(n)]
        for u, v, w in edges:
            adj[u].append((v, w))
            
        pq = [(0, 0, labels[0], 1)]
        dist = {}
        
        while pq:
            w, u, last_char, count = heapq.heappop(pq)
            
            if u == n - 1:
                return w
                
            state = (u, last_char, count)
            if state in dist and dist[state] < w:
                continue
            dist[state] = w
            
            for v, weight in adj[u]:
                nxt_char = labels[v]
                if nxt_char == last_char:
                    nxt_count = count + 1
                else:
                    nxt_count = 1
                    
                if nxt_count <= k:
                    nxt_state = (v, nxt_char, nxt_count)
                    nxt_w = w + weight
                    if nxt_state not in dist or nxt_w < dist[nxt_state]:
                        dist[nxt_state] = nxt_w
                        heapq.heappush(pq, (nxt_w, v, nxt_char, nxt_count))
                        
        return -1