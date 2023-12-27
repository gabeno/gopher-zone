package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Height: 10.0, Width: 10.0}, 40.0},
		{Circle{Radius: 10.0}, 62.83185307179586},
		{Triangle{Base: 10.0, Height: 5.0}, 26.18033988749895},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Height: 10.0, Width: 10.0}, 100.0},
		{Circle{Radius: 10.0}, 314.1592653589793},
		{Triangle{Base: 10.0, Height: 5.0}, 25.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
