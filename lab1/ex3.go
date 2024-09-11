package main

import "fmt"

func main() {
	// Инициализация переменных
	var userinit8 uint8 = 1
	var userinit16 uint16 = 2
	var userinit64 int64 = -3
	var userautoinit = -4 // Такой вариант инициализации также возможен

	fmt.Println("Values: ", userinit8, userinit16, userinit64, userautoinit, "\n")

	// Краткая запись объявления переменной
	// только для новых переменных
	intVar := 10

	fmt.Printf("Value = %d Type = %T\n", intVar, intVar)

	// Задание 1. Вывести типы всех переменных
	fmt.Printf("\nЗавдання 1.\n")
	fmt.Printf("Тип userinit8 = %T\n", userinit8)
	fmt.Printf("Тип userinit16 = %T\n", userinit16)
	fmt.Printf("Тип userinit64 = %T\n", userinit64)
	fmt.Printf("Тип userautoinit = %T\n", userautoinit)

	// Задание 2. Присвоить переменной intVar переменные userinit16 и userautoinit. Результат вывести.
	fmt.Printf("\nЗавдання 2.\n")
	intVar = int(userinit16) // Преобразование из uint16 в int
	fmt.Printf("intVar після присвоєння userinit16 = %d\n", intVar)

	intVar = userautoinit // Присвоение значения userautoinit (уже int)
	fmt.Printf("intVar після присвоєння userautoinit = %d\n", intVar)
}
