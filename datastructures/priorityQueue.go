package datastructures

import "sync"

type ScoreValue struct {
	Score int
	Value any
}

type PriorityQueue struct {
	Elements []ScoreValue
	Length   int
	MinHeap  bool
	Mux      sync.Mutex
}

func (pq *PriorityQueue) Heapify() {
	pq.Mux.Lock()
	defer pq.Mux.Unlock()
}

func (pq *PriorityQueue) Pop() {
	pq.Mux.Lock()
	defer pq.Mux.Unlock()
}

func (pq *PriorityQueue) Append(element ScoreValue) {
	pq.Mux.Lock()
	defer pq.Mux.Unlock()
}
