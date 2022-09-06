package main
import "fmt"

/* 26. Разработать программу, которая проверяет, что все символы в строке уникальные (true — если уникальные, false etc). 
Функция проверки должна быть регистронезависимой.
Например: 
abcd — true
abCdefAaf — false
aabcd — false */

func unique(s string) bool {
	m := make(map[rune]bool)
	for _, i := range s {
		_, ok := m[i]
		if ok {
			return false
		}
		m[i] = true;
	}
	return true;
}

func main() {
	var s string
	fmt.Print("Enter string\n")
	fmt.Scan(&s)
	unique(s)
	ok := unique(s)
	switch ok {
	case true:
		fmt.Print("String is unique")
	case false:
		fmt.Print("String is not unique")
	}
}