package router

import (
	"authorizater/middleware"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// Router is exported and used in main.go
func Router() *mux.Router {

	router := mux.NewRouter()

	router.HandleFunc("/account/{id}", middleware.AccountIDGet).Methods("GET", "OPTIONS")
	router.HandleFunc("/accounts", middleware.AccountsGet).Methods("GET", "OPTIONS")
	router.HandleFunc("/account", middleware.AccountPost).Methods("POST", "OPTIONS")
	router.HandleFunc("/account/{id}", middleware.AccountIDPut).Methods("PUT", "OPTIONS")
	router.HandleFunc("/account/{id}", middleware.AccountIDDelete).Methods("DELETE", "OPTIONS")
	router.HandleFunc("/customer/account", middleware.CustomerAccountGet).Methods("GET", "OPTIONS")
	router.HandleFunc("/customer/account", middleware.CustomerAccountGet).Queries("username", "{username}").Methods("GET", "OPTIONS")

	sh := http.StripPrefix("/doc", http.FileServer(http.Dir("./docs/")))
	router.PathPrefix("/doc/").Handler(sh)

	router.PathPrefix("/swagger").Handler(httpSwagger.Handler(
		httpSwagger.URL("/doc/swagger.json"), //The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))

	return router
}
