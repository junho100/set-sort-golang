package case1

type Set struct {
	data map[int]struct{}
}

func NewSet() *Set {
	s := &Set{
		data: make(map[int]struct{}),
	}
	return s
}

func (s *Set) Add(element int) {
	s.data[element] = struct{}{}
}

func (s *Set) Remove(element int) {
	delete(s.data, element)
}

func (s *Set) Contains(element int) bool {
	_, exists := s.data[element]
	return exists
}

func (s *Set) Size() int {
	return len(s.data)
}

func (s *Set) ToSlice() []int {
	slice := make([]int, 0, len(s.data))
	for value := range s.data {
		slice = append(slice, value)
	}
	return slice
}
