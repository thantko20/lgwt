package helloworld

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to Marco", func(t *testing.T) {
		got := Hello("Marco", "")
		want := "Hello, Marco"

		assertCorrectMessage(t, got, want)
	})
	t.Run("saying hello to the world!", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Maria", "Spanish")
		want := "Hola, Maria"

		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Marie", "French")
		want := "Bonjour, Marie"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
