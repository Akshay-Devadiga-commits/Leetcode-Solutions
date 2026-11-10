class Solution {
    public int climbStairs(int n, int[] costs) {
        int[] dp = new int[n + 1];
        for (int i = 1; i <= n; i++) {
            int minPrevCost = dp[i - 1] + 1;
            if (i >= 2) {
                minPrevCost = Math.min(minPrevCost, dp[i - 2] + 4);
            }
            if (i >= 3) {
                minPrevCost = Math.min(minPrevCost, dp[i - 3] + 9);
            }
            dp[i] = costs[i - 1] + minPrevCost;
        }
        return dp[n];
    }
}