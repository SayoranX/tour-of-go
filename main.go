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
	//fmt.Println(math.Pi)
	//fmt.Println(add(42, 15))

	//x := 41
	//y := 13
	//c := add(x, y)
	//fmt.Println(c)

	//a, b := swap("Smoky", "\nPanda") // задание с функцией swap вывод 2 строк
	//fmt.Println(a, b)

	//fmt.Println(split(20)) // задание с функцией split

	//var i, j int = 2, 4
	//v, python, java := false, true, "yes!"

	//fmt.Println(i, j, v, python, java) //  присвоение данных переменным влючая типы bool

	//fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)    // вывод типов данных %T выводит тип данных, %v значение переменной
	//fmt.Printf("Type: %T Value %v\n", MaxInt, MaxInt) // вывод типов данных %T выводит тип данных, %v значение переменной
	//fmt.Printf("Type: %T Value %v\n", z, z)           // вывод типов данных %T выводит тип данных, %v значение переменной

	//var i int
	//var f float64
	//var b bool
	//var s string
	//fmt.Printf("%v %v %v %q\n", i, f, b, s) // вывод значений переменных, %q используется для вывода строкового типа

	var x, y int = 3, 4                           // преобразование переменных из float в int
	var f float64 = math.Sqrt(float64(x*x + y*y)) // преобразование переменных из float в int
	var z uint = uint(f)                          // преобразование переменных из float в int
	fmt.Println(x, y, z)                          // преобразование переменных из float в int

	v := 42                            //%T выводит тип данных, его можно задать числом или указать через var, %T выведет тип указанного числа
	fmt.Printf("v is of type %T\n", v) //%T выводит тип данных, его можно задать числом или указать через var, %T выведет тип указанного числа

	//const World = "Мир" // Константы могут быть символьными, строковыми, булевыми или числовыми значениями. Константы нельзя объявлять с использованием синтаксиса :=.
	//
	//const Log = true                      // Константы могут быть символьными, строковыми, булевыми или числовыми значениями. Константы нельзя объявлять с использованием синтаксиса :=.
	//fmt.Println("Hello", World)           // Константы могут быть символьными, строковыми, булевыми или числовыми значениями. Константы нельзя объявлять с использованием синтаксиса :=.
	//fmt.Println("Happy", Pi, "Zdes", Log) // Константы могут быть символьными, строковыми, булевыми или числовыми значениями. Константы нельзя объявлять с использованием синтаксиса :=.
	//
	//const Truth = false
	//fmt.Println("Go rules?", Truth)

	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))

	//sum := 0                  //цикл for
	//for i := 0; i < 10; i++ { //цикл for
	//	sum += i //цикл for
	//}
	//fmt.Println(sum) //цикл for

	//sum := 1         // структура типа while
	//for sum < 1000 { // структура типа while
	//	sum += sum // структура типа while
	//}
	//fmt.Println(sum) // структура типа while

	//for { 	 //бесконечный цикл, если тело цикла отсутствует
	//}

	fmt.Println(sqrt(-2), sqrt(3), sqrt(-4), sqrt(5))

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

const Pi = 3.14
const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
