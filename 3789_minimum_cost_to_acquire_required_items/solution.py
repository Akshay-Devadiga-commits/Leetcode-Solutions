import math

class Solution:
    def minimumCost(self, cost1: int, cost2: int, costBoth: int, need1: int, need2: int) -> int:
        min_total_cost = math.inf

        # Iterate through all possible numbers of 'type 3' items we can buy.
        # The maximum useful number of 'type 3' items to consider is max(need1, need2).
        # If we buy more than max(need1, need2) of type 3, it doesn't help satisfy
        # any *additional* requirement beyond what's already covered by max(need1, need2).
        # However, it's possible that buying more than max(need1, need2) of type 3 items
        # is still cheaper if costBoth is very low and individual costs are high.
        # A safe upper bound for 'i' is max(need1, need2) + 1 to cover cases where
        # buying one extra 'type 3' item might be beneficial if it's cheaper than
        # individual items. Given the constraints, max(need1, need2) is at most 10.
        # So, iterating up to 11 or 12 is perfectly fine.
        # Let's consider up to max(need1, need2) + 1 to be safe.
        upper_bound_i = max(need1, need2) + 1

        for i in range(upper_bound_i):
            # Cost of buying 'i' items of type 3
            current_cost_both = i * costBoth

            # Remaining needs after buying 'i' items of type 3
            remaining_need1 = max(0, need1 - i)
            remaining_need2 = max(0, need2 - i)

            # Cost to fulfill remaining needs with type 1 and type 2 items
            current_cost_individual = remaining_need1 * cost1 + remaining_need2 * cost2

            # Total cost for this combination
            total_cost = current_cost_both + current_cost_individual

            # Update minimum cost found so far
            min_total_cost = min(min_total_cost, total_cost)
            
        # Consider the case where we only buy type 1 and type 2 items, without any type 3.
        # This is covered by i=0 in the loop, but explicitly stating it can be clearer.
        # However, the loop handles it correctly.
        
        # Also, consider the case where it might be cheaper to buy 'need1' items of type 1
        # and 'need2' items of type 2, even if costBoth is low. This is also covered by i=0.
        
        # The loop covers all combinations where we use 'i' items of type 3 and then
        # fulfill the rest with type 1 and type 2. The maximum number of type 3 items
        # we'd ever *need* to consider is related to need1 and need2. If we buy
        # k items of type 3, we satisfy k of need1 and k of need2. The remaining
        # needs are max(0, need1-k) and max(0, need2-k). If k > max(need1, need2),
        # then at least one of need1-k or need2-k will be negative, meaning we've
        # already satisfied both needs. So, iterating up to max(need1, need2) is sufficient.
        # The upper_bound_i = max(need1, need2) + 1 ensures we check up to and including
        # the point where we might have satisfied both needs with type 3 items.

        return min_total_cost