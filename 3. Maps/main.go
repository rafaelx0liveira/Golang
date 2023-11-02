// *****
// Maps:
// *****
// Maps são estruturas de dados que armazenam dados em pares chave e valor
// Mapas não são ordenados
// Maps são conhecidos em outras linguagens como dicionários
// Sintaxe: var nomeDaVariavel map[tipoDaChave]tipoDoValor

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

	fmt.Println("Maps:")

	// Imprimindo o map
	fmt.Println(mapas)

	fmt.Println()

	// Acessando um valor do map
	fmt.Println("Acessando um valor do map:")
	fmt.Println(mapas["nome"])

	fmt.Println()

	// Deletando um valor do map
	fmt.Println("Deletando um valor do map:")
	delete(mapas, "nome")
	fmt.Println(mapas)

	fmt.Println()

	// Criando um map com valores
	var mapas2 = map[string]string{
		"nome": "João",
		"sobrenome": "da Silva",
	}

	// Imprimindo o map
	fmt.Println("Criando um map com valores:")
	fmt.Println(mapas2)

	fmt.Println()
}