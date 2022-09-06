package main
import "fmt"

// 17. Реализовать бинарный поиск встроенными методами языка.

func binary_search(array []int, target int) bool {
	left := 0
	right := len(array) - 1
	for left <= right {
		middle := (left + right) / 2
		if (array[middle] == target) {
			return true
		} else if (array[middle] > target) {
			right = middle - 1
		} else {
			left = middle + 1
		}
	}
	return false
}

func main() {
	array := []int{1, 4, 7, 9, 10, 13, 15, 18, 19, 22}
	res1 := binary_search(array, 2)
	res2 := binary_search(array, 5)
	res3 := binary_search(array, 10)
	res4 := binary_search(array, 13)
	res5 := binary_search(array, 20)
	res6 := binary_search(array, 22)
	fmt.Print(res1, res2, res3, res4, res5, res6)
}