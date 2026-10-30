# 3970. Shortest Path With At Most K Consecutive Identical Characters

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://github.com/Akshay-Devadiga-commits/Leetcode-Solutions)

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

## 💡 Approach: Dijkstra with State Expansion

To solve this problem, we can use Dijkstra's algorithm. However, the standard graph state is just the current node. Since the problem imposes a constraint on the number of consecutive identical characters, we must augment the state to include the last visited character and its current run-length count. Therefore, each state in our priority queue will be represented as a tuple: (current_weight, current_node, last_character, run_length). As we traverse edges, we update the last character and run-length. If the new character is the same as the last character, we increment the run-length (which must not exceed k). If it is different, the run-length resets to 1. Because all edge weights are positive, Dijkstra's algorithm guarantees the shortest path once we reach node n - 1.

## ⏱️ Complexity

- **Time Complexity:** `O((E * K) log(V * K))`
- **Space Complexity:** `O(V * K)`
