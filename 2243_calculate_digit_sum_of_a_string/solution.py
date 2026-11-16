class Solution:
  def digitSum(self, s: str, k: int) -> str:
    while len(s) > k:
      next_s = []
      for i in range(0, len(s), k):
        current_sum = 0
        for j in range(i, min(len(s), i + k)):
          current_sum += int(s[j])
        next_s.append(str(current_sum))
      s = ''.join(next_s)
    return s