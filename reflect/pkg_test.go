package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	u := User{
		Name:   "yuchanns",
		Gender: 1,
		Age:    27,
	}

	to := reflect.TypeOf(u)
	vo := reflect.ValueOf(u)
	pto := reflect.TypeOf(&u)
	pvo := reflect.ValueOf(&u)
	fmt.Println("type is ", to)
	fmt.Println("value is ", vo)

	for i := 0; i < to.NumField(); i++ {
		fieldStruct := to.Field(i)
		value := vo.Field(1).Interface()
		fmt.Printf("%s: %v = %v\n", fieldStruct.Name, fieldStruct.Type, value)
	}

	for i := 0; i < pto.NumMethod(); i++ {
		m := pto.Method(i)
		fmt.Printf("%s: %v\n", m.Name, m.Type)
	}

	m := vo.MethodByName("Test")
	m.Call([]reflect.Value{})

	m2 := pvo.MethodByName("String")
	fmt.Println(m2.Call([]reflect.Value{})[0])
}
