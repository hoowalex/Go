package main

import "fmt"

func main() {
	variable8 := int8(127)
	variable16 := int16(16383)

	fmt.Println("Приведение типов\n")

	fmt.Printf("variable8         = %-5d = %.16b\n", variable8, variable8)
	fmt.Printf("variable16        = %-5d = %.16b\n", variable16, variable16)
	fmt.Printf("uint16(variable8) = %-5d = %.16b\n", uint16(variable8), uint16(variable8))
	fmt.Printf("uint8(variable16) = %-5d = %.16b\n", uint8(variable16), uint8(variable16))

	//Задание.
	//1. Создайте 2 переменные  разных типов. Выпоните арифметические операции. Результат вывести

	// Створюємо 2 змінні різних типів
	var a int32 = 1000
	var b int64 = 50000

	// Виконання арифметичних операцій з приведенням типів
	sum := int64(a) + b // Приведення int32 до int64
	difference := b - int64(a)
	product := int64(a) * b

	fmt.Println("\nАрифметичні операції з різними типами")
	fmt.Printf("Сума:         %d + %d = %d\n", a, b, sum)
	fmt.Printf("Різниця:      %d - %d = %d\n", b, a, difference)
	fmt.Printf("Добуток:      %d * %d = %d\n", a, b, product)
}
