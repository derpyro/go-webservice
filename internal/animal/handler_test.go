package animal

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestGetAnimals(t *testing.T) {

	store := NewStore()
	service := NewService(store)
	handler := NewHandler(service)

	store.Save(Animal{
		ID:   "1",
		Name: "Mira",
	})

	request := httptest.NewRequest(http.MethodGet, "/animals", nil)

	recorder := httptest.NewRecorder()

	handler.getAll(recorder, request)

	if recorder.Code != http.StatusOK {
		t.Fatalf("expected status 200 but got %d", recorder.Code)
	}

	body := recorder.Body.Bytes()

	if len(body) == 0 {
		t.Fatal("expected response body")
	}

	var result []Animal
	json.Unmarshal(body, &result)

	assertEquals(t, "length", "1", strconv.Itoa(len(result)))
	assertEquals(t, "name", "Mira", result[0].Name)
	assertEquals(t, "id", "1", result[0].ID)
}

func assertEquals(t *testing.T, prefix string, expected string, actual string) {
	if expected != actual {
		t.Fatal("unexpected " + prefix + ". Expected '" + expected + "' but was '" + actual + "'.")
	}
}
