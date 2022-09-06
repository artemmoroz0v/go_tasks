package main
import "fmt"
import "errors"

// 23. Удалить i-ый элемент из слайса.

func remove(array []int, target int) (error, []int) {
	if target < 0 || target > len(array) - 1 {
		return errors.New("Please enter correct data!\n"), nil
	}
	return nil, append(array[:target], array[target + 1:]...)
}
func main() {
	array := []int{3, 5, 6, 8, 22, 32, 45, 4, 20, 121, 31, 232, 44, 214}
	err, array := remove(array, 8)
	if err == nil {
		fmt.Print("Here is our slice:\n", array)
	}
}