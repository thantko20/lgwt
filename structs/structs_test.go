package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.00

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("works for rectangles", func(t *testing.T) {
		rectangle := Rectangle{5.0, 10.0}
		got := rectangle.Area()
		want := 50.00

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("works for circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}
