package sum

import "testing"

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
