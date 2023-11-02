// *****************
// Comma ok em maps:
// *****************
// O comma ok é um recurso da linguagem Go que permite verificar se uma chave existe em um map
// Pois se tentarmos acessar uma chave que não existe em um map, o Go retorna o valor zero do tipo do valor do map
// E há casos em que o valor zero do tipo do valor do map é um valor válido
// O comma ok retorna true se a chave existir no map e false se a chave não existir no map
// Sintaxe: valor, ok := nomeDoMapa[chave]
// O comma ok é usado em conjunto com o if
// if valor, ok := nomeDoMapa[chave]; ok {}

package main

import "fmt"

func main() {
	// Criando um map
	var mapas map[string]string

	// Inicializando um map
	mapas = make(map[string]string)

	// Adicionando valores ao map
	mapas["nome"] = "João"
	mapas["sobrenome"] = "da Silva"

	fmt.Println()

	// Comma ok em maps:
	fmt.Println("Comma ok em maps:")

	// Exemplo 1:
	fmt.Println("A chave nome existe no map?")
	if valor, ok := mapas["nome"]; ok {
		fmt.Println("A chave existe no map e o valor é:", valor)
	} else {
		fmt.Println("A chave não existe no map")
	}

	fmt.Println()

	// Exemplo 2:
	fmt.Println("A chave idade existe no map?:")
	if valor, ok := mapas["idade"]; ok {
		fmt.Println("A chave existe no map e o valor é:", valor)
	} else {
		fmt.Println("A chave não existe no map")
	}

	fmt.Println()
}