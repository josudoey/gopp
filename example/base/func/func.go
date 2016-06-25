//ref http://dancallahan.info/journal/go-overview/
package main

import "fmt"

func double(x int) int {
	return x + x
}

func square(x int) int {
	return x * x
}

func apply(f func(int) int, x int) int {
	return f(x)
}

func main() {
	fmt.Println(apply(double, 5))
	fmt.Println(apply(square, 5))
}
