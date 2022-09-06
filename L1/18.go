package main
import "fmt"
import "sync"

/* 18. Реализовать структуру-счетчик, которая будет инкрементироваться в конкурентной среде. 
По завершению программа должна выводить итоговое значение счетчика. */

type Counter struct {
	counter int
	mutex sync.RWMutex
}

func (c *Counter) Increment(wg *sync.WaitGroup) {
	// блокируем мьютекс на запись
	c.mutex.Lock()
	c.counter++
	c.mutex.Unlock()
	wg.Done()
}

func (c *Counter) Result() int {
	// блокируем мьютекc на чтение
	c.mutex.RLock()
	c.mutex.RUnlock()
	return c.counter
}

func main() {
	var c Counter
	fmt.Print("Before: ", c.counter, "\n")
	var wg sync.WaitGroup
	goroutines_count := 1000
	wg.Add(goroutines_count)
	for i := 0; i < goroutines_count; i++ {
		go c.Increment(&wg)
	}
	wg.Wait()
	result := c.Result()
	fmt.Print("After: ", result)
}