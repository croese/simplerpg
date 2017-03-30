package main

import (
	"fmt"
	"log"
	"os"
	"simplerpg/engine"
)

func createEntityManager() *engine.EntityManager {
	em := engine.NewEntityManager()
	f, err := os.Open("engine/assets/items.json")
	if err != nil {
		log.Fatalln(err)
	}
	em.LoadItemsJSON(f)

	return em
}

func main() {
	fmt.Println("Welcome to Simple RPG")
	em := createEntityManager()
	fmt.Printf("Loaded %d entities\n", em.Count())
}
