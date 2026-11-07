/**
 * @param {number[]} servers
 * @param {number[]} tasks
 * @return {number[]}
 */
var assignTasks = function(servers, tasks) {
    const n = servers.length;
    const m = tasks.length;
    const ans = new Array(m);

    // PriorityQueue from @datastructures-js/priority-queue (available globally on LeetCode)
    const available = new PriorityQueue({
        compare: (a, b) => a[0] !== b[0] ? a[0] - b[0] : a[1] - b[1]
    });

    for (let i = 0; i < n; i++) {
        available.enqueue([servers[i], i]);
    }

    const busy = new PriorityQueue({
        compare: (a, b) => {
            if (a[0] !== b[0]) return a[0] - b[0];
            if (a[1] !== b[1]) return a[1] - b[1];
            return a[2] - b[2];
        }
    });

    let currentTime = 0;
    let taskIdx = 0;

    while (taskIdx < m) {
        // 1. Free up servers that have finished their tasks by currentTime
        while (!busy.isEmpty() && busy.front()[0] <= currentTime) {
            const [freeTime, weight, index] = busy.dequeue();
            available.enqueue([weight, index]);
        }

        // 2. Assign tasks if there are available servers
        if (!available.isEmpty()) {
            const [weight, index] = available.dequeue();
            ans[taskIdx] = index;
            const freeTime = currentTime + tasks[taskIdx];
            busy.enqueue([freeTime, weight, index]);
            taskIdx++;
        } else {
            // 3. Advance time to next busy server's completion time
            currentTime = busy.front()[0];
        }

        // 4. Advance time to the arrival time of the next task if necessary
        if (taskIdx < m) {
            currentTime = Math.max(currentTime, taskIdx);
        }
    }

    return ans;
};