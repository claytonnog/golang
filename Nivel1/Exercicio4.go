package main

import "fmt"

type blabla int

var x blabla

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x) //o %T monstra o tipo da variavel x
	x = 42
	fmt.Println(x)
}
