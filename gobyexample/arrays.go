package main

import (
	"fmt"
)

func main() {

	slice := make([]int, 5)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	slice[0] = 1
	slice[1] = 2
	slice[2] = 3
	slice[3] = 4
	slice[4] = 5

	slice = append(slice, 6)

	fmt.Println(len(slice))
	fmt.Println(cap(slice))

	l := slice[1:3:4]
	fmt.Println(l)
}
