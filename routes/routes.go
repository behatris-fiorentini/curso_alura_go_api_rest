package routes

import (
	"net/http"

	"github.com/behatris-fiorentini/curso_alura_go_api_rest/controllers"
	"github.com/behatris-fiorentini/curso_alura_go_api_rest/middlewares"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middlewares.ContentTypeMiddleware)

	r.HandleFunc("/", controllers.Home).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.GetAll).Methods("GET")
	r.HandleFunc("/api/personalities/{id}", controllers.GetByID).Methods("GET")
	r.HandleFunc("/api/personalities", controllers.Create).Methods("POST")
	r.HandleFunc("/api/personalities/{id}", controllers.Delete).Methods("DELETE")
	r.HandleFunc("/api/personalities/{id}", controllers.Update).Methods("PATCH")

	startServer(r)
}

func startServer(r *mux.Router) {
	err := http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"http://localhost:3000"}))(r))
	if err != nil {
		panic(err.Error())
	}
}
