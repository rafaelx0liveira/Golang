// ******************************
// Structs: Agrupamento de dados:
// ******************************
// Structs (abreviação de "structure")  são tipos de dados compostos que agrupam dados de tipos diferentes
// Em resumo, structs são agrupamentos de dados de tipos diferentes
// Structs são usados para criar estruturas de dados complexas que representam registros, objetos ou qualquer outra entidade que tenha propriedades distintas.
// Sintaxe: var nomeDaVariavel struct {
// 	nomeDaVariavel tipoDaVariavel
// }

package main

import "fmt"

// Definindo um struct
type usuario struct {
	nome string
	idade int
}

// Criando um método para o struct
func (u usuario) salvar() {
	fmt.Println("Salvando os dados do usuário:", u.nome)
}

func main() {

	// Instanciando um struct pela forma literal (Sem o new)
	usuario1 := usuario{nome: "João", idade: 10}

	// Instanciando um struct ocultando os nomes dos campos
	usuario2 := usuario{"Maria", 20}

	fmt.Println("Structs:")
	fmt.Println(usuario1)
	fmt.Println(usuario2)

	fmt.Println()

	// Acessando um campo do struct
	fmt.Println("Acessando um campo do struct:")
	fmt.Println(usuario1.nome)
	fmt.Println(usuario2.idade)

	fmt.Println()

	// Alterando um campo do struct
	fmt.Println("Alterando um campo do struct:")
	usuario1.nome = "João da Silva"
	usuario2.idade = 30
	fmt.Println(usuario1)
	fmt.Println(usuario2)

	fmt.Println()

	// *****************
	// Structs anônimos:
	// *****************
	// Structs anônimos são structs sem o identificador (nome)
	// Structs anônimos são usados para criar structs temporárias
	// Sintaxe: var nomeDaVariavel struct {nomeDaVariavel tipoDaVariavel}

	// Criando um struct anônimo
	usuario3 := struct {
		nome string
		idade int
	}{
		nome: "João",
		idade: 10,
	}

	fmt.Println("Structs anônimos:")
	fmt.Println(usuario3)

	fmt.Println()

	// ***********************
	// Composition: Composição
	// ***********************
	// Composition é um recurso da linguagem Go que permite criar structs compostas por campos de uma ou mais structs existentes
	// A struct composta herda os campos das structs existentes
	// A composição de structs é uma técnica poderosa para criar estruturas de dados modulares e reutilizáveis
	// Sintaxe: var nomeDaVariavel struct {nomeDaVariavel tipoDaVariavel; nomeDaVariavel tipoDaVariavel}

	// Criando um struct
	type endereco struct {
		rua string
		numero int
	}

	// Criando um struct
	type cliente struct {
		nome string
		idade int
		endereco endereco
	}

	// Instanciando um struct
	cliente1 := cliente{
		nome: "João",
		idade: 10,
		endereco: endereco{
			rua: "Rua 1",
			numero: 10,
		},
	}

	fmt.Println("Composição de structs:")
	fmt.Println(cliente1)

	fmt.Println()

	// ******************************
	// Atrelando métodos a um struct:
	// ******************************
	// Em Go, podemos atrelar métodos a um struct
	// Em Java, dizemos que um método é um comportamento de uma classe. Em Go, dizemos que um método é um comportamento de um struct
	// Sintaxe: func (nomeDaVariavel nomeDoStruct) nomeDoMetodo() tipoDoRetorno {}

	// Criando um struct
	usuario4 := usuario{nome: "Rafael", idade: 10}

	fmt.Println("Atrelando métodos a um struct:")
	usuario4.salvar()

	fmt.Println()
}