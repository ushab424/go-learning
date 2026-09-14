package main

import (
	"testing"
)

func TestCountWordsTable(t *testing.T) {
	tests := []struct {
		nameTest string
		s        string
		want     int
	}{
		{"ZeroString", "", 0},
		{"OneWord", "test", 1},
		{"MoreWords", "test dog writing", 3},
		{"MoreFields", "test dog  ", 2},
	}
	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			result := CountWords(test.s)
			if result != test.want {
				t.Errorf("CountWord(%s) = %d, want: %d", test.s, result, test.want)
			}
		})
	}
}

func TestFiltertable(t *testing.T) {
	isEven := func(n int) bool { return n%2 == 0 }
	// filter function by tests

	tests := []struct {
		nameTest string
		nums     []int
		want     []int
	}{
		{"ZeroSlice", []int{}, []int{}},
		{"AllPass", []int{2, 4, 6}, []int{2, 4, 6}},
		{"NonePass", []int{1, 3, 5}, []int{}},
		{"Mixed", []int{1, 2, 3, 4}, []int{2, 4}},
	}
	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			result := Filter(test.nums, isEven)
			if len(result) != len(test.want) {
				t.Errorf("got %v, want %v", result, test.want)
				return
				// we can used t.Fatalf

			}
			for i := range result {
				if result[i] != test.want[i] {
					t.Errorf("got %v, want %v", result, test.want)
				}
			}
		})
	}
}

func TestMergeTable(t *testing.T) {
	tests := []struct {
		nameTest string
		a, b     map[string]int
		want     map[string]int
	}{
		{"TwoZero", map[string]int{}, map[string]int{}, map[string]int{}},
		{"OneZero", map[string]int{}, map[string]int{"a": 1}, map[string]int{"a": 1}},
		{"NoIntersection", map[string]int{"a": 1}, map[string]int{"b": 2}, map[string]int{"a": 1, "b": 2}},
		{"Intersection", map[string]int{"a": 1}, map[string]int{"a": 2}, map[string]int{"a": 2}},
		// variation test map
	}
	for _, test := range tests {
		t.Run(test.nameTest, func(t *testing.T) {
			result := Merge(test.a, test.b)
			if len(result) != len(test.want) {
				t.Errorf("got %v, want %v", result, test.want)
				return
				// we can used t.Fatalf
			}
			for k, v := range test.want {
				if result[k] != v {
					t.Errorf("keys %s: got %d, want %d", k, result[k], v)
				}
			}
		})
	}
}
