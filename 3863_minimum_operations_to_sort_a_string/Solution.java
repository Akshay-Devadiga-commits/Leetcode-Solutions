import java.util.*;

class Solution {
    public int minOperations(String s) {
        if (isSorted(s)) return 0;
        
        int n = s.length();
        Set<String> visited = new HashSet<>();
        Queue<Pair> queue = new LinkedList<>();
        
        visited.add(s);
        queue.offer(new Pair(s, 0));
        
        while (!queue.isEmpty()) {
            Pair curr = queue.poll();
            
            for (int i = 0; i < n; i++) {
                for (int j = i + 1; j <= n; j++) {
                    if (i == 0 && j == n) continue;
                    
                    char[] sub = curr.str.substring(i, j).toCharArray();
                    Arrays.sort(sub);
                    String nextStr = curr.str.substring(0, i) + new String(sub) + curr.str.substring(j);
                    
                    if (isSorted(nextStr)) return curr.ops + 1;
                    
                    if (!visited.contains(nextStr)) {
                        visited.add(nextStr);
                        queue.offer(new Pair(nextStr, curr.ops + 1));
                    }
                }
            }
        }
        
        return -1;
    }
    
    private boolean isSorted(String s) {
        for (int i = 1; i < s.length(); i++) {
            if (s.charAt(i) < s.charAt(i - 1)) return false;
        }
        return true;
    }
    
    private static class Pair {
        String str;
        int ops;
        Pair(String str, int ops) {
            this.str = str;
            this.ops = ops;
        }
    }
}