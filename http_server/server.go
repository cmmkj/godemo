package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//type HandlerFunc func(ResponseWriter, *Request)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
	GetLeague() []Player
}

type PlayerServer struct {
	store PlayerStore
	//router *http.ServeMux
	http.Handler
}

type Player struct {
	Name string
	Wins int
}

func NewPlayerServer(store PlayerStore) *PlayerServer {
	p := new(PlayerServer)
	p.store = store
	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	p.Handler = router
	return p
}

// func (p *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	// router := http.NewServeMux()
// 	// router.Handle("/league", http.HandlerFunc(p.leagueHandler))

// 	// router.Handle("/players/", http.HandlerFunc(p.playersHandler))
// 	//p.router.ServeHTTP(w, r)
// }

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	//json.NewEncoder(w).Encode(p.getLeagueTable())
	json.NewEncoder(w).Encode(p.store.GetLeague())
	w.Header().Set("content-type", "application/json")
	//w.WriteHeader(http.StatusOK)
}

// func (p *PlayerServer) getLeagueTable() []Player {
// 	return []Player{
// 		{"Chris", 20},
// 	}
// }

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]
	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {
	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func GetPlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}
	if name == "Floyd" {
		return "10"
	}
	return ""
}
