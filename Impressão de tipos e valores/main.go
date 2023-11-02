// *******************
// Impressão de tipos:
// *******************
// Em Go, podemos imprimir o tipo de uma variável
// %T: Tipo

// *********************
// Impressão de valores:
// *********************
// %v: Valor

package main

import "fmt"

func main() {

	var variavel string = "João"

	fmt.Println("Impressão de tipos:")
	fmt.Printf("Tipo da variável: %T\n", variavel)

	fmt.Println()

	fmt.Println("Impressão de valores:")
	fmt.Printf("Valor da variável: %v\n", variavel)

	fmt.Println()
}