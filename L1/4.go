package main
import "fmt"
import "os"
import "os/signal"
import "sync"
import "syscall"

/* 4. Реализовать постоянную запись данных в канал (главный поток). 
Реализовать набор из N воркеров, которые читают произвольные данные из канала и выводят в stdout.  
Необходима возможность выбора количества воркеров при старте.
Программа должна завершаться по нажатию Ctrl+C. Выбрать и обосновать способ завершения работы всех воркеров. */

func worker (index int, channel chan int, wg sync.WaitGroup) {
	for i := range channel {
		fmt.Printf("Worker №%d has received next message: %d\n", index + 1, i)
	}
	wg.Done()
}

func main() {
	var n int
	fmt.Print("Enter number of workers\n")
	fmt.Scan(&n)
	channel := make(chan int)
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go worker(i, channel, wg)
	}
	i := 0
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT)
	for {
		select {
		case <- sigs:
			close(channel)
			wg.Wait()
			return
		default:
			channel <- i
			i++
		}
	}
}