# 3693. Climbing Stairs II

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=O4wZcValuRk)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=O4wZcValuRk)**

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

You are climbing a staircase with 
n + 1
 steps, numbered from 0 to 
n
.


You are also given a 
1-indexed
 integer array 
costs
 of length 
n
, where 
costs[i]
 is the cost of step 
i
.


From step 
i
, you can jump 
only
 to step 
i + 1
, 
i + 2
, or 
i + 3
. The cost of jumping from step 
i
 to step 
j
 is defined as: 
costs[j] + (j - i)
2


You start from step 0 with 
cost = 0
.


Return the 
minimum
 total cost to reach step 
n
.

## 💡 Approach: Dynamic Programming

We can solve this problem using Dynamic Programming. Let `dp[i]` represent the minimum cost to reach step `i`. Starting at step 0 with `dp[0] = 0`, to reach any step `i` (for `1 <= i <= n`), the previous step `j` must be one of `i - 1`, `i - 2`, or `i - 3` (provided `j >= 0`). The cost of jumping from `j` to `i` is `costs[i - 1] + (i - j)^2`. Therefore, the recurrence relation is `dp[i] = costs[i - 1] + min(dp[i - 1] + 1, dp[i - 2] + 4, dp[i - 3] + 9)` (checking bounds for `i - 2` and `i - 3`). Computing this iteratively from `1` to `n` takes linear time $O(n)$ and guarantees the global optimal solution.

## ⏱️ Complexity

- **Time Complexity:** `O(n)`
- **Space Complexity:** `O(n)`
