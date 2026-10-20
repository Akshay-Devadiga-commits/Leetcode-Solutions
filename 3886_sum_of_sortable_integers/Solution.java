class Solution {
    public int sortableIntegers(int[] nums) {
        int n = nums.length;
        int totalSum = 0;
        
        for (int k = 1; k <= n; k++) {
            if (n % k == 0) {
                if (canSortWithK(nums, k)) {
                    totalSum += k;
                }
            }
        }
        return totalSum;
    }
    
    private boolean canSortWithK(int[] nums, int k) {
        int n = nums.length;
        int prevMax = -1;
        
        for (int i = 0; i < n; i += k) {
            int drops = 0;
            int split = 0;
            for (int j = 0; j < k; j++) {
                if (nums[i + j] > nums[i + (j + 1) % k]) {
                    drops++;
                    split = j + 1;
                }
            }
            if (drops > 1) return false;
            
            int minVal = nums[i + split];
            int maxVal = nums[i + (split - 1 + k) % k];
            
            if (prevMax != -1 && prevMax > minVal) {
                return false;
            }
            prevMax = maxVal;
        }
        return true;
    }
}