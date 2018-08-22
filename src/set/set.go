package set

type set struct {
	s map[interface{}]struct{}
}

// NewSet new a set
func NewSet() *set {
	return &set{
		s: make(map[interface{}]struct{}),
	}
}

// Contains whether a element in the set?
func (s *set) Contains(e interface{}) bool {
	_, ok := s.s[e]
	return ok
}

// Add a element to set
func (s *set) Add(e interface{}) {
	if s.Contains(e) {
		return
	}

	s.s[e] = struct{}{}
}

// Remove element from set
func (s *set) Remove(e interface{}) {
	if !s.Contains(e) {
		return
	}

	delete(s.s, e)
}

// Card return numbers of members of set
func (s *set) Card() int {
	return len(s.s)
}

// Members return all members of set
func (s *set) Members() []interface{} {
	members := make([]interface{}, s.Card())
	i := 0
	for k, _ := range s.s {
		members[i] = k
		i++
	}
	return members
}

// Diff Returns the members of the set resulting from the difference between the two sets
func (s *set) Diff(s2 *set) []interface{} {
	diff := make([]interface{}, s.Card())

	i := 0
	for k, _ := range s.s {
		if !s2.Contains(k) {
			diff[i] = k
			i++
		}
	}

	return diff[:i]
}

// Iterate return a channel from which every item in set could be read
func (s *set) Iterate() <-chan interface{} {
	itr := make(chan interface{})

	go func() {
		for k, _ := range s.s {
			itr <- k
		}
		close(itr)

	}()

	return itr
}

// Inter Returns the members of the set resulting from the intersection of the arguments
func Inter(s, s2 *set) *set {
	s3 := NewSet()
	for k, _ := range s.s {
		if s2.Contains(k) {
			s3.Add(k)
		}
	}
	return s3
}

func Union(s1, s2 *set) *set {
	s3 := NewSet()
	for k, _ := range s1.s {
		s3.Add(k)
	}
	for k, _ := range s2.s {
		s3.Add(k)
	}
	return s3
}
