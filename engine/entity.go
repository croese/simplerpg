package engine

type Entity struct {
	Id string
}

type Item struct {
	Entity
	Name        string
	Description string
}

type Armor struct {
	Item
	Defense int
}

type Weapon struct {
	Item
	Damage int
}
