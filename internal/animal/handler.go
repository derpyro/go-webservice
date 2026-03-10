package animal

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Handler struct {
	service *Service
}

func NewHandler(service *Service) *Handler {
	return &Handler{service: service}
}

func (handler *Handler) Routes(router chi.Router) {
	router.Get("/animals", handler.getAll)
	router.Post("/animals", handler.create)
	router.Get("/animals/{id}", handler.get)
	router.Put("/animals/{id}", handler.update)
	router.Delete("/animals/{id}", handler.delete)
}

func (handler *Handler) getAll(writer http.ResponseWriter, request *http.Request) {

	animals := handler.service.GetAll()

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(animals)
}

func (handler *Handler) create(writer http.ResponseWriter, request *http.Request) {
	var animal Animal

	if err := json.NewDecoder(request.Body).Decode(&animal); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	animal = handler.service.Create(animal)

	writer.Header().Set("Content-Type", "application/json")
	json.NewEncoder(writer).Encode(animal)
}

func (handler *Handler) get(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	animal, ok := handler.service.Get(id)

	if !ok {
		http.Error(writer, "no animal with id "+id+" found", http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode(animal)
}

func (handler *Handler) update(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")
	var animal Animal

	if err := json.NewDecoder(request.Body).Decode(&animal); err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}

	animal, ok := handler.service.Update(id, animal)

	if !ok {
		http.Error(writer, "no animal with id "+id+" found", http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode(animal)
}

func (handler *Handler) delete(writer http.ResponseWriter, request *http.Request) {
	id := chi.URLParam(request, "id")

	animal, ok := handler.service.Delete(id)

	if !ok {
		http.Error(writer, "no animal with id "+id+" found", http.StatusNotFound)
		return
	}

	json.NewEncoder(writer).Encode(animal)
}
