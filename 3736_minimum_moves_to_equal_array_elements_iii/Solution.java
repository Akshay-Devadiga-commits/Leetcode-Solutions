import java.util.List;

class Solution {
    /**
     * Calculates the minimum number of moves to make all elements in an array equal by only increasing elements.
     *
     * @param nums A list of integers.
     * @return The minimum total number of moves required.
     */
    public int minMoves(int[] nums) {
        if (nums == null || nums.length == 0) {
            return 0;
        }
        
        int maxVal = nums[0];
        long totalSum = 0; // Use long to prevent potential overflow for sum
        
        for (int num : nums) {
            if (num > maxVal) {
                maxVal = num;
            }
            totalSum += num;
        }
        
        // The total number of moves is the sum of (max_val - num) for each num in nums.
        // This can be rewritten as (n * max_val) - sum(nums).
        // Cast nums.length to long before multiplication to prevent overflow if maxVal is large.
        return (int) ((long)nums.length * maxVal - totalSum);
    }
}