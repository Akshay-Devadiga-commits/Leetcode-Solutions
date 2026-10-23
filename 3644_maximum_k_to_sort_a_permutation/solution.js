class Solution {
    sortPermutation(nums) {
        const n = nums.length;
        let isSorted = true;
        for (let i = 0; i < n - 1; i++) {
            if (nums[i] > nums[i + 1]) {
                isSorted = false;
                break;
            }
        }
        if (isSorted) return 0;

        let maxVal = 0;
        for (const num of nums) {
            if (num > maxVal) maxVal = num;
        }

        for (let k = maxVal; k >= 0; k--) {
            const parent = Array.from({ length: n }, (_, i) => i);
            
            const find = (i) => {
                let root = i;
                while (root !== parent[root]) {
                    root = parent[root];
                }
                let curr = i;
                while (curr !== root) {
                    let nxt = parent[curr];
                    parent[curr] = root;
                    curr = nxt;
                }
                return root;
            };

            const union = (i, j) => {
                const rootI = find(i);
                const rootJ = find(j);
                if (rootI !== rootJ) {
                    parent[rootI] = rootJ;
                }
            };

            for (let i = 0; i < n; i++) {
                for (let j = i + 1; j < n; j++) {
                    if ((nums[i] & nums[j]) === k) {
                        union(i, j);
                    }
                }
            }

            const components = new Map();
            for (let i = 0; i < n; i++) {
                const root = find(i);
                if (!components.has(root)) {
                    components.set(root, []);
                }
                components.get(root).push(i);
            }

            let possible = true;
            for (const indices of components.values()) {
                const currentValues = indices.map(i => nums[i]).sort((a, b) => a - b);
                const targetValues = [...indices].sort((a, b) => a - b);
                
                for (let i = 0; i < currentValues.length; i++) {
                    if (currentValues[i] !== targetValues[i]) {
                        possible = false;
                        break;
                    }
                }
                if (!possible) break;
            }

            if (possible) return k;
        }

        return 0;
    }
}