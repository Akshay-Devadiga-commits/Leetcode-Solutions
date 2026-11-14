class Solution {
    /**
     * @param {number[]} nums
     * @param {number} l
     * @param {number} r
     * @return {number}
     */
    minimumSumSubarray(nums, l, r) {
        const n = nums.length;

        // Calculate prefix sums
        // prefixSum[k] stores the sum of nums[0...k-1]
        // prefixSum[0] = 0
        const prefixSum = new Array(n + 1).fill(0);
        for (let i = 0; i < n; i++) {
            prefixSum[i + 1] = prefixSum[i] + nums[i];
        }

        let minPositiveSum = Number.POSITIVE_INFINITY;

        // Iterate over all possible start indices i
        for (let i = 0; i < n; i++) {
            // Iterate over all possible end indices j
            // j must be at least i to form a non-empty subarray
            for (let j = i; j < n; j++) {
                // Calculate the length of the current subarray nums[i...j]
                const currentLength = j - i + 1;

                // Check if the length constraint is met
                if (currentLength >= l && currentLength <= r) {
                    // Calculate the sum of the subarray nums[i...j] using prefix sums
                    const currentSum = prefixSum[j + 1] - prefixSum[i];

                    // Check if the sum is positive
                    if (currentSum > 0) {
                        // Update the minimum positive sum found so far
                        minPositiveSum = Math.min(minPositiveSum, currentSum);
                    }
                }
            }
        }

        // If minPositiveSum is still infinity, it means no such subarray was found
        if (minPositiveSum === Number.POSITIVE_INFINITY) {
            return -1;
        } else {
            return minPositiveSum;
        }
    }
}