package maps

import "testing"

func assertStrings(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func assertError(t testing.TB, wantErr bool, errGot, errWant error) {
	t.Helper()
	if errGot != nil && !wantErr {
		t.Errorf("got unexpected error: %v", errGot)
		return
	}
	if errGot == nil && !wantErr {
		return
	}
	if errGot == nil && wantErr {
		t.Fatal("expected an error, but got non")
	}
	if errGot.Error() != errWant.Error() && wantErr {
		t.Errorf("wantErr: %q, gotErr: %q", errWant, errGot)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()
	got, err := dictionary.Search(word)
	assertError(t, false, err, nil)
	assertStrings(t, got, definition)
}

func TestSearch(t *testing.T) {
	t.Run("find existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, err := dictionary.Search("test")
		want := "this is just a test"
		assertError(t, false, err, nil)
		assertStrings(t, got, want)

	})
	t.Run("find non-existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("tes")

		assertError(t, true, err, ErrNoSuchWord)
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding new word", func(t *testing.T) {
		dictionary := Dictionary{}

		dictionary.Add("test", "this is just a test")
		want := "this is just a test"
		got, err := dictionary.Search("test")

		assertError(t, false, err, nil)
		assertStrings(t, got, want)
	})
	t.Run("adding existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		_, err := dictionary.Add(word, "new line")

		assertError(t, true, err, ErrWordAlreadyExists)
		assertDefinition(t, dictionary, word, definition)
	})
}
