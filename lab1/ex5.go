package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Синонимы целых типов\n")

	fmt.Println("byte    - int8")
	fmt.Println("rune    - int32")
	fmt.Println("int     - int32, или int64, в зависимости от ОС")
	fmt.Println("uint    - uint32, или uint64, в зависимости от ОС")

	//Задание.
	//1. Определить разрядность ОС
	fmt.Println("Архитектура ОС: ", runtime.GOARCH)
}
