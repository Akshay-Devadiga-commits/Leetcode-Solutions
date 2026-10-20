import (
    "sort"
    "strconv"
    "strings"
)

func sortableIntegers(nums []int) int {
    n := len(nums)
    target := make([]int, n)
    copy(target, nums)
    sort.Ints(target)

    isCyclicShift := func(sub, orig []int) bool {
        if len(sub) != len(orig) {
            return false
        }
        doubleOrig := append(orig, orig...)
        
        subStrs := make([]string, len(sub))
        for i, v := range sub {
            subStrs[i] = strconv.Itoa(v)
        }
        subStr := strings.Join(subStrs, ",")

        doubleOrigStrs := make([]string, len(doubleOrig))
        for i, v := range doubleOrig {
            doubleOrigStrs[i] = strconv.Itoa(v)
        }
        doubleOrigStr := strings.Join(doubleOrigStrs, ",")

        return strings.Contains(doubleOrigStr, subStr)
    }

    ans := 0
    for k := 1; k <= n; k++ {
        if n%k == 0 {
            possible := true
            for i := 0; i < n; i += k {
                sub := nums[i : i+k]
                orig := target[i : i+k]
                if !isCyclicShift(sub, orig) {
                    possible = false;
                    break
                }
            }
            if possible {
                ans += k
            }
        }
    }
    return ans
}
