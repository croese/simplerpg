package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"simplerpg/engine"
)

func loadAssetFile(path string, readerFunc func(io.Reader) error) {
	f, err := os.Open(path)

	if err != nil {
		log.Fatalln(err)
	}

	err = readerFunc(f)
	if err != nil {
		log.Fatalln(err)
	}

	f.Close()
}

func createEntityManager() *engine.EntityManager {
	em := engine.NewEntityManager()
	loadAssetFile("engine/assets/items.json", func(r io.Reader) error { return em.LoadItemsJSON(r) })
	loadAssetFile("engine/assets/armor.json", func(r io.Reader) error { return em.LoadArmorJSON(r) })
	loadAssetFile("engine/assets/weapons.json", func(r io.Reader) error { return em.LoadWeaponJSON(r) })
	return em
}

func main() {
	fmt.Println("Welcome to Simple RPG")
	em := createEntityManager()
	fmt.Printf("Loaded %d entities\n", em.Count())
	fmt.Println(em)
}
