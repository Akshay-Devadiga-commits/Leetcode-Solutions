func canAliceWin(n int) bool {
	k := int((21 - math.Sqrt(float64(441-8*n))) / 2)
	return k%2 == 1
}