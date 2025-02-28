package main

import "fmt"

func main() {
	a, err := fmt.Println("Hello Smoky Panda")
	if err != nil {
		panic(err)
	}

	fmt.Println("var A=", a)
}
