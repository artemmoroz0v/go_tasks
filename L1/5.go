package main
import "fmt"
import "time"

/* 5. Разработать программу, которая будет последовательно отправлять значения в канал, а с другой стороны канала — читать. 
По истечению N секунд программа должна завершаться. */

func send (channel chan string) {
	for {
		channel <- "Hello"
	}
}
func main() {
	fmt.Print("Enter seconds\n")
	var n int
	fmt.Scan(&n)
	channel := make(chan string)
	go send(channel)
	countdown := time.After(time.Duration(n) * time.Second)
	for {
		select {
		case <- countdown:
			return
		default:
			fmt.Print("Received a message: " , <- channel, "\n")
		}
	}
}