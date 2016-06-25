package main

// #include <stdio.h>
// void hello()
// {
//	    printf("hello\n");
// }
import "C"

func main() {
	C.hello()
}
