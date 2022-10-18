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
	w.Write([]byte("Hello from Snipperbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	/* Extract the value of the id parameter from the query string
	convert it to an integer using the strconv.Atoi() function.
	be converted to an integer, or the value is less than 1, we return
	a not found response
	*/
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}
	/* Use the fmt.Fprintf() function to interopolate the id value
	and write it to the http.ResponseWriter
	*/
	fmt.Fprintf(w, "Display a specific snippet with ID %d...", id)
}

func snippetCreate(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.Header().Set("Allow", http.MethodPost)
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(` {"name":"Malik"}`))
	w.Write([]byte("Create a new snippet..."))
}

func main() {
	// register the handlers with thier corresponding URLS
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)

	log.Println("Starting Server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
