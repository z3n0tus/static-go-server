package main

import (
	"net/http"
	"encoding/json"
	"fmt"
)

type author struct {
	Name string
	FavouriteBook string
}

var favouriteAuthors [6]author = [6]author{
	{ Name: "N. K. Jemisin", FavouriteBook: "The Fifth Season" },
	{ Name: "Noam Chomsky", FavouriteBook: "Who Rules the World?" },
	{ Name: "Nnedi Okorafor", FavouriteBook: "Who Fears Death?" },
	{ Name: "Robert Jordan", FavouriteBook: "The Dragon Reborn" },
	{ Name: "Frederick Forsyth", FavouriteBook: "The Day of the Jackal" },
	{ Name: "Octavia Butler", FavouriteBook: "Kindred" },
}

func main() {
	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/*", fs)

	http.HandleFunc("/", spaHandler);

	http.HandleFunc("/authors", authorHandler);

	http.ListenAndServe(":8080", nil)
}

// This returns a list of authors.
func authorHandler(responseWriter http.ResponseWriter, request *http.Request) {
	favouriteAuthorsJSON, err := json.Marshal(favouriteAuthors)

	if err != nil {
		panic("Could not marshal json.")
	}

	fmt.Fprintf(responseWriter, string(favouriteAuthorsJSON))
}

func spaHandler(responseWriter http.ResponseWriter, request *http.Request) {
	http.ServeFile(responseWriter, request, "./static/index.html")
}