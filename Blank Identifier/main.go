// *******************
// Blank identifier: _
// *******************
// O blank identifier é usado para ignorar valores de retorno de funções

package main

import "fmt"

func devolveNomeCompleto() (string, string) {
	return "João", "da Silva"
}


func main() {
	nome, _ := devolveNomeCompleto()
	fmt.Println("Blank identifier:")
	fmt.Println(nome)
}