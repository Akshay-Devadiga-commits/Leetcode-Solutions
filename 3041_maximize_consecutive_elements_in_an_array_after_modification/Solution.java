import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

class Solution {
    public int maxSelectedElements(int[] nums) {
        Arrays.sort(nums);
        Map<Integer, Integer> dp = new HashMap<>();
        int maxLength = 0;
        
        for (int x : nums) {
            int dpX = dp.getOrDefault(x, 0);
            int dpXMinus1 = dp.getOrDefault(x - 1, 0);
            
            int newDpXPlus1 = dpX + 1;
            int newDpX = dpXMinus1 + 1;
            
            dp.put(x + 1, newDpXPlus1);
            dp.put(x, newDpX);
            
            maxLength = Math.max(maxLength, Math.max(newDpXPlus1, newDpX));
        }
        
        return maxLength;
    }
}