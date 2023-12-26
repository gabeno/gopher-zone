package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("add two positive integers", func(t *testing.T) {
		sum := Add(2, 2)
		want := 4

		if sum != want {
			t.Errorf("expected %d but got %d", sum, want)
		}
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
