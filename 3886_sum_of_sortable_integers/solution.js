class Solution {
    sortableIntegers(nums) {
        const n = nums.length;
        let totalSum = 0;
        const divisors = [];
        for (let i = 1; i <= n; i++) {
            if (n % i === 0) divisors.push(i);
        }

        for (const k of divisors) {
            let possible = true;
            const subarrays = [];
            for (let i = 0; i < n; i += k) {
                const sub = nums.slice(i, i + k);
                let drops = 0;
                for (let j = 0; j < k; j++) {
                    if (sub[j] > sub[(j + 1) % k]) {
                        drops++;
                    }
                }
                if (drops > 1) {
                    possible = false;
                    break;
                }
                let sortedSub;
                if (drops === 1) {
                    let startIdx = 0;
                    for (let j = 0; j < k; j++) {
                        if (sub[j] > sub[(j + 1) % k]) {
                            startIdx = (j + 1) % k;
                            break;
                        }
                    }
                    sortedSub = sub.slice(startIdx).concat(sub.slice(0, startIdx));
                } else {
                    sortedSub = sub;
                }
                subarrays.push(sortedSub);
            }
            if (!possible) continue;
            const flatArray = [].concat(...subarrays);
            for (let i = 0; i < n - 1; i++) {
                if (flatArray[i] > flatArray[i + 1]) {
                    possible = false;
                    break;
                }
            }
            if (possible) {
                totalSum += k;
            }
        }
        return totalSum;
    }
}