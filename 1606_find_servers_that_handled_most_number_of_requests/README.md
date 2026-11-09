# 1606. Find Servers That Handled Most Number of Requests

![Difficulty](https://img.shields.io/badge/Difficulty-Hard-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=UzaDVyBqJ00)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=UzaDVyBqJ00)**

📂 **YouTube Playlist:** [Hard LeetCode Problems Playlist](https://www.youtube.com/playlist?list=PLWnvKaA3LMvs)

## 💻 Solutions in 5 Major Languages

| Language | Solution File |
|---|---|
| 🐍 **Python 3** | [`solution.py`](solution.py) |
| ☕ **Java** | [`Solution.java`](Solution.java) |
| ⚡ **C++** | [`solution.cpp`](solution.cpp) |
| 🟨 **JavaScript** | [`solution.js`](solution.js) |
| 🐹 **Go** | [`solution.go`](solution.go) |

## 📝 Problem Statement

You have 
k
 servers numbered from 
0
 to 
k-1
 that are being used to handle multiple requests simultaneously. Each server has infinite computational capacity but 
cannot handle more than one request at a time
. The requests are assigned to servers according to a specific algorithm:




The 
i
th
 (0-indexed) request arrives.


If all servers are busy, the request is dropped (not handled at all).


If the 
(i % k)
th
 server is available, assign the request to that server.


Otherwise, assign the request to the next available server (wrapping around the list of servers and starting from 0 if necessary). For example, if the 
i
th
 server is busy, try to assign the request to the 
(i+1)
th
 server, then the 
(i+2)
th
 server, and so on.




You are given a 
strictly increasing
 array 
arrival
 of positive integers, where 
arrival[i]
 represents the arrival time of the 
i
th
 request, and another array 
load
, where 
load[i]
 represents the load of the 
i
th
 request (the time it takes to complete). Your goal is to find the 
busiest server(s)
. A server is considered 
busiest
 if it handled the most number of requests successfully among all the servers.


Return 
a list containing the IDs (0-indexed) of the 
busiest server(s)
. You may return the IDs in any order.

## 💡 Approach: Priority Queue and Sorted Sets

The problem requires us to simulate the request assignment process and track the number of requests handled by each server. Since requests arrive in strictly increasing order of time, we can process them one by one. For each request `i` arriving at `arrival[i]` with load `load[i]`, we need to find an available server. The target server is `target_server = i % k`. If `target_server` is available, we assign the request to it. If not, we search for the next available server in a circular manner. A server is available if its current task (if any) finishes before or at the arrival time of the current request. 

To efficiently find the next available server, we can maintain two data structures: 
1. `available_servers`: A sorted set (or a min-heap) storing the IDs of servers that are currently free. Initially, all servers are free.
2. `busy_servers`: A min-heap storing tuples `(finish_time, server_id)` for servers that are currently busy. The heap is ordered by `finish_time`.

When a request `i` arrives at `arrival[i]`: 
1. First, we check the `busy_servers` heap. Any server whose `finish_time` is less than or equal to `arrival[i]` becomes available. We pop these servers from `busy_servers` and add their `server_id` back to `available_servers`.
2. We then determine the `target_server = i % k`. 
3. If `target_server` is in `available_servers`, we assign the request to it. We remove `target_server` from `available_servers`, add `(arrival[i] + load[i], target_server)` to `busy_servers`, and increment the request count for `target_server`.
4. If `target_server` is not available, we need to find the next available server. Since `available_servers` is sorted, we can efficiently find the smallest server ID greater than or equal to `target_server`. If such a server exists, we use it. If not, we wrap around and find the smallest server ID in `available_servers` (which will be the smallest available server ID overall). If `available_servers` is empty, the request is dropped.
5. If we find an available server (either `target_server` or a subsequent one), we assign the request, update `available_servers` and `busy_servers`, and increment the request count for that server.

After processing all requests, we find the maximum request count and return the IDs of servers with that count.

For the constraints given (k <= 10, arrival/load length <= 10), a simpler approach using just lists and linear scans for available servers would also pass. However, the priority queue and sorted set approach is more scalable and demonstrates a robust solution for larger constraints. Given the small constraints, we can simplify the search for the next available server. Instead of a sorted set for `available_servers`, we can use a min-heap. When a server becomes free, we add it to the `available_servers` heap. When we need to find an available server for request `i % k`, we first check if `i % k` is available. If not, we look for the smallest server ID greater than `i % k` in the `available_servers` heap. If no such server exists, we look for the smallest server ID in the `available_servers` heap (effectively wrapping around). This can be done by peeking at the top of the `available_servers` heap. If the top element is less than `i % k`, we pop it and check again until we find a server >= `i % k` or the heap is empty. If the heap is empty, the request is dropped. If we find a server, we use it, update its finish time, and add it to the `busy_servers` heap. 

To handle the circular search efficiently, we can use two heaps: `available_servers` (min-heap of server IDs) and `busy_servers` (min-heap of `(finish_time, server_id)`). 

For each request `i` at `arrival[i]` with load `load[i]`: 
1. Free up servers from `busy_servers` whose `finish_time <= arrival[i]` and add them back to `available_servers`.
2. Calculate `target_server_idx = i % k`.
3. If `available_servers` is empty, drop the request.
4. Find the smallest server ID `s` in `available_servers` such that `s >= target_server_idx`. This can be done by repeatedly popping from `available_servers` until the condition is met or the heap is empty. If a server is popped and doesn't meet the condition, it needs to be temporarily stored and re-added later if not used.
5. If no such server `s` is found (i.e., all available servers have IDs less than `target_server_idx`), then wrap around and pick the smallest server ID from `available_servers` (which is the top of the heap if it's not empty). 
6. If a server `assigned_server_id` is found: 
   - Increment `request_counts[assigned_server_id]`.
   - Add `(arrival[i] + load[i], assigned_server_id)` to `busy_servers`.
   - If `assigned_server_id` was temporarily popped from `available_servers` and not used, re-add it. If it was used, it's already removed.

This approach uses heaps for efficient management of available and busy servers. The key insight is to use heaps to quickly find the next available server and to manage server availability based on finish times.

## ⏱️ Complexity

- **Time Complexity:** `O(N * log K), where N is the number of requests and K is the number of servers. For each request, we perform heap operations (push, pop, peek) on `available_servers` and `busy_servers`. In the worst case, we might iterate through all K servers in `available_servers` to find the correct one. Each heap operation takes O(log K) time. Thus, the total time complexity is O(N * K * log K) if we naively search, but with the heap and temporary storage, it becomes O(N * log K) because each server ID is pushed and popped from `available_servers` at most once per request cycle (when it becomes free and then when it's assigned). The `busy_servers` heap also has at most K elements. So, each request processing involves a few heap operations, leading to O(log K) per request.`
- **Space Complexity:** `O(K), where K is the number of servers. This is for storing the `available_servers` heap, `busy_servers` heap, and `server_request_counts` array, all of which are proportional to K.`
