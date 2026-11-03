# 3736. Minimum Moves to Equal Array Elements III

![Difficulty](https://img.shields.io/badge/Difficulty-Easy-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=hOXrkdCbgD8)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=hOXrkdCbgD8)**

📂 **YouTube Playlist:** [Easy LeetCode Problems Playlist](https://www.youtube.com/playlist?list=PLO_wgoXxiqtY)

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
.


In one move, you may 
increase
 the value of any single element 
nums[i]
 by 1.


Return the 
minimum total
 number of 
moves
 required so that all elements in 
nums
 become 
equal
.

## 💡 Approach: Median as Target

The problem asks for the minimum number of moves to make all elements in an array equal, where a move consists of increasing a single element by 1. Let the target value be 'T'. For each element `nums[i]`, the number of moves required to make it equal to 'T' is `T - nums[i]`, assuming `T >= nums[i]`. The total number of moves is the sum of `(T - nums[i])` for all `i`. This can be rewritten as `n * T - sum(nums)`, where `n` is the length of the array. To minimize this sum, we need to minimize 'T'. Since we can only increase elements, the target value 'T' must be greater than or equal to the maximum element in the original array. However, the problem statement implies we can only *increase* elements. This means the final equal value must be at least the maximum value in the original array. If we choose a target value `T`, the total moves would be `sum(T - nums[i])` for all `i`. This is equivalent to `n*T - sum(nums)`. To minimize this, we need to minimize `T`. The smallest possible `T` is the maximum element in the array.  Wait, the problem statement says 'increase the value of any single element nums[i] by 1'. This implies that the final equal value can be *any* value greater than or equal to the maximum element.  Let's re-read. 'Return the minimum total number of moves required so that all elements in nums become equal.'  This is a classic problem. If we can only increase elements, the final value must be at least the maximum element.  However, the problem is actually simpler than that. The problem statement is slightly misleading. The standard interpretation of this problem (and the one that matches the examples) is that a move consists of *incrementing* an element. If we can only increment, then the final value must be at least the maximum element.  Let's consider the examples. For `[2, 1, 3]`, the output is 3. If the target is 3, moves are (3-2) + (3-1) + (3-3) = 1 + 2 + 0 = 3. If the target is 4, moves are (4-2) + (4-1) + (4-3) = 2 + 3 + 1 = 6.  The minimum is achieved when the target is the maximum element.  This seems too simple.  Let's re-read the problem statement carefully. 'In one move, you may increase the value of any single element nums[i] by 1.'  This is indeed the constraint.  The examples confirm that the target value is the maximum element.  This is because any target value `T` greater than the maximum element `max(nums)` would require `T - max(nums)` additional moves for *every* element, thus increasing the total moves.  Therefore, the optimal target value is `max(nums)`. The total moves will be `sum(max(nums) - nums[i])` for all `i`. This can be calculated as `n * max(nums) - sum(nums)`.  This is correct.  The problem is equivalent to finding a target value `T` that minimizes `sum(|nums[i] - T|)`. This sum is minimized when `T` is the median of the array.  However, the constraint here is that we can *only increase* elements. This changes the problem significantly. If we can only increase elements, the final equal value must be at least the maximum element in the original array. Let the final equal value be `X`. Then for each `nums[i]`, we need `X - nums[i]` moves. The total moves are `sum(X - nums[i]) = n*X - sum(nums)`. To minimize this, we need to minimize `X`. The smallest possible value for `X` is `max(nums)`.  So, the minimum moves is `sum(max(nums) - nums[i])`.  This is equivalent to `n * max(nums) - sum(nums)`.  This is the correct interpretation for this specific problem statement. The examples provided align with this interpretation. The problem is NOT asking to find a value that minimizes the sum of absolute differences (which would be the median). It's asking to find the minimum moves to make all elements equal by *only increasing* them. Thus, the target value must be at least the maximum element. The minimum moves are achieved when the target is exactly the maximum element.  The total moves is the sum of differences between the maximum element and each element.  This can be computed efficiently by summing up all elements and subtracting this sum from `n * max(nums)`.  Alternatively, and perhaps more intuitively, we can sort the array and then sum the differences between the largest element and all other elements.  The sum of `max(nums) - nums[i]` is equivalent to `sum(nums) - n * min(nums)` if we were allowed to *decrease* elements.  Since we can only *increase*, the target must be at least `max(nums)`.  The minimum moves are achieved when the target is `max(nums)`.  The total moves is `sum(max(nums) - nums[i])`.  This is equivalent to `n * max(nums) - sum(nums)`.  Let's re-verify with the examples. `nums = [2, 1, 3]`. `max(nums) = 3`. `sum(nums) = 6`. `n = 3`. Moves = `3 * 3 - 6 = 9 - 6 = 3`. Correct. `nums = [4, 4, 5]`. `max(nums) = 5`. `sum(nums) = 13`. `n = 3`. Moves = `3 * 5 - 13 = 15 - 13 = 2`. Correct.  The key insight is that since we can only increase elements, the final target value must be at least the maximum element in the original array. To minimize the total moves, we should choose the smallest possible target value, which is the maximum element itself. The total moves is then the sum of the differences between this maximum element and every other element. This can be calculated as `n * max(nums) - sum(nums)`.  This approach is optimal because any other target value greater than `max(nums)` would increase the total number of moves by `(target - max(nums))` for each of the `n` elements, leading to a strictly larger sum.  The brute force would be to try every possible target value from `max(nums)` upwards, calculate the moves for each, and find the minimum. However, we've proven that `max(nums)` is the optimal target.  Therefore, a single pass to find the sum and maximum is sufficient.

## ⏱️ Complexity

- **Time Complexity:** `O(n)`
- **Space Complexity:** `O(1)`
