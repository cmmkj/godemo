package main

import (
	"io"
)

type FileSystemStore struct {
	database io.Reader
}

func (f *FileSystemStore) GetLeague() []Player {
	// var league []Player
	// json.NewDecoder(f.database).Decode(&league)
	league, _ := NewLeague(f.database)
	return league
}
