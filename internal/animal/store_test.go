package animal

import "testing"

func TestNewStore(t *testing.T) {

	store := NewStore()

	if store == nil {
		t.Fatal("expected store to be initialized")
	}

	if len(store.animals) != 0 {
		t.Fatal("expected store to start empty")
	}
}

func TestStoreSaveAndGet(t *testing.T) {

	store := NewStore()

	animal := Animal{
		ID:     "123",
		Type:   "cat",
		Gender: "F",
		Name:   "Mira",
		Weight: 3.3,
	}

	store.Save(animal)

	result, ok := store.Get("123")

	if !ok {
		t.Fatal("expected animal to exist")
	}

	if result.Name != "Mira" {
		t.Fatalf("expected name Mira but got %s", result.Name)
	}
}

func TestStoreGetAll(t *testing.T) {

	store := NewStore()

	store.Save(Animal{ID: "1", Name: "Mira"})
	store.Save(Animal{ID: "2", Name: "Pommes"})

	animals := store.GetAll()

	if len(animals) != 2 {
		t.Fatalf("expected 2 animals but got %d", len(animals))
	}
}

func TestStoreDelete(t *testing.T) {

	store := NewStore()

	animal := Animal{
		ID:   "1",
		Name: "Mira",
	}

	store.Save(animal)

	deleted, ok := store.Delete("1")

	if !ok {
		t.Fatal("expected delete to succeed")
	}

	if deleted.ID != "1" {
		t.Fatal("unexpected deleted animal")
	}

	_, ok = store.Get("1")

	if ok {
		t.Fatal("animal should be deleted")
	}
}
