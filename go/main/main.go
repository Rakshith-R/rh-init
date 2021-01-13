package main

import (
	"fmt"

	"./submodule"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	ptr := &arr
	fmt.Println(ptr[0])
	fmt.Println(submodule.Reverse("Hello world!"))
}
