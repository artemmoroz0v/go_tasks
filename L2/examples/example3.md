# 3. Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.
```go
package main
 
import (
    "fmt"
    "os"
)
 
func Foo() error {
    var err *os.PathError = nil
    return err
}
 
func main() {
    err := Foo()
    fmt.Println(err)
    fmt.Println(err == nil)
}
```
Будет выведено ```<nil> false```. Указатель в го - пара, содержащая тип, на который указывает указатель, а также адрес в памяти. Значение интерфейса равно `nil` только когда и значение, и динамический тип - `nil`.
В примере `Foo()` возвращает `[nil, *os.PathError]`, а сравниваем мы это с `[nil, nil]`.
Для корректного сравнения можно было бы возвращать `*os.PathError` из функции.
