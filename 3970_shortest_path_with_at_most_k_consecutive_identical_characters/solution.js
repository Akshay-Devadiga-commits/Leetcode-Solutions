const { MinPriorityQueue } = require('@datastructures-js/priority-queue');

class Solution {
    shortestPath(n, edges, labels, k) {
        const adj = Array.from({ length: n }, () => []);
        for (const [u, v, w] of edges) {
            adj[u].push([v, w]);
        }

        const pq = new MinPriorityQueue({ priority: x => x[0] });
        pq.enqueue([0, 0, labels[0], 1]);

        const dist = new Map();

        while (!pq.isEmpty()) {
            const [w, u, lastChar, count] = pq.dequeue().element;

            if (u === n - 1) {
                return w;
            }

            const stateKey = `${u},${lastChar},${count}`;
            if (dist.has(stateKey) && dist.get(stateKey) < w) {
                continue;
            }
            dist.set(stateKey, w);

            for (const [v, weight] of adj[u]) {
                const nxtChar = labels[v];
                const nxtCount = (nxtChar === lastChar) ? count + 1 : 1;

                if (nxtCount <= k) {
                    const nxtW = w + weight;
                    const nxtStateKey = `${v},${nxtChar},${nxtCount}`;
                    if (!dist.has(nxtStateKey) || nxtW < dist.get(nxtStateKey)) {
                        dist.set(nxtStateKey, nxtW);
                        pq.enqueue([nxtW, v, nxtChar, nxtCount]);
                    }
                }
            }
        }

        return -1;
    }
}