package datastructures

import "sync"

type ScoreValue struct {
	Score int
	Value any
}

type PriorityQueue struct {
	Elements []ScoreValue
	Length   int
	Mux      sync.Mutex
}
