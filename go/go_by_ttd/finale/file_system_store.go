package main

import (
	"encoding/json"
	"io"
	"log"
)

type FileSystemPlayerStore struct {
	database io.Writer
	league   League
}

func NewFileSystemPlayerStore(database io.ReadWriteSeeker) *FileSystemPlayerStore {
	database.Seek(0, 0)
	var league League
	err := json.NewDecoder(database).Decode(&league)
	if err != nil {
		log.Fatal("Json err:", err)
	}

	return &FileSystemPlayerStore{
		database: &tape{database},
		league:   league,
	}
}

func (f *FileSystemPlayerStore) GetLeague() League {
	return f.league
}

func (f *FileSystemPlayerStore) GetPlayerScore(name string) int {
	player := f.league.Find(name)

	if player != nil {
		return player.Wins
	}

	return 0
}

func (f *FileSystemPlayerStore) RecordWin(name string) {
	player := f.league.Find(name)
	if player != nil {
		player.Wins++

	} else {
		f.league = append(f.league, Player{name, 1})

	}

	json.NewEncoder(f.database).Encode(f.league)
}
