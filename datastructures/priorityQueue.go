package datastructures

import (
	"errors"
	"sync"
)

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

func (pq *PriorityQueue) Pop() (ScoreValue, error) {
	pq.Mux.Lock()
	defer pq.Mux.Unlock()

	if pq.Length == 0 {
		return ScoreValue{}, errors.New("The PQ is empty")
	} else if pq.Length == 1 {
		pq.Length--
		item := pq.Elements[0]
		pq.Elements = pq.Elements[:0]
		return *item, nil
	}

	item := pq.Elements[0]
	pq.Elements[0] = pq.Elements[pq.Length-1]
	pq.Elements = pq.Elements[:pq.Length-1]
	pq.Length--
	pq.heapifyDown()

	return *item, nil

}

func (pq *PriorityQueue) Append(value any, score int) {
	pq.Mux.Lock()
	defer pq.Mux.Unlock()

	newElement := &ScoreValue{Value: value, Score: score}
	pq.Elements = append(pq.Elements, newElement)
	pq.Length++
	pq.heapifyUp()

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

	currIndex := pq.Length - 1
	flag := true

	for flag {

		var parentIndex int
		var parentElement *ScoreValue
		currElement := pq.Elements[currIndex]

		if currIndex == 0 {
			flag = false
			break
		}

		if currIndex%2 == 0 {
			parentIndex = currIndex/2 - 1
		} else {
			parentIndex = currIndex / 2
		}

		parentElement = pq.Elements[parentIndex]

		currScore := currElement.Score
		parentScore := parentElement.Score

		if pq.MinHeap {
			if currScore > parentScore {
				flag = false
				break
			}
		} else {
			if currScore < parentScore {
				flag = false
				break
			}
		}

		pq.Elements[currIndex] = parentElement
		pq.Elements[parentIndex] = currElement
		currIndex = parentIndex

	}

}

func (pq *PriorityQueue) heapifyDown() {

	currIndex := 0
	flag := true

	for flag {

		leftIndex := 2*currIndex + 1
		rightIndex := 2*currIndex + 2

		var leftElement *ScoreValue
		var rightElement *ScoreValue

		rightExists := false

		if leftIndex > pq.Length-1 && rightIndex > pq.Length-1 {
			flag = false
			break
		} else if leftIndex > pq.Length-1 {
			flag = false
			break
		}

		leftElement = pq.Elements[leftIndex]

		if rightIndex < pq.Length {
			rightElement = pq.Elements[rightIndex]
			rightExists = true
		}

		if pq.MinHeap {

			if rightExists {
				if pq.Elements[currIndex].Score <= leftElement.Score && pq.Elements[currIndex].Score <= rightElement.Score {
					flag = false
					break
				}
			} else {
				if pq.Elements[currIndex].Score <= leftElement.Score {
					flag = false
					break
				}
			}

		} else {

			if rightExists {
				if pq.Elements[currIndex].Score >= leftElement.Score && pq.Elements[currIndex].Score >= rightElement.Score {
					flag = false
					break
				}
			} else {
				if pq.Elements[currIndex].Score >= leftElement.Score {
					flag = false
					break
				}
			}

		}

		var swapIndex int
		var swapElement *ScoreValue

		if rightExists {

			if pq.MinHeap {

				if leftElement.Score < rightElement.Score {
					swapElement = leftElement
					swapIndex = leftIndex
				} else {
					swapElement = rightElement
					swapIndex = rightIndex
				}

			} else {

				if leftElement.Score > rightElement.Score {
					swapElement = leftElement
					swapIndex = leftIndex
				} else {
					swapElement = rightElement
					swapIndex = rightIndex
				}

			}

		} else {

			swapElement = leftElement
			swapIndex = leftIndex

		}

		pq.Elements[swapIndex] = pq.Elements[currIndex]
		pq.Elements[currIndex] = swapElement
		currIndex = swapIndex

	}

}
