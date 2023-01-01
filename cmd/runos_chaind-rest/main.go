package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"runos_chain/cmd/runos_chaind-rest/internal/handlers"
)

func main() {
	// mux handlers
	fmt.Println("Started")
	handlers.Init()
	r := mux.NewRouter()
	r.HandleFunc("/getPort", handlers.GetPortHandler).Methods("GET")
	r.HandleFunc("/setPort", handlers.SetPortHandler).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(handlers.Config.Port, r))
}
