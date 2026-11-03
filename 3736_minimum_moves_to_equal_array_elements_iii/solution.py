from typing import List

class Solution:
    def minMoves(self, nums: List[int]) -> int:
        """
        Calculates the minimum number of moves to make all elements in an array equal by only increasing elements.
        
        Args:
            nums: A list of integers.
            
        Returns:
            The minimum total number of moves required.
        """
        
        if not nums:
            return 0
            
        # The optimal strategy is to make all elements equal to the maximum element in the array.
        # This is because we can only increase elements. Any target value greater than the maximum
        # would require additional moves for every element.
        
        max_val = nums[0]
        total_sum = 0
        
        for num in nums:
            if num > max_val:
                max_val = num
            total_sum += num
            
        # The total number of moves is the sum of (max_val - num) for each num in nums.
        # This can be rewritten as (n * max_val) - sum(nums).
        return len(nums) * max_val - total_sum