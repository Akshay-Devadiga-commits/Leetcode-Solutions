import java.util.*;

class Solution {
    public int smallestUniqueSubarray(int[] nums) {
        int n = nums.length;
        
        for (int length = 1; length <= n; length++) {
            Map<List<Integer>, Integer> subarrayCounts = new HashMap<>();
            
            for (int i = 0; i <= n - length; i++) {
                List<Integer> subarray = new ArrayList<>();
                for (int j = i; j < i + length; j++) {
                    subarray.add(nums[j]);
                }
                subarrayCounts.put(subarray, subarrayCounts.getOrDefault(subarray, 0) + 1);
            }
            
            for (int count : subarrayCounts.values()) {
                if (count == 1) {
                    return length;
                }
            }
        }
        
        return n;
    }
}