package set

import (
	"testing"
)

func sameSet[T comparable](s1 Set[T], s2 Set[T]) bool {
	if len(s1.m) != len(s2.m) {
		return false
	}
	for k := range s1.m {
		if _, ok := s2.m[k]; !ok {
			return false
		}
	}
	return true
}

func TestOfInt(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args []int
		want Set[int]
	}{
		"one element": {
			args: []int{1},
			want: Set[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
		},
		"several elements": {
			args: []int{1, 2, 3},
			want: Set[int]{
				m: map[int]struct{}{
					1: {},
					2: {},
					3: {},
				},
			},
		},
		"empty": {
			args: []int{},
			want: Set[int]{
				m: map[int]struct{}{},
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			got := Of(tt.args...)
			if !sameSet(got, tt.want) {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddInt(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args  []int
		start Set[int]
		want  Set[int]
	}{
		"initialization and one element": {
			args:  []int{1},
			start: Set[int]{},
			want: Set[int]{
				m: map[int]struct{}{
					1: {},
				},
			},
		},
		"initialization and several elements": {
			args:  []int{1, 2, 3},
			start: Set[int]{},
			want: Set[int]{
				m: map[int]struct{}{
					1: {},
					2: {},
					3: {},
				},
			},
		},
		"initialization and empty": {
			args:  []int{},
			start: Set[int]{},
			want: Set[int]{
				m: map[int]struct{}{},
			},
		},
		"no initialization and one element": {
			args:  []int{1},
			start: Of(-1, -2),
			want: Set[int]{
				m: map[int]struct{}{
					-1: {},
					-2: {},
					1:  {},
				},
			},
		},
		"no initialization and several elements": {
			args:  []int{1, 2, 3},
			start: Of(-1, -2),
			want: Set[int]{
				m: map[int]struct{}{
					-1: {},
					-2: {},
					1:  {},
					2:  {},
					3:  {},
				},
			},
		},
		"no initialization and empty": {
			args:  []int{},
			start: Of(-1, -2),
			want: Set[int]{
				m: map[int]struct{}{
					-1: {},
					-2: {},
				},
			},
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			tt.start.Add(tt.args...)
			if !sameSet(tt.start, tt.want) {
				t.Fatalf("got %v, want %v", tt.start, tt.want)
			}
		})
	}
}

func TestAddSet(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		args  Set[int]
		start Set[int]
		want  Set[int]
	}{
		"initialization and one element": {
			args:  Of(1),
			start: Set[int]{},
			want:  Of(1),
		},
		"initialization and several elements": {
			args:  Of(1, 2, 3),
			start: Set[int]{},
			want:  Of(1, 2, 3),
		},
		"initialization and empty": {
			args:  Of[int](),
			start: Set[int]{},
			want:  Of[int](),
		},
		"no initialization and one element": {
			args:  Of(1),
			start: Of(-1, -2),
			want:  Of(-1, -2, 1),
		},
		"no initialization and several elements": {
			args:  Of(1, 2, 3),
			start: Of(-1, -2),
			want:  Of(-1, -2, 1, 2, 3),
		},
		"no initialization and empty": {
			args:  Of[int](),
			start: Of(-1, -2),
			want:  Of(-1, -2),
		},
	}

	for name, tt := range tests {
		tt := tt
		t.Run(name, func(t *testing.T) {
			tt.start.AddSet(tt.args)
			if !sameSet(tt.start, tt.want) {
				t.Fatalf("got %v, want %v", tt.start, tt.want)
			}
		})
	}
}
