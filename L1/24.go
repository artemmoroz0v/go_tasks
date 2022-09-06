package main
import "fmt"
import "math"

/* 24. Разработать программу нахождения расстояния между двумя точками, которые представлены в виде структуры Point 
с инкапсулированными параметрами x,y и конструктором. */

type Point struct {
	x int
	y int
}
func Constructor(x int, y int) Point {
	return Point{x: x, y: y}
}
func (p1* Point) Distance(p2 Point) float64 {
	return math.Sqrt(float64((p1.x - p2.x) * (p1.x - p2.x) + (p1.x - p2.x) * (p1.y - p2.y)))
}
func main() {
	var x1, x2 int
	var y1, y2 int
	fmt.Print("Input first point\n")
	fmt.Scan(&x1, &y1)
	fmt.Print("Input second point\n")
	fmt.Scan(&x2, &y2)
	p1 := Constructor(x1, y1)
	p2 := Constructor(x2, y2)
	distance := p1.Distance(p2)
	fmt.Print("Distance is:\n", distance)
}