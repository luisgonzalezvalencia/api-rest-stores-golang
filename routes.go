package main

import (
	"encoding/json"
	"net/http"

	"fake.com/apirest/config"
	"github.com/gorilla/mux"
)

func setupRoutesForStores(router *mux.Router) {
	// First enable CORS. If you don't need cors, comment the next line
	enableCORS(router)

	router.HandleFunc("/stores/{tipotienda}", func(w http.ResponseWriter, r *http.Request) {
		tipoTiendaAsString := mux.Vars(r)["tipotienda"]
		tipoTienda, err := stringToInt64(tipoTiendaAsString)

		if err != nil {
			respondWithError(err, w)
			// We return, so we stop the function flow
			return
		}

		stores, err := getStores(tipoTienda)
		if err == nil {
			respondWithSuccess(stores, w)
		} else {
			respondWithError(err, w)
		}
	}).Methods(http.MethodGet)

	// router.HandleFunc("/videogame/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	idAsString := mux.Vars(r)["id"]
	// 	id, err := stringToInt64(idAsString)
	// 	if err != nil {
	// 		respondWithError(err, w)
	// 		// We return, so we stop the function flow
	// 		return
	// 	}
	// 	videogame, err := getVideoGameById(id)
	// 	if err != nil {
	// 		respondWithError(err, w)
	// 	} else {
	// 		respondWithSuccess(videogame, w)
	// 	}
	// }).Methods(http.MethodGet)

	// router.HandleFunc("/videogame", func(w http.ResponseWriter, r *http.Request) {
	// 	// Declare a var so we can decode json into it
	// 	var videoGame VideoGame
	// 	err := json.NewDecoder(r.Body).Decode(&videoGame)
	// 	if err != nil {
	// 		respondWithError(err, w)
	// 	} else {
	// 		err := createVideoGame(videoGame)
	// 		if err != nil {
	// 			respondWithError(err, w)
	// 		} else {
	// 			respondWithSuccess(true, w)
	// 		}
	// 	}
	// }).Methods(http.MethodPost)

	// router.HandleFunc("/videogame", func(w http.ResponseWriter, r *http.Request) {
	// 	// Declare a var so we can decode json into it
	// 	var videoGame VideoGame
	// 	err := json.NewDecoder(r.Body).Decode(&videoGame)
	// 	if err != nil {
	// 		respondWithError(err, w)
	// 	} else {
	// 		err := updateVideoGame(videoGame)
	// 		if err != nil {
	// 			respondWithError(err, w)
	// 		} else {
	// 			respondWithSuccess(true, w)
	// 		}
	// 	}
	// }).Methods(http.MethodPut)
	// router.HandleFunc("/videogame/{id}", func(w http.ResponseWriter, r *http.Request) {
	// 	idAsString := mux.Vars(r)["id"]
	// 	id, err := stringToInt64(idAsString)
	// 	if err != nil {
	// 		respondWithError(err, w)
	// 		// We return, so we stop the function flow
	// 		return
	// 	}
	// 	err = deleteVideoGame(id)
	// 	if err != nil {
	// 		respondWithError(err, w)
	// 	} else {
	// 		respondWithSuccess(true, w)
	// 	}
	// }).Methods(http.MethodDelete)
}

func enableCORS(router *mux.Router) {
	router.PathPrefix("/").HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", config.AllowedCORSDomain)
	}).Methods(http.MethodOptions)
	router.Use(middlewareCors)
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, req *http.Request) {
			// Just put some headers to allow CORS...
			w.Header().Set("Access-Control-Allow-Origin", config.AllowedCORSDomain)
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			// and call next handler!
			next.ServeHTTP(w, req)
		})
}

// Helper functions for respond with 200 or 500 code
func respondWithError(err error, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(err.Error())
}

func respondWithSuccess(data interface{}, w http.ResponseWriter) {

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}
