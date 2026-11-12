class Solution {
    longestMonotonicSubarray(nums) {
        let ans = 1;
        let increasing = 1;
        let decreasing = 1;

        for (let i = 1; i < nums.length; i++) {
            if (nums[i] > nums[i - 1]) {
                increasing++;
                decreasing = 1;
            } else if (nums[i] < nums[i - 1]) {
                decreasing++;
                increasing = 1;
            } else {
                increasing = 1;
                decreasing = 1;
            }
            ans = Math.max(ans, increasing, decreasing);
        }

        return ans;
    }
}