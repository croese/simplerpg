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

type Armor struct {
	Item
	Defense int
}

type Weapon struct {
	Item
	Damage int
}

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

func (em *EntityManager) Count() int {
	return len(em.entities)
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

func (em *EntityManager) LoadArmorJSON(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	m := map[string]struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Defense     int    `json:"defense"`
	}{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	for k, v := range m {
		em.entities[k] = &Armor{
			Item: Item{
				Name:        v.Name,
				Description: v.Description,
			},
			Defense: v.Defense,
		}
	}

	return err
}

func (em *EntityManager) LoadWeaponJSON(r io.Reader) error {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}

	m := map[string]struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Damage      int    `json:"damage"`
	}{}
	err = json.Unmarshal(b, &m)
	if err != nil {
		return err
	}

	for k, v := range m {
		em.entities[k] = &Weapon{
			Item: Item{
				Name:        v.Name,
				Description: v.Description,
			},
			Damage: v.Damage,
		}
	}

	return err
}
