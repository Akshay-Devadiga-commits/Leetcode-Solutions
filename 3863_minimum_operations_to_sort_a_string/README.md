# 3863. Minimum Operations to Sort a String

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=E0hCAq_pvso)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=E0hCAq_pvso)**

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

## 💡 Approach: Breadth-First Search

Since the string length is very small (up to 10), we can treat this problem as finding the shortest path on an unweighted graph where each node is a string and edges represent valid substring sorting operations. We can use Breadth-First Search (BFS) to find the minimum number of operations to reach a sorted state. If the string is already sorted, 0 operations are needed. If BFS completes without reaching a sorted state, it is impossible, so we return -1.

## ⏱️ Complexity

- **Time Complexity:** `O(n^3 * n! / prod(freq_i!))`
- **Space Complexity:** `O(n! / prod(freq_i!))`
