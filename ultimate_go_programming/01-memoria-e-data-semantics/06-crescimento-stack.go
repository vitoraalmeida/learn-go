package main

import (
	"fmt"
	"time"
)

/*
   a cada chamada de função encadeada (sem que haja um retorno antes
   de chamar a próxima) ocorre o crescimento da stack.
   O tamanho inicial da stack de uma goroutine é 2KB, mas se o crescimento
   da stack chegar num ponto de ultrapassar esse valor, então outra goroutine
   e criada com o dobro de tamanho e os valores são copiados
*/

// tamanho do slice de inteiros que está sendo copiado em cada
// stack frame de chamadas de função
//const size = 10
const size = 9000

// valores baixos não vão causar um uso grande da stack, então outra não precisará
// ser criada. Aumentar o tamanho fará com que uma quantidade grande de memória
// tenha que ser alocada em cada chamada de função, causando crescimento da stack
// e necessidade de criar outra goroutine, podendo perceber que o endereço da
// string compartilhada vai mudar, indicando a criação de outra stack;

func main() {
	s := "HELLO"
	stackCopy(&s, 0, [size]int{})
}

func stackCopy(s *string, c int, a [size]int) {
	fmt.Println(c, s, *s)

	c++
	if c == 10 {
		return
	}
	time.Sleep(1 * time.Second)

	stackCopy(s, c, a)
}
