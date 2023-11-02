// ********************************************
// Tipos de dados em Go: primitivos e compostos
// ********************************************

// *******************************
// Tipos de dados primitivos em Go
// *******************************
// Tipos primitivos em Go são tipos básicos que não são compostos de outros tipos.
// São eles:
// - bool
// - string
// - todos os tipos inteiros (int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, uintptr)
// - todos os tipos float (float32, float64)
// - complex64, complex128
// - byte (uint8)
// - rune (int32)
// - error
// Sintaxe: var nomeDaVariavel tipoDaVariavel = valorDaVariavel

// *******************************
// Tipos de dados Compostos em Go:
// *******************************
// Em Go, os tipos de dados compostos são tipos que consistem em uma combinação de valores de outros tipos. 
// O ato de definir, criar, estruturar tipos de dados compostos se chama composição
// Os tipos de dados compostos mais comuns em Go são:
// Arrays, Slices, Maps, Structs, Channels, Slices de Slices, Maps de Maps, etc.

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


package main

import "fmt"

func main() {
	var nome string
	var idade int
	var versao float32
	var ehMaior bool

	nome = "João da Silva"
	idade = 10
	versao = 1.1
	ehMaior = idade >= 18

	fmt.Println("Nome:", nome)
	fmt.Println("Idade:", idade)
	fmt.Println("Versão:", versao)
	fmt.Println("É maior de idade?", ehMaior)
}