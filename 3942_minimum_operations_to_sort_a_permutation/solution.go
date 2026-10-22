import (
    "container/list"
    "fmt"
)

type State struct {
    arr  []int
    dist int
}

func minOperations(nums []int) int {
    n := len(nums)
    target := make([]int, n)
    for i := 0; i < n; i++ {
        target[i] = i
    }

    targetKey := fmt.Sprint(target)
    startKey := fmt.Sprint(nums)

    if startKey == targetKey {
        return 0
    }

    queue := list.New()
    queue.PushBack(State{arr: nums, dist: 0})

    visited := make(map[string]bool)
    visited[startKey] = true

    for queue.Len() > 0 {
        elem := queue.Front()
        queue.Remove(elem)
        currState := elem.Value.(State)
        curr := currState.arr
        dist := currState.dist

        nextRev := make([]int, n)
        for i := 0; i < n; i++ {
            nextRev[i] = curr[n-1-i]
        }
        revKey := fmt.Sprint(nextRev)
        if revKey == targetKey {
            return dist + 1
        }
        if !visited[revKey] {
            visited[revKey] = true
            queue.PushBack(State{arr: nextRev, dist: dist + 1})
        }

        nextRot := make([]int, n)
        for i := 0; i < n-1; i++ {
            nextRot[i] = curr[i+1]
        }
        nextRot[n-1] = curr[0]
        rotKey := fmt.Sprint(nextRot)
        if rotKey == targetKey {
            return dist + 1
        }
        if !visited[rotKey] {
            visited[rotKey] = true
            queue.PushBack(State{arr: nextRot, dist: dist + 1})
        }
    }

    return -1
}