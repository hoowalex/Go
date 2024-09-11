package main

import "fmt"

func main() {
	var chartype int8 = 'R'

	fmt.Printf("Code '%c' - %d\n", chartype, chartype)

	//Задание.
	//1. Вывести украинскую букву 'Ї'
	//2. Пояснить назначение типа "rune"

	// Виведення української букви 'Ї' та її Unicode коду
	fmt.Printf("\nЗавдання 1.\n")
	var ukrainianChar rune = 'Ї'
	fmt.Printf("Code '%c' - %U\n", ukrainianChar, ukrainianChar)

	fmt.Printf("\nЗавдання 2.\n")
	// Пояснення типу "rune"
	fmt.Println("\nПояснення:")
	fmt.Println("Тип 'rune' використовується для зберігання символів Unicode.")
	fmt.Println("Це синонім для 'int32', що дозволяє зберігати символи, які займають більше ніж 1 байт.")
	fmt.Println("Наприклад, буква 'Ї' має Unicode код, який займає більше 1 байта, тому для її зберігання використовують 'rune'.")
}
