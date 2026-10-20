import java.util.ArrayList;
import java.util.Arrays;
import java.util.Collections;
import java.util.List;

class Solution {
    public int sortableIntegers(int[] nums) {
        int n = nums.length;
        int[] target = Arrays.copyOf(nums, n);
        Arrays.sort(target);

        boolean isCyclicShift(List<Integer> sub, List<Integer> orig) {
            if (sub.size() != orig.size()) {
                return false;
            }
            List<Integer> doubleOrig = new ArrayList<>(orig);
            doubleOrig.addAll(orig);
            
            String subStr = sub.toString().replace(" ", "");
            String doubleOrigStr = doubleOrig.toString().replace(" ", "");
            
            return doubleOrigStr.substring(1, doubleOrigStr.length() - 1).contains(subStr.substring(1, subStr.length() - 1));
        }

        int ans = 0;
        for (int k = 1; k <= n; k++) {
            if (n % k == 0) {
                boolean possible = true;
                for (int i = 0; i < n; i += k) {
                    List<Integer> sub = new ArrayList<>();
                    for(int j=0; j<k; j++) sub.add(nums[i+j]);
                    
                    List<Integer> orig = new ArrayList<>();
                    for(int j=0; j<k; j++) orig.add(target[i+j]);

                    if (!isCyclicShift(sub, orig)) {
                        possible = false;
                        break;
                    }
                }
                if (possible) {
                    ans += k;
                }
            }
        }
        return ans;
    }
}