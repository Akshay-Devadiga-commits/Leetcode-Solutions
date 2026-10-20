class Solution:
    def sortableIntegers(self, nums: list[int]) -> int:
        n = len(nums)
        total_sum = 0
        divisors = [i for i in range(1, n + 1) if n % i == 0]
        for k in divisors:
            possible = True
            subarrays = []
            for i in range(0, n, k):
                sub = nums[i:i+k]
                drops = 0
                for j in range(k):
                    if sub[j] > sub[(j + 1) % k]:
                        drops += 1
                if drops > 1:
                    possible = False
                    break
                if drops == 1:
                    start_idx = 0
                    for j in range(k):
                        if sub[j] > sub[(j + 1) % k]:
                            start_idx = (j + 1) % k
                            break
                    sorted_sub = sub[start_idx:] + sub[:start_idx]
                else:
                    sorted_sub = sub
                subarrays.append(sorted_sub)
            if not possible:
                continue
            flat_array = [elem for sub in subarrays for elem in sub]
            for i in range(n - 1):
                if flat_array[i] > flat_array[i + 1]:
                    possible = False
                    break
            if possible:
                total_sum += k
        return total_sum