/**
 * @param {number[]} nums
 * @return {number}
 */
var maxLength = function(nums) {
    const gcd = (a, b) => b === 0 ? a : gcd(b, a % b);
    const lcm = (a, b) => (a === 0 || b === 0) ? 0 : Math.abs(a * b) / gcd(a, b);
    let maxLen = 0;
    for (let i = 0; i < nums.length; i++) {
        let currProd = 1;
        let currGcd = 0;
        let currLcm = 1;
        for (let j = i; j < nums.length; j++) {
            let val = nums[j];
            currProd *= val;
            currGcd = gcd(currGcd, val);
            currLcm = lcm(currLcm, val);
            if (currProd === currLcm * currGcd) {
                maxLen = Math.max(maxLen, j - i + 1);
            }
        }
    }
    return maxLen;
};