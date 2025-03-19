package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// handlers

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		http.NotFound(w, r)
		return
	}
	msg := fmt.Sprintf("Show snippet with id: %d", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("snippet create"))
}

func snippetCreatePost(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("save a new snippet ..."))
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{id}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
