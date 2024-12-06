package parsers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Artist struct {
	// ID           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
	// Locations    []Location `json:"-"`
	// ConcertDates Date     `json:"concertDates"`
	Relations map[string][]string `json:"-"`
}

func GetArtist(id string) Artist {
	result, err := http.Get(ArtistsURL + "/" + id)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}

	var artist Artist
	err = json.Unmarshal(body, &artist)
	if err != nil {
		log.Fatal(err)
	}

	return artist
}