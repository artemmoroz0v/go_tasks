package main
import "fmt"

/* 9. Разработать конвейер чисел. Даны два канала: в первый пишутся числа (x) из массива, 
во второй — результат операции x*2, после чего данные из второго канала должны выводиться в stdout. */

func first (array[] int, first_channel chan int) {
	for _, i := range array {
		first_channel <- i
	}
	close(first_channel)
}
func second (array[] int, second_channel chan int, first_channel chan int) {
	for i := range first_channel {
		second_channel <- i * 2
	}
	close(second_channel)
}
func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	first_channel := make(chan int)
	second_channel := make(chan int)
	go first(array, first_channel)
	go second(array, second_channel, first_channel)
	for i := range second_channel {
		fmt.Print(i, " ")
	}
}