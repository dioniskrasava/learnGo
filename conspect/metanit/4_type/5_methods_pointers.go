package types_ex

import "fmt"

////////////////////////////////////////////////////////////
//                                                        //
//                                                        //
//                   М  Е  Т  О  Д   Ы                    //
//                                                        //
//                                                        //
////////////////////////////////////////////////////////////

////////////////////////////////////////////////////////////
//                                                        //
//                                                        //
//                У  К  А  З  А  Т  Е  Л  Е  Й            //
//                                                        //
//                                                        //
////////////////////////////////////////////////////////////

type person2 struct {
	name string
	age  int
}

// МЕТОД НЕ МЕНЯЮЩИЙ ПОЛЕ ВОЗРАСТА
// потому что получает копию объекта и работает с ней
func (p person2) updateAge(newAge int) {
	p.age = newAge
}

// МЕТОД МЕНЯЮЩИЙ ПОЛЕ ВОЗРАСТА
// потому что получает сс
func (p *person2) updateByPointerAge(newAge int) {
	(*p).age = newAge
}

func ExampleMethodsPointers() {

	var tom = person2{name: "Tom", age: 24}
	fmt.Println("before", tom.age) // 24
	tom.updateAge(33)
	fmt.Println("after", tom.age) // НЕ ИЗМЕНИТСЯ. Останется 24

	// можно его вызвать для самого объекта
	tom.updateByPointerAge(33)
	fmt.Println("after by pointers", tom.age)

	// а можно, как и положено, для указателя
	var tomPointer *person2 = &tom
	tomPointer.updateByPointerAge(35)
	fmt.Println("refresh through the pointer at the pointer ", tom.age)

}
