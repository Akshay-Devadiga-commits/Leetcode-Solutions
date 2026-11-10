/**
 * @param {number} n
 * @param {number[]} costs
 * @return {number}
 */
var climbStairs = function(n, costs) {
    const dp = new Array(n + 1).fill(0);
    for (let i = 1; i <= n; i++) {
        let minPrevCost = dp[i - 1] + 1;
        if (i >= 2) {
            minPrevCost = Math.min(minPrevCost, dp[i - 2] + 4);
        }
        if (i >= 3) {
            minPrevCost = Math.min(minPrevCost, dp[i - 3] + 9);
        }
        dp[i] = costs[i - 1] + minPrevCost;
    }
    return dp[n];
};