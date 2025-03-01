package main

import (
	"fmt"
	"math"
	//"math/rand"
	"math/cmplx"
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

	x := 41
	y := 13
	c := add(x, y)
	fmt.Println(c)

	a, b := swap("Smoky", "\nPanda") // задание с функцией swap вывод 2 строк
	fmt.Println(a, b)

	fmt.Println(split(20)) // задание с функцией split

	var i, j int = 2, 4
	v, python, java := false, true, "yes!"

	fmt.Println(i, j, v, python, java) //  присвоение данных переменным влючая типы bool

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)    // вывод типов данных %T выводит тип данных, %v значение переменной
	fmt.Printf("Type: %T Value %v\n", MaxInt, MaxInt) // вывод типов данных %T выводит тип данных, %v значение переменной
	fmt.Printf("Type: %T Value %v\n", z, z)           // вывод типов данных %T выводит тип данных, %v значение переменной
}
func add(x, y int) int {
	c := x + y
	return c
}
func swap(x, y string) (string, string) {
	return x, y
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

var (
	ToBe   int        = 10
	MaxInt uint32     = 1<<32 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)
