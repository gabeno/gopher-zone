package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Perimeter()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkPerimeter(t, rectangle, 40.0)
	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkPerimeter(t, circle, 62.83185307179586)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangle", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}
		checkArea(t, rectangle, 100.0)

	})

	t.Run("circle", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
