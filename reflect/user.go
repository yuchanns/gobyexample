package reflect

import (
	"fmt"
)

type User struct {
	Name   string
	Gender int8
	Age    int
}

func (u User) Test() {
	fmt.Println("call test")
}

func (u *User) String() string {
	return fmt.Sprintf("a person name %s who aged %d", u.Name, u.Age)
}
