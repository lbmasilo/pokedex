package main

import (
	"net/http"
	"os"
	_ "pokedex-api-v1/docs"
	"pokedex-api-v1/handler"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// @title Swagger Pokedex API
// @version 1.0
// @description Pokedex api with golang.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
func main() {
	setVariables()
	requestHandler()
}

func requestHandler() {
	router := mux.NewRouter()
	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET"})
	origins := handlers.AllowedOrigins([]string{"*"})

	// Read
	router.HandleFunc(os.Getenv("pokemonListEndPoint"), handler.GetAllPokemons).Methods("GET")
	router.HandleFunc(os.Getenv("pokekomByNameEndPoint"), handler.GetPokemonInfo).Methods("GET")
	router.PathPrefix(os.Getenv("swaggerUrl")).Handler(httpSwagger.WrapHandler)

	http.ListenAndServe(os.Getenv("port"), handlers.CORS(headers, methods, origins)(router))
}

func setVariables() {
	os.Setenv("port", ":8080")
	os.Setenv("pokemonListEndPoint", "/api/v1/pokemons")
	os.Setenv("pokekomByNameEndPoint", "/api/v1/pokemon/{name}")
	os.Setenv("swaggerUrl", "/swagger/")
}
