package base

import "fmt"

func ExampleDefetPanic() {
	// Оператор defer позволяет выполнить определенную функцию в конце программы
	line("DEFER")
	defer fmt.Println("E")
	defer fmt.Println("D")
	defer fmt.Println("C")
	fmt.Println("A")
	fmt.Println("B")

	//
	//
	//
	//Оператор panic позволяет сгенерировать ошибку и выйти из программы:
	line("PANIC")
	fmt.Println(divide(15, 5))
	fmt.Println(divide(4, 0)) // вызовет панику и выход из программы
	fmt.Println("Program has been finished")

}

func divide(x, y float64) float64 {
	if y == 0 {
		panic("Division by zero!") // БЕЗ ПАНИКИ программа при делении на ноль выдает "Inf+"
	}
	return x / y
}
