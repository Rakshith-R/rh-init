package main

import (
	"fmt"

	"github.com/rakshith-r/rh-init/go/src/main/submodule"
)

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	ptr := &arr
	fmt.Scanln(&arr[0])
	fmt.Println(ptr[0])
	fmt.Println(submodule.Reverse("Hello world!"))
}

//Main ..
func Main() {
	main()
}
