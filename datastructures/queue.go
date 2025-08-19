package datastructures

import "sync"

type Queue struct {
	Elements []string
	Length int
	Mux sync.Mutex
}

func (q *Queue) Enqueue(element string) {
	q.Mux.Lock()
	defer q.Mux.Unlock()

	q.Elements = append(q.Elements, element)
	q.Length++
}

func (q *Queue) Dequeue() string {
	q.Mux.Lock()
	defer q.Mux.Unlock()

	url := q.Elements[0]
	q.Elements = q.Elements[1:]
	q.Length--
	return url
}

func (q *Queue) Size() int {
	q.Mux.Lock()
	defer q.Mux.Unlock()

	return q.Length
}

func (q* Queue) IsEmpty() bool {
	q.Mux.Lock()
	defer q.Mux.Unlock()

	return q.Length == 0
}