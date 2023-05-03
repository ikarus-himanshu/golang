package controller

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/himanshukukreja/mongoapis/dbactions"
	"github.com/himanshukukreja/mongoapis/model"
)

//Actual controllers

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contemt-Type", "applicatiom/x-www-form-urlemcode")
	allMovies := dbactions.GetAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contemt-Type", "applicatiom/x-www-form-urlemcode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie model.Netflix
	_ = json.NewDecoder(r.Body).Decode(&movie)
	dbactions.InsertOneMovie(movie)
	json.NewEncoder(w).Encode(movie)
}

func MarkedAsWatched(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contemt-Type", "applicatiom/x-www-form-urlemcode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	params := mux.Vars(r)
	dbactions.UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}

func DeleteAMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contemt-Type", "applicatiom/x-www-form-urlemcode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	params := mux.Vars(r)
	dbactions.DeleteOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}

func DeleteAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Contemt-Type", "applicatiom/x-www-form-urlemcode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	count := dbactions.DeleteAllMovie()
	json.NewEncoder(w).Encode(count)
}
