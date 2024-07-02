package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("defined word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("undefined word", func(t *testing.T) {
		_, err := dictionary.Search("burger")

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("adding new word", func(t *testing.T) {
		dictionary := Dictionary{}

		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("adding existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}

		err := dictionary.Add(word, "this is blabla")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("updating an existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{word: definition}
		newDefinition := "this is weeeee"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("updating a non-existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"

		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("deleting an existing value", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test blabla"}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q, given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got error, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word string, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition)
}
