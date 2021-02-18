package abstract_class

import "fmt"

func PlayAnimalWith(a IAnimal) {
	fmt.Println(a.Describe())
	fmt.Println(a.Move())
}
