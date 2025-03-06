package points

import "fmt"

func ExamplePointsFunctions() {
	line("Передача переменной")

	{
		// функция получает копию переменной
		changeValue := func(x int) {
			x = x * x
		}

		d := 5
		fmt.Println("d before:", d) // 5
		changeValue(d)              // изменяем значение
		fmt.Println("d after:", d)  // 5 - значение не изменилось
	}

	line("Передача указателя")

	{
		// функция получает копию переменной
		changeValue := func(x *int) {
			*x = (*x) * (*x)
		}

		d := 5
		fmt.Println("d before:", d) // 5
		changeValue(&d)             // изменяем значение // передавая адрес переменной
		fmt.Println("d after:", d)  // 5 - значение не изменилось
	}

	line("Возвращает указатель")

	// функция возвращает указатель
	{
		createPointer := func(x int) *int {
			p := new(int)
			*p = x
			return p
		}

		p1 := createPointer(7)
		fmt.Println("p1:", *p1, ". Adress: ", p1) // p1: 7
		p2 := createPointer(10)
		fmt.Println("p2:", *p2, ". Adress: ", p2) // p2: 10
		p3 := createPointer(28)
		fmt.Println("p3:", *p3, ". Adress: ", p3) // p3: 28

	}
}
