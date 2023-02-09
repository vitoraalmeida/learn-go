package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

//o dado que tera o acesso sincronizado
var data []string

//multiplas leituras e uma escrita. O RWMutex é mais lento que o normal
var rwMutex sync.RWMutex

//conta o numero de leituras
var readCount int64

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for i := 0; i < 5; i++ {
			time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			writer(i)
		}
		wg.Done()
	}()

	for i := 0; i < 8; i++ {
		go func(id int) {
			for {
				reader(id)
			}
		}(i)
	}
	wg.Wait()
	fmt.Println("Program Complete")

}

func writer(i int) {
	//Apenas permite que uma goroutine leia/escreva no slice num mesmo momento
	//Solicita o lock, se tiver alguma goroutine fazendo leitura, espera terminar
	//e tranca. Quando termina a escrita, libera a leitura
	rwMutex.Lock()
	{
		//Load atomico permite garantir que não haverá outras leituras, mas
		//o rwmutex já garante isso, então pode ler normalmente
		//rc := atomic.LoadInt64(&readCount)
		rc := readCount

		fmt.Printf("****=> : Performing Write: RCount[%d]\n", rc)
		data = append(data, fmt.Sprintf("String: %d", i))
	}
	rwMutex.Unlock()
}

func reader(id int) {
	// qualquer goroutine pode ler se não estiver acontecendo escrita
	rwMutex.RLock()
	{
		//Adiciona de forma atômica indicando que tem uma goroutine lendo
		rc := atomic.AddInt64(&readCount, 1)

		time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)
		fmt.Printf("%d : Performing Read : Length[%d] RCount[%d]\n", id, len(data), rc)
		//terminou de ler
		atomic.AddInt64(&readCount, -1)
	}
	rwMutex.RUnlock()
}
