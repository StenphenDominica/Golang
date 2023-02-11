package main

import (
	"fmt"
)

// func main() {
// 	fmt.Printf("Hello Golang \n")
// 	printTest()
// }

func printTest() {
	var name, age, address string
	fmt.Printf("Enter your name: \n")
	fmt.Scanln(&name)
	fmt.Printf("Enter your age: \n")
	fmt.Scanln(&age)
	fmt.Printf("Enter your address: \n")
	fmt.Scanln(&address)

	fmt.Printf("Your name is %v and %v years old, you are live in %v", name, age, address)
}
