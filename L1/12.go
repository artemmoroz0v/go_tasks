package main
import "fmt"

// 12. Имеется последовательность строк - (cat, cat, dog, cat, tree) создать для нее собственное множество.

func set(array []string) []string {
	strings := make(map[string]bool)
	for _, i := range array {
		strings[i] = true
	}
	set := make([]string, 0)
	for j := range strings {
		set = append(set, j)
	}
	return set
}
func main() {
	array := []string{"cat", "cat", "dog", "cat", "tree"}
	array = set(array)
	fmt.Print(array)
}