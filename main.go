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

	fmt.Println(q.elements)

	s := Set{elements: make(map[string]bool, 0), length: 0}

	s.add("testUrl")
	s.add("testUrl")

	fmt.Println(s.elements)

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

type Set struct {
	elements map[string]bool
	length int
	mu sync.Mutex
}

func (s *Set) add(url string) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.elements[url]

	if ok {
		return
	} else {
		s.elements[url] = true
		s.length++
	}
}

func (s *Set) contains(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.elements[url]
	return ok
}

func (s *Set) Size() int {
	return s.length
}



