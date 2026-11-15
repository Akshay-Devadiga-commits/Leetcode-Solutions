# 1524. Number of Sub-arrays With Odd Sum

![Difficulty](https://img.shields.io/badge/Difficulty-Medium-blue?style=flat-square)

### Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=TdjVLQSVAtg)

**[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=TdjVLQSVAtg)**

**YouTube Playlist:** [Medium LeetCode Problems Playlist](https://www.youtube.com/playlist?list=PLDwpLRpPSYZI)

## Solutions in 5 Major Languages

| Language | Solution File |
|---|---|
| Python 3 | [`solution.py`](solution.py) |
| Java | [`Solution.java`](Solution.java) |
| C++ | [`solution.cpp`](solution.cpp) |
| JavaScript | [`solution.js`](solution.js) |
| Go | [`solution.go`](solution.go) |

## Problem Statement

Given an array of integers  arr , return  the number of subarrays with an  odd  sum .


Since the answer can be very large, return it modulo  10^9 + 7 .

## Approach: Prefix Sum with Odd-Even Counting

Instead of computing the sum of every possible subarray from scratch, we can leverage prefix sums to speed things up. The key insight is that the sum of any subarray can be determined by the difference between two prefix sums. To understand this, consider the prefix sum at an index, which represents the cumulative sum of elements from the start of the array up to that index. The sum of a subarray starting at index `i` and ending at `j` is simply the difference between the prefix sum at `j` and the prefix sum at `i - 1`. This means that whether a subarray sum is odd or even depends only on the 

## Complexity

- **Time Complexity:** `O(n)`
- **Space Complexity:** `O(1)`
