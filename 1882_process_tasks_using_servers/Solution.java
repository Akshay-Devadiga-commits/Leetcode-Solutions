import java.util.PriorityQueue;

class Solution {
    public int[] assignTasks(int[] servers, int[] tasks) {
        int n = servers.length;
        int m = tasks.length;
        int[] ans = new int[m];

        // Priority queue for available servers: [weight, index]
        PriorityQueue<int[]> available = new PriorityQueue<>((a, b) -> {
            if (a[0] != b[0]) return Integer.compare(a[0], b[0]);
            return Integer.compare(a[1], b[1]);
        });

        for (int i = 0; i < n; i++) {
            available.offer(new int[]{servers[i], i});
        }

        // Priority queue for busy servers: [free_time (long), weight (int), index (int)]
        PriorityQueue<long[]> busy = new PriorityQueue<>((a, b) -> {
            if (a[0] != b[0]) return Long.compare(a[0], b[0]);
            if (a[1] != b[1]) return Long.compare(a[1], b[1]);
            return Long.compare(a[2], b[2]);
        });

        long currentTime = 0;
        int taskIdx = 0;

        while (taskIdx < m) {
            // 1. Free up servers that have finished their tasks
            while (!busy.isEmpty() && busy.peek()[0] <= currentTime) {
                long[] server = busy.poll();
                available.offer(new int[]{(int) server[1], (int) server[2]});
            }

            // 2. Assign task
            if (!available.isEmpty()) {
                int[] server = available.poll();
                int weight = server[0];
                int index = server[1];
                ans[taskIdx] = index;

                long freeTime = currentTime + tasks[taskIdx];
                busy.offer(new long[]{freeTime, weight, index});
                taskIdx++;
            } else {
                // 3. Advance time to next free server
                currentTime = busy.peek()[0];
            }

            // 4. Ensure currentTime is at least taskIdx
            if (taskIdx < m) {
                currentTime = Math.max(currentTime, (long) taskIdx);
            }
        }

        return ans;
    }
}