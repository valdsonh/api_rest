package routes

import (
	"log"
	"net/http"
	"valdson/api_rest/controllers"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", r))
}
