package main

import (
	"fmt"
	"os"
	"sync"
)

//interfaces são tipos de 2 words. quando são acessadas,

type Speaker interface {
	Speak() bool
}

type Ben struct {
	name string
}

func (b *Ben) Speak() bool {
	if b.name != "Ben" {
		fmt.Printf("Ben says, \"Hello my name is %s\"\n", b.name)
		return false
	}

	return true
}

type Jerry struct {
	name string
}

func (r *Jerry) Speak() bool {
	if r.name != "Jerry" {
		fmt.Printf("Jerry says, \"Hello my name is %s\"\n", r.name)
		return false
	}

	return true
}

func main() {
	ben := Ben{"Ben"}
	jerry := Jerry{"Jerry"}

	//interface apontando para ben
	person := Speaker(&ben)

	//
	// 	Não ha garantia de qual goroutine será executada primeiro.
	// 	Supondo que a gorutine jerry seja primeiro, ela faz a atribuição
	//do ponteiro de jerry para person. Mas como person é uma interface e
	//interfaces são estruturas de 2 palavras, pode acontecer de só atribuir
	//a primeira palavra que aponta para a implementação do método em jerry,
	//mas antes atribuir a segunda que apronta para o dado concreto de jerry,
	//ocorre uma mudança de contexto e a goroutine de ben é executada.
	//	Antes que person faça a chamada do metodo a partir de Ben, ocorre mudança
	//de contexto e volta para jerry, que vai atribuir a outra parte de person ao
	//valor de jerry, enquando a primeira metade, que aponta para a implementação da
	//função, está apontando para Ben. Ou seja, person vai chamar a função imple
	//mentada em Ben, mas a partir do valor de jerry, usando o name jerry.
	//a função de ben compara o nome jerry com a string ben e causa a diferença

	//go build -race -> vai mostrar a data race
	go func() {
		for {
			person = &ben
			if !person.Speak() {
				os.Exit(1)
			}
		}
	}()

	go func() {
		for {
			person = &jerry
			if !person.Speak() {
				os.Exit(1)
			}
		}
	}()

	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
