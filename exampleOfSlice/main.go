package main

import (
	"fmt"
	"reflect"
)

func main() {

	var slice []int

	fmt.Println(reflect.ValueOf(slice))
	fmt.Println(reflect.ValueOf(slice).Kind())
	fmt.Println("uninit:", slice, slice == nil, len(slice) == 0)

	slice = make([]int, 10, 10)
	fmt.Println("emp:", slice, "len:", len(slice), "cap:", cap(slice))

	slice[0] = 100
	slice[1] = 200

	for index, value := range slice {
		fmt.Println("index:", index, "value:", value)
	}

	fmt.Println("emp:", slice, "len:", len(slice), "cap:", cap(slice))

	b := append(slice, 300)
	b[0] = 1
	fmt.Println(b)

}
