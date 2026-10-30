import java.util.*;

class Solution {
    public int shortestPath(int n, int[][] edges, String labels, int k) {
        List<int[]>[] adj = new ArrayList[n];
        for (int i = 0; i < n; i++) {
            adj[i] = new ArrayList<>();
        }
        for (int[] edge : edges) {
            adj[edge[0]].add(new int[]{edge[1], edge[2]});
        }

        char initChar = labels.charAt(0);
        if (1 > k) return -1;

        PriorityQueue<long[]> pq = new PriorityQueue<>(Comparator.comparingLong(a -> a[0]));
        pq.offer(new long[]{0, 0, initChar, 1});

        Map<String, Integer> dist = new HashMap<>();
        dist.put(0 + "," + initChar + "," + 1, 0);

        while (!pq.isEmpty()) {
            long[] curr = pq.poll();
            int d = (int) curr[0];
            int u = (int) curr[1];
            char lastChar = (char) curr[2];
            int count = (int) curr[3];

            if (u == n - 1) return d;

            String key = u + "," + lastChar + "," + count;
            if (d > dist.getOrDefault(key, Integer.MAX_VALUE)) continue;

            for (int[] neighbor : adj[u]) {
                int v = neighbor[0];
                int weight = neighbor[1];
                char vChar = labels.charAt(v);
                int newCount = (vChar == lastChar) ? count + 1 : 1;

                if (newCount <= k) {
                    int newDist = d + weight;
                    String nextKey = v + "," + vChar + "," + newCount;
                    if (newDist < dist.getOrDefault(nextKey, Integer.MAX_VALUE)) {
                        dist.put(nextKey, newDist);
                        pq.offer(new long[]{newDist, v, vChar, newCount});
                    }
                }
            }
        }

        return -1;
    }
}