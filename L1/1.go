package main
import "fmt"

/* 1. Дана структура Human (с произвольным набором полей и методов). 
Реализовать встраивание методов в структуре Action от родительской структуры Human (аналог наследования). */

type Human struct {
	Height int
	Weight int
}
func (human* Human) ShowPhysicalData() {
	fmt.Printf("Info: Height = %d, Weight = %d\n", human.Height, human.Weight)
}
type Action struct {
	Human
	Gender string
}
func (action* Action) GetInfo() {
	action.ShowPhysicalData()
	if action.Gender == "M" {
		fmt.Print("Male\n")
	} else {
		fmt.Print("Female\n")
	}
}

func main() {
	ArtemMorozov := Action {
					Human: Human {
							Height: 181,
							Weight: 79,
					},
					Gender: "M",
	}
	ArtemMorozov.GetInfo() 
	ArtemMorozov.ShowPhysicalData()
}
