class Solution {
    /**
     * @param {number[]} nums
     * @return {number}
     */
    minMoves(nums) {
        const mn = Math.min(...nums);
        let sum = 0;
        for (const num of nums) {
            sum += num - mn;
        }
        return sum;
    }
}