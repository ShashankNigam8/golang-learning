package main

import "fmt"

func main() {
	num := 5

	//address of num
	fmt.Println("address of num: ", &num)

	name := "shashank"
	var ptr *string

	ptr = &name

	// * to get the value pointed by ptr
	fmt.Println(*ptr)
	fmt.Println("address of name: ", &num)
}
