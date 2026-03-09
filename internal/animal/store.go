package animal

import "sync"

type Store struct {
	mutex   sync.Mutex
	animals map[string]Animal
}

func NewStore() *Store {
	return &Store{
		animals: map[string]Animal{},
	}
}

func (store *Store) GetAll() []Animal {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	result := make([]Animal, 0, len(store.animals))

	for _, animal := range store.animals {
		result = append(result, animal)
	}

	return result
}

func (store *Store) Get(id string) (Animal, bool) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	animal, ok := store.animals[id]
	return animal, ok
}

func (store *Store) Save(animal Animal) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	store.animals[animal.ID] = animal
}

func (store *Store) Delete(id string) (Animal, bool) {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	animal, ok := store.animals[id]

	if ok {
		delete(store.animals, id)
	}

	return animal, ok
}
