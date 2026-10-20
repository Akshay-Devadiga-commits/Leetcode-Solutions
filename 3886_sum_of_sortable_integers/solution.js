class Solution {
    sortableIntegers(nums) {
        const n = nums.length;
        
        const canSortWithK = (k) => {
            const subarrays = [];
            for (let i = 0; i < n; i += k) {
                const sub = nums.slice(i, i + k);
                let drops = 0;
                for (let j = 0; j < k; j++) {
                    if (sub[j] > sub[(j + 1) % k]) {
                        drops++;
                    }
                }
                if (drops > 1) return false;
                
                let rotatedSub;
                if (drops === 1) {
                    let split = 0;
                    for (let j = 0; j < k; j++) {
                        if (sub[j] > sub[(j + 1) % k]) {
                            split = j + 1;
                            break;
                        }
                    }
                    rotatedSub = sub.concat(sub.slice(0, split));
                    rotatedSub.splice(0, split);
                } else {
                    rotatedSub = sub;
                }
                subarrays.push(rotatedSub);
            }
            
            for (let i = 0; i < subarrays.length - 1; i++) {
                if (subarrays[i][subarrays[i].length - 1] > subarrays[i + 1][0]) {
                    return false;
                }
            }
            
            return true;
        };

        let totalSum = 0;
        for (let k = 1; k <= n; k++) {
            if (n % k === 0) {
                if (canSortWithK(k)) {
                    totalSum += k;
                }
            }
        }
        return totalSum;
    }
}