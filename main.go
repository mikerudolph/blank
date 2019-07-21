package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Success")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Starting server")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
