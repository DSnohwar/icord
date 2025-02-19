package main

import (
	"fmt"
	"log"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to icord Backend"))
}

func main() {
	http.HandleFunc("/", homeHandler)

	port := "8080"
	fmt.Println("Backend running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
