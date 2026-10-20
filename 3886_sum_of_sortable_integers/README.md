# 3886. Sum of Sortable Integers

![Difficulty](https://img.shields.io/badge/Difficulty-Hard-blue?style=flat-square)

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

## 💡 Approach: Divisor Check with Subarray Rotation Simulation

Given the constraints where n <= 10, we can directly test all divisors k of n. For each valid divisor k, we partition the array into consecutive subarrays of length k. For a subarray to be sortable into a non-decreasing sequence through cyclic rotations, it must be possible to cyclically shift it such that it forms a contiguous non-decreasing sequence. Furthermore, when we concatenate all these sorted subarrays from left to right, the entire array must end up in non-decreasing order (meaning the maximum of the i-th subarray is less than or equal to the minimum of the (i+1)-th subarray). We simulate this condition for each divisor k that divides n, and sum up the valid k values.

## ⏱️ Complexity

- **Time Complexity:** `O(n * d(n)) where n is the length of the array and d(n) is the number of divisors of n. Given n <= 10, this is extremely fast and well within time limits.`
- **Space Complexity:** `O(n) to store subarrays and flattened arrays for each divisor.`
