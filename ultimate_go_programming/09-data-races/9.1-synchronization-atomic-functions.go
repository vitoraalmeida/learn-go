package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//var counter int

//usando atomic, o tipo do dado que será compartilhado tem que ser definido
//especificamente para ter garantia de ser consistente com a plataforma.
var counterAtomic int64

func main1() {
	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				//assim, as duas goroutines agem no mesmo endereço, e é o hard
				//ware quem garante a sincronização
				atomic.AddInt64(&counterAtomic, 1)

				/*
					value := counter

					//value++ é uma instrucao de leitura, modificção e escrita
					//em assembly, então mesmo nessa linha de código podem ocorrer
					//trocas de contexto.
					value++

					//fmt.Println faz uma system call, o que deixa a goroutine
					//em estado de espera e ocorre uma troca de contexto.
					//isso pode causar resultados diferentes do esperado, pois
					//antes de atualizar o valor, a outra goroutine lê o valor
					//da variável global no valor antigo.
					//fmt.Println(value)

					counter = value
				*/
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("\nFinal counter:", counter)
}
