package main

import "fmt"

var x int = 42
var y string = "James Bond"
var z bool = true

func main() {

	s := fmt.Sprintf("%v\t%v\t%v", x, y, z) //atribuimos a variavel s o resultado do sprintf das variaveis x y z
	fmt.Println(s)
}
