// *******
// Slices:
// *******
// Uma coleção de dados de tipo primitivo, ou composto, de tamanho dinâmico
// Ou seja, Slices são arrays com tamanho dinâmico
// Sintaxe: var nomeDaVariavel []tipo

package main

import "fmt"

func main() {
	var slice []int

	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)

	fmt.Println("Slices:")
	fmt.Println(slice)

	fmt.Println()

	
	// ******
	// Range:
	// ******
	// Range é um laço de repetição que itera sobre itens de uma coleção de dados, como arrays, slices, maps, strings, etc.
	// Sintaxe: for indice, valor := range nomeDaVariavel {}
	fmt.Println("Range:")
	for indice, valor := range slice {
		fmt.Println("Indice:", indice, "Valor:", valor)
	}

	fmt.Println()

	// ****************
	// Fatiando slices:
	// ****************
	// Podemos fatiar um slice usando o operador de fatiamento, que é o dois pontos (:). 
	// Para acessar um item do slice, usamos o índice do item, que começa em 0
	// Exemplo: Se temos um slice com 5 itens, o índice do primeiro item é 0 e o índice do último item é 4.
	// Para acessar o terceiro item, usamos o índice 2, ou seja, slice[2]
	// Analogia: "Quero que comece no índice 2 e vá até o índice 4". Obs.: O índice 4 não é incluído
	// Sintaxe: slice[inicio:fim]
	fmt.Println("Fatiando slices:")

	fmt.Println("Fatiando 0:2:")
	fmt.Println(slice[0:2])

	fmt.Println("Fatiando 1:3:")
	fmt.Println(slice[1:3])

	fmt.Println("Fatiando 2:4:")
	fmt.Println(slice[2:4])

	fmt.Println("Acessando o último elemento:")
	fmt.Println(slice[len(slice) - 1])

	fmt.Println("Acessando todos os itens do Slice sem usar o for range:")
	fmt.Println(slice[:])

	fmt.Println()

	// ****************************
	// Deletando itens de um slice:
	// ****************************
	// Para deletar um item de um slice, usamos o append para criar um novo slice sem o item que queremos deletar
	// Sintaxe: slice = append(slice[:indice], slice[indice + 1:]...)
	// Deletando o primeiro item
	// Obs.: Os três pontos (...) são chamados de ellipsis
	fmt.Println("Deletando slice:")

	fmt.Println("Deletando o primeiro item:")
	slice = append(slice[:0], slice[1:]...)
	fmt.Println(slice)

	// Deletando o último item
	fmt.Println("Deletando o último item:")
	slice = append(slice[:len(slice) - 1], slice[len(slice):]...)
	fmt.Println(slice)

	fmt.Println()

	// *****************
	// Agrupando slices:
	// *****************
	// Podemos agrupar dois ou mais slices usando o append
	// Sintaxe: slice = append(slice1, slice2...)
	fmt.Println("Agrupando slices:")
	slice1 := []int{1, 2, 3}

	slice2 := append(slice1, 4, 5, 6)

	fmt.Println("Primeiro slice:")
	fmt.Println(slice1)

	fmt.Println("Slice agrupado:")
	fmt.Println(slice2)

	fmt.Println()

	// ***********************
	// Make: Criação de slices
	// ***********************
	// O make é uma função que cria slices
	// o make é frequentemente utilizado para criar slices, maps e channels.
	// O make é usado para criar slices com tamanho e capacidade definidos
	// Ou seja, o make aloca espaço para o array subjacente, inicializa o slice e define sua capacidade inicial.
	// Sintaxe: nomeDaVariavel := make([]tipo, tamanho, capacidade)
	// tamanho: Tamanho do slice
	// capacidade: Tamanho do array interno, ou seja, o tamanho do array interno que o slice criado pelo make faz referência
	// O make é mais performático porque já cria um array interno com o tamanho e capacidade definidos, sem a necessidade de realocar memória sempre que o slice crescer
	fmt.Println("Make:")
	slice3 := make([]int, 5, 10)
	slice3[0] = 1
	slice3[1] = 2
	slice3[2] = 3
	slice3[3] = 4
	slice3[4] = 5
	fmt.Println(slice3)

	fmt.Println()

	// ************************
	// Slices multidimensionais
	// ************************
	// Slices multidimensionais são slices de slices, ou seja, um slice que contém outros slices
	// Sintaxe: var nomeDaVariavel [][]tipo
	sliceMultidimensional := [][]int{
		{1, 2, 3},
		{4, 5, 6},
	}
	
	fmt.Println("Slices multidimensionais:")
	fmt.Println(sliceMultidimensional)

	fmt.Println("Acessando o primeiro slice:")
	fmt.Println(sliceMultidimensional[0])

	fmt.Println("Acessando o primeiro item do primeiro slice:")
	fmt.Println(sliceMultidimensional[0][0])

	fmt.Println("Acessando o primeiro item do segundo slice:")
	fmt.Println(sliceMultidimensional[1][0])

	fmt.Println()

	// ****************************************
	// A surpresa do array subjacente no slice:
	// ****************************************
	// Quando criamos um slice, o Go cria um array interno para armazenar os dados do slice
	// Ou seja, arrays subjacentes são a referencia de memória usada pelos slices. 
	// Como se o array subjacente fosse uma caixa com tamanho definido, e o slice é uma caixa com tamanho variável que aponta para a referencia de memória do array subjacente

	// Quando criamos um slice a partir de outro slice, o Go cria um novo slice que faz referência ao mesmo array interno do slice original
	// Ou seja, se alterarmos um item do slice original, o item do slice criado a partir do slice original também será alterado
	// Isso acontece porque o slice criado a partir do slice original faz referência ao mesmo array interno do slice original
	// Para resolver esse problema, podemos usar o make para criar um novo slice com um novo array interno
	fmt.Println("A surpresa do array subjacente no slice:")
	
	// Criando um array
	array_novo := [3]int{1, 2, 3}
	fmt.Println("Array original:", array_novo)

	// Criando duas slices que compartilham o mesmo array subjacente
	primeiro_slice := array_novo[0:2]
	fmt.Println("Recortando os dois primeiros itens do array original:")
	fmt.Println("Slice1:", primeiro_slice)

	segundo_slice := array_novo[1:3]
	fmt.Println("Recortando os dois últimos itens do array original:")
	fmt.Println("Slice2:", segundo_slice)

	// Modificando um elemento em slice1]
	fmt.Println("Modificando o índice zero em slice1:")
	primeiro_slice[0] = 99
	fmt.Println("Slice1:", primeiro_slice)

	// Imprimindo o array original e as duas slices
	fmt.Println("Array original:", array_novo)
	fmt.Println("Slice1:", primeiro_slice)
	fmt.Println("Slice2:", segundo_slice)

	fmt.Println()
}
