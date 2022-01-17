package main

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("Collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("Sum 2 slices", func(t *testing.T) {
		s1 := []int{1, 2}
		s2 := []int{0, 9}

		got := SumAll(s1, s2)
		want := []int{3, 9}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v, %v", got, want, s1, s2)
		}
	})

	t.Run("Sum 1 slice", func(t *testing.T) {
		s1 := []int{1, 2}

		got := SumAll(s1)
		want := []int{3}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %d want %d given, %v", got, want, s1)
		}
	})
}
