package main

import (
	"fmt"
	"simplerpg/engine"
)

func main() {
	e := engine.Item{
		Entity:      engine.Entity{Id: "id"},
		Description: "desc",
		Name:        "name",
	}
	fmt.Println(e)
}
