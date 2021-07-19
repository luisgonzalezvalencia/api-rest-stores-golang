package main

import (
	"log"
	"net/http"
	"time"

	"fake.com/apirest/config"
	"github.com/gorilla/mux"
)

// list of tasks example

func main() {
	db, err := config.GetDB()
	if err != nil {
		log.Printf("Error with database" + err.Error())
		return
	} else {
		err = db.Ping()
		if err != nil {
			log.Printf("Error making db connection. Please check your configuration. Error: " + err.Error())
			return
		}

		//define routes
		router := mux.NewRouter()
		setupRoutesForStores(router)
		// .. here you can define more routes
		// ...
		// for example setupRoutesForGenres(router)

		// Setup and start server
		port := ":8000"

		server := &http.Server{
			Handler: router,
			Addr:    port,
			// timeouts so the server never waits forever...
			WriteTimeout: 15 * time.Second,
			ReadTimeout:  15 * time.Second,
		}
		log.Printf("Server started at %s", port)
		log.Fatal(server.ListenAndServe())
	}

}
