import (
    "container/heap"
    "fmt"
)

type Item struct {
    w         int64
    u         int
    lastChar  byte
    count     int
    index     int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].w < pq[j].w }
func (pq PriorityQueue) Swap(i, j int) {
    pq[i], pq[j] = pq[j], pq[i]
    pq[i].index = i
    pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
    item := x.(*Item)
    item.index = len(*pq)
    *pq = append(*pq, item)
}
func (pq *PriorityQueue) Pop() interface{} {
    old := *pq
    n := len(old)
    item := old[n-1]
    old[n-1] = nil
    item.index = -1
    *pq = old[0 : n-1]
    return item
}

type StateKey struct {
    u        int
    lastChar byte
    count    int
}

func shortestPath(n int, edges [][]int, labels string, k int) int {
    adj := make([][]pair, n)
    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        adj[u] = append(adj[u], pair{v, w})
    }

    pq := &PriorityQueue{}
    heap.Init(pq)
    heap.Push(pq, &Item{w: 0, u: 0, lastChar: labels[0], count: 1})

    dist := make(map[StateKey]int64)

    for pq.Len() > 0 {
        item := heap.Pop(pq).(*Item)
        w, u, lastChar, count := item.w, item.u, item.lastChar, item.count

        if u == n-1 {
            return int(w)
        }

        state := StateKey{u, lastChar, count}
        if existingW, ok := dist[state]; ok && existingW < w {
            continue
        }
        dist[state] = w

        for _, edge := range adj[u] {
            v, weight := edge.v, edge.w
            nxtChar := labels[v]
            nxtCount := 1
            if nxtChar == lastChar {
                nxtCount = count + 1
            }

            if nxtCount <= k {
                nxtW := w + int64(weight)
                nxtState := StateKey{v, nxtChar, nxtCount}
                if existingW, ok := dist[nxtState]; !ok || nxtW < existingW {
                    dist[nxtState] = nxtW
                    heap.Push(pq, &Item{w: nxtW, u: v, lastChar: nxtChar, count: nxtCount})
                }
            }
        }
    }

    return -1
}

type pair struct {
    v int
    w int
}