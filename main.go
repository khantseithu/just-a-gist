package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("Hello, World!"))
}

func SnippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		// w.WriteHeader(405)
		// w.Write([]byte("Method not allowed!"))
		//
		// or

		w.Header().Set("Allow", http.MethodPost)

		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Write([]byte("Creating a new snippet"))
}

func ViewSnippet(w http.ResponseWriter, r *http.Request) {
	// check for id
	// if id is not present, show all the snippet
	// if id is present, show the snippet with that id

	id, err := strconv.Atoi(r.URL.Query().Get("id"))

	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	} else {
		fmt.Fprintf(w, "Displaying a snippet with that ID %d", id)
	}

	// w.Write([]byte("Viewing all the snippet..."))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/create", SnippetCreate)
	mux.HandleFunc("/snippet/view", ViewSnippet)

	log.Print("starting server on :4000")

	err := http.ListenAndServe(":4000", mux)

	log.Fatal(err)
}
