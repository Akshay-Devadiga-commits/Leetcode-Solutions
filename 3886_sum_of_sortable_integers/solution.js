class Solution {
    sortableIntegers(nums) {
        const n = nums.length;
        const target = [...nums].sort((a, b) => a - b);

        const isCyclicShift = (sub, orig) => {
            if (sub.length !== orig.length) {
                return false;
            }
            const doubleOrig = [...orig, ...orig];
            const subStr = sub.join(',');
            const doubleOrigStr = doubleOrig.join(',');
            return doubleOrigStr.includes(subStr);
        };

        let ans = 0;
        for (let k = 1; k <= n; k++) {
            if (n % k === 0) {
                let possible = true;
                for (let i = 0; i < n; i += k) {
                    const sub = nums.slice(i, i + k);
                    const orig = target.slice(i, i + k);
                    if (!isCyclicShift(sub, orig)) {
                        possible = false;
                        break;
                    }
                }
                if (possible) {
                    ans += k;
                }
            }
        }
        return ans;
    }
}