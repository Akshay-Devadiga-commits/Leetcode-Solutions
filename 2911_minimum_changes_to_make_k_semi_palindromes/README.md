# 2911. Minimum Changes to Make K Semi-palindromes

![Difficulty](https://img.shields.io/badge/Difficulty-Hard-blue?style=flat-square)

### 📺 Video Walkthrough & Visual Execution

[![Video Walkthrough & Visual Execution](thumbnail.png)](https://www.youtube.com/watch?v=RrQhl-v-M2o)

👉 **[Watch the complete step-by-step visual explanation on YouTube](https://www.youtube.com/watch?v=RrQhl-v-M2o)**

📂 **YouTube Playlist:** [Hard LeetCode Problems Playlist](https://www.youtube.com/playlist?list=PLWnvKaA3LMvs)

## 💻 Solutions in 5 Major Languages

| Language | Solution File |
|---|---|
| 🐍 **Python 3** | [`solution.py`](solution.py) |
| ☕ **Java** | [`Solution.java`](Solution.java) |
| ⚡ **C++** | [`solution.cpp`](solution.cpp) |
| 🟨 **JavaScript** | [`solution.js`](solution.js) |
| 🐹 **Go** | [`solution.go`](solution.go) |

## 📝 Problem Statement

Given a string 
s
 and an integer 
k
, partition 
s
 into 
k
 
substrings
 such that the letter changes needed to make each substring a 
semi-palindrome
 are minimized.


Return the 
minimum
 number of letter changes
 required
.


A 
semi-palindrome
 is a special type of string that can be divided into 
palindromes
 based on a repeating pattern. To check if a string is a semi-palindrome:​




Choose a positive divisor 
d
 of the string's length. 
d
 can range from 
1
 up to, but not including, the string's length. For a string of length 
1
, it does not have a valid divisor as per this definition, since the only divisor is its length, which is not allowed.


For a given divisor 
d
, divide the string into groups where each group contains characters from the string that follow a repeating pattern of length 
d
. Specifically, the first group consists of characters at positions 
1
, 
1 + d
, 
1 + 2d
, and so on; the second group includes characters at positions 
2
, 
2 + d
, 
2 + 2d
, etc.


The string is considered a semi-palindrome if each of these groups forms a palindrome.




Consider the string 
"abcabc"
:




The length of 
"abcabc"
 is 
6
. Valid divisors are 
1
, 
2
, and 
3
.


For 
d = 1
: The entire string 
"abcabc"
 forms one group. Not a palindrome.


For 
d = 2
:
	


Group 1 (positions 
1, 3, 5
): 
"acb"


Group 2 (positions 
2, 4, 6
): 
"bac"


Neither group forms a palindrome.






For 
d = 3
:
	


Group 1 (positions 
1, 4
): 
"aa"


Group 2 (positions 
2, 5
): 
"bb"


Group 3 (positions 
3, 6
): 
"cc"


All groups form palindromes. Therefore, 
"abcabc"
 is a semi-palindrome.

## 💡 Approach: Dynamic Programming with Precomputation

The problem asks us to partition a string `s` into `k` substrings with the minimum total changes to make each substring a semi-palindrome. This problem has optimal substructure and overlapping subproblems, making it suitable for dynamic programming.

First, let's define what it means to make a substring a semi-palindrome with minimum changes. For a substring `sub` of length `L`, we need to find a divisor `d` (1 <= d < L) such that the cost to make each of the `d` groups palindromic is minimized. The cost for a single group (e.g., characters at indices `i`, `i+d`, `i+2d`, ...) is the number of characters that need to be changed to make that group a palindrome. This can be calculated by comparing characters symmetrically from the ends of the group. For example, if a group is `c1, c2, ..., cm`, the cost is the number of pairs `(c_j, c_{m-j+1})` where `c_j != c_{m-j+1}`.

We can precompute the minimum cost to make any substring `s[i:j+1]` a semi-palindrome. Let `cost[i][j]` be this minimum cost. To calculate `cost[i][j]`, we iterate through all possible divisors `d` of the substring length `L = j - i + 1`, where `1 <= d < L`. For each `d`, we calculate the cost of making the `d` groups palindromic. The minimum among these costs for all valid `d` will be `cost[i][j]`.

Once we have the `cost` table, we can use dynamic programming to solve the main problem. Let `dp[i][p]` be the minimum changes required to partition the prefix `s[0:i]` into `p` semi-palindromic substrings. Our goal is to find `dp[n][k]`, where `n` is the length of `s`.

The recurrence relation is: `dp[i][p] = min(dp[j][p-1] + cost[j][i-1])` for `0 <= j < i`. This means we consider all possible split points `j` for the `p`-th substring, which would be `s[j:i]`. The base cases are `dp[0][0] = 0` and `dp[i][0] = infinity` for `i > 0`. We also need to handle the case where `p=1`, where `dp[i][1] = cost[0][i-1]`.

The state transition involves iterating through `i` from 1 to `n`, `p` from 1 to `k`, and `j` from 0 to `i-1`. The precomputation of `cost[i][j]` takes O(N^3) time (N^2 substrings, each checking O(N) divisors, and each divisor check takes O(N) time). The DP part takes O(k * N^2) time. The overall time complexity is dominated by the precomputation, making it O(N^3). The space complexity is O(N^2) for the `cost` table and O(k*N) for the DP table.

## ⏱️ Complexity

- **Time Complexity:** `O(N^3 + K*N^2)`
- **Space Complexity:** `O(N^2)`
