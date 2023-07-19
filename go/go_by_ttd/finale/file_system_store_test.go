package main

import (
	"io"
	"log"
	"os"
	"strconv"
	"testing"
)

func TestFileSystem(t *testing.T) {

	t.Run("League from a reader", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name": "John","Wins": 2},
            {"Name": "Joe","Wins": 1}
            ]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)
		got := store.GetLeague()
		want := []Player{
			{"John", 2},
			{"Joe", 1},
		}

		assertLeague(t, got, want)
	})

	t.Run("Get Player Score", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name":"John","Wins":2},
            {"Name":"Joe","Wins":1}
        ]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)
		got := store.GetPlayerScore("John")

		want := 2

		assertScore(t, strconv.Itoa(want), strconv.Itoa(got))

	})

	t.Run("Test RecordWin of player in db", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name":"John","Wins":2},
            {"Name":"Joe","Wins":1}
        ]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)
		store.RecordWin("Joe")

		got := store.GetPlayerScore("Joe")
		want := 2

		assertScore(t, strconv.Itoa(got), strconv.Itoa(want))

	})

	t.Run("Add player to DB", func(t *testing.T) {
		database, cleanDatabase := createTempFile(t, `[
            {"Name":"John","Wins":2},
            {"Name":"Joe","Wins":1}
        ]`)
		defer cleanDatabase()

		store := NewFileSystemPlayerStore(database)
		store.RecordWin("Johny")

		got := store.GetPlayerScore("Johny")
		want := 1

		assertScore(t, strconv.Itoa(got), strconv.Itoa(want))

	})
}

func createTempFile(t testing.TB, initialData string) (io.ReadWriteSeeker, func()) {
	t.Helper()

	tmpFile, err := os.CreateTemp("", "db")
	if err != nil {
		log.Fatal(err)
	}

	tmpFile.Write([]byte(initialData))

	removeFile := func() {
		tmpFile.Close()
		os.Remove(tmpFile.Name())
	}

	return tmpFile, removeFile

}
