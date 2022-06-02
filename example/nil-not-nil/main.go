package main

// see https://youtu.be/ynoY2xz-F8s?t=885
import "fmt"

type doError struct{}

func (_ *doError) Error() string {
	return ""
}

func do() *doError { // nil of type *doError
	return nil
}

func wrapDo() error { // error (*doError,nil)
	return do() // nil of type *doError
}

func main() {
	err := wrapDo()         // error (*doError,nil)
	fmt.Println(err == nil) // false
}
