// ********************************
// Ponteiros: Referência de memória
// ********************************
// Ponteiros são referências de memória
// Ponteiros são usados para acessar o endereço de memória de uma variável
// Em Go, você não precisa alocar memória manualmente, pois o Go possui um coletor de lixo (garbage collector) que faz isso automaticamente
// Sintaxe: var nomeDaVariavel *tipoDaVariavel

package main

import "fmt"

// Função para trocar valores
func trocaValores(valor *int) {
	*valor = 20
}

func main() {
	// Criando uma variável
	var nome string = "João"

	// Criando um ponteiro
	var ponteiro *string

	// Atribuindo o endereço de memória da variável ao ponteiro
	ponteiro = &nome

	fmt.Println("Ponteiros:")
	fmt.Println("Endereço de memória da variável:", &nome)
	fmt.Println("Valor da variável:", nome)
	fmt.Println("Endereço de memória do ponteiro:", ponteiro)
	fmt.Println("Valor do ponteiro:", *ponteiro)

	fmt.Println()

	// Alterando o valor da variável através do ponteiro
	fmt.Println("Alterando o valor da variável através do ponteiro:")
	*ponteiro = "João da Silva"
	fmt.Println("Valor da variável:", nome)
	fmt.Println("Valor do ponteiro:", *ponteiro)

	fmt.Println()

	// ***********************************
	// Passagem de Ponteiros para funções:
	// ***********************************
	// Em Go, os argumentos são passados por valor por padrão, mas você pode passar um ponteiro para uma função se desejar que a função modifique o valor original.
	// Sintaxe: func nomeDaFuncao(nomeDoPonteiro *tipoDoPonteiro) {}
	fmt.Println("Passagem de Ponteiros para funções:")
	valor := 10
	fmt.Println("Valor original:", valor)

	trocaValores(&valor)
	fmt.Println("Valor modificado:", valor)

	fmt.Println()
}