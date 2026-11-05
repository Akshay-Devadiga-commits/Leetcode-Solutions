import bisect

class Solution:
    def findTheDistanceValue(self, arr1: list[int], arr2: list[int], d: int) -> int:
        arr2.sort()
        distance_value = 0
        
        for num1 in arr1:
            # We need to check if there is any num2 in arr2 such that |num1 - num2| <= d
            # This is equivalent to checking if there is any num2 such that num1 - d <= num2 <= num1 + d
            
            # Find the index of the first element in arr2 that is >= num1 - d
            # bisect_left returns the insertion point for num1 - d to maintain order.
            # If num1 - d is present, it points to the first occurrence.
            # If num1 - d is not present, it points to the first element greater than num1 - d.
            left_bound_index = bisect.bisect_left(arr2, num1 - d)
            
            # Check if this index is valid and if the element at this index satisfies the condition
            # The element arr2[left_bound_index] is the smallest element >= num1 - d.
            # We need to check if this element is also <= num1 + d.
            
            found_close_element = False
            if left_bound_index < len(arr2):
                if arr2[left_bound_index] <= num1 + d:
                    found_close_element = True
            
            # If no element in arr2 is found within the range [num1 - d, num1 + d],
            # then num1 contributes to the distance value.
            if not found_close_element:
                distance_value += 1
                
        return distance_value