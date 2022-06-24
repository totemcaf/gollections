package repositories

import (
	"errors"
	"sync"

	"github.com/totemcaf/gollections/lists"
	"github.com/totemcaf/gollections/maps"
	"github.com/totemcaf/gollections/types"
)

var invalidKey = errors.New("invalid key, nil")
var notFound = errors.New("not found")
var duplicateKey = errors.New("key is duplicated")

type InMemoryRepository[Key comparable, Entity any] struct {
	elementsById  map[Key]Entity
	emptyKey      Key
	lock          sync.RWMutex
	AllowEmptyKey bool
	GetKey        func(Entity) Key
}

func (r *InMemoryRepository[Key, Entity]) init() {
	if r.elementsById == nil {
		r.elementsById = make(map[Key]Entity, 16)
	}
}

func (r *InMemoryRepository[Key, Entity]) Create(entity Entity) (Entity, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.init()

	key := r.GetKey(entity)
	if !r.AllowEmptyKey && key == r.emptyKey {
		return entity, invalidKey
	}

	_, alreadyInMap := r.elementsById[key]
	if alreadyInMap {
		return entity, duplicateKey
	}

	r.elementsById[key] = entity
	return entity, nil
}

func (r *InMemoryRepository[Key, Entity]) Update(entity Entity) (Entity, error) {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.init()

	key := r.GetKey(entity)
	if !r.AllowEmptyKey && key == r.emptyKey {
		return entity, invalidKey
	}

	_, alreadyInMap := r.elementsById[key]
	if !alreadyInMap {
		return entity, notFound
	}

	r.elementsById[key] = entity
	return entity, nil
}

func (r *InMemoryRepository[Key, Entity]) Delete(key Key) error {
	r.lock.Lock()
	defer r.lock.Unlock()
	r.init()

	if !r.AllowEmptyKey && key == r.emptyKey {
		return invalidKey
	}

	_, alreadyInMap := r.elementsById[key]

	if !alreadyInMap {
		return notFound
	}

	delete(r.elementsById, key)

	return nil
}

func (r *InMemoryRepository[Key, Entity]) FindById(key Key) (Entity, error) {
	r.lock.RLock()
	defer r.lock.RUnlock()
	r.init()

	if !r.AllowEmptyKey && key == r.emptyKey {
		var entity Entity
		return entity, invalidKey
	}

	if entity, found := r.elementsById[key]; found {
		return entity, nil
	}
	var empty Entity
	return empty, notFound
}

func (r *InMemoryRepository[Key, Entity]) FindBy(predicate types.Predicate[Entity]) []Entity {
	r.lock.RLock()
	defer r.lock.RUnlock()

	// This is not the most efficient way to do it, but this repository is meant for tests
	return lists.Filter(maps.Values(r.elementsById), predicate)
}

func (r *InMemoryRepository[Key, Entity]) TotalCount() int {
	return len(r.elementsById)
}
