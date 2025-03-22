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
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		assertDefinition(t, dictionary, word, definition)

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
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)
		assertError(t, false, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})
	t.Run("adding existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "new line")

		assertError(t, true, err, ErrWordAlreadyExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("updating existing word", func(t *testing.T) {
		word := "test"
		dictionaty := Dictionary{word: "old definition"}
		newDefinition := "new definition"

		err := dictionaty.Update(word, newDefinition)
		assertError(t, false, err, nil)
		assertDefinition(t, dictionaty, word, newDefinition)
	})
	t.Run("updating non-existing word", func(t *testing.T) {
		word := "test"
		dictionaty := Dictionary{}
		newDefinition := "new definition"

		err := dictionaty.Update(word, newDefinition)
		assertError(t, true, err, ErrNoSuchWord)
	})
}
