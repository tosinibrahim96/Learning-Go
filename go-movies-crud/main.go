package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}

var movies []Movie

func getMovies(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	json.NewEncoder(res).Encode(movies)
}

func deleteMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(res).Encode(movies)
}

func getMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for _, item := range movies {

		if item.ID == params["id"] {
			json.NewEncoder(res).Encode(item)
			return
		}
	}
}

func createMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	var movie Movie

	_ = json.NewDecoder(req.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)

	json.NewEncoder(res).Encode(movie)
}

func updateMovie(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")

	params := mux.Vars(req)

	for index, item := range movies {

		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(req.Body).Decode(&movie)
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(res).Encode(movie)
			return
		}
	}

}

func main() {

	r := mux.NewRouter()

	movies = append(
		movies,
		Movie{
			ID:       "1",
			Isbn:     "438227",
			Title:    "Movie 1",
			Director: &Director{Firstname: "John", Lastname: "Doe"},
		},
		Movie{
			ID:       "2",
			Isbn:     "438228",
			Title:    "Movie 2",
			Director: &Director{Firstname: "Jane", Lastname: "Doe"},
		},
	)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at PORT 8001")
	log.Fatal(http.ListenAndServe(":8001", r))

}
