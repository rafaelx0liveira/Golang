// **********
// JSON em Go
// **********
// JSON (JavaScript Object Notation) é um formato de troca de dados entre sistemas e programas muito utilizado atualmente. É um formato de texto simples, baseado em um subconjunto da linguagem de programação JavaScript, porém independente de linguagem. JSON é um formato de dados muito utilizado em APIs RESTful.
// Em Go, a manipulação de dados JSON é muito fácil devido à presença da biblioteca padrão "encoding/json". 
// Essa biblioteca oferece suporte para codificação (marshal) e decodificação (unmarshal) de dados JSON, permitindo que você converta dados em Go em formato JSON e vice-versa. 
// O pacote "encoding/json" é amplamente utilizado em Go para interagir com serviços da web, armazenar configurações ou dados estruturados em JSON e muito mais.

package main

import (
	"encoding/json"
	"fmt"
)

// Definindo uma estrutura
type Pessoa struct {
	Nome  string 
	Idade int
}

func main() {
	
	fmt.Println("JSON em Go")
	fmt.Println()

	// Criando uma instância da estrutura Pessoa
	pessoa1 := Pessoa{"Gopher", 10}

	fmt.Println("Instância da estrutura Pessoa criada:")
	fmt.Println(pessoa1)

	fmt.Println()

	// Convertendo a instância da estrutura Pessoa para JSON
	p1, err := json.Marshal(pessoa1)
	if err != nil {
		fmt.Println("Erro ao transformar a struct em JSON:", err)
		return
	}

	fmt.Println("Instância da estrutura Pessoa convertida para JSON:")
	// Exibindo o JSON
	fmt.Println(p1)

	fmt.Println()

	fmt.Println("O JSON é um array de bytes, por isso é exibido em números.")

	fmt.Println()

	fmt.Println("Para exibir o JSON como string, basta converter o array de bytes para string.")

	fmt.Println()

	// Convertendo o array de bytes para string
	fmt.Println("Convertendo o array de bytes para string:")
	fmt.Println(string(p1))

	fmt.Println()

	// **************
	// Tags de campos
	// **************
	// As tags de campos são usadas para definir metadados para campos de uma estrutura, como o nome do campo, se o campo é exportado ou não, etc.
	// As tags de campos são usadas principalmente para codificar e decodificar dados JSON em Go.
	// As tags de campos são definidas dentro de backticks (Crase) (`) imediatamente após o nome do campo.
	// As tags de campos são definidas no formato "key:value".

	// Definindo uma estrutura com tags de campos
	type Pessoa2 struct {
		Nome  string `json:"nome"`
		Idade int    `json:"idade"`
	}

	pessoa2 := Pessoa2{"Gopher", 10}

	fmt.Println("Instância da estrutura Pessoa criada para exemplo com Tags:")
	fmt.Println(pessoa2)

	// Convertendo a instância da estrutura Pessoa para JSON
	p2, err := json.Marshal(pessoa2)
	if err != nil {
		fmt.Println("Erro ao transformar a struct em JSON:", err)
		return
	}

	fmt.Println()

	fmt.Println("Instância da estrutura Pessoa convertida para JSON:")

	// Exibindo o JSON
	fmt.Println(string(p2))

	fmt.Println()

	// *********************************
	// Hidratando uma estrutura com JSON
	// *********************************
	// Hidratar uma estrutura significa converter um JSON em um struct.
	// Para hidratar uma estrutura com JSON, usamos a função "json.Unmarshal()".
	// A função "json.Unmarshal()" recebe dois parâmetros: o JSON e um ponteiro para a estrutura que será hidratada.
	// É necessário passar o endereço da estrutura para a função "json.Unmarshal()" para que a estrutura seja hidratada no endereço de memória correto, e não somente do escopo da função.
	// A função "json.Unmarshal()" retorna um erro caso ocorra algum problema durante a hidratação da estrutura.

	json1 := `{"nome":"Rafael","idade":23}`
	fmt.Println("JSON para hidratar a estrutura:")
	fmt.Println(json1)

	fmt.Println()

	// Hidratando a estrutura com JSON
	json.Unmarshal([]byte(json1), &pessoa2)

	fmt.Println("Estrutura hidratada com JSON:")
	fmt.Println(pessoa2)

	fmt.Println()
}