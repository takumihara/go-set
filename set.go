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
	for _, v := range v {
		s.m[v] = struct{}{}
	}
	return s
}
