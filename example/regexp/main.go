//ref https://golang.org/pkg/regexp/
//ref go doc regexp/syntax
package main

import (
	"fmt"
	"regexp"
)

func Example1() {
	re := regexp.MustCompile("a(x*)b(y|z)c")
	fmt.Printf("%q\n", re.FindStringSubmatch("-axxxbyc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abzc-"))
	fmt.Printf("%q\n", re.FindStringSubmatch("-abcdef-"))
}

func Example2() {
	re := regexp.MustCompile("(?P<first>[a-zA-Z]+) (?P<last>[a-zA-Z]+)")
	fmt.Println(re.MatchString("Alan Turing"))
	fmt.Printf("%q\n", re.SubexpNames())
	reversed := fmt.Sprintf("${%s} ${%s}", re.SubexpNames()[2], re.SubexpNames()[1])
	fmt.Println(reversed)
	fmt.Println(re.ReplaceAllString("Alan Turing", reversed))
}

func main() {
	fmt.Println("=Example1=")
	Example1()
	fmt.Println("=Example2=")
	Example2()
}
