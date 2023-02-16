package main

import (
	"fmt"
	_ "net/http/pprof"
	// _ "chapter2"
)

func myfunc() (int, string) {
	a := 10
	b := "golang"
	return a, b
}

type Foo interface {
	Say()
}

type Dog struct {
	name string
}

func (d Dog) Say() {
	fmt.Println(d.name + " say hi")
}

// func main() {
// 	a, _ := myfunc()
// 	fmt.Printf("Fetch function myfunc's only first result %d: \n", a)
// 	var dog Foo = Dog{"Black dog"}
// 	dog.Say()

// 	fmt.Printf("Call custom package's parameter: %s", chapter2.A_string)
// }

type Student struct {
	Name  string
	Age   int
	Score float32
}
