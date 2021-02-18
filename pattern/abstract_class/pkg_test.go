package abstract_class

import "testing"

func TestPlayAnimal(t *testing.T) {
	cat := NewCat()
	PlayAnimalWith(cat)
	eagle := NewEagle()
	PlayAnimalWith(eagle)
	monkey := NewMonkey()
	PlayAnimalWith(monkey)
	parrot := NewParrot()
	PlayAnimalWith(parrot)
}
