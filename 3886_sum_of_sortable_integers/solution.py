class Solution:
    def sortableIntegers(self, nums: list[int]) -> int:
        n = len(nums)
        target = sorted(nums)
        
        def is_cyclic_shift(sub, orig):
            # Check if sub is a cyclic shift of orig
            if len(sub) != len(orig):
                return False
            double_orig = orig + orig
            sub_str = ",".join(map(str, sub))
            double_orig_str = ",".join(map(str, double_orig))
            return sub_str in double_orig_str
        
        ans = 0
        for k in range(1, n + 1):
            if n % k == 0:
                possible = True
                for i in range(0, n, k):
                    sub = nums[i:i+k]
                    orig = target[i:i+k]
                    if not is_cyclic_shift(sub, orig):
                        possible = False
                        break
                if possible:
                    ans += k
        return ans