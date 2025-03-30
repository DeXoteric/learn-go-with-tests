package main

import (
	"fmt"
	"testing"
)

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		if got == nil {
			t.Fatal("expected to get an error")
		}

		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		definition := "this is just a test"

		err := dictionary.Add(word, definition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, definition)
	})

	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{word: definition}
		newDefinition := "new definition"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		word := "test"
		definition := "this is just a test"
		dictionary := Dictionary{}

		err := dictionary.Update(word, definition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{word: "test definition"}

		err := dictionary.Delete(word)

		assertError(t, err, nil)

		_, err = dictionary.Search(word)

		assertError(t, err, ErrNotFound)
	})

	t.Run("non-existing word", func(t *testing.T) {
		word := "test"
		dictionary := Dictionary{}

		err := dictionary.Delete(word)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func assertStrings(t testing.TB, got string, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
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

func BenchmarkSearch(b *testing.B) {
	dictionary := Dictionary{"key": "value"}
	for b.Loop() {
		dictionary.Search("key")
	}
}

func BenchmarkAdd(b *testing.B) {
	for b.Loop() {
		dictionary := Dictionary{}
		dictionary.Add("key", "value")
	}
}

func BenchmarkUpdate(b *testing.B) {
	dictionary := Dictionary{"key": "value"}
	for b.Loop() {
		dictionary.Update("key", "new value")
	}
}

func BenchmarkDelete(b *testing.B) {
	for b.Loop() {
		dictionary := Dictionary{"key": "value"}
		dictionary.Delete("key")
	}
}

func ExampleDictionary_Search() {
	dictionary := Dictionary{"key": "value"}
	result, err := dictionary.Search("key")
	fmt.Println(result)
	fmt.Println(err)
	// Output:
	// value
	// <nil>
}

func ExampleDictionary_Add() {
	dictionary := Dictionary{}
	dictionary.Add("key", "value")
	fmt.Println(dictionary)
	// Output:
	// map[key:value]
}

func ExampleDictionary_Update() {
	dictionary := Dictionary{"key": "value"}
	dictionary.Update("key", "new value")
	fmt.Println(dictionary)
	// Output:
	// map[key:new value]
}

func ExampleDictionary_Delete() {
	dictionary := Dictionary{"key": "value"}
	dictionary.Delete("key")
	fmt.Println(dictionary)
	// Output:
	// map[]
}
