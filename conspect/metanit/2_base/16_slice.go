package base

import "fmt"

func ExampleSlices() {
	var users []string
	var users2 = []string{"Tom", "Alice", "Kate"}
	// или так
	users3 := []string{"Tom", "Alice", "Kate"}

	fmt.Println(users)
	fmt.Println(users2)
	fmt.Println(users3)

	//------------------------------------
	// изменение элементов среза
	//

	line("")
	fmt.Println(users3[2]) // Kate
	users3[2] = "Katherine"

	//------------------------------------
	// цикл по элементам среза
	//

	line("")
	for _, value := range users3 {
		fmt.Println(value)
	}

	//+++++++++++++++++++++++++++
	//                          +
	//        m a k e ( )       +
	//                          +
	//+++++++++++++++++++++++++++

	var users4 []string = make([]string, 3)
	users4[0] = "Tom"
	users4[1] = "Alice"
	users4[2] = "Bob"

	//+++++++++++++++++++++++++++++++++++++++++++
	//                                          +
	//     д о б а в л е н и е   в    с р е з   +
	//                                          +
	//+++++++++++++++++++++++++++++++++++++++++++

	users5 := []string{"Tom", "Alice", "Kate"}
	users5 = append(users5, "Bob")

	for _, value := range users5 {
		fmt.Println(value)
	}

	//+++++++++++++++++++++++++++++++++++++++++++
	//                                          +
	//    о п е р а т о р   с р е з а           +
	//                                          +
	//+++++++++++++++++++++++++++++++++++++++++++

	line("оператор среза")

	initialUsers := [8]string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	users6 := initialUsers[2:6] // с 3-го по 6-й
	users7 := initialUsers[:4]  // с 1-го по 4-й
	users8 := initialUsers[3:]  // с 4-го до конца

	fmt.Println(users6) // ["Kate", "Sam", "Tom", "Paul"]
	fmt.Println(users7) // ["Bob", "Alice", "Kate", "Sam"]
	fmt.Println(users8) // ["Sam", "Tom", "Paul", "Mike", "Robert"]

	//+++++++++++++++++++++++++++++++++++++++++++
	//                                          +
	//    у д а л е н и е    э л е м е н т а    +
	//                                          +
	//+++++++++++++++++++++++++++++++++++++++++++

	line("удаление элемента")

	/*
		Что делать, если необходимо удалить какой-то определенный элемент? В этом случае
		 мы можем комбинировать функцию append и оператор среза:
	*/

	users9 := []string{"Bob", "Alice", "Kate", "Sam", "Tom", "Paul", "Mike", "Robert"}
	// удаляем 4-й элемент
	var n = 3
	users9 = append(users9[:n], users9[n+1:]...)
	fmt.Println(users9) //["Bob", "Alice", "Kate", "Tom", "Paul", "Mike", "Robert"]

}
