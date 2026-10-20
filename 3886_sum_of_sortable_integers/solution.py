class Solution:
    def sortableIntegers(self, nums: list[int]) -> int:
        n = len(nums)
        
        def can_sort_with_k(k):
            subarrays = []
            for i in range(0, n, k):
                sub = nums[i:i+k]
                drops = 0
                for j in range(k):
                    if sub[j] > sub[(j + 1) % k]:
                        drops += 1
                if drops > 1:
                    return False
                
                if drops == 1:
                    split = 0
                    for j in range(k):
                        if sub[j] > sub[(j + 1) % k]:
                            split = j + 1
                            break
                    rotated_sub = sub[split:] + sub[:split]
                else:
                    rotated_sub = sub
                subarrays.append(rotated_sub)
            
            for i in range(len(subarrays) - 1):
                if subarrays[i][-1] > subarrays[i+1][0]:
                    return False
            
            return True

        total_sum = 0
        for k in range(1, n + 1):
            if n % k == 0:
                if can_sort_with_k(k):
                    total_sum += k
                    
        return total_sum