package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	log.Println("Got ID", id)
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	msg := fmt.Sprint("Display a specific snippet with ID: %d", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Create a specific snippet..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	log.Print("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
