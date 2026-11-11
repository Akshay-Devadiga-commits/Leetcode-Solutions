class Solution {
    public int minimumCost(int cost1, int cost2, int costBoth, int need1, int need2) {
        int minTotalCost = Integer.MAX_VALUE;

        // Iterate through all possible numbers of 'type 3' items we can buy.
        // The maximum useful number of 'type 3' items to consider is max(need1, need2).
        // If we buy more than max(need1, need2) of type 3, it doesn't help satisfy
        // any *additional* requirement beyond what's already covered by max(need1, need2).
        // However, it's possible that buying more than max(need1, need2) of type 3 items
        // is still cheaper if costBoth is very low and individual costs are high.
        // A safe upper bound for 'i' is max(need1, need2) + 1 to cover cases where
        // buying one extra 'type 3' item might be beneficial if it's cheaper than
        // individual items. Given the constraints, max(need1, need2) is at most 10.
        // So, iterating up to 11 or 12 is perfectly fine.
        // Let's consider up to max(need1, need2) + 1 to be safe.
        int upperBoundI = Math.max(need1, need2) + 1;

        for (int i = 0; i < upperBoundI; i++) {
            // Cost of buying 'i' items of type 3
            int currentCostBoth = i * costBoth;

            // Remaining needs after buying 'i' items of type 3
            int remainingNeed1 = Math.max(0, need1 - i);
            int remainingNeed2 = Math.max(0, need2 - i);

            // Cost to fulfill remaining needs with type 1 and type 2 items
            int currentCostIndividual = remainingNeed1 * cost1 + remainingNeed2 * cost2;

            // Total cost for this combination
            int totalCost = currentCostBoth + currentCostIndividual;

            // Update minimum cost found so far
            minTotalCost = Math.min(minTotalCost, totalCost);
        }

        // The loop covers all combinations where we use 'i' items of type 3 and then
        // fulfill the rest with type 1 and type 2. The maximum number of type 3 items
        // we'd ever *need* to consider is related to need1 and need2. If we buy
        // k items of type 3, we satisfy k of need1 and k of need2. The remaining
        // needs are max(0, need1-k) and max(0, need2-k). If k > max(need1, need2),
        // then at least one of need1-k or need2-k will be negative, meaning we've
        // already satisfied both needs. So, iterating up to max(need1, need2) is sufficient.
        // The upperBoundI = Math.max(need1, need2) + 1 ensures we check up to and including
        // the point where we might have satisfied both needs with type 3 items.

        return minTotalCost;
    }
}