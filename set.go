package set

import "sync"

type Set[Elem comparable] struct {
	// contains filtered or unexported fields
	initOnce sync.Once
	m        map[Elem]struct{}
}

// Of returns a new set containing the listed elements.
func Of[Elem comparable](v ...Elem) Set[Elem] {
	s := Set[Elem]{
		m: make(map[Elem]struct{}, len(v)),
	}
	// TOOD: what is the best way to use sync.Once
	s.initOnce.Do(func() {})
	for _, v := range v {
		s.m[v] = struct{}{}
	}
	return s
}

// Add adds elements to a set.
func (s *Set[Elem]) Add(v ...Elem) {
	s.initOnce.Do(s.init)
	for _, v := range v {
		s.m[v] = struct{}{}
	}
}

// AddSet adds the elements of set s2 to s.
func (s *Set[Elem]) AddSet(s2 Set[Elem]) {
	s.initOnce.Do(s.init)
	for k := range s2.m {
		s.m[k] = struct{}{}
	}
}

// Remove removes elements from a set.
// Elements that are not present are ignored.
func (s *Set[Elem]) Remove(v ...Elem) {
	// TODO: is initialization needed?
	for _, v := range v {
		delete(s.m, v)
	}
}

// RemoveSet removes the elements of set s2 from s.
// Elements present in s2 but not s are ignored.
func (s *Set[Elem]) RemoveSet(s2 Set[Elem]) {
	for k := range s2.m {
		delete(s.m, k)
	}
}

// Contains reports whether v is in the set.
func (s *Set[Elem]) Contains(v Elem) bool {
	_, ok := s.m[v]
	return ok
}



func (s *Set[Elem]) init() {
	s.m = map[Elem]struct{}{}
}
