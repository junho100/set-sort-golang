package case4

import "sync"

type Set struct {
	data map[int]struct{}
	m    sync.Mutex
}

func NewSet() *Set {
	s := &Set{
		data: make(map[int]struct{}),
		m:    sync.Mutex{},
	}
	return s
}

func (s *Set) Add(element int) {
	s.m.Lock()
	s.data[element] = struct{}{}
	s.m.Unlock()
}

func (s *Set) AddAll(elements []int) {
	s.m.Lock()
	for element := range elements {
		s.data[element] = struct{}{}
	}
	s.m.Unlock()
}

func (s *Set) Remove(element int) {
	s.m.Lock()
	delete(s.data, element)
	s.m.Unlock()
}

func (s *Set) Contains(element int) bool {
	s.m.Lock()
	_, exists := s.data[element]
	s.m.Unlock()
	return exists
}

func (s *Set) Size() int {
	s.m.Lock()
	defer s.m.Unlock()
	return len(s.data)
}

func (s *Set) ToSlice() []int {
	s.m.Lock()
	slice := make([]int, 0, len(s.data))
	for value := range s.data {
		slice = append(slice, value)
	}
	s.m.Unlock()
	return slice
}
