import java.util.*;

class Solution {
    private int[] parent;

    private int find(int i) {
        if (parent[i] == i) {
            return i;
        }
        return parent[i] = find(parent[i]);
    }

    private void union(int i, int j) {
        int rootI = find(i);
        int rootJ = find(j);
        if (rootI != rootJ) {
            parent[rootI] = rootJ;
        }
    }

    public int sortPermutation(int[] nums) {
        int n = nums.length;
        boolean isSorted = true;
        for (int i = 0; i < n - 1; i++) {
            if (nums[i] > nums[i + 1]) {
                isSorted = false;
                break;
            }
        }
        if (isSorted) return 0;

        int maxVal = 0;
        for (int num : nums) {
            maxVal = Math.max(maxVal, num);
        }

        for (int k = maxVal; k >= 0; k--) {
            parent = new int[n];
            for (int i = 0; i < n; i++) {
                parent[i] = i;
            }

            for (int i = 0; i < n; i++) {
                for (int j = i + 1; j < n; j++) {
                    if ((nums[i] & nums[j]) == k) {
                        union(i, j);
                    }
                }
            }

            Map<Integer, List<Integer>> components = new HashMap<>();
            for (int i = 0; i < n; i++) {
                int root = find(i);
                components.computeIfAbsent(root, x -> new ArrayList<>()).add(i);
            }

            boolean possible = true;
            for (Map.Entry<Integer, List<Integer>> entry : components.entrySet()) {
                List<Integer> indices = entry.getValue();
                List<Integer> currentValues = new ArrayList<>();
                for (int idx : indices) {
                    currentValues.add(nums[idx]);
                }
                Collections.sort(currentValues);
                List<Integer> targetValues = new ArrayList<>(indices);
                Collections.sort(targetValues);

                if (!currentValues.equals(targetValues)) {
                    possible = false;
                    break;
                }
            }

            if (possible) return k;
        }

        return 0;
    }
}