import java.util.ArrayList;
import java.util.List;

class Solution {
    public List<Integer> solve(List<Integer> data) {
        if (data.size() <= 1) {
            return data;
        }
        
        int mid = data.size() / 2;
        List<Integer> leftHalf = solve(data.subList(0, mid));
        List<Integer> rightHalf = solve(data.subList(mid, data.size()));
        
        return merge(leftHalf, rightHalf);
    }
    
    private List<Integer> merge(List<Integer> left, List<Integer> right) {
        List<Integer> sortedList = new ArrayList<>();
        int i = 0, j = 0;
        
        while (i < left.size() && j < right.size()) {
            if (left.get(i) < right.get(j)) {
                sortedList.add(left.get(i));
                i++;
            } else {
                sortedList.add(right.get(j));
                j++;
            }
        }
        
        while (i < left.size()) {
            sortedList.add(left.get(i));
            i++;
        }
        
        while (j < right.size()) {
            sortedList.add(right.get(j));
            j++;
        }
        
        return sortedList;
    }
}