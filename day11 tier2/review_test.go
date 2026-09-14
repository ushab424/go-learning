package main

import (
	"testing"
)

func TestReverseTable(t *testing.T) {
	tests := []struct {
		nameCase    string
		input, want string
	}{
		{"Normal", "cucumber", "rebmucuc"},
		{"ZeroString", "", ""},
		{"1Value", "q", "q"},
		{"Kirilian", "собака", "акабос"},
	}
	for _, test := range tests {
		t.Run(test.nameCase, func(t *testing.T) {
			result := Reverse(test.input)
			if result != test.want {
				t.Errorf("Error: Reverse(%q) = %q, want %q", test.input, result, test.want)
			}
		})
	}
}

func TestMaxTable(t *testing.T) {
	tests := []struct {
		nameCase string
		nums     []int
		want     int
		wantErr  bool
	}{
		{"ZeroSlice", []int{}, 0, true},
		{"OneElement", []int{1}, 1, false},
		{"Normal", []int{1, -1, 2, 5, 0}, 5, false},
		{"Negative", []int{-1, -2, -10, -4}, -1, false},
	}
	for _, test := range tests {
		t.Run(test.nameCase, func(t *testing.T) {
			result, err := Max(test.nums)
			if test.wantErr {
				if err == nil {
					t.Errorf("expected error, got nil")
				}
				return
			}
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != test.want {
				t.Errorf("Max() = %d, want %d", result, test.want)
			}
		})
	}
}

func TestContainsTable(t *testing.T) {
	tests := []struct {
		nameCase string
		slice    []string
		target   string
		want     bool
	}{
		{"NilSlice", []string{}, "dog", false},
		{"NilTarget", []string{"dog"}, "", false},
		{"Normal", []string{"dog"}, "dog", true},
		{"Cirilic", []string{"собака"}, "собака", true},
	}
	for _, test := range tests {
		t.Run(test.nameCase, func(t *testing.T) {
			result := Contains(test.slice, test.target)
			if result != test.want {
				t.Errorf("Contains() = %t, want %t", result, test.want)
			}
		})
	}
}
