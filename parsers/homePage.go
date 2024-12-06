package parsers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Artists struct {
	ID    int    `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}

func GetArtists() []Artists {
	result, err := http.Get(ArtistsURL)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}

	var artists []Artists
	err = json.Unmarshal(body, &artists)
	if err != nil {
		log.Fatal(err)
	}

	return artists
}
