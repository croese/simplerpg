package engine

import (
	"strings"
	"testing"
)

func TestNewEntityManagerInitializesMap(t *testing.T) {
	em := NewEntityManager()

	if em.entities == nil {
		t.Error("entities map wasn't initialized")
	}
}

func TestLoadItemsJSON(t *testing.T) {
	data := `{
	"item_gold_coin": {
		"name": "Gold Coin",
		"description": "A small disc made of lustrous metal"
	},

	"item_iron_key": {
		"name": "Iron Key",
		"description": "A heavy iron key with a simple cut"
	}
}`

	reader := strings.NewReader(data)
	em := NewEntityManager()
	err := em.LoadItemsJSON(reader)

	if err != nil {
		t.Errorf("returned an error %s", err)
	}

	if len(em.entities) != 2 {
		t.Errorf("wrong number of entities. should be 2. got=%d",
			len(em.entities))
	}

	item1, item2 := em.entities["item_gold_coin"].(*Item),
		em.entities["item_iron_key"].(*Item)

	if item1.Name != "Gold Coin" &&
		item1.Description != "A small disc made of lustrous metal" {
		t.Errorf("incorrect parsed item1 data. got=%v", item1)
	}

	if item2.Name != "Iron Key" &&
		item2.Description != "A heavy iron key with a simple cut" {
		t.Errorf("incorrect parsed item1 data. got=%v", item1)
	}
}

func TestTestLoadItemsJSONWithEOF(t *testing.T) {
	reader := strings.NewReader("")
	em := NewEntityManager()
	err := em.LoadItemsJSON(reader)

	if err == nil {
		t.Error("LoadItemsJSON should have returned error")
	}
}
