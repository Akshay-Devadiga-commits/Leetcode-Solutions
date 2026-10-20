class Solution {
    sortableIntegers(nums) {
        const n = nums.length;
        
        const isSortedRotation = (sub) => {
            const m = sub.length;
            if (m <= 1) return true;
            let drops = 0;
            for (let i = 0; i < m; i++) {
                if (sub[i] > sub[(i + 1) % m]) {
                    drops++;
                }
            }
            return drops <= 1;
        };

        let totalSupply = 0;
        for (let k = 1; k <= n; k++) {
            if (n % k !== 0) continue;
            
            let valid = true;
            const sortedSubarrays = [];
            
            for (let i = 0; i < n; i += k) {
                const sub = nums.slice(i, i + k);
                if (!isSortedRotation(sub)) {
                    valid = false;
                    break;
                }
                const m = sub.length;
                let startIdx = 0;
                for (let j = 0; j < m; j++) {
                    if (sub[j] > sub[(j + 1) % m]) {
                        startIdx = (j + 1) % m;
                        break;
                    }
                }
                const rotated = sub.slice(startIdx).concat(sub.slice(0, startIdx));
                sortedSubarrays.push(rotated);
            }
            
            if (!valid) continue;
            
            let flat = [];
            for (const sub of sortedSubarrays) {
                flat = flat.concat(sub);
            }
            
            let isNonDecreasing = true;
            for (let i = 0; i < n - 1; i++) {
                if (flat[i] > flat[i + 1]) {
                    isNonDecreasing = false;
                    break;
                }
            }
            
            if (isNonDecreasing) {
                totalSupply += k;
            }
        }
        
        return totalSupply;
    }
}