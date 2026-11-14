import java.util.Arrays;

class Solution {
    public int minimumSumSubarray(int[] nums, int l, int r) {
        int n = nums.length;

        // Calculate prefix sums
        // prefixSum[k] stores the sum of nums[0...k-1]
        // prefixSum[0] = 0
        // Use long for prefix sums to prevent potential overflow if sums can be large
        long[] prefixSum = new long[n + 1];
        for (int i = 0; i < n; i++) {
            prefixSum[i + 1] = prefixSum[i] + nums[i];
        }

        int minPositiveSum = Integer.MAX_VALUE;

        // Iterate over all possible start indices i
        for (int i = 0; i < n; i++) {
            // Iterate over all possible end indices j
            // j must be at least i to form a non-empty subarray
            for (int j = i; j < n; j++) {
                // Calculate the length of the current subarray nums[i...j]
                int currentLength = j - i + 1;

                // Check if the length constraint is met
                if (currentLength >= l && currentLength <= r) {
                    // Calculate the sum of the subarray nums[i...j] using prefix sums
                    long currentSumLong = prefixSum[j + 1] - prefixSum[i];
                    
                    // Check if the sum is positive and fits into an int
                    if (currentSumLong > 0 && currentSumLong <= Integer.MAX_VALUE) {
                        int currentSum = (int) currentSumLong;
                        // Update the minimum positive sum found so far
                        minPositiveSum = Math.min(minPositiveSum, currentSum);
                    }
                }
            }
        }

        // If minPositiveSum is still Integer.MAX_VALUE, it means no such subarray was found
        if (minPositiveSum == Integer.MAX_VALUE) {
            return -1;
        } else {
            return minPositiveSum;
        }
    }
}