import java.util.Arrays;

class Solution {
    public int minMoves(int[] nums) {
        int mn = nums[0];
        for (int num : nums) {
            mn = Math.min(mn, num);
        }
        int sum = 0;
        for (int num : nums) {
            sum += num - mn;
        }
        return sum;
    }
}