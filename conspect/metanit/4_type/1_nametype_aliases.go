package types_ex

import "fmt"

func ExampleTypesAliases() {

	{
		// создаем новый тип на основе другого
		type mile uint

		var distance mile = 5
		fmt.Println(distance)
		distance += 5
		fmt.Println(distance)
	}

	{
		// ТАКОГО РОДА ПОЛЬЗОВАТЕЛЬСКИЕ ТИПЫ НУЖНЫ ДЛЯ
		// УМЕНЬШЕНИЯ ВЕРОЯТНОСТИ ПЕРЕДАЧИ НЕКОРРЕКТНЫХ ДАННЫХ
		//
		// поэтому mile, kilometer считаются разными типами
		// хотя основаны на общем типе
		type mile uint
		type kilometer uint

		distanceToEnemy := func(distance mile) {

			fmt.Println("расстояние для противника:")
			fmt.Println(distance, "миль")
		}

		var distance mile = 5
		distanceToEnemy(distance)

		//var distance1 uint = 5
		// distanceToEnemy(distance1)   // !Ошибка

		var distance2 kilometer = 5
		// distanceToEnemy(distance2)   // ! ошибка
		fmt.Println(distance2) // чтобы IDLE не ругалась
	}

	/*
		И М Е Н О В А Н Н Ы Е    Т И П Ы    П О З В О Л Я Ю Т    П Р И Д А Т Ь
		Т И П У    Д О П О Л Н И Т Е Л Ь Ы Й   С М Ы С Л

		П О З В О Л Я Е Т    У К А З А Т Ь    Н А    П Р Е Д Н А З Н А Ч Е Н И Е
		П Е Р М Е Н Н О Й   (Б О Л Е Е   О П И С А Т Е Л Ь Н О  Н Е Ж Е Л И   UINT)
	*/

	line("Без использования дополнительных типов")

	//
	//
	//
	// ПРИМЕР БЕЗ ИСПОЛЬЗОВАНИЯ ДОПОЛНИТЕЛЬНЫХ ТИПОВ
	//
	//
	//

	{
		action := func(n1 int, n2 int, op func(int, int) int) {
			result := op(n1, n2)
			fmt.Println(result)
		}

		add := func(x int, y int) int {
			return x + y
		}

		var myOperation func(int, int) int = add
		action(10, 25, myOperation) // 35

	}

	line("С дополнительным типом")

	//
	//
	//
	// С ДОПОЛНИТЕЛЬНЫМ ТИПОМ
	//
	//
	//

	{
		type BinaryOp func(int, int) int // вводим новый тип для удобства

		action := func(n1 int, n2 int, op BinaryOp) {

			result := op(n1, n2)
			fmt.Println(result)
		}

		add := func(x int, y int) int {

			return x + y
		}

		var myOperation BinaryOp = add
		action(10, 35, myOperation) // 45

	}

	////////////////////////////////////////////
	//                                        //
	//                                        //
	//          П С Е В Д О Н И М Ы           //
	//                                        //
	//                                        //
	////////////////////////////////////////////

	{
		// НЕ ОПРЕДЕЛЯЕТ НОВОГО ТИПА
		type mile = uint
		type kilometer = uint

		distanceToEnemy := func(distance mile) {

			fmt.Println("расстояние для противника:")
			fmt.Println(distance, "миль")
		}

		var distance mile = 5
		distanceToEnemy(distance)

		var distance1 uint = 5
		distanceToEnemy(distance1) // норм

		var distance2 kilometer = 5
		distanceToEnemy(distance2) // норм

		/*
			Обычо псевдонимы применяются для сокращения названий других
			типов или для определения более описательного имени
		*/

	}
}

func line(str string) {
	fmt.Println("--------------------------------------------------------------------------")
	fmt.Println("           ", str, "             ")
	fmt.Println("--------------------------------------------------------------------------")
}
