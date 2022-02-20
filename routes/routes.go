package routes

import (
	"go-rest-api/controllers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {

	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalities", controllers.AllPersonalities).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetPersonalityById).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.CreatePersonality).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.DeletePersonalityById).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}", controllers.EditPersonality).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8000", r))
}
