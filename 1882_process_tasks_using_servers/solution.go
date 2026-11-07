package main

import (
	"container/heap"
)

type Server struct {
	weight int
	index  int
}

type AvailableHeap []Server

func (h AvailableHeap) Len() int           { return len(h) }
func (h AvailableHeap) Less(i, j int) bool {
	if h[i].weight != h[j].weight {
		return h[i].weight < h[j].weight
	}
	return h[i].index < h[j].index
}
func (h AvailableHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *AvailableHeap) Push(x interface{}) {
	*h = append(*h, x.(Server))
}
func (h *AvailableHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type BusyServer struct {
	freeTime int64
	weight   int
	index    int
}

type BusyHeap []BusyServer

func (h BusyHeap) Len() int           { return len(h) }
func (h BusyHeap) Less(i, j int) bool {
	if h[i].freeTime != h[j].freeTime {
		return h[i].freeTime < h[j].freeTime
	}
	if h[i].weight != h[j].weight {
		return h[i].weight < h[j].weight
	}
	return h[i].index < h[j].index
}
func (h BusyHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *BusyHeap) Push(x interface{}) {
	*h = append(*h, x.(BusyServer))
}
func (h *BusyHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func assignTasks(servers []int, tasks []int) []int {
	n := len(servers)
	m := len(tasks)
	ans := make([]int, m)

	available := &AvailableHeap{}
	heap.Init(available)
	for i := 0; i < n; i++ {
		heap.Push(available, Server{weight: servers[i], index: i})
	}

	busy := &BusyHeap{}
	heap.Init(busy)

	var currentTime int64 = 0
	taskIdx := 0

	for taskIdx < m {
		// 1. Free up servers
		for busy.Len() > 0 && (*busy)[0].freeTime <= currentTime {
			srv := heap.Pop(busy).(BusyServer)
			heap.Push(available, Server{weight: srv.weight, index: srv.index})
		}

		// 2. Assign task
		if available.Len() > 0 {
			srv := heap.Pop(available).(Server)
			ans[taskIdx] = srv.index

			freeTime := currentTime + int64(tasks[taskIdx])
			heap.Push(busy, BusyServer{freeTime: freeTime, weight: srv.weight, index: srv.index})
			taskIdx++
		} else {
			// 3. Jump to next free server
			currentTime = (*busy)[0].freeTime
		}

		// 4. Ensure currentTime >= taskIdx
		if taskIdx < m {
			if int64(taskIdx) > currentTime {
				currentTime = int64(taskIdx)
			}
		}
	}

	return ans
}