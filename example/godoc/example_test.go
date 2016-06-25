package example_test

import (
	"./"
)

func ExampleHello() {
	example.Hello()
	// Output:
	// Hello, 世界
}

func ExampleHello_two() {
	example.Hello()
	example.Hello()
	// Output:
	// Hello, 世界
	// Hello, 世界
}

func ExampleFoo() {
	example.Foo()
	// Output:
}
