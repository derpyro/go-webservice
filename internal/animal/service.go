package animal

import "github.com/google/uuid"

type Service struct {
	store *Store
}

func NewService(store *Store) *Service {
	return &Service{store: store}
}

func (service *Service) GetAll() []Animal {
	return service.store.GetAll()
}

func (service *Service) Get(id string) (Animal, bool) {
	return service.store.Get(id)
}

func (service *Service) Create(animal Animal) Animal {
	animal.ID = uuid.New().String()
	service.store.Save(animal)
	return animal
}

func (service *Service) Update(id string, animal Animal) (Animal, bool) {

	_, exists := service.store.Get(id)
	if !exists {
		return Animal{}, false
	}

	animal.ID = id
	service.store.Save(animal)

	return animal, true
}

func (service *Service) Delete(id string) (Animal, bool) {
	return service.store.Delete(id)
}
