package main

import "fmt"

func main() {
	//declaração "zero value"
	// inicializados com os valores zero de cada tipo
	// ints = 0, floats = 0.0, string = "", bool = false
	var a int8 // para quando precisamos ser específicos por restrições de memória
	var b int16
	var c int32
	var d int64
	var e int  //tamanho da arquitetura 32 ou 64. Mais eficiente para cada arquitetura. O tamanho de um word (o tamanho de um endereço de memória)
	var f uint // int sem sinal

	var g float32
	var h float64

	var i bool

	// strings são tipos de 2 words (ponteiro para o local em memória onde estão os valors,  e outro para armazenar o número de bytes
	// o zero value de uma string é [nil, 0]
	var j string

	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b int \t %T [%v]\n", b, b)
	fmt.Printf("var c int \t %T [%v]\n", c, c)
	fmt.Printf("var d int \t %T [%v]\n", d, d)
	fmt.Printf("var e int \t %T [%v]\n", e, e)
	fmt.Printf("var f int \t %T [%v]\n", f, f)
	fmt.Printf("var g int \t %T [%v]\n", g, g)
	fmt.Printf("var h int \t %T [%v]\n", h, h)
	fmt.Printf("var i int \t %T [%v]\n", i, i)
	fmt.Printf("var j int \t %T [%v]\n", j, j)

	// declara e inicializa
	aa := 10
	bb := "hello"
	cc := 3.14159
	dd := true

	fmt.Printf("var aa int \t %T [%v]\n", aa, aa)
	// não é possível referenciar o endereço específico de um item da string, pois são imutáveis
	fmt.Printf("var bb int \t %T [%v] \t endereço para o primeiro caractere[%v]\n", bb, bb, &bb)
	fmt.Printf("var cc int \t %T [%v]\n", cc, cc)
	fmt.Printf("var dd int \t %T [%v]\n", dd, dd)

	aaa := int32(10) // conversão de um int [int64] em int32
}
