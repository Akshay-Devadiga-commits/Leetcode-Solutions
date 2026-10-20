class Solution:
    def sortableIntegers(self, nums: list[int]) -> int:
        n = len(nums)
        
        def is_sorted_rotation(sub):
            m = len(sub)
            if m <= 1:
                return True
            drops = 0
            for i in range(m):
                if sub[i] > sub[(i + 1) % m]:
                    drops += 1
            return drops <= 1

        total_sum = 0
        for k in range(1, n + 1):
            if n % k != 0:
                continue
            
            subarrays = [nums[i:i + k] for i in range(0, n, k)]
            
            valid = True
            for sub in subarrays:
                if not is_sorted_rotation(sub):
                    valid = False
                    break
            
            if not valid:
                continue
                
            sorted_subarrays = []
            for sub in subarrays:
                m = len(sub)
                start_idx = 0
                for i in range(m):
                    if sub[i] > sub[(i + 1) % m]:
                        start_idx = (i + 1) % m
                        break
                rotated = sub[start_idx:] + sub[:start_idx]
                sorted_subarrays.append(rotated)
            
            flat = []
            for sub in sorted_subarrays:
                flat.extend(sub)
                
            is_non_decreasing = all(flat[i] <= flat[i+1] for i in range(n - 1))
            if is_non_decreasing:
                total_sum += k
                
        return total_sum