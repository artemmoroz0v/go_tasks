package main
import "fmt"

// 19. Разработать программу, которая переворачивает подаваемую на ход строку (например: «главрыба — абырвалг»). Символы могут быть unicode.

func reverseString(s string) string {
	chars := []rune(s)
	result := ""
	for i := len(chars) - 1; i >= 0; i-- {
		result += string(chars[i])
	}
	return result
}
func main() {
	s := reverseString("главрыба")
	fmt.Print(s)
}