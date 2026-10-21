import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<Integer> solve(List<Integer> data) {
        if (data == null || data.size() <= 1) {
            return data;
        }
        int pivot = data.get(data.size() / 2);
        List<Integer> left = new ArrayList<>();
        List<Integer> middle = new ArrayList<>();
        List<Integer> right = new ArrayList<>();
        
        for (int x : data) {
            if (x < pivot) {
                left.add(x);
            } else if (x == pivot) {
                middle.add(x);
            } else {
                right.add(x);
            }
        }
        
        List<Integer> result = new ArrayList<>();
        result.addAll(solve(left));
        result.addAll(middle);
        result.addAll(solve(right));
        return result;
    }
}