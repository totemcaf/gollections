package repositories

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

const Key1 = "key-1"

type entity struct {
	Id    string
	Value int
}

func TestNew_Is_is_empty(t *testing.T) {
	repo := newRepo()

	count := repo.TotalCount()

	assert.Equal(t, 0, count)
}

func newRepo() *InMemoryRepository[string, *entity] {
	return &InMemoryRepository[string, *entity]{
		GetKey: func(e *entity) string { return e.Id },
	}
}

func Test_added_element_it_is_found(t *testing.T) {
	repo := newRepo()

	toStore := &entity{Key1, 42}

	_, _ = repo.Create(toStore)

	entity, err := repo.FindById(Key1)
	assert.Nil(t, err)
	assert.Equal(t, toStore, entity)
}

func Test_added_element_is_returned(t *testing.T) {
	repo := newRepo()

	toStore := &entity{Key1, 42}

	stored, err := repo.Create(toStore)

	if !assert.Nil(t, err) {
		assert.Equal(t, toStore, stored)
	}
}

func Test_added_element_is_counted(t *testing.T) {
	repo := newRepo()

	toStore := &entity{Key1, 42}
	_, _ = repo.Create(toStore)

	count := repo.TotalCount()

	assert.Equal(t, 1, count)
}

func Test_added_elements_are_counted(t *testing.T) {
	repo := newRepo()

	for idx := 1; idx <= 10; idx++ {
		_, _ = repo.Create(&entity{fmt.Sprintf("key-%d", idx), 42*1000 + idx})
	}

	count := repo.TotalCount()

	assert.Equal(t, 10, count)
}

func Test_cannot_add_element_with_same_key(t *testing.T) {
	repo := newRepo()
	_, _ = repo.Create(&entity{Key1, 42})

	_, err := repo.Create(&entity{Key1, 42})

	assert.ErrorIs(t, err, duplicateKey)
}

func Test_cannot_add_element_with_empty_key(t *testing.T) {
	repo := newRepo()

	_, err := repo.Create(&entity{"", 42})

	assert.ErrorIs(t, err, invalidKey)
}

func Test_Update_returns_replaced_entity(t *testing.T) {
	repo := newRepo()
	_, _ = repo.Create(&entity{Key1, 42})

	entity, err := repo.Update(&entity{Key1, 4242})

	assert.Nil(t, err)
	assert.Equal(t, 4242, entity.Value)
}

func Test_Update_replace_entity(t *testing.T) {
	repo := newRepo()
	_, _ = repo.Create(&entity{Key1, 42})
	_, _ = repo.Update(&entity{Key1, 4242})

	entity, err := repo.FindById(Key1)

	assert.Nil(t, err)
	assert.Equal(t, Key1, entity.Id)
	assert.Equal(t, 4242, entity.Value)
}

func Test_FindBy_founds_entities(t *testing.T) {
	repo := newRepo()
	_, _ = repo.Create(&entity{"a-key-001", 4200})
	_, _ = repo.Create(&entity{"a-key-002", 42})
	_, _ = repo.Create(&entity{"a-key-003", 35})
	_, _ = repo.Create(&entity{"a-key-004", 179})

	greaterThan100 := func(e *entity) bool { return e.Value > 100 }

	entities := repo.FindBy(greaterThan100)

	assert.Len(t, entities, 2)

	expected := []*entity{
		{"a-key-001", 4200},
		{"a-key-004", 179},
	}

	allEquals(t, expected, entities)
}

func allEquals[T any](t *testing.T, expected []T, entities []T) bool {
	if !assert.Len(t, entities, len(expected)) {
		return false
	}

	for idx, e := range entities {
		if !assert.Equal(t, expected[idx], e) {
			return false
		}
	}
	return true
}

func Test_Delete_reduce_count(t *testing.T) {
	repo := newRepo()

	for idx := 1; idx <= 10; idx++ {
		_, _ = repo.Create(&entity{fmt.Sprintf("key-%d", idx), 42*1000 + idx})
	}

	previousCount := repo.TotalCount()

	_ = repo.Delete("key-3")
	_ = repo.Delete("key-5")

	count := repo.TotalCount()

	assert.Equal(t, previousCount-2, count)
}
