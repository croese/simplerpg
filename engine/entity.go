package engine

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

// Entity represents any game entity
type Entity interface {
	Id() string
}

// Item is a specific type of Entity representing
// an in-game item
type Item struct {
	id          string
	Name        string
	Description string
}

// Id returns the entity ID
func (i *Item) Id() string { return i.id }

// EntityManager manages entities used
// by the game engine
type EntityManager struct {
	entities map[string]Entity
}

// NewEntityManager constructs a new
// EntityManager
func NewEntityManager() *EntityManager {
	return &EntityManager{entities: make(map[string]Entity)}
}

// LoadItemsJSON reads data to populate Items
// from the JSON in r
func (em *EntityManager) LoadItemsJSON(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	m := map[string]map[string]string{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	for k, v := range m {
		em.entities[k] = &Item{
			Name:        v["name"],
			Description: v["description"],
		}
	}

	return err
}
