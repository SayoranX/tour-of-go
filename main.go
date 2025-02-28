package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println("Hello Smoky Panda")
	//fmt.Println("The time is", time.Now())
	//fmt.Println("My favorite number is", rand.Intn(10))
	//fmt.Println("My favorite number is", rand.Int())
	//fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
	//fmt.Println("My favorite \n number")
	//fmt.Printf("Now you have %g problems gfhfnh bgfn %g hkhjkhjk dhgdgd %d dhhhnfj.\n", math.Sqrt(7), 25, math.Sqrt(7))
	//fmt.Printf("Now you have %g problems gfhfnh bgfn %d hkhjkhjk dhgdgd %g dhhhnfj.\n", math.Sqrt(7), rand.Intn(25), math.Sqrt(7))
	fmt.Println(math.Pi)
	fmt.Println(add(42, 15))

	x := 42
	y := 13
	c := add(x, y)
	fmt.Println(c)
}
func add(x int, y int) int {
	c := x + y
	return c
}
