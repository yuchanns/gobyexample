package abstract_class

import (
	"strings"
)

// IAnimal is the base interface of all animals
type IAnimal interface {
	Move() string
	GetName() string
	Describe() string
}

// WalkAnimal is the base abstract class of walk animals
type WalkAnimal struct {
	IAnimal
}

// NewWalkAnimalWith returns a animal implemented based on WalkAnimal abstraction
func NewWalkAnimalWith(a IAnimal) *WalkAnimal {
	return &WalkAnimal{IAnimal: a}
}

// Move is the implementation of IAnimal by WalkAnimal which defines that all walk animals move by walk
func (a *WalkAnimal) Move() string {
	return strings.Join([]string{a.IAnimal.GetName(), "moves by walk."}, " ")
}

// FlyAnimal is the base abstract class of fly animals
type FlyAnimal struct {
	IAnimal
}

// NewFlyAnimalWith returns a animal implemented based on FlyAnimal abstraction
func NewFlyAnimalWith(a IAnimal) *FlyAnimal {
	return &FlyAnimal{IAnimal: a}
}

// Move is the implementation of IAnimal by FlyAnimal which defines that all fly animals move by fly
func (a *FlyAnimal) Move() string {
	return strings.Join([]string{a.GetName(), "moves by fly."}, " ")
}

// Cat is a kind of walk animals
type Cat struct {
	*WalkAnimal
}

func NewCat() IAnimal {
	return NewWalkAnimalWith(&Cat{})
}

func (a *Cat) GetName() string {
	return "Cat"
}

func (a *Cat) Describe() string {
	return "The cat is a domestic species of small carnivorous mammal."
}

// Monkey is a kind of walk animals
type Monkey struct {
	*WalkAnimal
}

func NewMonkey() IAnimal {
	return NewWalkAnimalWith(&Monkey{})
}

func (a *Monkey) GetName() string {
	return "Monkey"
}

func (a *Monkey) Describe() string {
	return "Monkey is a common name that may refer to groups or species of mammals, in part, the simians of infraorder Simiiformes."
}

// Eagle is a kind of fly animals
type Eagle struct {
	*FlyAnimal
}

func NewEagle() IAnimal {
	return NewFlyAnimalWith(&Eagle{})
}

func (a *Eagle) GetName() string {
	return "Eagle"
}

func (a *Eagle) Describe() string {
	return "Eagle is the common name for many large birds of prey of the family Accipitridae."
}

// Parrot is a kind of fly animals
type Parrot struct {
	*FlyAnimal
}

func NewParrot() IAnimal {
	return NewFlyAnimalWith(&Parrot{})
}

func (a *Parrot) GetName() string {
	return "Parrot"
}

func (a *Parrot) Describe() string {
	return "Parrots, also known as psittacines, are birds of the roughly 398 species in 92 genera comprising the order Psittaciformes, found mostly in tropical and subtropical regions."
}
