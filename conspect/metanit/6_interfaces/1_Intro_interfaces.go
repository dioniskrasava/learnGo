package interfaces

import (
	"fmt"
)

// Интерфейсы представляют абстракцию поведения других типов.
// Интерфейсы позволяют определять функции, которые не привязаны
// к конкретной реализации. То есть интерфейсы определяют
// некоторый функционал, но не реализуют его.

// В Go интерфейс реализуется НЕЯВНО. Нам не надо специально указывать,
// что структуры применяют определенный интерфейс, как в некоторых других
// языках программирования. Для реализации типу данных достаточно реализовать
// методы, которые определяет интерфейс.

type Vehicle interface {
	move()
}

// структура "Автомобиль"
type Car struct{}

// структура "Самолет"
type Aircraft struct{}

func (c Car) move() {
	fmt.Println("Автомобиль едет")
}

func (a Aircraft) move() {
	fmt.Println("Самолет летит")
}

func ExampleIntroInterfaces() {
	var tesla Vehicle = Car{}
	var boing Vehicle = Aircraft{}
	tesla.move()
	boing.move()
}
