func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func maxLength(nums []int) int {
	n := len(nums)
	maxLen := 0
	for i := 0; i < n; i++ {
		currProd := 1
		currGcd := 0
		currLcm := 1
		for j := i; j < n; j++ {
			val := nums[j]
			currProd *= val
			currGcd = gcd(currGcd, val)
			currLcm = (currLcm / gcd(currLcm, val)) * val
			if currProd == currLcm*currGcd {
				if j-i+1 > maxLen {
					maxLen = j - i + 1
				}
			}
		}
	}
	return maxLen
}