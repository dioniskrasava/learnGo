package types_ex

import "fmt"

////////////////////////////////////////////////////////////
//                                                        //
//                                                        //
//      В Л О Ж Е Н Н Ы Е      С Т Р У К Т У Р Ы          //
//                                                        //
//                                                        //
////////////////////////////////////////////////////////////

func ExampleNestedStructures() {
	line("Простой пример")
	{
		type contact struct {
			email string
			phone string
		}

		type person struct {
			name        string
			age         int
			contactInfo contact
		}

		var tom = person{
			name: "Tom",
			age:  24,
			contactInfo: contact{
				email: "tom@gmail.com",
				phone: "+1234567899",
			},
		}
		tom.contactInfo.email = "supertom@gmail.com"

		fmt.Println(tom.contactInfo.email) // supertom@gmail.com
		fmt.Println(tom.contactInfo.phone) // +1234567899

	}

	line("Сокращенный пример")

	// МОЖНО СОКРАТИТЬ ОПРЕДЕЛЕНИЕ ПОЛЯ ТАК :
	{
		type contact struct {
			email string
			phone string
		}

		type person struct {
			name string
			age  int
			contact
		}

		/* // Эквивалентно
		type person struct {
			name    string
			age     int
			contact contact
		}
		*/

		var tom = person{
			name: "Tom",
			age:  24,
			contact: contact{
				email: "tom@gmail.com",
				phone: "+1234567899",
			},
		}
		tom.email = "supertom@gmail.com"

		fmt.Println(tom.email) // supertom@gmail.com
		//fmt.Println(tom.contact.email) // а можно и так
		fmt.Println(tom.phone) // +1234567899

	}

	////////////////////////////////////////////////////////////
	//      Хранение ссылки на структуру того же типа         //
	////////////////////////////////////////////////////////////

	{
		/* // ОШИБКА
		type node struct {
			value int
			next  node
		}
		*/

		example1()
	}

}

////////////////////////////////////////////////////////////
//                                                        //
//                                                        //
//            ПРИМЕР РЕКУРСИИ В СТРУКТУРАХ                //
//                                                        //
//                                                        //
////////////////////////////////////////////////////////////

func example1() {
	first := node{value: 4}
	second := node{value: 5}
	third := node{value: 6}

	first.next = &second
	second.next = &third

	var current *node = &first
	for current != nil {
		fmt.Println(current.value)
		current = current.next
	}
}

type node struct {
	value int
	next  *node
}

// рекурсивный вывод списка
func printNodeValue(n *node) {

	fmt.Println(n.value)
	if n.next != nil {
		printNodeValue(n.next)
	}
}
