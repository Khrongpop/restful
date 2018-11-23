package main

import (
	"fmt"
	"net/http"
	"os"

	mux "github.com/gorilla/mux"
)

func main() {
	port := os.Getenv("PORT")
	mux := mux.NewRouter()
	mux.HandleFunc("/", homepage)
	mux.HandleFunc("/greets/{name}", greeting)
	fmt.Println("Start Serve http://localhost:8080/ ")
	// http.ListenAndServe(":8080", mux)
	http.ListenAndServe(":"+port, mux)
}

func homepage(w http.ResponseWriter, req *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if req.URL.Path != "/" {
		http.NotFound(w, req)
		return
	}
	w.Header().Set("Content-type", "appliation/json")
	fmt.Fprintf(w, `{"message":"wellcome to homepage"}`)
}

func greeting(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	name := vars["name"]

	w.Header().Set("Content-type", "appliation/json")
	fmt.Fprintf(w, `{"message":"Hello %s"}`, name)
}
