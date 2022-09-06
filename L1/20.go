package main
import "fmt"
import "strings"

/* 20. Разработать программу, которая переворачивает слова в строке. 
Пример: «snow dog sun — sun dog snow». */

func reverseWords(s string) string {
	string_array := strings.Split(s, " ")
	result := ""
	for i := len(string_array) - 1; i >= 1; i-- {
		result += string_array[i] + " "
	}
	result += string_array[0]
	return result
}
func main() {
	result := reverseWords("snow dog sun")
	fmt.Print(result)
}