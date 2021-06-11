package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"pokedex-api-v1/models"

	"github.com/gorilla/mux"
)

// GetAllPokemons godoc
// @Summary Get a list of pokemons
// @Description Get a list of pokemons
// @Tags List of Pokemons
// @Accept  json
// @Produce  json
// @Param limit query int true "Number of pokemons"
// @Param offset query int true "Number of pokemons to exclude"
// @Success 200 {array} models.Pokemons
// @Router /api/v1/pokemons [get]
func GetAllPokemons(w http.ResponseWriter, r *http.Request) {
	limit := r.FormValue("limit")
	offset := r.FormValue("offset")
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=" + limit + "&offset=" + offset)

	if err != nil {
		fmt.Print("Error " + err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject models.Results
	json.Unmarshal(responseData, &responseObject)

	pokemonList := make([]models.Pokemons, len(responseObject.Pokemon))

	for i := 0; i < len(responseObject.Pokemon); i++ {
		pokemonList[i] = getPokemonByName(responseObject.Pokemon[i].Name)
	}

	json.NewEncoder(w).Encode(pokemonList)
}

func getPokemonByName(name string) models.Pokemons {
	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + name)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObj models.Pokemons
	json.Unmarshal(responseData, &responseObj)
	return responseObj
}

// GetPokemonInfo godoc
// @Summary Get pokemon's information
// @Description Get pokemon's information
// @Tags Pokemon's Info
// @Accept  json
// @Produce  json
// @Param name path string true "pokemon name"
// @Success 200 {array} models.PokemonInfo
// @Router /api/v1/pokemon/{name} [get]
func GetPokemonInfo(w http.ResponseWriter, r *http.Request) {
	name := mux.Vars(r)["name"]

	response, err := http.Get("https://pokeapi.co/api/v2/pokemon/" + name)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var responseObj models.PokemonInfo
	json.Unmarshal(responseData, &responseObj)
	json.NewEncoder(w).Encode(responseObj)
}
