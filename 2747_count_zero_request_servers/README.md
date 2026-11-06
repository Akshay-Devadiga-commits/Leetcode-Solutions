# 2747. Count Zero Request Servers

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=fWCZjZ_3LmU)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=fWCZjZ_3LmU)**

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

You are given an integer 
n
 denoting the total number of servers and a 
2D
 
0-indexed 
integer array 
logs
, where 
logs[i] = [server_id, time]
 denotes that the server with id 
server_id
 received a request at time 
time
.


You are also given an integer 
x
 and a 
0-indexed
 integer array 
queries
.


Return 
a 
0-indexed
 integer array
 
arr
 
of length
 
queries.length
 
where
 
arr[i]
 
represents the number of servers that 
did not receive
 any requests during the time interval
 
[queries[i] - x, queries[i]]
.


Note that the time intervals are inclusive.

## 💡 Approach: Array Approach

This solution uses the Array Approach with time complexity O(textttsort(n) + textttsort(q) + qlog n) and auxiliary space O(n + q). It achieves optimal performance by taking advantage of Array, Hash Table, Sliding Window, Sorting.

## ⏱️ Complexity

- **Time Complexity:** `O(textttsort(n) + textttsort(q) + qlog n)`
- **Space Complexity:** `O(n + q)`
