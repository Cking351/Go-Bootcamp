package main

import (
	"fmt"
	"net/http"
	"encoding/json"
	"log"
)

type GameOfThrones []struct {
	URL         string   `json:"url"`
	Name        string   `json:"name"`
	Gender      string   `json:"gender"`
	Culture     string   `json:"culture"`
	Born        string   `json:"born"`
	Died        string   `json:"died"`
	Titles      []string `json:"titles"`
	Aliases     []string `json:"aliases"`
	Father      string   `json:"father"`
	Mother      string   `json:"mother"`
	Spouse      string   `json:"spouse"`
	Allegiances []string `json:"allegiances"`
	Books       []string `json:"books"`
	PovBooks    []string `json:"povBooks"`
	TvSeries    []string `json:"tvSeries"`
	PlayedBy    []string `json:"playedBy"`
}

func main() {

	url := "https://anapioficeandfire.com/api/characters/583"

	method := "GET"

	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		log.Panic(err)
	}

	response, err := client.Do(req)
	if err != nil {
		log.Panic(err)
	}

	defer response.Body.Close()


	var data GameOfThrones
	if err := json.NewDecoder(response.Body).Decode(&data); err != nil {
		log.Println(err)
	}

	fmt.Println(data)
}


