package types_ex

import "fmt"

////////////////////////////////////////////////////////////
//                                                        //
//                                                        //
//                   М  Е  Т  О  Д   Ы                    //
//                                                        //
//                                                        //
////////////////////////////////////////////////////////////

// определяем свой тип срезов строк
type library []string

// прикрепляем к нему метод
//
// для вывода всех элементов из среза
// (l library)  - определение получателя
func (l library) print() {

	for _, val := range l {
		fmt.Println(val)
	}
}

////////////////////////////////////////////////////////////
//                                                        //
//         М Е Т О Д Ы   Д Л Я    С Т Р У К Т У Р         //
//                                                        //
////////////////////////////////////////////////////////////

type person struct {
	name string
	age  int
}

func (p person) print() {
	fmt.Println("Имя:", p.name)
	fmt.Println("Возраст:", p.age)
}

func (p person) eat(meal string) {
	fmt.Println(p.name, "ест", meal)
}

// //////////////////////////////////////////////////////////
//
// Вызовы (для просмотра работы кода)
func ExampleMethods() {
	var lib library = library{"Book1", "Book2", "Book3"}
	lib.print()

	line("Методы структур")

	var tom = person{name: "Tom", age: 24}
	tom.print()
	tom.eat("борщ с капустой, но не красный")
}
