# 1882. Process Tasks Using Servers

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=8ksUsRU5SUM)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=8ksUsRU5SUM)**

📂 **YouTube Playlist:** [Medium LeetCode Problems Playlist](https://www.youtube.com/playlist?list=PLDwpLRpPSYZI)

## 💻 Solutions in 5 Major Languages

| Language | Solution File |
|---|---|
| 🐍 **Python 3** | [`solution.py`](solution.py) |
| ☕ **Java** | [`Solution.java`](Solution.java) |
| ⚡ **C++** | [`solution.cpp`](solution.cpp) |
| 🟨 **JavaScript** | [`solution.js`](solution.js) |
| 🐹 **Go** | [`solution.go`](solution.go) |

## 📝 Problem Statement

You are given two 
0-indexed
 integer arrays 
servers
 and 
tasks
 of lengths 
n
​​​​​​ and 
m
​​​​​​ respectively. 
servers[i]
 is the 
weight
 of the 
i
​​​​​​th
​​​​ server, and 
tasks[j]
 is the 
time needed
 to process the 
j
​​​​​​th
​​​​ task 
in seconds
.


Tasks are assigned to the servers using a 
task queue
. Initially, all servers are free, and the queue is 
empty
.


At second 
j
, the 
j
th
 task is 
inserted
 into the queue (starting with the 
0
th
 task being inserted at second 
0
). As long as there are free servers and the queue is not empty, the task in the front of the queue will be assigned to a free server with the 
smallest weight
, and in case of a tie, it is assigned to a free server with the 
smallest index
.


If there are no free servers and the queue is not empty, we wait until a server becomes free and immediately assign the next task. If multiple servers become free at the same time, then multiple tasks from the queue will be assigned 
in order of insertion
 following the weight and index priorities above.


A server that is assigned task 
j
 at second 
t
 will be free again at second 
t + tasks[j]
.


Build an array 
ans
​​​​ of length 
m
, where 
ans[j]
 is the 
index
 of the server the 
j
​​​​​​th
 task will be assigned to.


Return 
the array 
ans
​​​​.

## 💡 Approach: Two Priority Queues and Event Simulation

This problem can be modeled as a discrete event simulation. We need to keep track of two main things: available servers and tasks waiting to be processed. 

Available servers need to be prioritized by weight, then by index. A min-heap (priority queue) is ideal for this, storing tuples of (weight, index). 

Tasks that are waiting to be processed also need to be handled in chronological order of their arrival. Since tasks arrive at specific times (second `j` for task `j`), we can think of this as events happening at different times. 

We can maintain a 'current time' that advances as events occur. 

When a server becomes free, it's added back to the available servers priority queue. 
When a task arrives, if there's an available server, it's immediately assigned. The server's 'free time' is calculated as `current_time + task_duration`. This 'free time' is a future event. 
If no servers are available when a task arrives, the task must wait. We can use another priority queue to store these waiting tasks, prioritized by their arrival time. 

To manage the simulation efficiently, we can use two priority queues:
1. `available_servers`: Stores `(weight, index)` of free servers, ordered by weight then index.
2. `busy_servers`: Stores `(free_time, weight, index)` of servers currently processing tasks, ordered by `free_time`. This helps us know when servers become free.

We also need a way to manage tasks that arrive when all servers are busy. A simple queue or a list can store tasks that are waiting for a server to become free. However, since tasks are processed in the order they are *inserted* into the queue, and we are simulating time, we can simply iterate through tasks and assign them as time progresses.

Let's refine the simulation logic:

Initialize `available_servers` with all servers `(weight, index)`.
Initialize `busy_servers` as an empty min-heap.
Initialize `current_time = 0`.
Initialize `task_idx = 0`.
Initialize `ans` array of size `m`.

We will loop until all tasks are assigned (`task_idx == m`).

Inside the loop:
1. **Advance Time and Free Servers:** While `busy_servers` is not empty and the earliest `free_time` in `busy_servers` is less than or equal to `current_time`, pop the server `(free_time, weight, index)` from `busy_servers` and push `(weight, index)` back into `available_servers`.

2. **Assign Tasks:** While `available_servers` is not empty and `task_idx < m`:
   - Pop the best server `(weight, index)` from `available_servers`.
   - Assign `tasks[task_idx]` to this server.
   - Record `ans[task_idx] = index`.
   - Calculate the server's `free_time = current_time + tasks[task_idx]`.
   - Push `(free_time, weight, index)` into `busy_servers`.
   - Increment `task_idx`.

3. **Handle Time Advancement if No Tasks Assigned:** If no tasks were assigned in step 2 (meaning `available_servers` was empty or `task_idx == m` but we still need to wait for servers to free up), and `task_idx < m`:
   - If `busy_servers` is empty, this scenario should not happen if `m > 0` and `n > 0` as tasks will eventually be assigned. If `m=0`, the loop terminates.
   - If `busy_servers` is not empty, advance `current_time` to the `free_time` of the earliest busy server. This ensures we jump to the next event where a server becomes free.

4. **Increment Time:** If tasks were assigned in step 2, and `task_idx < m`, we need to ensure `current_time` advances. If `current_time` is still less than the arrival time of the next task (which is `task_idx`), we should advance `current_time` to `task_idx`. Otherwise, `current_time` remains the same for the next iteration if tasks were assigned at `current_time`.

A simpler way to manage `current_time` is to always consider the current time as the arrival time of the next task if no servers are available, or the time a server becomes free. 

Let's refine the time management:

Initialize `available_servers` with `(weight, index)` for all servers.
Initialize `busy_servers` as an empty min-heap.
Initialize `current_time = 0`.
Initialize `task_idx = 0`.
Initialize `ans` array of size `m`.

Loop while `task_idx < m`:

1. **Process Server Free-ups:** While `busy_servers` is not empty and `busy_servers[0][0] <= current_time`:
   - `free_time, weight, index = heapq.heappop(busy_servers)`
   - `heapq.heappush(available_servers, (weight, index))`

2. **Assign Task:** If `available_servers` is not empty:
   - `weight, index = heapq.heappop(available_servers)`
   - `ans[task_idx] = index`
   - `free_time = current_time + tasks[task_idx]`
   - `heapq.heappush(busy_servers, (free_time, weight, index))`
   - `task_idx += 1`

3. **Advance Time:** If `available_servers` is empty (and `task_idx < m`):
   - This means we must wait for a server to become free. The next event time is the `free_time` of the earliest busy server.
   - `current_time = busy_servers[0][0]`
   - Then, go back to step 1 to free up servers that become available at this new `current_time`.

4. **Increment Time (if tasks were assigned):** If tasks were assigned in step 2, and `task_idx < m`, we need to ensure `current_time` is at least the arrival time of the next task. The arrival time of task `task_idx` is `task_idx` seconds. So, `current_time = max(current_time, task_idx)`. This ensures that if we assigned a task at time `t` and the next task arrives at `t+1`, we correctly consider time `t+1` for the next assignment.

This refined logic seems more robust. The key insight is to manage events (task arrivals and server free-ups) chronologically using priority queues.

Let's consider the constraints: `n, m <= 2 * 10^5`. This means an O(m log n) or O(m log m) solution is likely required. Using priority queues will give us this complexity.

Edge Cases:
- Empty `tasks` array: The loop `while task_idx < m` will not execute, and an empty `ans` will be returned (or `ans` of size 0 if `m=0`).
- All servers have the same weight: The tie-breaking rule (smallest index) is handled by the tuple comparison in the priority queue.
- All tasks have the same duration: This doesn't fundamentally change the logic, just the `free_time` calculations.
- `n=1`: Only one server is available, tasks are processed sequentially on it.
- `m=1`: A single task is assigned to the best available server.
- Large number of tasks and few servers: `available_servers` will often be empty, forcing `current_time` to advance based on `busy_servers`.
- Large number of servers and few tasks: All tasks will likely be assigned quickly to the best servers.

Complexity Analysis:
- Time Complexity: Each task is pushed into `busy_servers` once and popped once. Each server is pushed into `available_servers` once initially, and potentially multiple times if it becomes free and is reassigned. However, each server is only *assigned* a task once per time step it's available. The total number of pushes and pops across both heaps is proportional to `m` (for tasks being assigned) and `n` (for servers being initially added and potentially freed). Each heap operation (push/pop) takes O(log n) time for `available_servers` and O(log n) for `busy_servers` (since the size of `busy_servers` is at most `n`). Therefore, the total time complexity is O(m log n).
- Space Complexity: `available_servers` stores up to `n` elements. `busy_servers` stores up to `n` elements. The `ans` array stores `m` elements. Thus, the space complexity is O(n + m). Given the constraints, this is acceptable.


## ⏱️ Complexity

- **Time Complexity:** `O(m log n)`
- **Space Complexity:** `O(n + m)`
