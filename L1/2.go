package main
import "fmt"
import "sync"

// 2. Написать программу, которая конкурентно рассчитает значение квадратов чисел взятых из массива (2,4,6,8,10) и выведет их квадраты в stdout.

func square (number int, wg* sync.WaitGroup) {
	result := number * number
	fmt.Print(result, " ")
	wg.Done()
}

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var wg sync.WaitGroup
	wg.Add(len(array))
	for _, i := range array {
		go square(i, &wg)
	}
	wg.Wait()
}