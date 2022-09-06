package main
import "fmt"

// 13. Поменять местами два числа без создания временной переменной.

func main() {
	a := 5
	b := 3
	fmt.Printf("Before swapping: a = %d, b = %d\n", a, b)
	a, b = b, a
	fmt.Printf("Before swapping: a = %d, b = %d\n", a, b)
}