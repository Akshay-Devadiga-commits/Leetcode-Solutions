import java.util.*;

class Solution {
    public int minOperations(String s) {
        char[] sortedChars = s.toCharArray();
        Arrays.sort(sortedChars);
        String target = new String(sortedChars);
        if (s.equals(target)) return 0;

        int n = s.length();
        Queue<Pair> queue = new LinkedList<>();
        Set<String> visited = new HashSet<>();

        queue.offer(new Pair(s, 0));
        visited.add(s);

        while (!queue.isEmpty()) {
            Pair currPair = queue.poll();
            String curr = currPair.s;
            int ops = currPair.ops;

            for (int i = 0; i < n; i++) {
                for (int j = i + 2; j <= n; j++) {
                    if (i == 0 && j == n) continue;

                    char[] sub = curr.substring(i, j).toCharArray();
                    Arrays.sort(sub);
                    String nextS = curr.substring(0, i) + new String(sub) + curr.substring(j);

                    char[] nextSorted = nextS.toCharArray();
                    Arrays.sort(nextSorted);
                    if (nextS.equals(new String(nextSorted))) {
                        return ops + 1;
                    }

                    if (!visited.contains(nextS)) {
                        visited.add(nextS);
                        queue.offer(new Pair(nextS, ops + 1));
                    }
                }
            }
        }

        return -1;
    }

    private static class Pair {
        String s;
        int ops;
        Pair(String s, int ops) {
            this.s = s;
            this.ops = ops;
        }
    }
}