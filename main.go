package main

import (
	"fmt"
)

// Criando um struct Carro
type Carro struct {
	marca string
	modelo string
}

// Criando um método para o struct pessoa
func (c Carro) Andar() {
	fmt.Println("O carro", c.marca, c.modelo, "está andando...")
}

func trocaValores(ptr *int) {
	*ptr = 20
}

func main() {
	// Hello World! simples 
	var hello string = "Hello World!"
	fmt.Println(hello)

	fmt.Println()

	// *******************************
	// Tipos de dados primitivos em Go
	// *******************************
	// var nomeDaVariavel tipoDaVariavel
	var nome string
	var idade int
	var versao float32
	var ehMaior bool

	// *******************************
	// Tipos de dados Compostos em Go:
	// *******************************
	// Em Go, os tipos de dados compostos são tipos que consistem em uma combinação de valores de outros tipos. 
	// O ato de definir, criar, estruturar tipos de dados compostos se chama composição
	// Os tipos de dados compostos mais comuns em Go são:
	// Arrays, Slices, Maps, Structs, Channels, Slices de Slices, Maps de Maps, etc.

	// Atribuindo valores
	nome = "João"
	idade = 10
	versao = 1.1
	ehMaior = idade > 18

	// Imprimindo valores
	fmt.Println("Variáveis:")
	fmt.Println(nome, idade, versao, ehMaior)

	fmt.Println()

	// *****************************
	// Operador curto de declaração:
	// *****************************
	// O operador curto de declaração é usado para declarar e atribuir valores a variáveis
	// nomeDaVariavel := valorDaVariavel
	// Operador curto de declaração só pode ser usado dentro de funções por conta do escopo
	nomeDaVariavel := "João"
	fmt.Println("Operador curto de declaração:")
	fmt.Println(nomeDaVariavel)

	fmt.Println()

	// *******************
	// Impressão de tipos:
	// *******************
	// Em Go, podemos imprimir o tipo de uma variável 
	// %T: Tipo
	fmt.Println("Tipos:")
	fmt.Printf("Tipo de nomeDaVariavel: %T\n", nomeDaVariavel)

	fmt.Println()

	// *********************
	// Impressão de valores:
	// *********************
	// %v: Valor
	fmt.Println("Valores:")
	fmt.Printf("Valor de nomeDaVariavel: %v\n", nomeDaVariavel)

	fmt.Println()

	// *******************
	// Blank identifier: _
	// *******************
	// O blank identifier é usado para ignorar valores de retorno de funções
	_, erro := fmt.Println("Blank identifier", hello)
	fmt.Println(erro)

	fmt.Println()

	// ***********
	// Constantes:
	// ***********
	// Constantes são valores imutáveis, ou seja, não podem ser alterados
	// const nomeDaConstante tipoDaConstante = valorDaConstante
	const constante string = "Constante"
	fmt.Println(constante)

	fmt.Println()

	// ****************
	// Tipos numéricos:
	// ****************
	// int8, int16, int32, int64, int
	// uint8, uint16, uint32, uint64, uint
	// byte: alias para uint8
	// rune: alias para int32
	// uint8: 0 a 255
	// uint16: 0 a 65535
	// uint32: 0 a 4294967295
	// uint64: 0 a 18446744073709551615
	// uint: 0 a 18446744073709551615
	// int8: -128 a 127
	// int16: -32768 a 32767
	// int32: -2147483648 a 2147483647
	// int64: -9223372036854775808 a 9223372036854775807
	// int: -9223372036854775808 a 9223372036854775807
	// float32: 1.18E-38 a 3.4E+38
	// float64: 2.23E-308 a 1.8E+308
	// complex64: 32 bits para a parte real e 32 bits para a parte imaginária
	// complex128: 64 bits para a parte real e 64 bits para a parte imaginária

	// *******
	// Arrays:
	// *******
	// Arrays são coleções de dados de tipo primitivo, ou composto, de tamanho fixo.
	// Ou seja, Arrays são coleções de dados de tipo primitivo, ou composto, de tamanho fixo
	// Sintaxe: var nomeDaVariavel [tamanho]tipo
	var array [5]int
	array[0] = 1
	array[1] = 2
	array[2] = 3
	array[3] = 4
	array[4] = 5
	fmt.Println("Arrays:")
	fmt.Println(array)

	fmt.Println()

	// *******
	// Slices:
	// *******
	// Uma coleção de dados de tipo primitivo, ou composto, de tamanho dinâmico
	// Ou seja, Slices são arrays com tamanho dinâmico
	// Sintaxe: var nomeDaVariavel []tipo
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
	fmt.Println("Slices for range:")
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
	// slice[inicio:fim]
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

	// *****
	// Maps:
	// *****
	// Maps são estruturas de dados que armazenam dados em pares chave e valor
	// Mapas não são ordenados
	// Maps são conhecidos em outras linguagens como dicionários
	// Sintaxe: var nomeDaVariavel map[tipoDaChave]tipoDoValor
	var mapa map[string]string
	mapa = make(map[string]string) // Inicializa o mapa do tipo string para string
	mapa["nome"] = "João"
	mapa["idade"] = "10"
	fmt.Println("Maps:")
	fmt.Println(mapa)

	fmt.Println()

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
	fmt.Println("Comma ok em maps:")
	sera, ok := mapa["nome"]

	fmt.Println("A chave nome existe no map?")
	fmt.Println(sera, ok)

	sera, ok = mapa["sobrenome"]
	fmt.Println("A chave sobrenome existe no map?")
	fmt.Println(sera, ok)

	fmt.Println("Utilizando o comma ok em conjunto com o if:")
	if valor, ok := mapa["nome"]; ok {
		fmt.Println("A chave nome existe no map?")
		fmt.Println(valor, ok)
	}

	fmt.Println()

	// ******************************
	// Structs: Agrupamento de dados:
	// ******************************
	// Structs (abreviação de "structure")  são tipos de dados compostos que agrupam dados de tipos diferentes
	// Em resumo, structs são agrupamentos de dados de tipos diferentes
	// Structs são usados para criar estruturas de dados complexas que representam registros, objetos ou qualquer outra entidade que tenha propriedades distintas.
	// Sintaxe: var nomeDaVariavel struct {
	// 	nomeDaVariavel tipoDaVariavel
	// }

	// Definindo um struct
	type usuario struct {
		nome string
		idade int
	}
	// Instanciando um struct pela forma literal (Sem o new)
	usuario1 := usuario{nome: "João", idade: 10}

	// Instanciando um struct ocultando os nomes dos campos
	usuario2 := usuario{"Maria", 20}

	fmt.Println("Structs:")
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
	// type Carro struct {
	// 	marca string
	// 	modelo string
	// }

	// Criando um método para o struct pessoa
	// func (c Carro) Andar() {
	// 	fmt.Println("O carro", c.marca, c.modelo, "está andando...")
	// }

	// Instanciando um struct
	carro1 := Carro{
		marca: "Fiat",
		modelo: "Uno",
	}

	fmt.Println("Atrelando métodos a um struct:")
	carro1.Andar()

	fmt.Println()

	// *************************
	// Goroutines: Threads leves
	// *************************
	// Uma goroutine é uma função que é executada concorrentemente com outras goroutines. 
	// É uma unidade de execução leve que é gerenciada pelo runtime de Go, o que significa que você pode criar um grande número de goroutines em um programa Go sem consumir muita memória.
	// Execução Concorrente: As goroutines permitem a execução concorrente de tarefas. Isso é útil para lidar com tarefas que podem ser executadas de forma independente, como requisições de rede, cálculos paralelizáveis, ou qualquer tarefa que precise ser executada em segundo plano sem bloquear a execução principal.
	// Comunicação entre Goroutines: As goroutines podem se comunicar entre si por meio de canais (channels), que são uma forma segura de trocar dados entre goroutines. Os canais facilitam a coordenação e a sincronização entre as goroutines.
	// Sincronização: Go fornece mecanismos de sincronização, como canais, mutexes e primitivas de sincronização, para garantir que as goroutines cooperem de maneira segura e coordenada. Isso ajuda a evitar problemas como condições de corrida e torna a programação concorrente mais robusta.
	// Fechamento de goroutines: As goroutines são automaticamente encerradas quando a função que estão executando retorna. Isso significa que você não precisa gerenciar manualmente o ciclo de vida de uma goroutine.
	// As goroutines são uma parte fundamental da concorrência em Go e são uma das razões pelas quais a linguagem é tão eficaz no desenvolvimento de sistemas concorrentes e paralelos. Elas facilitam a criação de programas que podem tirar proveito de múltos núcleos de CPU e lidar com tarefas simultâneas de forma elegante e eficiente.
	// Goroutine são uma forma de execução concorrente ou paralela de funções
	// Goroutines são usadas para executar funções em paralelo
	// Em comparativo, Threads pesam 1MB e Goroutines pesam 2KB
	// Sintaxe: go nomeDaFuncao()

	// Criando uma goroutine
	go func() {
		fmt.Println("Goroutines:")
		fmt.Println("Hello World!")
	}()

	fmt.Println()

	// **************************************
	// Channels: Comunicação entre goroutines
	// **************************************
	// Channels são canais de comunicação entre goroutines, que são threads leves
	// Channels são usados para enviar e receber dados entre goroutines
	// Channels são tipos de dados compostos
	// Channels são criados com a função make
	// Sintaxe: var nomeDaVariavel chan tipoDaVariavel
	var canal chan string
	canal = make(chan string)

	go func() {
		canal <- "Hello World!"
	}()

	mensagem := <- canal
	fmt.Println("Channels:")
	fmt.Println(mensagem)

	fmt.Println()

	// ********************************
	// Ponteiros: Referência de memória
	// ********************************
	// Ponteiros são referências de memória
	// Ponteiros são usados para acessar o endereço de memória de uma variável
	// Em Go, você não precisa alocar memória manualmente, pois o Go possui um coletor de lixo (garbage collector) que faz isso automaticamente
	// Sintaxe: var nomeDaVariavel *tipoDaVariavel

	// Declaração de ponteiros
	var ponteiro *int // Ponteiro para um inteiro

	// Atribuindo o endereço de memória de uma variável para um ponteiro
	var variavel int = 10
	ponteiro = &variavel

	// Imprimindo o endereço de memória de uma variável
	fmt.Println("Ponteiros:")
	fmt.Println("Endereço de memória da variável:", &variavel)

	// Imprimindo o endereço de memória de um ponteiro
	fmt.Println("Endereço de memória do ponteiro:", ponteiro)

	// Desreferenciando um ponteiro
	fmt.Println("Desreferenciando um ponteiro:", *ponteiro)

	fmt.Println()

	// ***********************************
	// Passagem de Ponteiros para funções:
	// ***********************************
	// Em Go, os argumentos são passados por valor por padrão, mas você pode passar um ponteiro para uma função se desejar que a função modifique o valor original.
	// Sintaxe: func nomeDaFuncao(nomeDoPonteiro *tipoDoPonteiro) {}

	// Função que recebe um ponteiro
	// func trocaValores(ptr *int) {
	// 	fmt.Println("Valor original:", valor)
	// 	*ptr = 20
	// }

	// Função que recebe um ponteiro
	fmt.Println("Passagem de Ponteiros para funções:")
	valor := 10
	fmt.Println("Valor original:", valor)

	trocaValores(&valor)
	fmt.Println("Valor modificado:", valor)

	fmt.Println()

}