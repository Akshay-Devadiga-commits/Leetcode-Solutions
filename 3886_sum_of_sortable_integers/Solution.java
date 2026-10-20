import java.util.*;

class Solution {
    public int sortableIntegers(int[] nums) {
        int n = nums.length;
        int totalSupply = 0;

        for (int k = 1; k <= n; k++) {
            if (n % k != 0) continue;

            boolean valid = true;
            List<int[]> sortedSubarrays = new ArrayList<>();

            for (int i = 0; i < n; i += k) {
                int[] sub = new int[k];
                System.arraycopy(nums, i, sub, 0, k);

                if (!isSortedRotation(sub)) {
                    valid = false;
                    break;
                }

                int m = sub.length;
                int startIdx = 0;
                for (int j = 0; j < m; j++) {
                    if (sub[j] > sub[(j + 1) % m]) {
                        startIdx = (j + 1) % m;
                        break;
                    }
                }

                int[] rotated = new int[m];
                for (int j = 0; j < m; j++) {
                    rotated[j] = sub[(startIdx + j) % m];
                }
                sortedSubarrays.add(rotated);
            }

            if (!valid) continue;

            int[] flat = new int[n];
            int idx = 0;
            for (int[] sub : sortedSubarrays) {
                for (int val : sub) {
                    flat[idx++] = val;
                }
            }

            boolean isNonDecreasing = true;
            for (int i = 0; i < n - 1; i++) {
                if (flat[i] > flat[i + 1]) {
                    isNonDecreasing = false;
                    break;
                }
            }

            if (isNonDecreasing) {
                totalSupply += k;
            }
        }

        return totalSupply;
    }

    private boolean isSortedRotation(int[] sub) {
        int m = sub.length;
        if (m <= 1) return true;
        int drops = 0;
        for (int i = 0; i < m; i++) {
            if (sub[i] > sub[(i + 1) % m]) {
                drops++;
            }
        }
        return drops <= 1;
    }
}