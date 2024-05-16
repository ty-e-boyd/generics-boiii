package main

import "fmt"

func main() {
	knightA := newKnight[string]("Jeffrey", "Godfried", "Broadsword", Goodly)
	knightB := newKnight[int64]("Francios", "Bellafor", "Great Axe", Badly)
	knightC := newKnight[float64]("William", "Wallace", "Flail", Goodly)

	knightA.PresentThyself()
	knightB.PresentThyself()
	knightC.PresentThyself()
}

// assigning the type of standing
type Standing int64

const (
	Goodly Standing = iota
	Badly
)

func (s Standing) String() string {
	switch s {
	case Goodly:
		return "Goodly"
	case Badly:
		return "Badly"
	}
	return "Unknown status"
}

// creating the Knight Struct, with a generic used to assign knight type
type Knight[T any] struct {
	Name       string
	House      string
	Weapon     string
	Standing   Standing
	knightType T
}

func (k Knight[T]) PresentThyself() {
	fmt.Printf("I am %v, a %T knight of house %v. I use a %v, and I am of %v standing.\n", k.Name, k.knightType, k.House, k.Weapon, k.Standing)
}

func newKnight[T string | int64 | float64](name string, house string, weapon string, standing Standing) *Knight[T] {
	knight := Knight[T]{
		Name:     name,
		House:    house,
		Weapon:   weapon,
		Standing: standing,
	}

	return &knight
}
