class Solution {
    minOperations(s) {
        const sortedStr = s.split('').sort().join('');
        if (s === sortedStr) return 0;

        const n = s.length;
        const queue = [[s, 0]];
        const visited = new Set([s]);

        while (queue.length > 0) {
            const [curr, ops] = queue.shift();

            for (let i = 0; i < n; i++) {
                for (let j = i + 2; j <= n; j++) {
                    if (i === 0 && j === n) continue;

                    const sub = curr.substring(i, j);
                    const sortedSub = sub.split('').sort().join('');
                    const nextS = curr.substring(0, i) + sortedSub + curr.substring(j);

                    const nextSorted = nextS.split('').sort().join('');
                    if (nextS === nextSorted) {
                        return ops + 1;
                    }

                    if (!visited.has(nextS)) {
                        visited.add(nextS);
                        queue.push([nextS, ops + 1]);
                    }
                }
            }
        }

        return -1;
    }
}