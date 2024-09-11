package main

//Импорт нескольких пакетов
import (
	"fmt"
	"math"
)

func main() {
	var defaultFloat float32
	var defaultDouble float64 = 5.5

	fmt.Println("defaultfloat       = ", defaultFloat)
	fmt.Printf("defaultDouble (%T) = %f\n\n", defaultDouble, defaultDouble)

	fmt.Println("MAX float32        = ", math.MaxFloat32)
	fmt.Println("MIN float32        = ", math.SmallestNonzeroFloat32, "\n")

	fmt.Println("MAX float64        = ", math.MaxFloat64)
	fmt.Println("MIN float64        = ", math.SmallestNonzeroFloat64, "\n")

	//Задание.
	//1. Создайте переменные разных типов, используя краткую запись и инициализацию по-умолчанию. Результат вывести

	// Створюємо змінні різних типів, використовуючи коротку форму запису і ініціалізацію за замовчуванням
	intVar := 10             // int
	floatVar := 3.14         // float64
	stringVar := "Some_Text" // string
	boolVar := true          // bool

	// Виведення результатів
	fmt.Println("Значення змінних:")
	fmt.Printf("intVar (int)      = %d\n", intVar)
	fmt.Printf("floatVar (float64)= %f\n", floatVar)
	fmt.Printf("stringVar (string)= %s\n", stringVar)
	fmt.Printf("boolVar (bool)    = %t\n", boolVar)
}
