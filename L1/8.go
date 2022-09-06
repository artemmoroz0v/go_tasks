package main
import "fmt"
import "strconv"

// 8. Дана переменная int64. Разработать программу которая устанавливает i-й бит в 1 или 0.

func set (i int, bit int, number int64) int64 {
	if bit == 1 {
		return number | (1 << i)
	} else {
		return number &^ (1 << i)
	}
}

func main() {
	var i, bit int
	var number int64
	fmt.Print("Enter number\n")
	fmt.Scan(&number)
	fmt.Print("Enter bit\n")
	fmt.Scan(&bit)
	fmt.Print("Enter i\n")
	fmt.Scan(&i)
	fmt.Printf("Before: %s\n", strconv.FormatInt(number, 2))
	number = set(i, bit, number)
	fmt.Printf("After: %s\n", strconv.FormatInt(number, 2))
}