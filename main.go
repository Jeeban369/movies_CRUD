package main

import (
	"fmt"
	"log"
	"math/rand"
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

//all the functions
func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
	
}

func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range movies{
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
   		http.Error(w, err.Error(), http.StatusBadRequest)
    	return
	}
	movie.ID = strconv.Itoa(rand.Intn(10000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	//set json content type
	w.Header().Set("Content-Type", "application/json")
	//params
	params := mux.Vars(r)
	//loop over the movies, range
	//delete the movie with the id that i have sent
	//add a new movie - the movie that we send in the body of postman
	for index, item:= range movies{
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			var movie Movie
			if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
   				http.Error(w, err.Error(), http.StatusBadRequest)
    			return
			}
			movie.ID = params["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)


	for index, item := range movies {
    	if item.ID == params["id"] {
    	    movies = append(movies[:index], movies[index+1:]...)
    	    var movie Movie
    	    if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
    	        http.Error(w, err.Error(), http.StatusBadRequest)
    	        return
    	    }
    	    movie.ID = params["id"]
    	    movies = append(movies, movie)
    	    json.NewEncoder(w).Encode(movie)
    	    return
    	}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)

}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// for index, item := range movies{
	// 	if item.ID == params["id"]{
	// 		movies = append(movies[:index], movies[index+1:]...)
	// 		break
	// 	}
	// }
	// json.NewEncoder(w).Encode(movies)
	// http.Error(w, "Movie not found", http.StatusNotFound)

	for index, item := range movies {
    	if item.ID == params["id"] {
    	    movies = append(movies[:index], movies[index+1:]...)
    	    json.NewEncoder(w).Encode(movies)
    	    return
    	}
	}
	http.Error(w, "Movie not found", http.StatusNotFound)

}



func main(){
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn:"4587227", Title:"movie 01", Director: &Director{Firstname:"John", Lastname:"Cameron"}})
	movies = append(movies, Movie{ID: "2", Isbn:"4587228", Title:"movie 02", Director: &Director{Firstname:"Tim", Lastname:"Smith"}})

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}",updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")


	fmt.Println("server is running on port 8000...")
	log.Fatal(http.ListenAndServe(":8000", r))
}
