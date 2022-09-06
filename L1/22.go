package main
import "fmt"
import "os"
import "bufio"
import "errors"
import "math/big"

// 22. Разработать программу, которая перемножает, делит, складывает, вычитает две числовых переменных a,b, значение которых > 2^20.

func main() {
	var n1 string
	var n2 string
	fmt.Print("Enter first number\n")
	fmt.Scan(&n1)
	fmt.Print("Enter second number\n")
	fmt.Scan(&n2)
	var op rune
	fmt.Print("Enter operation: ")
	reader := bufio.NewReader(os.Stdin)
	op, _, err := reader.ReadRune()
	if err != nil {
		fmt.Print(errors.New("Something went wrong...\n"))
	}
	a := new(big.Float)
	a.SetString(n1)
	b := new(big.Float)
	b.SetString(n2)
	c := big.NewFloat(0)
	switch op {
	case '*':
		fmt.Printf("%f * %f = %f\n", a, b, c.Mul(a, b))
	case '/':
		fmt.Printf("%f / %f = %f\n", a, b, c.Quo(a, b))
	case '+':
		fmt.Printf("%f + %f = %f\n", a, b, c.Add(a, b))
	case '-':
		fmt.Printf("%f - %f = %f\n", a, b, c.Sub(a, b))
	default:
		fmt.Print("Unknown rune\n")
	}
}