class Solution {
    /**
     * @param {number[]} nums
     * @param {number} l
     * @param {number} r
     * @return {number}
     */
    minimumSumSubarray(nums, l, r) {
        const n = nums.length;
        let minPositiveSum = Infinity;
        let found = false;

        // Iterate through all possible start indices of the subarray
        for (let i = 0; i < n; i++) {
            let currentSum = 0;
            // Iterate through all possible end indices of the subarray starting from i
            for (let j = i; j < n; j++) {
                currentSum += nums[j];
                const subarrayLength = j - i + 1;

                // Check if the subarray length is within the given range [l, r]
                if (l <= subarrayLength && subarrayLength <= r) {
                    // Check if the sum is positive
                    if (currentSum > 0) {
                        minPositiveSum = Math.min(minPositiveSum, currentSum);
                        found = true;
                    }
                }
            }
        }

        // If no such subarray was found, return -1, otherwise return the minimum positive sum
        return found ? minPositiveSum : -1;
    }
}