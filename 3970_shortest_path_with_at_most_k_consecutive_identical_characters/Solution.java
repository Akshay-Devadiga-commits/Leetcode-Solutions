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

        PriorityQueue<long[]> pq = new PriorityQueue<>(Comparator.comparingLong(a -> a[0]));
        // Queue stores: (weight, node, last_char_index, count)
        pq.offer(new long[]{0, 0, labels.charAt(0) - 'a', 1});

        Map<String, Long> dist = new HashMap<>();

        while (!pq.isEmpty()) {
            long[] curr = pq.poll();
            long w = curr[0];
            int u = (int) curr[1];
            int lastChar = (int) curr[2];
            int count = (int) curr[3];

            if (u == n - 1) {
                return (int) w;
            }

            String state = u + "," + lastChar + "," + count;
            if (dist.getOrDefault(state, Long.MAX_VALUE) < w) {
                continue;
            }
            dist.put(state, w);

            for (int[] neighbor : adj[u]) {
                int v = neighbor[0];
                int weight = neighbor[1];
                int nxtChar = labels.charAt(v) - 'a';
                int nxtCount = (nxtChar == lastChar) ? count + 1 : 1;

                if (nxtCount <= k) {
                    long nxtW = w + weight;
                    String nxtState = v + "," + nxtChar + "," + nxtCount;
                    if (nwtDistCheck(dist, nxtState, nxtW)) {
                        dist.put(nxtState, nxtW);
                        pq.offer(new long[]{nxtW, v, nxtChar, nxtCount});
                    }
                }
            }
        }

        return -1;
    }

    private boolean nwtDistCheck(Map<String, Long> dist, String state, long w) {
        return !dist.containsKey(state) || w < dist.get(state);
    }
}