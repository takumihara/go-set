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
	s.initOnce.Do(func() {})
	for _, v := range v {
		if v != v {
			panic("element in set has to be equal to itself")
		}
		s.m[v] = struct{}{}
	}
	return s
}

// WithCap returns a new set with the given capacity.
func WithCap[Elem comparable](cap int) Set[Elem] {
	s := Set[Elem]{
		m: make(map[Elem]struct{}, cap),
	}
	s.initOnce.Do(func() {})
	return s
}

// Add adds elements to a set.
func (s *Set[Elem]) Add(v ...Elem) {
	s.initOnce.Do(s.init)
	for _, v := range v {
		if v != v {
			panic("element in set has to be equal to itself")
		}
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

// ContainsAny reports whether any of the elements in s2 are in s.
func (s *Set[Elem]) ContainsAny(s2 Set[Elem]) bool {
	for k := range s2.m {
		if s.Contains(k) {
			return true
		}
	}
	return false
}

// ContainsAll reports whether all of the elements in s2 are in s.
func (s *Set[Elem]) ContainsAll(s2 Set[Elem]) bool {
	for k := range s2.m {
		if !s.Contains(k) {
			return false
		}
	}
	return true
}

// ToSlice returns the elements in the set s as a slice.
// The values will be in an indeterminate order.
func (s *Set[Elem]) ToSlice() []Elem {
	r := make([]Elem, 0, len(s.m))
	for k := range s.m {
		r = append(r, k)
	}
	return r
}

// Equal reports whether s and s2 contain the same elements.
func (s *Set[Elem]) Equal(s2 Set[Elem]) bool {
	if s.Len() != s2.Len() {
		return false
	}
	for k := range s.m {
		if _, ok := s2.m[k]; !ok {
			return false
		}
	}
	return true
}

// Clear removes all elements from s, leaving it empty.
func (s *Set[Elem]) Clear() {
	s.m = map[Elem]struct{}{}
}

// Clone returns a copy of s.
// The elements are copied using assignment,
// so this is a shallow clone.
func (s *Set[Elem]) Clone() Set[Elem] {
	r := WithCap[Elem](len(s.m))
	for k := range s.m {
		r.m[k] = struct{}{}
	}
	return r
}

// Retain deletes any elements from s for which keep returns false.
func (s *Set[Elem]) Retain(keep func(Elem) bool) {
	for k := range s.m {
		if !keep(k) {
			delete(s.m, k)
		}
	}
}

// Len returns the number of elements in s.
func (s *Set[Elem]) Len() int {
	return len(s.m)
}

// Do calls f on every element in the set s,
// stopping if f returns false.
// f should not change s.
// f will be called on values in an indeterminate order.
func (s *Set[Elem]) Do(f func(Elem) bool) {
	for k := range s.m {
		if !f(k) {
			break
		}
	}
}

// Pop removes and returns an arbitrary element from s.
func (s *Set[Elem]) Pop() (Elem, bool) {
	for k := range s.m {
		delete(s.m, k)
		return k, true
	}
	var r Elem
	return r, false
}

func (s *Set[Elem]) init() {
	s.m = map[Elem]struct{}{}
}

// Union constructs a new set containing the union of s1 and s2.
func Union[Elem comparable](s1, s2 Set[Elem]) Set[Elem] {
	r := WithCap[Elem](len(s1.m) + len(s2.m))
	for k := range s1.m {
		r.m[k] = struct{}{}
	}
	for k := range s2.m {
		r.m[k] = struct{}{}
	}
	return r
}

// Intersect constructs a new set containing the intersection of s1 and s2.
func Intersect[Elem comparable](s1, s2 Set[Elem]) Set[Elem] {
	if len(s1.m) > len(s2.m) {
		s1, s2 = s2, s1
	}
	r := WithCap[Elem](len(s1.m))
	for k := range s1.m {
		if s2.Contains(k) {
			r.m[k] = struct{}{}
		}
	}
	return r
}

// Difference constructs a new set containing the elements of s1 that are not in s2.
func Difference[Elem comparable](s1, s2 Set[Elem]) Set[Elem] {
	r := WithCap[Elem](len(s1.m))
	for k := range s1.m {
		if _, ok := s2.m[k]; !ok {
			r.m[k] = struct{}{}
		}
	}
	return r
}
