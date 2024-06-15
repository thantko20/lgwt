package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	t.Run("correctly computes (2 + 2)", func(t *testing.T) {
		sum := Add(2, 2)
		expected := 4

		assertComputation(t, sum, expected)
	})

	t.Run("correctly computes (4 + 10)", func(t *testing.T) {
		sum := Add(4, 10)
		expected := 14

		assertComputation(t, sum, expected)
	})

	t.Run("adds correctly for negative integers", func(t *testing.T) {
		sum := Add(-6, 32)
		expected := 26

		assertComputation(t, sum, expected)
	})
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}

func assertComputation(t testing.TB, sum, expected int) {
	t.Helper()

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
