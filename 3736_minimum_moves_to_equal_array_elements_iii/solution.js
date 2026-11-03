class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    minMoves(nums) {
        if (!nums || nums.length === 0) {
            return 0;
        }
        
        let maxVal = nums[0];
        let totalSum = 0;
        
        for (const num of nums) {
            if (num > maxVal) {
                maxVal = num;
            }
            totalSum += num;
        }
        
        return nums.length * maxVal - totalSum;
    }
}