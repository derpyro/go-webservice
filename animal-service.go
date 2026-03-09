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

	writer.Header().Set("Location", "/animals/"+animal.Id)
	writer.WriteHeader(http.StatusCreated)
	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(animal)
}

func updateAnimalHandler(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
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

	for index, element := range animals {
		if element.Id == id {
			animals[index] = animal
			mutex.Unlock()
			writer.WriteHeader(http.StatusOK)
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(animal)
			return
		}
	}
	mutex.Unlock()
	http.Error(writer, "no animal with id "+id+" found", http.StatusNotFound)
}

func deleteAnimalHandler(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	mutex.Lock()

	var animal Animal
	for index, element := range animals {
		if element.Id == id {
			animal = animals[index]
			animals = append(animals[:index], animals[index+1:]...)
			mutex.Unlock()
			writer.WriteHeader(http.StatusOK)
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(animal)
			return
		}
	}
	mutex.Unlock()
	http.Error(writer, "no animal with id "+id+" found", http.StatusNotFound)
}

func getAnimalDetailHandler(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	mutex.Lock()

	var animal Animal
	for index, element := range animals {
		if element.Id == id {
			animal = animals[index]
			mutex.Unlock()
			writer.WriteHeader(http.StatusOK)
			writer.Header().Set("Content-Type", "application/json")
			json.NewEncoder(writer).Encode(animal)
			return
		}
	}
	mutex.Unlock()
	http.Error(writer, "no animal with id "+id+" found", http.StatusNotFound)
}

func main() {
	router := chi.NewRouter()

	router.Get("/animals", getAllAnimalsHandler)
	router.Post("/animals", createNewAnimalHandler)

	router.Put("/animals/{id}", updateAnimalHandler)
	router.Delete("/animals/{id}", deleteAnimalHandler)
	router.Get("/animals/{id}", getAnimalDetailHandler)

	http.ListenAndServe(":8080", router)
}
