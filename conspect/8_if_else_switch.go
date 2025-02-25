package base

import "fmt"

func ExampleIfElseSwitch() {

	//-------------------------------------

	a := 6
	b := 7
	if a < b {
		fmt.Println("a меньше b")
	}

	//-------------------------------------

	c := 6
	d := 7
	if c < d {
		fmt.Println("c меньше d")
	} else {
		fmt.Println("c больше d")
	}

	//-------------------------------------

	if a == 9 {
		fmt.Println("a = 9")
	} else if a == 8 {
		fmt.Println("a = 8")
	} else if a == 7 {
		fmt.Println("a == 7")
	}

	//-------------------------------------

	//            S   W   I   T   C   H
	a = 8
	switch a {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	}
	//--------------------------------------
	a = 7
	switch a + 2 { // можно выражение
	case 9:
		fmt.Println("9")
	case 8:
		fmt.Println("8")
	case 7:
		fmt.Println("7")
	}
	//---------------------------------------
	a = 87
	switch a {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	default:
		fmt.Println("значение переменной a не определено")
	}
	//----------------------------------------
	a = 5
	switch a {
	case 9:
		fmt.Println("a = 9")
	case 8:
		fmt.Println("a = 8")
	case 7:
		fmt.Println("a = 7")
	case 6, 5, 4: // можно несколько значений
		fmt.Println("a = 6 или 5 или 4, но это не точно")
	default:
		fmt.Println("значение переменной a не определено")
	}
}
