class Solution {
    minimizeSum(nums) {
        nums.sort((a, b) => a - b);
        // Can always change the number to any other number in `nums`, so `low` becomes 0.
        // Thus, rephrase the problem as finding the minimum `high`.
        const highOfChangingTwoMins = nums[nums.length - 1] - nums[2];
        const highOfChangingTwoMaxs = nums[nums.length - 3] - nums[0];
        const highOfChangingMinAndMax = nums[nums.length - 2] - nums[1];
        return Math.min(highOfChangingTwoMins, highOfChangingTwoMaxs, highOfChangingMinAndMax);
    }
}