package main

import (
	"net/http"
	"net/http/httptest"
	"pokedex-api-v1/handler"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/api/v1/pokemons", handler.GetAllPokemons).Methods("GET")
	router.HandleFunc("/api/v1/pokemon/{name}", handler.GetPokemonInfo).Methods("GET")
	return router
}
func TestGetAllPokemons(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/pokemons", nil)
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Expected Response")
}
func TestGetPokemonInfo(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/v1/pokemon/{name}", nil)
	response := httptest.NewRecorder()

	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "Expected Response")
}
