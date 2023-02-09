package main

import (
	"fmt"
	"sync"
)

var counter int

//mutex define uma seção critica no codigo, uma seção que só pode ser acessada
//por uma goroutine por vez
var mutex sync.Mutex

//não usar variaveis globais.

func main2() {
	const grs = 2

	var wg sync.WaitGroup
	wg.Add(grs)

	for i := 0; i < grs; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				//lock e unlock devem estar no mesmo escopo
				//deve-se usar mutex com seções que podem ser operações atomicas
				//não usar em trechos muito longos.
				//não é necessário usar chaves, só é mais claro
				mutex.Lock()
				{
					value := counter

					value++
					//fmt.Println geraria uma syscall que provavelmente causaria
					//troca de contexto, mas o mutex não permite
					//NÃO USAR Prints DENTRO DE MUTEX, USAR O MINIMO POSSIVEL
					fmt.Println(value)

					counter = value
				}
				mutex.Unlock()
			}

			wg.Done()
		}()
	}

	wg.Wait()

	fmt.Println("\nFinal counter:", counter)
}
