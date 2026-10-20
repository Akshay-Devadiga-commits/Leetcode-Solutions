# 3886. Sum of Sortable Integers

![Difficulty](https://img.shields.io/badge/Difficulty-Hard-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution
Watch the complete step-by-step visual explanation and algorithmic intuition on YouTube:
👉 **[Watch on YouTube](https://www.youtube.com/watch?v=EV45m2VoD84)**

## 💻 Solutions in 5 Major Languages

| Language | Solution File |
|---|---|
| 🐍 **Python 3** | [`solution.py`](solution.py) |
| ☕ **Java** | [`Solution.java`](Solution.java) |
| ⚡ **C++** | [`solution.cpp`](solution.cpp) |
| 🟨 **JavaScript** | [`solution.js`](solution.js) |
| 🐹 **Go** | [`solution.go`](solution.go) |

## 📝 Problem Statement

You are given an integer array 
nums
 of length 
n
.


An integer 
k
 is called 
sortable
 if 
k
 
divides
 
n
 and you can sort 
nums
 in 
non-decreasing
 order by sequentially performing the following operations:




Partition 
nums
 into 
consecutive 
subarrays
 of length 
k
.


Cyclically rotate each subarray independently
 any number of times to the left or to the right.




Return an integer denoting the sum of all possible sortable integers 
k
.

## 💡 Approach: Divisor Check with Subarray Rotation Matching

For a given divisor k of n, we can partition nums into n/k subarrays of length k. Each subarray can be independently rotated cyclically. For the entire array to be sortable, each subarray must be formed by a contiguous chunk of the final sorted array, and its elements must be a cyclic shift of that chunk. Furthermore, the relative order of the sorted chunks must be preserved. We can verify this by sorting the entire array to get the target state, dividing it into chunks of length k, and checking if each chunk in nums is a cyclic shift of the corresponding target chunk. Since n is very small (up to 10 based on constraints), we can iterate over all divisors k of n, and for each divisor, perform the verification efficiently.

## ⏱️ Complexity

- **Time Complexity:** `O(n^2 log n)`
- **Space Complexity:** `O(n)`
