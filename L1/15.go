package main

import (
	"fmt"
	"math/rand"
)

/* 15. К каким негативным последствиям может привести данный фрагмент кода, и как это исправить? Приведите корректный пример реализации.

var justString string
func someFunc() {
  v := createHugeString(1 << 10)
  justString = v[:100]
}
func main() {
  someFunc()
} */

func createHugeString(count int) string {
	array := []rune("0123456789абвгдеёжзийклмнопрстуфхцчшщъыьэюяАБВГДЕЁЖЗИЙКЛМНОПРСТУФХЦЧШЩЪЫЬЭЮЯabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	huge_string := ""
	for i := 0; i < count; i++ {
		huge_string += string(array[rand.Intn(len(array))])
	}
	return huge_string
}

func someFunc() string {
	return createHugeString(1 << 10)[:100]
}

func main() {
	fmt.Print(someFunc())
}

/*что плохого в изначальном коде:
1. ни к чему глобальные переменные
2. после этих строк v := createHugeString(1 << 10)
  justString = v[:100] 924 элемента будут неиспользуемыми, проще взять 100 символов от возвращаемого результата функции
*/