package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func articlesReturn(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Movie 1", Desc: "This movie is not a movie", Content: "This movie has no content in it"},
		Article{"Movie 2", "Movie 2 desc", "Movie 2 content"},
	}

	json.NewEncoder(w).Encode(articles)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "EndPointHit")
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/articles", articlesReturn)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
