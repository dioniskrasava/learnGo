package base

import "fmt"

func ExampleReturnResultFunction() {

	/*
	   func имя_функции (список_параметров) тип_возвращаемого_значения {
	       выполняемые_операторы
	       return возвращаемое_значение
	   }
	*/

	answ := add2(5, 7)
	fmt.Println("Результат ф-ии add2 :", answ)

	answ = add3(5, 7)
	fmt.Println("Результат ф-ии add3 :", answ)

	answ = add4(5, 7)
	fmt.Println("Результат ф-ии add4 :", answ)

	// В О З В Р А Щ Е Н И Е
	// Н Е С К О Л Ь К И Х
	// З Н А Ч Е Н И Й

	age, name := add5(4, 5, "Tom", "Simpson")
	fmt.Println(age)  // 9
	fmt.Println(name) // Tom Simpson

	age, name = add6(4, 5, "Tom", "Simpson")
	fmt.Println(age)  // 9
	fmt.Println(name) // Tom Simpson

}

// возвращает int
func add2(x, y int) int {
	return x + y
}

// Возвращаемый результат может быть именован:
func add3(x, y int) (z int) {
	z = x + y
	return
}

// эквивалент ф-ии
func add4(x, y int) int {
	var z int = x + y
	return z
}

// В О З В Р А Щ Е Н И Е
// Н Е С К О Л Ь К И Х
// З Н А Ч Е Н И Й

func add5(x, y int, firstName, lastName string) (int, string) {
	var z int = x + y
	var fullName = firstName + " " + lastName
	return z, fullName
}

// или именованные результаты
func add6(x, y int, firstName, lastName string) (z int, fullName string) {
	z = x + y
	fullName = firstName + " " + lastName
	return
}
