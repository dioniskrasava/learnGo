package base

import "fmt"

func ExampleRecursion() {
	/*
	   Рекурсивная функция представляет такую функцию,
	   которая вызывает саму себя. Рекурсивные функции
	   представляют мощный инструмент для обработки рекурсивных
	   структур данных, например, различных деревьев.
	*/
	line("Factorial")
	fmt.Println(factorial(4)) // 24
	fmt.Println(factorial(5)) // 120
	fmt.Println(factorial(6)) // 720
	line("fibbonachi")
	fmt.Println(fibbonachi(4))  // 3
	fmt.Println(fibbonachi(5))  // 5
	fmt.Println(fibbonachi(6))  // 8
	fmt.Println(fibbonachi(10)) // 55

}

func factorial(n uint) uint {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func fibbonachi(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return fibbonachi(n-1) + fibbonachi(n-2)
}
