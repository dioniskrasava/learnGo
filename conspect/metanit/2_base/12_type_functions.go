package base

import "fmt"

func ExampleTypeFunction() {
	fmt.Println(adding(5, 4))
	fmt.Println(multiply(5, 4))
	display("Lol !")

	line("exaple1")
	example1()
	line("exaple2")
	example2()
	line("exaple3")
	example3()
	line("exaple4")
	example4()
	line("exaple5")
	example5()

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

//
//
// Ф У Н К Ц И И   К А К    П А Р А М Е Т Р Ы     Д Р У Г И Х    Ф У Н К Ц И Й
//
//

func example3() {
	// В Н И М А Н И Е
	// Для наглядности забегу немного вперед и буду испольовать анонимные функции

	add := func(x int, y int) int { return x + y }

	multiply := func(x int, y int) int { return x * y }

	action := func(n1 int, n2 int, operation func(int, int) int) {
		result := operation(n1, n2)
		fmt.Println(result)
	}

	action(10, 25, add)    // 35
	action(5, 6, multiply) // 30

}

func example4() {

	// ФУНКЦИЯ ОПРЕДЕЛЕНИЯ ЧЕТНОСТИ ЧИСЛА
	isEven := func(n int) bool {
		return n%2 == 0
	}

	// ФУНКЦИЯ ОПРЕДЕЛЕНИЯ ПОЛОЖИТЕЛЬНОСТИ ЧИСЛА
	isPositive := func(n int) bool {
		return n > 0
	}

	// Ф-Я СЛОЖЕНИЯ СРЕЗА ПО КРИТЕРИЮ
	sum := func(numbers []int, criteria func(int) bool) int {

		result := 0
		for _, val := range numbers {
			if criteria(val) {
				result += val
			}
		}
		return result
	}

	// САМ ПРИМЕР

	slice := []int{-2, 4, 3, -1, 7, -4, 23}

	sumOfEvens := sum(slice, isEven) // сумма четных чисел
	fmt.Println(sumOfEvens)          // -2

	sumOfPositives := sum(slice, isPositive) // сумма положительных чисел
	fmt.Println(sumOfPositives)              // 37

}

//
//
//  Ф У Н К Ц И Я   К А К   Р Е З У Л Ь Т А Т    Д Р У Г О Й     Ф У Н К Ц И И
//
//

func example5() {
	add := func(x int, y int) int { return x + y }
	subtract := func(x int, y int) int { return x - y }
	multiply := func(x int, y int) int { return x * y }

	selectFn := func(n int) func(int, int) int {
		if n == 1 {
			return add
		} else if n == 2 {
			return subtract
		} else {
			return multiply
		}
	}

	f := selectFn(1)
	fmt.Println(f(3, 4)) // 7

	f = selectFn(3)
	fmt.Println(f(3, 4)) // 12
}

// вспомогательная ф-я для красивого вывода
//
//	------------------------ ПРИМЕР -----------------------
func line(name string) {
	fmt.Println()
	fmt.Println("------------------", name, "------------------")
	fmt.Println()
}
