package main
import "fmt"

// 3. Дана последовательность чисел: 2,4,6,8,10. Найти сумму их квадратов(22+32+42….) с использованием конкурентных вычислений.

func square(number int, channel chan int) {
	channel <- number * number
}
func main() {
	array := []int{1, 2, 3, 4}
	result := 0
	channel := make(chan int)
	for _, i := range array {
		go square(i, channel)
	}
	for i := 0; i < len(array); i++ {
		result += <- channel
	}
	fmt.Print(result, "\n")
}