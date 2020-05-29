package main

import "fmt"

type blabla int

var x blabla

var y int

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x) //mostra o tipo da variavel x
	x = 42                //agora que definimos um valor para a variavel X
	fmt.Println(x)

	fmt.Println("valor da variavel X ACIMA\n")
	fmt.Println("valor da variavel Y ABAIXO")

	y = int(x) //y recebe o valor da variavel x tipo inteira (int)

	fmt.Println(y)
	fmt.Printf("%T\n", y) // verifique que o tipo da variavel mudou ...

}
