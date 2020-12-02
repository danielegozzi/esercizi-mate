package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	mux := http.NewServeMux()

	// mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "Welcome to my website!")
	// })

	mux.HandleFunc("/frac", frac)

	fs := http.FileServer(http.Dir("static/"))
	mux.Handle("/", fs)
	// mux.Handle("/", http.StripPrefix("/static/", fs))

	log.Fatal(http.ListenAndServe(":"+port, mux))
}
