class Solution {
    minOperations(nums) {
        const n = nums.length;
        const start = nums.join(',');
        const targetArr = Array.from({length: n}, (_, i) => i);
        const target = targetArr.join(',');

        if (start === target) return 0;

        const queue = [[start, 0]];
        const visited = new Set([start]);

        while (queue.length > 0) {
            const [currStr, dist] = queue.shift();
            const curr = currStr.split(',').map(Number);

            const nextRev = [...curr].reverse().join(',');
            if (nextRev === target) return dist + 1;
            if (!visited.has(nextRev)) {
                visited.add(nextRev);
                queue.push([nextRev, dist + 1]);
            }

            const nextRot = [...curr.slice(1), curr[0]].join(',');
            if (nextRot === target) return dist + 1;
            if (!visited.has(nextRot)) {
                visited.add(nextRot);
                queue.push([nextRot, dist + 1]);
            }
        }
        return -1;
    }
}