package hello

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Liz", "English")
		want := "Hello, Liz"
		assertCorrect(t, got, want)
	})
	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrect(t, got, want)
	})
	t.Run("Hello in German", func(t *testing.T) {
		got := Hello("Liz", "German")
		want := "Hallo, Liz"
		assertCorrect(t, got, want)
	})
	t.Run("Hello in Spanish", func(t *testing.T) {
		got := Hello("Liz", "Spanish")
		want := "Hola, Liz"
		assertCorrect(t, got, want)
	})
	t.Run("Hello in Russian", func(t *testing.T) {
		got := Hello("Лиз", "Russian")
		want := "Привет, Лиз"
		assertCorrect(t, got, want)
	})
}

func assertCorrect(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
