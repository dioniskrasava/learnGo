package base

import "fmt"

// УСЛОВНЫЕ ВЫРАЖЕНИЯ
func ExampleConditionalExpressions() {
	// выражение возвращающее или true или false

	//
	//
	// ==
	//
	// Операция "равно". Возвращает true, если оба операнда равны, и false, если они не равны:
	var a int = 8
	var b int = 3
	var c bool = a == b
	fmt.Println(c) // false

	//
	//
	// >
	//
	// Операция "больше чем". Возвращает true, если первый операнд больше второго,
	// и false, если первый операнд меньше второго:
	var d int = 8
	var e int = 3
	var f bool = d > e
	fmt.Println(f) // true

	//
	//
	// <
	//
	// Операция "меньше чем". Возвращает true, если первый операнд меньше второго,
	//  и false, если первый операнд больше второго:
	var a2 int = 8
	var b2 int = 3
	var c2 bool = a2 < b2
	fmt.Println(c2) // false

	//
	//
	// <=
	//
	// Операция "меньше или равно". Возвращает true, если первый операнд меньше
	// или равен второму, и false, если первый операнд больше второго:
	var a3 int = 8
	var b3 int = 3
	var c3 bool = a3 <= b3
	fmt.Println(c3) // false

	//
	//
	// >=
	//
	// Операция "больше или равно". Возвращает true, если первый операнд больше или
	// равен второму, и false, если первый операнд меньше второго:
	var a4 int = 8
	var b4 int = 3
	var c4 bool = a4 >= b4
	fmt.Println(c4) // true

	//
	//
	// !=
	//
	// Операция "не равно". Возвращает true, если первый операнд не равен второму,
	// и false, если оба операнда равны:
	var a5 int = 8
	var b5 int = 3
	var c5 bool = a5 != b5
	var d5 bool = a5 != 8
	fmt.Println(c5) // true
	fmt.Println(d5) // false

	//
	//
	// Л О Г И Ч Е С К И Е          О П Е Р А Ц И И
	//
	//

	//
	//
	// ! (операция отрицания)
	//
	// Инвертирует значение. Если операнд равен true, то возвращает false,
	//  иначе возвращает true.
	var a6 bool = true
	var b6 bool = !a6 // false
	var c6 bool = !b6 // true
	fmt.Println(b6)
	fmt.Println(c6)

	//
	//
	// && (конъюнкция, логическое умножение)
	//
	// Возвращает true, если оба операнда не равны false. Возвращает false,
	//  если хотя бы один операнд равен false.
	var b7 bool = 4 > 5 && 6 > 8   //false
	var c7 bool = 3 <= 5 && 10 > 8 // true
	fmt.Println(b7)
	fmt.Println(c7)

	//
	//
	// || (дизъюнкция, логическое сложение)
	//
	// Возвращает true, если хотя бы один операнд не равен false. Возвращает
	// false, если оба операнда равны false.
	var b8 bool = 4 > 5 || 6 > 8   //false
	var c8 bool = 3 == 5 || 10 > 8 // true
	fmt.Println(b8)
	fmt.Println(c8)

}
