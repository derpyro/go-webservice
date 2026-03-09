package main

import (
	"encoding/json"
	"net/http"
	"strings"
	"sync"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Animal struct {
	Id     string  `json:"id"`
	Type   string  `json:"type"`
	Gender string  `json:"gender"`
	Name   string  `json:"name"`
	Weight float32 `json:"weight"`
}

type ValidationError struct {
	ErrorCode int
	Message   string
}

var (
	animals = []Animal{}
	mutex   sync.Mutex
)

func validate(animal Animal) ValidationError {
	var validationError ValidationError
	validationError.ErrorCode = 0
	validationError.Message = ""

	if animal.Gender != "M" && animal.Gender != "F" && animal.Gender != "X" {
		validationError.ErrorCode = 400
		validationError.Message = "animal must have a gender of M, F or X"
	} else if strings.TrimSpace(animal.Name) == "" {
		validationError.ErrorCode = 400
		validationError.Message = "animal must have a name"
	} else if animal.Weight <= 0 {
		validationError.ErrorCode = 400
		validationError.Message = "animal must have a positive weight"
	} else if strings.TrimSpace(animal.Type) == "" {
		validationError.ErrorCode = 400
		validationError.Message = "animal must have a type"
	}

	return validationError
}

func getAllAnimalsHandler(writer http.ResponseWriter, request *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(animals)
}

func createNewAnimalHandler(writer http.ResponseWriter, request *http.Request) {
	var animal Animal

	err := json.NewDecoder(request.Body).Decode(&animal)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	validationError := validate(animal)
	if validationError.ErrorCode != 0 {
		http.Error(writer, validationError.Message, validationError.ErrorCode)
		return
	}

	mutex.Lock()
	animal.Id = uuid.New().String()
	animals = append(animals, animal)
	mutex.Unlock()

	writer.Header().Set("Location", "/animal/"+animal.Id)
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(animal)
}

func updateAnimalHandler(writer http.ResponseWriter, request *http.Request) {
	var animal Animal

	err := json.NewDecoder(request.Body).Decode(&animal)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	if animal.Id == "" {
		http.Error(writer, "an id has to be provided to edit an animal", http.StatusBadRequest)
		return
	}
	validationError := validate(animal)
	if validationError.ErrorCode != 0 {
		http.Error(writer, validationError.Message, validationError.ErrorCode)
		return
	}

	mutex.Lock()

	var foundMatching bool = false
	for index, element := range animals {
		if element.Id == animal.Id {
			animals[index] = animal
			foundMatching = true
			break
		}
	}
	mutex.Unlock()
	if !foundMatching {
		http.Error(writer, "no animal with id "+animal.Id+" found", http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(animal)
}

func deleteAnimalHandler(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	mutex.Lock()

	var animal Animal
	var foundMatching bool = false
	for index, element := range animals {
		if element.Id == id {
			animal = animals[index]
			animals = append(animals[:index], animals[index+1:]...)
			foundMatching = true
			break
		}
	}
	mutex.Unlock()
	if !foundMatching {
		http.Error(writer, "no animal with id "+id+" found", http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(animal)
}

func getAnimalDetailHandler(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	mutex.Lock()

	var animal Animal
	var foundMatching bool = false
	for index, element := range animals {
		if element.Id == id {
			animal = animals[index]
			foundMatching = true
			break
		}
	}
	mutex.Unlock()
	if !foundMatching {
		http.Error(writer, "no animal with id "+id+" found", http.StatusNotFound)
		return
	}
	writer.WriteHeader(http.StatusOK)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(animal)
}

func main() {
	router := chi.NewRouter()

	router.Get("/animal", getAllAnimalsHandler)
	router.Post("/animal", createNewAnimalHandler)
	router.Put("/animal", updateAnimalHandler)

	router.Delete("/animal/{id}", deleteAnimalHandler)
	router.Get("/animal/{id}", getAnimalDetailHandler)

	http.ListenAndServe(":8080", router)
}
