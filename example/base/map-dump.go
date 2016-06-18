package main

//ref https://blog.golang.org/go-maps-in-action
import "fmt"

func main() {
	var m = make(map[string]string)
	m = map[string]string{
		"hello": "world",
	}
	n := len(m)
	fmt.Println("map len =", n)
	for key, value := range m {
		fmt.Println("Key:", key, "Value:", value)
	}
}
