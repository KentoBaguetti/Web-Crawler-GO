package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Init")

	q := Queue{elements: make([] string, 0), length: 0}

	q.Enqueue("Hello")
	q.Enqueue("Kentaro")
	q.Dequeue()

	fmt.Println(q)
}

type Queue struct {
	elements []string
	length int
	mu sync.Mutex
}

func (q *Queue) Enqueue(element string) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.elements = append(q.elements, element)
	q.length++
}

func (q *Queue) Dequeue() string {
	q.mu.Lock()
	defer q.mu.Unlock()

	url := q.elements[0]
	q.elements = q.elements[1:]
	q.length--
	return url
}

func (q *Queue) Size() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return q.length
}



