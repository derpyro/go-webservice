package restclient

import (
	"io"
	"log"
	"net/http"
)

func CallGetAnimals() ([]byte, error) {
	url := "http://localhost:8080/animals"
	response, error := http.Get(url)
	if error != nil {
		log.Fatal(error)
		return []byte{}, error
	}
	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatal("Status was not OK but " + response.Status)
		return []byte{}, nil
	}
	data, error := io.ReadAll(response.Body)
	if error != nil {
		log.Fatal(error)
		return []byte{}, error
	}
	return data, nil
}
