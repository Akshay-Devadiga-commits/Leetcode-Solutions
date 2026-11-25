/**
 * @param {number[]} nums
 * @return {number}
 */
var smallestUniqueSubarray = function(nums) {
    const n = nums.length;
    
    for (let length = 1; length <= n; length++) {
        const subarrayCounts = new Map();
        
        for (let i = 0; i <= n - length; i++) {
            const subarray = nums.slice(i, i + length).join(',');
            subarrayCounts.set(subarray, (subarrayCounts.get(subarray) || 0) + 1);
        }
        
        for (const count of subarrayCounts.values()) {
            if (count === 1) {
                return length;
            }
        }
    }
    
    return n;
};