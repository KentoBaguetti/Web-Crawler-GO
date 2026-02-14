package datastructures

import "sync"

// switch to use RWMutex as reads dominate writes

type Set struct {
	Elements map[string]bool
	Length   int
	Mux      sync.Mutex
}

func (s *Set) Add(url string) {
	s.Mux.Lock()
	defer s.Mux.Unlock()

	_, ok := s.Elements[url]

	if ok {
		return
	} else {
		s.Elements[url] = true
		s.Length++
	}
}

func (s *Set) Contains(url string) bool {
	s.Mux.Lock()
	defer s.Mux.Unlock()

	_, ok := s.Elements[url]
	return ok
}

func (s *Set) Size() int {
	s.Mux.Lock()
	defer s.Mux.Unlock()
	return s.Length
}
