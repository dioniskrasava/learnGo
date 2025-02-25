package base

import "fmt"

// Поразрядные операции
func ExampleBitwiseOperations() {
	// ********************************************************************************************************************************
	// *                                                                                                                              *
	// ********************************************************************************************************************************
	//
	//
	//	<<
	//
	// Сдвигает битовое представление числа, представленного первым операндом,
	// влево на определенное количество разрядов, которое задается вторым операндом.

	//
	//
	//  >>
	//
	// Сдвигает битовое представление числа вправо на определенное количество разрядов.

	var b int = 2 << 2  // 10  на два разрядов влево = 1000 - 8
	var c int = 16 >> 3 // 10000 на три разряда вправо = 10 - 2

	fmt.Println(b)
	fmt.Println(c)

	// ********************************************************************************************************************************
	// *                                                                                                                              *
	// ********************************************************************************************************************************
	//
	//  П О Р А З Р Я Д Н Ы Е      О П Е Р А Ц И И
	//
	//
	//------------------------------------------------------------------------------------------
	// &: поразрядная конъюнкция (операция И или поразрядное умножение). Возвращает 1,         -
	// если оба из соответствующих разрядов обоих чисел равны 1. Возвращает 0, если            -
	// разряд хотя бы одного числа равен 0                                                     -
	//------------------------------------------------------------------------------------------
	// |: поразрядная дизъюнкция (операция ИЛИ или поразрядное сложение). Возвращает 1,        -
	// если хотя бы один из соответствующих разрядов обоих чисел равен 1                       -
	//------------------------------------------------------------------------------------------
	// ^: поразрядное исключающее ИЛИ. Возвращает 1, если только один из                       -
	// соответствующих разрядов обоих чисел равен 1                                            -
	//------------------------------------------------------------------------------------------
	// &^: сброс бита (И НЕ). В выражении z = x &^ y каждый бит z равен 0,                     -
	// если соответствующий бит y равен 1. Если бит в y равен 0, то берется                    -
	// значение соответствующего бита из x.                                                    -
	//------------------------------------------------------------------------------------------

	var d int = 5 | 2  // 101 | 010 = 111  - 7
	var e int = 6 & 2  // 110 & 010 = 010  - 2
	var f int = 5 ^ 2  // 101 ^ 010 = 111 - 7
	var g int = 5 &^ 6 // 101 &^ 110 = 001 - 1

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)

}
