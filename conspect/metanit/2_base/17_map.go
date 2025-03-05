package base

import "fmt"

func ExampleMap() {
	// это хеш-таблица
	{
		line("Пустая карта")
		var people map[string]int // Ключи представляют тип string, значения - тип int
		fmt.Println(people)
	}

	{
		line("Инициализированная значениями")
		var people = map[string]int{
			"Tom":   1,
			"Bob":   2,
			"Sam":   4,
			"Alice": 8,
		}
		fmt.Println(people) // map[Tom:1 Bob:2 Sam:4 Alice:8]
	}

	{
		line("Изменение значения")

		// для обращения к элементам нужно применять ключи
		var people = map[string]int{
			"Tom":   1,
			"Bob":   2,
			"Sam":   4,
			"Alice": 8,
		}
		fmt.Println(people["Alice"]) // 8
		fmt.Println(people["Bob"])   // 2
		people["Bob"] = 32
		fmt.Println(people["Bob"]) // 32

	}

	{
		line("Проверка наличия элемента по ключу")
		// для проверки наличия элемента по определенному ключу можно применять:
		var people = map[string]int{
			"Tom":   1,
			"Bob":   2,
			"Sam":   4,
			"Alice": 8,
		}
		// если пара с таким ключом есть, то ok будет содержать значение true
		if val, ok := people["Tom"]; ok {
			fmt.Println(val)
		}
	}

	{
		line("Перебор")
		// перебор элементов
		var people = map[string]int{
			"Tom":   1,
			"Bob":   2,
			"Sam":   4,
			"Alice": 8,
		}
		for key, value := range people {
			fmt.Println(key, value)
		}

	}

	{
		line("Создание через make")
		// make - функция для создания хэш-таблицы
		people := make(map[string]int)
		fmt.Println(people)
	}

	{
		line("Добавление значения")
		// для добавления элемента просто присваеваем новому (несуществующему ключу)
		// таблицы какое-то значение
		var people = map[string]int{"Tom": 1, "Bob": 2}
		people["Kate"] = 128
		fmt.Println(people) // map[Tom:1  Bob:2  Kate:128]
	}

	{
		line("Удаление значения")
		var people = map[string]int{"Tom": 1, "Bob": 2, "Sam": 8}
		fmt.Print("До : ")
		fmt.Println(people) // map[Tom:1  Sam:8]
		// Для удаления применяется встроенная функция delete(map, key)
		delete(people, "Bob")
		fmt.Print("После : ")
		fmt.Println(people) // map[Tom:1  Sam:8]
	}
}
