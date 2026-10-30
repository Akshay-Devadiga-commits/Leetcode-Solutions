const { MinPriorityQueue } = require('@datastructures-js/priority-queue');

class Solution {
    shortestPath(n, edges, labels, k) {
        const adj = Array.from({ length: n }, () => []);
        for (const [u, v, w] of edges) {
            adj[u].push([v, w]);
        }

        const initChar = labels.charCodeAt(0);
        if (1 > k) return -1;

        const pq = new MinPriorityQueue(x => x[0]);
        pq.enqueue([0, 0, initChar, 1]);

        const dist = new Map();
        const getStateKey = (u, c, count) => `${u},${c},${count}`;
        dist.set(getStateKey(0, initChar, 1), 0);

        while (!pq.isEmpty()) {
            const [d, u, lastChar, count] = pq.dequeue();

            if (u === n - 1) return d;

            const key = getStateKey(u, lastChar, count);
            if (d > (dist.get(key) ?? Infinity)) continue;

            for (const [v, weight] of adj[u]) {
                const vChar = labels.charCodeAt(v);
                const newCount = (vChar === lastChar) ? count + 1 : 1;

                if (newCount <= k) {
                    const newDist = d + weight;
                    const nextKey = getStateKey(v, vChar, newCount);
                    if (newDist < (dist.get(nextKey) ?? Infinity)) {
                        dist.set(nextKey, newDist);
                        pq.enqueue([newDist, v, vChar, newCount]);
                    }
                }
            }
        }

        return -1;
    }
}