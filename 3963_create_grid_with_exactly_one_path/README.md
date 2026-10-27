# 3963. Create Grid With Exactly One Path

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=jbRvlQBy_Ug)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=jbRvlQBy_Ug)**

## 💻 Solutions in 5 Major Languages

| Language | Solution File |
|---|---|
| 🐍 **Python 3** | [`solution.py`](solution.py) |
| ☕ **Java** | [`Solution.java`](Solution.java) |
| ⚡ **C++** | [`solution.cpp`](solution.cpp) |
| 🟨 **JavaScript** | [`solution.js`](solution.js) |
| 🐹 **Go** | [`solution.go`](solution.go) |

## 📝 Problem Statement

You are given two integers 
m
 and 
n
, representing the number of rows and columns of a grid.


Construct 
any
 
m x n
 grid consisting only of the characters 
'.'
 and 
'#'
, where:




'.'
 represents a free cell.


'#'
 represents an obstacle cell.




A 
valid path
 is a sequence of free cells that:




Starts at the top-left cell 
(0, 0)
.


Ends at the bottom-right cell 
(m - 1, n - 1)
.


Moves only:
	


Right, from 
(i, j)
 to 
(i, j + 1)
, or


Down, from 
(i, j)
 to 
(i + 1, j)
.








Return any grid such that there is 
exactly one valid path
 from the top-left cell to the bottom-right cell.

## 💡 Approach: Zigzag Snake Path with Obstacles

To ensure there is exactly one valid path from (0,0) to (m-1, n-1), we can construct a zigzag pattern of free cells ('.') where every other cell or alternating rows/columns are blocked by obstacles ('#'). A simple and effective strategy is to make the entire grid obstacles ('#') except for a specific sequence of cells forming a winding path or by placing obstacles such that only one continuous corridor remains. Alternatively, we can make the first row all '.' and then block everything else except a single vertical drop at the very end, or use a snake-like pattern. A very robust approach is to fill the grid with '#' and carve out a path, or fill it with '.' and place '#' strategically. For instance, we can alternate '.' and '#' in rows such that the path is forced to zigzag, guaranteeing a unique path.

## ⏱️ Complexity

- **Time Complexity:** `O(m * n)`
- **Space Complexity:** `O(m * n)`
