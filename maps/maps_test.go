package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, got, ErrNotFound)
	})
	
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}

func TestAdd(t *testing.T) {
	t.Run("test add definition", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is a test"
		dictionary.Add(word, definition)

		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("dont add an existing definition", func(t *testing.T) {
		word := "test"
		value := "this is a test"
		dictionary := Dictionary{word: value}

		err := dictionary.Add(word, "this is a test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, value)		

	})

	t.Run("update value", func(t *testing.T) {
		word := "test"
		value := "this is a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, value)

		assertError(t, err, ErrWordDoesNotExist)
		
	})

	t.Run("delete value", func(t *testing.T) {
		word := "test"
		value := "this is a test"
		dictionary := Dictionary{word: value}

		dictionary.Delete(word)

		_, err := dictionary.Search(word)
		assertError(t, err, ErrNotFound)
		
	})
	
}


func assertDefinition(t *testing.T, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word", err)
	}
	assertStrings(t, got, definition)
}