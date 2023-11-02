// **************************************
// Channels: Comunicação entre goroutines
// **************************************
// Channels são canais de comunicação entre goroutines, que são threads leves
// Channels são usados para enviar e receber dados entre goroutines
// Channels são tipos de dados compostos
// Channels são criados com a função make
// Sintaxe: var nomeDaVariavel chan tipoDaVariavel

package main

import "fmt"

// Função para enviar dados para o canal
func enviarDados(c chan int) {
    for i := 1; i <= 5; i++ {
        c <- i // Envia valores para o canal
    }
    close(c) // Fecha o canal após a conclusão do envio
}

// Função para receber dados do canal
func receberDados(c chan int, done chan bool) {
    for {
        valor, aberto := <-c // Recebe valores do canal
        if !aberto {
            break // O canal foi fechado, saia do loop
        }
        fmt.Println("Valor recebido:", valor)
    }
    done <- true // Sinaliza que a goroutine terminou
}

func main() {
    // Cria um canal para a comunicação entre goroutines
    canal := make(chan int)

    // Canal de sinalização para indicar que a goroutine de recebimento terminou
    sinalizador := make(chan bool)

    // Inicia uma goroutine para enviar dados para o canal
    go enviarDados(canal)

    // Inicia uma goroutine para receber dados do canal
    go receberDados(canal, sinalizador)

    // Aguarde até que a goroutine de recebimento termine
    <-sinalizador

    fmt.Println("Programa principal: Todas as goroutines terminaram.")
}
