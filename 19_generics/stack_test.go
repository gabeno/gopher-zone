package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("integer stack", func(t *testing.T) {
		stackOfInts := new(Stack[int])

		AssertTrue(t, stackOfInts.IsEmpty())

		stackOfInts.Push(1)
		AssertFalse(t, stackOfInts.IsEmpty())

		stackOfInts.Push(2)
		value, _ := stackOfInts.Pop()
		AssertEqual(t, value, 2)
		value, _ = stackOfInts.Pop()
		AssertEqual(t, value, 1)
		AssertTrue(t, stackOfInts.IsEmpty())

		// check we get same types from stack
		stackOfInts.Push(10)
		stackOfInts.Push(1)
		firstNum, _ := stackOfInts.Pop()
		secondNum, _ := stackOfInts.Pop()
		AssertEqual(t, firstNum+secondNum, 11)
	})

	t.Run("strings stack", func(t *testing.T) {
		stackOfStrings := new(Stack[string])

		AssertTrue(t, stackOfStrings.IsEmpty())

		stackOfStrings.Push("A")
		AssertFalse(t, stackOfStrings.IsEmpty())

		stackOfStrings.Push("B")
		value, _ := stackOfStrings.Pop()
		AssertEqual(t, value, "B")
		value, _ = stackOfStrings.Pop()
		AssertEqual(t, value, "A")
		AssertTrue(t, stackOfStrings.IsEmpty())
	})
}

func AssertTrue(t testing.TB, got bool) {
	if !got {
		t.Errorf("got %v want true", got)
	}
}

func AssertFalse(t testing.TB, got bool) {
	if got {
		t.Errorf("got %v wanted false", got)
	}
}

func AssertEqual[T comparable](t testing.TB, got, want T) {
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
