package main

import "fmt"

/*
Tipos:
	Sim tipos não há como ter, realmente, integridade. Não tem como entender
	realmente o custo das decisiões que estão sendo tomadas no código.

	Byte: unidade  de memória básica

	00001010 - qual o valor? não da pra saber a não ser que tenhamos a informação
			   do tipo que esse padrão de bits representa.
	* Tipos fornecem 2 tipos de informação: tamanho e representação

*/

//APAGAR "1" PARA RODAR
func main1() { //	     word	          word		   uint8 uint16 uint32 uint64
	var a int     //arch 32bits =  int32/ 64bits =  int64 (int8 int16 int32 int64)
	var b string  // uma estrutura de 2 words. 1º um ponteiro. 2º o nº de bytes
	var c float64 // representa decimal de 8 bytes
	var d bool    // 1 byte (1 bit usado)
	// declarar com var inicializa com o 'zero value' do tipo.
	// bool = false
	// int = 0        +-> nº bytes
	// float = 0.0    |
	// string = [nil][0]
	//            |
	//			  +-> ponteiro
	fmt.Printf("var a int \t %T [%v]\n", a, a)
	fmt.Printf("var b string \t %T [%v]\n", b, b)
	fmt.Printf("var c float64 \t %T [%v]\n", c, c)
	fmt.Printf("var d bool \t %T [%v]\n", d, d)
	//                +----------------------| bb
	aa := 10      //  |                    [ 0xa9c8 ]
	bb := "hello" //[ h | e | l | l | o]   [    5   ]
	cc := 3.14159 //  |___0xa9c8
	dd := true
	// declarar com var para zero values e := para existentes = consistencia
	fmt.Printf("var aa :=  10 \t %T [%v]\n", aa, aa)
	fmt.Printf("var bb := \"hello\" \t %T [%v]\n", bb, bb)
	fmt.Printf("var cc := 3.14159 \t %T [%v]\n", cc, cc)
	fmt.Printf("var dd := true \t %T [%v]\n", dd, dd)

	// Go não tem casting, tem conversão. Casting pode corromper memória.
	// Go procura integridade

}
