package main

import (
	"testing"
)

func Test_map(t *testing.T) {
	//#helpers
	check_expectations := func(t testing.TB, get, want string) {
		t.Helper()
		if get != want {
			t.Errorf("\nget: %v \n want: %v", get, want)
		}

	}

	check_definition := func(t *testing.T, dictionary Dictionary, key, def string) {
		t.Helper()
		word, err := dictionary.Search(key)
		if err != nil {
			t.Fatal("Word not found")
		}

		check_expectations(t, word, def)

	}

	text := "hi mom"
	dictionary := Dictionary{
		"greet": text,
	}

	t.Run("Search in dictionary", func(t *testing.T) {
		get, _ := dictionary.Search("greet")
		check_expectations(t, get, text)
	})

	t.Run("Word not in dictionary", func(t *testing.T) {
		_, err := dictionary.Search("goodbye")
		if err == nil {
			t.Fatal(err)
		}

		check_expectations(t, err.Error(), Err_not_found.Error())
	})

	t.Run("Add word: new word", func(t *testing.T) {
		dictionary := Dictionary{}
		dictionary.Add("key", "unlocks locked objects")

		check_definition(t, dictionary, "key", "unlocks locked objects")

	})

	t.Run("Add word: already exists", func(t *testing.T) {
		dictionary := Dictionary{"x": "hi"}
		key := "w"
		def := "cannot find x"

		err := dictionary.Add(key, def)
		if err == nil {
			check_definition(t, dictionary, key, def)
		}
	})

	t.Run("Update word", func(t *testing.T) {
		key := "x"
		def := "cannot find x"
		dictionary := Dictionary{key: def}

		err := dictionary.Update(key, "lool")
		if err != nil {
			check_definition(t, dictionary, key, "lool")
		}
	})

	t.Run("Delete definition", func(t *testing.T) {
		definition := Dictionary{"greet": "Hello"}
		err := definition.Delete("greet")
		if err == Err_delete {
			t.Fatalf("expected **%v** to be deleted", "greet")
		}

	})

}
