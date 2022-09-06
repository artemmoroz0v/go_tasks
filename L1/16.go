package main
import "fmt"

// 16. Реализовать быструю сортировку массива (quicksort) встроенными методами языка.

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p - 1)
		arr = quickSort(arr, p + 1, high)
	}
	return arr
}

func main() {
	array := []int{13, 121, 10, 4, 2, 6, 9, 11, 342, 124, 44, 21, 12, 19, 20, 24, 25, 28}
	fmt.Printf("Current array: %v\n", array)
	array = quickSort(array, 0, len(array) - 1)
	fmt.Printf("After quicksort: %v\n", array)
}