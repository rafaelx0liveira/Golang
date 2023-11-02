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

package main

import (
    "fmt"
    "sync"
)

func imprimirMensagem(prefixo string, wg *sync.WaitGroup) {
    defer wg.Done() // Sinaliza que a goroutine terminou ao final da função
    for i := 1; i <= 3; i++ {
        fmt.Printf("%s: Mensagem %d\n", prefixo, i)
    }
}

func main() {
    var wg sync.WaitGroup // Usada para aguardar todas as goroutines

    // Crie duas goroutines para imprimir mensagens
    wg.Add(2) // Informa que estamos esperando que duas goroutines terminem

		fmt.Println("Goroutines:")

    go imprimirMensagem("Goroutine 1", &wg)
    go imprimirMensagem("Goroutine 2", &wg)

    // Aguarde até que todas as goroutines terminem
    wg.Wait()

    fmt.Println("Programa principal: Todas as goroutines terminaram.")

		fmt.Println()
}
