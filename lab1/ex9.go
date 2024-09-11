package main

import "fmt"

func main() {
	// Логічні змінні з ініціалізацією за замовчуванням (false) та явно присвоєними значеннями
	var first, second bool // за замовчуванням значення false
	var third bool = true  // явно присвоєно true
	fourth := !third       // протилежне значення third, тобто false
	var fifth = true       // явно присвоєно true

	// Виведення значень змінних
	fmt.Println("first  = ", first)       // false
	fmt.Println("second = ", second)      // false
	fmt.Println("third  = ", third)       // true
	fmt.Println("fourth = ", fourth)      // false
	fmt.Println("fifth  = ", fifth, "\n") // true

	// Логічне заперечення (!)
	fmt.Println("!true  = ", !true)        // false: заперечення true = false
	fmt.Println("!false = ", !false, "\n") // true: заперечення false = true

	// Логічне "і" (&&) — істинно, якщо обидва значення істинні
	fmt.Println("true && true   = ", true && true)         // true: обидва значення true
	fmt.Println("true && false  = ", true && false)        // false: одне значення false
	fmt.Println("false && false = ", false && false, "\n") // false: обидва значення false

	// Логічне "або" (||) — істинно, якщо хоча б одне значення істинне
	fmt.Println("true || true   = ", true || true)         // true: обидва значення true
	fmt.Println("true || false  = ", true || false)        // true: одне значення true
	fmt.Println("false || false = ", false || false, "\n") // false: обидва значення false

	// Порівняння чисел
	fmt.Println("2 < 3  = ", 2 < 3)        // true: 2 менше за 3
	fmt.Println("2 > 3  = ", 2 > 3)        // false: 2 більше за 3 — неправда
	fmt.Println("3 < 3  = ", 3 < 3)        // false: 3 не менше за 3
	fmt.Println("3 <= 3 = ", 3 <= 3)       // true: 3 менше або дорівнює 3
	fmt.Println("3 > 3  = ", 3 > 3)        // false: 3 не більше за 3
	fmt.Println("3 >= 3 = ", 3 >= 3)       // true: 3 більше або дорівнює 3
	fmt.Println("2 == 3 = ", 2 == 3)       // false: 2 не дорівнює 3
	fmt.Println("3 == 3 = ", 3 == 3)       // true: 3 дорівнює 3
	fmt.Println("2 != 3 = ", 2 != 3)       // true: 2 не дорівнює 3
	fmt.Println("3 != 3 = ", 3 != 3, "\n") // false: 3 дорівнює 3
}
