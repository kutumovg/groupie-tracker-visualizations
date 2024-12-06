package parsers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
)

type Relations struct {
	// ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

func GetRelations(id string) Relations {
	result, err := http.Get(RelationURL + "/" + id)
	if err != nil {
		log.Fatal(err)
	}
	defer result.Body.Close()

	body, err := io.ReadAll(result.Body)
	if err != nil {
		log.Fatal(err)
	}

	var relations Relations
	err = json.Unmarshal(body, &relations)
	if err != nil {
		log.Fatal(err)
	}

	return relations
}
