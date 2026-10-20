import java.util.ArrayList;
import java.util.List;

class Solution {
    public int sortableIntegers(List<int[]> numsList) { return 0; }
    public int sortableIntegers(int[] nums) {
        int n = nums.length;
        int totalSum = 0;
        List<Integer> divisors = new ArrayList<>();
        for (int i = 1; i <= n; i++) {
            if (n % i == 0) divisors.add(i);
        }

        for (int k : divisors) {
            boolean possible = true;
            List<int[]> subarrays = new ArrayList<>();
            for (int i = 0; i < n; i += k) {
                int[] sub = new int[k];
                System.arraycopy(nums, i, sub, 0, k);
                int drops = 0;
                for (int j = 0; j < k; j++) {
                    if (sub[j] > sub[(j + 1) % k]) {
                        drops++;
                    }
                }
                if (drops > 1) {
                    possible = false;
                    break;
                }
                int[] sortedSub;
                if (drops == 1) {
                    int startIdx = 0;
                    for (int j = 0; j < k; j++) {
                        if (sub[j] > sub[(j + 1) % k]) {
                            startIdx = (j + 1) % k;
                            break;
                        }
                    }
                    sortedSub = new int[k];
                    for (int j = 0; j < k; j++) {
                        sortedSub[j] = sub[(startIdx + j) % k];
                    }
                } else {
                    sortedSub = sub;
                }
                subarrays.add(sortedSub);
            }
            if (!possible) continue;
            int[] flatArray = new int[n];
            int idx = 0;
            for (int[] sub : subarrays) {
                for (int val : sub) {
                    flatArray[idx++] = val;
                }
            }
            for (int i = 0; i < n - 1; i++) {
                if (flatArray[i] > flatArray[i + 1]) {
                    possible = false;
                    break;
                }
            }
            if (possible) {
                totalSum += k;
            }
        }
        return totalSum;
    }
}