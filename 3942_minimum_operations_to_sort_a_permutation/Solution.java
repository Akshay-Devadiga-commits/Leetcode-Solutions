import java.util.*;

class Solution {
    public int minOperations(int[] nums) {
        int n = nums.length;
        int[] targetArr = new int[n];
        for (int i = 0; i < n; i++) targetArr[i] = i;
        
        String target = Arrays.toString(targetArr);
        String start = Arrays.toString(nums);
        
        if (start.equals(target)) return 0;
        
        Queue<State> queue = new LinkedList<>();
        Set<String> visited = new HashSet<>();
        
        queue.offer(new State(nums, 0));
        visited.add(start);
        
        while (!queue.isEmpty()) {
            State currState = queue.poll();
            int[] curr = currState.arr;
            int dist = currState.dist;
            
            int[] nextRev = new int[n];
            for (int i = 0; i < n; i++) {
                nextRev[i] = curr[n - 1 - i];
            }
            String revStr = Arrays.toString(nextRev);
            if (revStr.equals(target)) return dist + 1;
            if (!visited.contains(revStr)) {
                visited.add(revStr);
                queue.offer(new State(nextRev, dist + 1));
            }
            
            int[] nextRot = new int[n];
            for (int i = 0; i < n - 1; i++) {
                nextRot[i] = curr[i + 1];
            }
            nextRot[n - 1] = curr[0];
            String rotStr = Arrays.toString(nextRot);
            if (rotStr.equals(target)) return dist + 1;
            if (!visited.contains(rotStr)) {
                visited.add(rotStr);
                queue.offer(new State(nextRot, dist + 1));
            }
        }
        return -1;
    }
    
    private static class State {
        int[] arr;
        int dist;
        State(int[] arr, int dist) {
            this.arr = arr;
            this.dist = dist;
        }
    }
}