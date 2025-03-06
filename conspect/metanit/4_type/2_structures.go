package types_ex

import "fmt"

////////////////////////////////////////////
//                                        //
//                                        //
//           С Т Р У К Т У Р Ы            //
//                                        //
//                                        //
////////////////////////////////////////////

func ExampleStructures() {
	// Структуры - тип данных для представления объекта
	//
	// структуры содержат набор полей, которые представляют атрибуты объекта
	//
	//

	// описание структуры
	type person struct {
		name string
		age  int
	}

	// создание структуры
	{
		var tom person
		fmt.Println(tom)
	}

	// создание и инициализация
	{
		var tom person = person{"Tom", 23}
		fmt.Println(tom)
	}

	// или ЯВНО указать какие значения каким полям передаются
	{
		var alice person = person{age: 23, name: "Alice"}
		fmt.Println(alice)
	}

	// сокращенная форма инициализации
	{
		bob := person{name: "Bob", age: 31}
		fmt.Println(bob)
	}

	// можно ничего не указывать (свойтва получат нулевые значения)
	{
		undefined := person{} // name - пустая строка, age - 0
		fmt.Println(undefined)
	}

	////////////////////////////////////////////
	//     обращение к полям структуры        //
	////////////////////////////////////////////

	{
		// обращение к полям структуры реализуется через точку. Пример :
		//
		// объект.поле

		type person struct {
			name string
			age  int
		}

		var tom = person{name: "Tom", age: 24}
		fmt.Println(tom.name) // Tom
		fmt.Println(tom.age)  // 24

		tom.age = 38                   // изменяем значение поля age
		fmt.Println(tom.name, tom.age) // Tom 38
	}

	////////////////////////////////////////////
	//       указатели на структуры           //
	////////////////////////////////////////////

	line("Указатели на структуры")

	{
		type person struct {
			name string
			age  int
		}

		tom := person{name: "Tom", age: 22}

		var tomPointer *person = &tom

		tomPointer.age = 29
		fmt.Println(tom.age) // 29

		(*tomPointer).age = 32
		fmt.Println(tom.age) // 32

		/*
			для обращения к полям структуры через указатель можно использовать
			операцию разыменования, а можно обращаться напрямую (как показано чуть выше)
		*/
	}

	// можно присвоить через оператор получения адреса    &
	// а можно через ф-ю new (которая ворачивает указатель на пустой объект)
	{
		var tom *person = &person{name: "Tom", age: 23}
		var bob *person = new(person)

		fmt.Println(tom)
		fmt.Println(bob)

	}

}
