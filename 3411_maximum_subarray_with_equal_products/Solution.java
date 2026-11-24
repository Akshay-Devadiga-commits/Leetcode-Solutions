class Solution {
    private long gcd(long a, long b) {
        while (b != 0) {
            a %= b;
            long temp = a;
            a = b;
            b = temp;
        }
        return a;
    }

    public int maxLength(int[] nums) {
        int n = nums.length;
        int maxLen = 0;
        for (int i = 0; i < n; i++) {
            long currProd = 1;
            long currGcd = 0;
            long currLcm = 1;
            for (int j = i; j < n; j++) {
                long val = nums[j];
                currProd *= val;
                currGcd = gcd(currGcd, val);
                currLcm = (currLcm / gcd(currLcm, val)) * val;
                if (currProd == currLcm * currGcd) {
                    maxLen = Math.max(maxLen, j - i + 1);
                }
            }
        }
        return maxLen;
    }
}