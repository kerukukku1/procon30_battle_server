package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kerukukku1/procon30_battle_server/internal/pkg/database"
)

func entryHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>game room</h1>"))
	vars := mux.Vars(r)
	database.SaveRoomState(vars["roomid"])
}

func main() {
	// http.HandleFunc("/entry", entryHandler)
	// http.ListenAndServe(":8080", nil)
	router := mux.NewRouter()
	router.HandleFunc("/api/game/{roomid}", entryHandler).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", router))
}
