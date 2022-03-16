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

func (s *Set[Elem]) init() {
	s.m = map[Elem]struct{}{}
}
