package main

import "fmt"

/* Alinhamento:
   A memória é divida em pedaços iguais para o uso. O montante disponível é
   dividido em pedaços de tamanho definido e valores só podem ser colocados
   de forma que caibam nesses pedaços respeitando os limites. Um valor não pode
   estar em dois pedaços ao mesmo tempo. Para ler cada pedaço, é preciso uma
   operação, se os valores fossem colocados em dois ao mesmo tempo, seriam duas
   operações para cada leitura e escrita.
*/

//o tamanho de cada campo define o local no alinhamento e só pode ser colocado
type example struct { // em endereços multiplos de 2
	flag    bool    // 1 byte ---+
	counter int16   // 2 bytes---|------+
	pi      float32 // 4 bytes---|------|-------|
	//							|-|    |--|  |--------|
} //                            [0][1][2][3][4][5][6][7]
//                               |______ Inutilizado

//é recomendável ordenar os campos dos structs do maior tamnho para o menor
//para fins de economia de espaço (SE tiver impacto na performace, se não,
//deve-se priorizar a legibilidade, a semantica do código.)

type example2 struct {
	flag    bool  // 1 byte---+
	counter int64 // 2 bytes--|------------------------+
} //                         |-|   não cabem 8bytes    |--------------------|
//							 [0][1][2][3][4][5][6][7] [0][1][2][3][4][5][6][7]
//                                    7 bytes desperdiçados

type bill struct { //----
	counter int64   //   |
	pi      float32 //   |
	flag    bool    //   |
} //                     |
//                       |
type alice struct { //----+--- Tipos diferentes ainda que idênticos e comparáveis
	counter int64
	pi      float32
	flag    bool
}

//tirar o "2" para funcionar
func main2() {
	var e1 example //inicia e1 com zero value em cada campo
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	var e3 struct { // struct anônimo
		counter int64
		pi      float32
		flag    bool
	}

	e4 := struct { // struct anônimo
		counter int64
		pi      float32
		flag    bool
	}{ // contrucao literal
		counter: 10,
		pi:      3.141592,
		flag:    true,
	}

	fmt.Println(e3, e4)

	var b bill
	var a alice
	//b = a //erro
	b = bill(a) //conversão
	fmt.Println(a, b)
	fmt.Printf("%+v\n", e1)
	fmt.Printf("%+v\n", e2)
}
