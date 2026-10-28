# 3882. Minimum XOR Path in a Grid

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

You are given a 2D integer array 
grid
 of size 
m * n
.


You start at the 
top-left
 cell 
(0, 0)
 and want to reach the 
bottom-right
 cell 
(m - 1, n - 1)
.


At each step, you 
may
 move either 
right or down
.


The 
cost
 of a path is defined as the 
bitwise XOR
 of all the values in the cells along that path, 
including
 the start and end cells.


Return the 
minimum
 possible XOR value among all valid paths from 
(0, 0)
 to 
(m - 1, n - 1)
.

## 💡 Approach: Dynamic Programming with Set of Possible XOR Values

Since grid elements are constrained between 0 and 1023 (meaning the maximum possible XOR sum is less than 1024), we can maintain the set of all possible XOR sums at each cell (r, c). When moving from top-left to bottom-right, the set of XOR sums at (r, c) is formed by taking the XOR of grid[r][c] with every possible XOR sum from the cell above (r-1, c) and the cell to the left (r, c-1). Finally, we return the minimum value in the set at the bottom-right cell.

## ⏱️ Complexity

- **Time Complexity:** `O(m * n * MAX_XOR)`
- **Space Complexity:** `O(n * MAX_XOR)`
