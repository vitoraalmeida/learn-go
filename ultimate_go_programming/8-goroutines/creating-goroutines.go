package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	//numero de processadores logicos que serão usados
	runtime.GOMAXPROCS(2)
	//se 2 forem usados, cada gorountine sera executada em processadores
	//diferentes, fazendo com que os outputs saiam misturados
}

const oneSec = time.Millisecond * 1000

func main() {
	//waitgroup é um semaforo de contagem de referencia. Gerencia concorrencia
	var wg sync.WaitGroup
	//adiciona dois contadores
	wg.Add(2)
	//se adicionar menos que o numero de goroutines iniciadas, a goroutine main
	//pode finalizar sem esperar todas terminarem

	fmt.Println("Start Goroutines")

	go func() { //inicia uma goroutine - o wg vai gerenciar/monitorar o estado
		uppercase()
		wg.Done() //avisa que o wg pode parar de gerenciar essa goroutine
		//se o wg estiver gerenciando duas goroutines wg.Add(2) e uma delas
		//não der Done, ocorre deadlock, a goroutine main não poderá continuar
		//já que wg.Wait foi usado.
		//o runtime do go tem um detector simples de deadlock, vai dar runtime
		//error. Não funciona em toda situação
	}()

	go func() {
		lowercase()
		wg.Done()
	}()

	fmt.Println("Waiting to finish")

	wg.Wait() //faz a goroutine main esperar para finalizar
	fmt.Println("\nTerminating program")
}

func lowercase() {
	for count := 0; count < 3; count++ {
		for r := 'a'; r <= 'z'; r++ {
			fmt.Printf("%c ", r)
			//a espera ajuda a vizualizar troca de contexto entre goroutine
			//time.Sleep(oneSec)
		}
	}
}

func uppercase() {
	for count := 0; count < 3; count++ {
		for r := 'A'; r <= 'Z'; r++ {
			fmt.Printf("%c ", r)
			//time.Sleep(oneSec)
		}
	}
}
