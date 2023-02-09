package main

import (
	"fmt"
	"sync"
)

//Apesar de acessarem o map atraves de keys diferentes, não significa que
//está insento de causar problemas de integridade. O runtime do go consegue
//detectar data races no acesso à maps quando ocorrem
var scores = make(map[string]int)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		for i := 0; i < 1000; i++ {
			scores["A"]++
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 1000; i++ {
			scores["B"]++
		}
		wg.Done()
	}()

	wg.Wait()

	fmt.Println("Final scores", scores)
}
