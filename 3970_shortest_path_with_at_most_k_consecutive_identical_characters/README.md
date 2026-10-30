# 3970. Shortest Path With At Most K Consecutive Identical Characters

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=OufLfaqKs3o)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=OufLfaqKs3o)**

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
 representing the number of nodes in a 
directed weighted
 graph, numbered from 0 to 
n - 1
. This is represented by a 2D integer array 
edges
, where 
edges[i] = [u
i
, v
i
, w
i
]
 represents a directed edge from node 
u
i
 to node 
v
i
 with weight 
w
i
.


You are also given a string 
labels
 of length 
n
, where 
labels[i]
 is the character assigned to node 
i
, and an integer 
k
.


Return the 
minimum
 
total
 edge weight of a path from node 0 to node 
n - 1
 such that the concatenation of the labels of the nodes along the path contains 
at most
 
k
 
consecutive
 
identical
 characters. If no valid path exists, return -1.

## 💡 Approach: Dijkstra with Extended State

To solve this problem, we can use Dijkstra's algorithm. Since the validity of a path depends not only on the current node but also on the last character visited and how many times it has appeared consecutively, we must include these attributes in our state. Thus, the state in our priority queue will be (current_distance, current_node, last_char, count). We can then run Dijkstra's algorithm to find the minimum distance to node n-1 with any valid last_char and count.

## ⏱️ Complexity

- **Time Complexity:** `O(E * k * log(V * k))`
- **Space Complexity:** `O(V * k)`
