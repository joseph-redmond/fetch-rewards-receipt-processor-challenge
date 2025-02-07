package db

import (
	"fmt"
	"sync"
)

type Entity interface {
	ID() string
}

type Store[K Entity] struct {
	data map[string]K
	mu   sync.RWMutex
}

// Function that creates a new dataset of your chosen type
func NewStore[K Entity]() *Store[K] {
	return &Store[K]{
		data: make(map[string]K),
	}
}

// saves a new entity to the dataset
func (store *Store[K]) Save(entity K) (K, error) {
	store.mu.Lock()
	defer store.mu.Unlock()

	id := entity.ID()
	if _, exists := store.data[id]; exists {
		var empty K
		return empty, fmt.Errorf("entity with ID %s already exists", id)
	}

	store.data[id] = entity

	return entity, nil
}

// finds a entity in the dataset by it's id
func (store *Store[K]) FindById(id string) (K, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	entity, exists := store.data[id]
	if !exists {
		var empty K
		return empty, fmt.Errorf("entity with ID %s was not found", id)
	}

	return entity, nil
}

// updates an entity in the dataset
func (store *Store[K]) Update(entity K) (K, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()

	entity, exists := store.data[entity.ID()]
	if !exists {
		var empty K
		return empty, fmt.Errorf("entity with ID %s was not found", entity.ID())
	}
	return store.Save(entity)
}

// deletes an entity from the dataset
func (store *Store[K]) DeleteById(id string) error {
	store.mu.Lock()
	defer store.mu.Unlock()

	if _, exists := store.data[id]; !exists {
		return fmt.Errorf("entity with ID %s was not found", id)
	}

	delete(store.data, id)
	return nil
}

// prints out the dataset
func (store *Store[K]) List() []K {
	store.mu.RLock()
	defer store.mu.RUnlock()

	entities := make([]K, 0, len(store.data))
	for _, entity := range store.data {
		entities = append(entities, entity)
	}
	return entities
}

// allows querying the dataset based on a predicate function
func (store *Store[K]) Query(predicate func(K) bool) []K {
	store.mu.RLock()
	defer store.mu.RUnlock()

	var result []K
	for _, entity := range store.data {
		if predicate(entity) {
			result = append(result, entity)
		}
	}
	return result
}
