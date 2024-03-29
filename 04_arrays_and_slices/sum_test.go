package sum

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("compute sum of numbers in a collection of any size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		sum := Sum(numbers)
		expected := 15

		if sum != expected {
			t.Errorf("got %d expected %d for numbers %v", sum, expected, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{3, 4, 2})
	expected := []int{3, 9}

	if !slices.Equal(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}

func TestSumAllTails(t *testing.T) {
	got := SumAllTails([]int{1, 2, 3}, []int{2, 3}, []int{1}, []int{})
	expected := []int{5, 3, 0, 0}

	if !slices.Equal(got, expected) {
		t.Errorf("got %v expected %v", got, expected)
	}
}

func TestReduce(t *testing.T) {
	t.Run("multiplication of all elements", func(t *testing.T) {
		multiply := func(x, y int) int {
			return x * y
		}

		AssertEqual(t, Reduce([]int{1, 2, 3}, multiply, 1), 6)
	})

	t.Run("concatenate strings", func(t *testing.T) {
		concatenator := func(a, b string) string {
			return a + b
		}

		AssertEqual(t, Reduce([]string{"a", "b", "c"}, concatenator, ""), "abc")
	})
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	if got != want {
		t.Errorf("got %#v want %#v", got, want)
	}
}
