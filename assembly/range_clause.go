package assembly

import "fmt"

func RangeClause() {
	arr := []int{1, 2, 4}
	var newArr []*int
	for _, v := range arr {
		newArr = append(newArr, &v)
	}
	for _, v := range newArr {
		fmt.Println(*v)
	}
}
