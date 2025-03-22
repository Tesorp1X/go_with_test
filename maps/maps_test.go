package maps

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, errGot, errWant error) {
	t.Helper()
	if errGot == nil {
		t.Fatal("expected an error, but got non")
	}
	if errGot.Error() != errWant.Error() {
		t.Errorf("wantErr: %q, gotErr: %q", errWant, errGot)
	}
}

func TestSearch(t *testing.T) {
	t.Run("find existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)

	})
	t.Run("find non-existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("tes")

		assertError(t, err, ErrNoSuchWord)
	})
}
