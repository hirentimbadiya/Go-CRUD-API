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
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func main() {
	//* append some movies in movies variable of Movies Struct type
	movies = append(movies, Movie{ID: "1", Isbn: "548329", Title: "Krrish", Director: &Director{FirstName: "Rakesh", LastName: "Roshan"}})
	movies = append(movies, Movie{ID: "2", Isbn: "677438", Title: "Iron Man 3", Director: &Director{FirstName: "Shane", LastName: "Black"}})
	movies = append(movies, Movie{ID: "3", Isbn: "234530", Title: "Pushpa - The Rise", Director: &Director{FirstName: "Sukumar", LastName: "Bandreddi"}})
	movies = append(movies, Movie{ID: "4", Isbn: "765983", Title: "Intersteller", Director: &Director{FirstName: "Christopher", LastName: "Nolan"}})
	movies = append(movies, Movie{ID: "5", Isbn: "567392", Title: "Theri", Director: &Director{FirstName: "Atlee", LastName: "Kumar"}})
	movies = append(movies, Movie{ID: "6", Isbn: "894321", Title: "Spiderman", Director: &Director{FirstName: "Sam", LastName: "Raimi"}})
	movies = append(movies, Movie{ID: "7", Isbn: "695556", Title: "Guardians Of The Galaxy", Director: &Director{FirstName: "James", LastName: "Gunn"}})
	movies = append(movies, Movie{ID: "8", Isbn: "675433", Title: "The Amazing Spiderman", Director: &Director{FirstName: "Marc", LastName: "Web"}})
	movies = append(movies, Movie{ID: "9", Isbn: "894532", Title: "Singham", Director: &Director{FirstName: "Rohit", LastName: "Shetty"}})

	r := mux.NewRouter()

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server at Port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	//* set the header as json application.
	w.Header().Set("Content-Type", "application/json")

	// Short way to display in json
	json.NewEncoder(w).Encode(movies)

	/**  //!ANOTHER WAY TO DISPLAY IN JSON
		jsonData, err := json.Marshal(movies)

		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
	    }

		w.Write(jsonData)
	*/
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// create a new movie
	var movie Movie
	// decode the request body into movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	// set the id of the movie
	movie.ID = strconv.Itoa(rand.Intn(1000000000))

	// append the movie to the movies slice
	movies = append(movies, movie)
	// return the movie in json format
	json.NewEncoder(w).Encode(movie)
}

// TODO: it is not an ideal way to update but we are not using
// TODO: any database so we are handling this way.
func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			//! delete that movie
			movies = append(movies[:index], movies[index+1:]...)
			//! create a new movie with given updates
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = item.ID

			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
		}
	}
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// get the id (as params) from the request
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}
