package main
import "fmt"

// 11. Реализовать пересечение двух неупорядоченных множеств.

func traversal(first_array []int, second_array[]int) []int {
	set := []int{}
	numbers := make(map[int]bool)
	for _, i := range first_array {
		numbers[i] = true
	}
	for _, i := range second_array {
		if numbers[i] {
			set = append(set, i)
		}
	}
	return set
}
func main() {
	first_array := []int{4, 3, 121, 434, 44, 1, 22, 34, 1212, 443434, 23, 21, 653}
	second_array := []int{12, 323, 3, 4, 122, 1212, 121, 33333}
	result := traversal(first_array, second_array)
	fmt.Print(result)
}