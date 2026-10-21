# 3863. Minimum Operations to Sort a String

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=SVUMHaYW7oE)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=SVUMHaYW7oE)**

## 💻 Solutions in 5 Major Languages

| Language | Solution File |
|---|---|
| 🐍 **Python 3** | [`solution.py`](solution.py) |
| ☕ **Java** | [`Solution.java`](Solution.java) |
| ⚡ **C++** | [`solution.cpp`](solution.cpp) |
| 🟨 **JavaScript** | [`solution.js`](solution.js) |
| 🐹 **Go** | [`solution.go`](solution.go) |

## 📝 Problem Statement

You are given a string 
s
 consisting of lowercase English letters.


In one operation, you can select any 
substring
 of 
s
 that is 
not
 the entire string and 
sort
 it in 
non-descending alphabetical
 order.


Return the 
minimum
 number of operations required to make 
s
 sorted in 
non-descending
 order. If it is not possible, return -1.

## 💡 Approach: Breadth-First Search (BFS)

Since the string length is extremely small (1 <= length <= 10), we can model this problem as finding the shortest path in an unweighted graph where nodes are strings and edges are valid substring sorting operations. We can use Breadth-First Search (BFS) to find the minimum number of operations to reach a sorted string. If the string is already sorted, 0 operations are needed. If BFS queue becomes empty without reaching a sorted string, it is impossible, so we return -1.

## ⏱️ Complexity

- **Time Complexity:** `O(N! * N^3)`
- **Space Complexity:** `O(N!)`
