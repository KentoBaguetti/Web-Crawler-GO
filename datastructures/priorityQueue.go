package datastructures

import "sync"

type ScoreValue struct {
	Score int
	Value any
}

type PriorityQueue struct {
	Elements []*ScoreValue
	Length   int
	MinHeap  bool
	Mux      sync.Mutex
}

func CreatePriorityQueue(isMinHeap bool) *PriorityQueue {
	pq := &PriorityQueue{Elements: make([]*ScoreValue, 0), Length: 0, MinHeap: isMinHeap, Mux: sync.Mutex{}}
	return pq
}

// Note: Since the PQ elements array has structs as elements, using Heapify on a general array might not be the play
func (pq *PriorityQueue) Heapify() {
	pq.Mux.Lock()
	defer pq.Mux.Unlock()
}

func (pq *PriorityQueue) Pop() ScoreValue {
	pq.Mux.Lock()
	defer pq.Mux.Unlock()

	item := pq.Elements[0]
	pq.Elements = pq.Elements[1:]

	pq.Length--
	return *item

}

func (pq *PriorityQueue) Append(value any, score int) {
	pq.Mux.Lock()
	defer pq.Mux.Unlock()

	newElement := &ScoreValue{Value: value, Score: score}
	pq.Elements = append(pq.Elements, newElement)
	pq.Length++

}

func (pq *PriorityQueue) Peek() ScoreValue {
	pq.Mux.Lock()
	defer pq.Mux.Unlock()

	return *pq.Elements[0]
}

func (pq *PriorityQueue) Size() int {
	pq.Mux.Lock()
	defer pq.Mux.Unlock()

	return pq.Length
}

/////////////////////////////////////////////////////////////////
// helper functions
/////////////////////////////////////////////////////////////////

func (pq *PriorityQueue) heapifyUp() {
	pq.Mux.Lock()
	defer pq.Mux.Unlock()
}

func (pq *PriorityQueue) heapifyDown() {
	pq.Mux.Lock()
	defer pq.Mux.Unlock()
}
