package main

import (
	"fmt"
	"os"
)

/*
var name type = expression
	       |      |
		   +------+-----> Podem sem omitidos, mas não ao mesmo tempo
*/

//variáveis no nivel do pacote são inicializadas antes da main.
//variaveis locais são inicializadas quando sua declaração é encontrada

//iniciados com o valor padrão de int
var i, j, k int

//tipo inferido pela expressão
var b, d, s = true, true, true //boolean

//variaveis podem ser inicializadas por funções com multiplos retornos
var f, err = os.Open("notes.txt")

func main() {
	//declara e inicializa ao mesmo tempo (precisa ter pelo menos uma variavel nova)
	//funciona em blocos de função
	t := 0.0
	i, j := 1.0, 0.0

	//declarando apenas u e redefinindo j. Só declara se não foi delcarado antes
	u, j := 3.0, j

	//swap
	i, j = j, i

	//pointers = endereços de variáveis
	p := &t         // p guarda o endereço de t / aponta para t / alias de t
	fmt.Println(*p) //output: 0
	*p = 2          // x = 2
	fmt.Println(*p) //output: 2
	fmt.Println(&t == &t, &t == &u, &t == &*p)
	p = pointer()
	fmt.Println(pointer() == pointer()) //cada chamada entrega um valor diferente

	v := 1
	incr(&v)              //v == 2
	fmt.Println(incr(&v)) // v == 3
}

func pointer() *float64 {
	v := 1.0
	return &v
}

func incr(p *int) int {
	*p++
	return *p
}
