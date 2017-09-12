package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getNews)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}

func getNews(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "News Project")
}
