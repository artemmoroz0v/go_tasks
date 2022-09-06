package main
import "fmt"
import "time"

// 25. Реализовать собственную функцию sleep.

func homemade_sleep(clock time.Duration) {
	<-time.After(clock)
}
func main() {
	var clock time.Duration
	fmt.Print("Enter time for sleeping\n")
	fmt.Scan(&clock)
	fmt.Print("Here we are before\n")
	homemade_sleep(clock * time.Second)
	fmt.Print("Here we are after\n")
}