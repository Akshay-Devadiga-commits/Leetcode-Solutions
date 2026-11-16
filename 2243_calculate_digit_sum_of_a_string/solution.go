func digitSum(s string, k int) string {
    for len(s) > k {
        var next_s strings.Builder
        for i := 0; i < len(s); i += k {
            current_sum := 0
            for j := i; j < min(len(s), i+k); j++ {
                current_sum += int(s[j] - '0')
            }
            next_s.WriteString(strconv.Itoa(current_sum))
        }
        s = next_s.String()
    }
    return s
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}