// *****************************
// Operador curto de declaração:
// *****************************
// O operador curto de declaração é usado para declarar e atribuir valores a variáveis
// nomeDaVariavel := valorDaVariavel
// Operador curto de declaração só pode ser usado dentro de funções por conta do escopo

package main

import "fmt"

func main() {
	nomeDaVariavel := "João"
	fmt.Println("Operador curto de declaração:")
	fmt.Println(nomeDaVariavel)
}