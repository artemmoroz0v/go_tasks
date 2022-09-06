package main
import "fmt"

// 14. Разработать программу, которая в рантайме способна определить тип переменной: int, string, bool, channel из переменной типа interface{}.

func check(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Print("This is string!\n")
	case int:
		fmt.Print("This is int!\n")
	case bool:
		fmt.Print("This is bool!\n")
	case chan int:
		fmt.Print("This is chan int!\n")
	}
}
func main() {
	check(interface{}("europe"))
	check(interface{}(100))
	check(interface{}(false))
	check(interface{}(make(chan int)))
}