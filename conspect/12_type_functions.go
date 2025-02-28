package base

import "fmt"

func ExampleTypeFunction() {
	fmt.Println(adding(5, 4))
	fmt.Println(multiply(5, 4))
	display("Lol !")

	example1()
	example2()

}

// Эта функция представляет тип
//
//	func(int, int) int.
func adding(x int, y int) int {
	return x + y
}

// И эта ф-я того же типа
//
//	func(int, int) int.
func multiply(x int, y int) int {
	return x * y
}

// Эта функция имеет тип
//
//	func(string)
func display(message string) {
	fmt.Println(message)
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
func example1() {

	// переменная может быть функцией

	//
	var f func(int, int) int = adding
	result := f(4, 5)
	fmt.Println(result)
}

// ----------------------------------------------------------------
// ----------------------------------------------------------------
// ----------------------------------------------------------------
func example2() {
	f := adding          //или так var f func(int, int) int = add
	fmt.Println(f(3, 4)) // 7

	f = multiply         // теперь переменная f указывает на функцию multiply
	fmt.Println(f(3, 4)) // 12

	// f = display      // ошибка, так как функция display имеет тип func(string)

	var f1 func(string) = display // норм
	f1("hello")
}
