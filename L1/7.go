package main
import "fmt"
import "sync"

// 7. Реализовать конкурентную запись данных в map.

type Map struct {
	things map[int]int
	mutex sync.RWMutex
}
func Create() Map {
	return Map{things: make(map[int]int)}
}
func (m* Map) ShowThings() {
	m.mutex.RLock()
	for i, j := range m.things {
		fmt.Print("Key: " , i, ", " , "Value: ", j, "\n")
	}
	m.mutex.RUnlock()
}
func (m* Map) Take(key int) int {
	m.mutex.RLock()
	m.mutex.RUnlock()
	return m.things[key]
}
func (m* Map) Insert(key int, value int, wg *sync.WaitGroup) {
	m.mutex.Lock()
	m.things[key] = value
	m.mutex.Unlock()
	wg.Done()
}

func main() {
	var n int
	fmt.Print("Enter N\n")
	fmt.Scan(&n)
	mymap := Create()
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i <= n; i++ {
		go mymap.Insert(i, i + 1, &wg)
	}
	wg.Wait()
	mymap.ShowThings()
}