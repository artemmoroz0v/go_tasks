package main
import "fmt"
import "context"
import "sync"
import "time"

// 6. Реализовать все возможные способы остановки выполнения горутины.

func channel_break(channel chan int, wg *sync.WaitGroup) {
	for {
		select {
		case <- channel:
			fmt.Print("Goodbye!\n")
			wg.Done()
			return
		default:
			fmt.Print("Waiting for a message...\n")
			time.Sleep(time.Second * 2)
		}
	}
}

func channel_close(channel chan int, wg *sync.WaitGroup) {
	for i := range channel {
		fmt.Print(i, "\n")
	}
	fmt.Print("Goodbye!\n")
	wg.Done()
	return
}

func context_break(context context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-context.Done():
			fmt.Print("Goodbye!\n")
			wg.Done()
			return
		default:
			fmt.Print("Waiting for a message...\n")
			time.Sleep(time.Second * 2)
		}
	}
}

func main() {
	/*channel break
	var wg sync.WaitGroup
	wg.Add(1)
	channel := make(chan int)
	go channel_break(channel, &wg)
	time.Sleep(time.Second * 2)
	channel <- 5
	wg.Wait() */


	/*channel close
	var wg sync.WaitGroup
	wg.Add(1)
	channel := make(chan int)
	go channel_close(channel, &wg)
	for i := 0; i < 5; i++ {
		channel <- i
	}
	close(channel)
	wg.Wait() */


	/*context break
	var wg sync.WaitGroup
	context, cancel := context.WithCancel(context.Background())
	wg.Add(1)
	go context_break(context, &wg)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait() */
}