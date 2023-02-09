package main

import (
	"fmt"
	"reflect"
)

// um tipo composto: baseado em outros tipos existentes
type example struct {
	flag    bool    // 1 byte
	pi      float64 // 4 bytes
	counter int16   // 2 bytes
}

//Estruturar o struct de forma que os campos maiores estejam listados
//primeiro ajuda a ter menos padding, pois será acrescentado apenas numa das pontas

/*

Alinhamento
                                                   8 bytes (word)
  b1   b2   b3   b4   b5   b6   b7   b8     |------------------------|------------------------|
+----+----+----+----+----+----+----+----+    +--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+--+
|flag|Padd| counter |          pi       |    |                    | a| b|
+----+----+----+----+--------------------    +-----------------------+-----------------------+

O struct inteiro tem que seguir o alinhamento do maior campo. O maior campo (pi) possui
4 bytes, então o cada bloco que compõe o struct vai ter 4 bytes.
Counter tem 2 bytes, então será necessário adicionar um preenchimento de 1 byte,
já que flag tem apenas 1.
Isso acontece para que não tenhamos os risco de que desperdicemos memória, e também
que não seja necessário realizar a leitura de blocos de memória diferentes para
ler um valor apenas.

*/

func main() {
	// declara uma variável do tipo example definindo com seu falor zero (false,0, 0.0)
	var e1 example
	fmt.Printf("%+v\n", e1)

	// declara e inicializa com valores determinados
	e2 := example{
		flag:    true,
		counter: 10,
		pi:      3.141592,
	}

	fmt.Println("Flag", e2.flag)
	fmt.Println("counter", e2.counter)
	fmt.Println("pi", e2.pi)

	// declara um struct anonimo com valores padrão
	var e3 struct {
		flag    bool
		counter int16
		pi      float32
	}

	fmt.Printf("%+v\n", e3)
	// declara e inicializa struct anonimo com valorse definidos
	e4 := struct { // tamanho 12, pois tem padding depois da flag e depois de counter
		flag    bool    // byte
		pi      float32 // 4 bytes -> alinhamento
		counter int16   // 2 bytes
	}{
		flag:    true,
		counter: 42,
		pi:      3.141592,
	}
	fmt.Printf("%+v \t tamanho: %d\n", e4, reflect.TypeOf(e4).Size())

	var b bill
	var a ana
	// b = a  // não compila pois apesar de serem identicos e compatíveis, são diferentes no nome di tipo.
	b = bill(a) // conversão explicita -> checagem pelo compilador em relação a compatiblidade
	fmt.Printf("%+v\n", b)

	var b_anon struct {
		name string
		age  int
	}

	b = b_anon //compila pois b_anon não tem um tipo nomeado, mas é compativel
	// é assim que as funções em go recebem valores literais nos seus argumentos

}

type bill struct {
	name string
	age  int
}
type ana struct {
	name string
	age  int
}
