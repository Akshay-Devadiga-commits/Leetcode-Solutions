import java.util.Arrays;

class Solution {
    public int minimizeSum(int[] nums) {
        Arrays.sort(nums);
        // Can always change the number to any other number in `nums`, so `low` becomes 0.
        // Thus, rephrase the problem as finding the minimum `high`.
        int highOfChangingTwoMins = nums[nums.length - 1] - nums[2];
        int highOfChangingTwoMaxs = nums[nums.length - 3] - nums[0];
        int highOfChangingMinAndMax = nums[nums.length - 2] - nums[1];
        return Math.min(highOfChangingTwoMins, Math.min(highOfChangingTwoMaxs, highOfChangingMinAndMax));
    }
}