package base

import "fmt"

func ExampleFunctionsParameters() {

	/*

			func имя_функции (список_параметров) (типы_возвращаемых_значений){
		    выполняемые_операторы
		}

	*/

	add(5, 4, 2.0, 3.0, 1.0)

	a := 8
	increment(a)
	fmt.Println("Значение переменной не изменилось и равно ", a)

	//
	//
	addManyParam(1)                             // 1
	addManyParam(1, 2, 3)                       // 6
	addManyParam(1, 2, 3, 4, 5, 6, 7, 8, 9, 10) // 55

	// вызов примеров передачи срезов в ф-ю которая
	// принимает неопред.количество параметров
	// (описано ниже в коде)
	functionCallAddManyParam()
}

//
//
//Название функции вместе с типами ее параметров и типами возвращаемых значений еще называют сигнатурой.
//
//

// Если несколько параметров подряд имеют один и тот же тип, то мы можем указать
// тип только для последнего параметра, а предыдущие параметры также будут представлять этот тип
func add(x, y int, a, b, c float32) {
	var z = x + y
	var d = a + b + c
	fmt.Println("x + y = ", z)
	fmt.Println("a + b + c = ", d)
}

// /
// /
// /
// Стоит помнить, что в функцию передается копия, а не сама переменная
func increment(x int) {

	fmt.Println("x before: ", x)
	x = x + 20
	fmt.Println("x after: ", x)
}

// /
// /
// /
// В функцию можно передать неограниченное количество параметров
func addManyParam(numbers ...int) {
	var sum = 0
	for _, number := range numbers {
		sum += number
	}
	fmt.Println("sum = ", sum)
}

//         ↑↑↑
//         ↑↑↑
//         ↑↑↑

// В данную ф-ю нельзя передавать срез
//
// данные вызовы этой ф-ии вызовут ошибки:
/*
add([]int{1, 2, 3})
add([]int{1, 2, 3, 4})
add([]int{5, 6, 7, 2, 3})
*/
// т.к. передача среза не эквивалентна передаче неопределенного количества параметров
// того же типа

// НО ЕСТЬ ВОЗМОЖНОСТЬ ПЕРЕДАТЬ СРЕЗ!

// ф-я для вызова примеров вызова пред.ф-ии  =)
func functionCallAddManyParam() {
	addManyParam([]int{1, 2, 3}...)

	addManyParam([]int{1, 2, 3, 4}...)

	var nums = []int{5, 6, 7, 2, 3}
	addManyParam(nums...)
}
