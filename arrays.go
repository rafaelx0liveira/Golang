// *******
// Arrays:
// *******
// Arrays são coleções de dados de tipo primitivo, ou composto, de tamanho fixo.
// Ou seja, Arrays são coleções de dados de tipo primitivo, ou composto, de tamanho fixo
// Sintaxe: var nomeDaVariavel [tamanho]tipo

package main

import "fmt"

func main() {
	var array [5]int

	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5
	
	fmt.Println("Arrays:")
	fmt.Println(array)

	fmt.Println()
}
