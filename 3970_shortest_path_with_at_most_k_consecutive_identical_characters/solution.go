import (
	"container/heap"
	"fmt"
)

type Item struct {
	dist      int
	node      int
	lastChar  byte
	count     int
	index     int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].dist < pq[j].dist }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}
func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
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
	node     int
	lastChar byte
	count    int
}

type Edge struct {
	to     int
	weight int
}

func shortestPath(n int, edges [][]int, labels string, k int) int {
	adj := make([][]Edge, n)
	for _, edge := range edges {
		adj[edge[0]] = append(adj[edge[0]], Edge{to: edge[1], weight: edge[2]})
	}

	initChar := labels[0]
	if 1 > k {
		return -1
	}

	pq := &PriorityQueue{&Item{dist: 0, node: 0, lastChar: initChar, count: 1}}
	heater.Init(pq)

	dist := make(map[StateKey]int)
	initState := StateKey{node: 0, lastChar: initChar, count: 1}
	dist[initState] = 0

	for pq.Len() > 0 {
		curr := heap.Pop(pq).(*Item)

		if curr.node == n-1 {
			return curr.dist
		}

		stateKey := StateKey{node: curr.node, lastChar: curr.lastChar, count: curr.count}
		if d, ok := dist[stateKey]; ok && curr.dist > d {
			continue
		}

		for _, edge := range adj[curr.node] {
			vChar := labels[edge.to]
			newCount := 1
			if vChar == curr.lastChar {
				newCount = curr.count + 1
			}

			if newCount <= k {
				newDist := curr.dist + edge.weight
				nextState := StateKey{node: edge.to, lastChar: vChar, count: newCount}
				if d, ok := dist[nextState]; !ok || newDist < d {
					dist[nextState] = newDist
					heater.Push(pq, &Item{dist: newDist, node: edge.to, lastChar: vChar, count: newCount})
				}
			}
		}
	}

	return -1
}