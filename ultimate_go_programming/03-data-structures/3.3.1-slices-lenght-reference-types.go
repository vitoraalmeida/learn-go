package main

import "fmt"

/*
 Slices, maps, channels, interfaces e funções são reference types.
 Reference types guardam um ponteiro e tem nil como zero value.
 Strings, apesar de guardar um ponteiro, tem zero value vazio, não nil, então
 não é reference type.
*/
func main() {
	/*
		slices são estruturas de dados de 3 words (24 bytes - arch 64bits):
		ponteiro para o array que baseia o slice, tamanho usado e capacidade.
		Podem crescer quando necessário.
		O slice é uma view, uma janela, da acesso à parte do backing array.
		Um slice menor de um slice pre-exsistente apontara p/ o mesmo array

		usar make (indicado quando souber o tamanho que será usado) vai criar
		um slice que aponta para um array com o zero value do tipo usado.

		var fruits []string -> criaria um slice com zero value, ou seja,
							   nil pointer, 0 len e 0 capacity
		fruits := []string{}  -> criaria um slice vazio, ou seja,
							   um ponteiro para empty struct, mas 0 len e 0 cap
		empty struct é um tipo de alocação zero.
	*/

	//make([]string, 4) -> criaria slice de tamanho 4 e capacidade 4
	fruits := make([]string, 3, 4)
	/*
		fruits  -> não é um slice vazio, é um slice que aponta para um array
		+---+     +---+---+---+---+  vazio
		| * ------>nil|nil|nil|nil|
		+---+	  +---+---+---+---+
		| 3 |	  | 0 | 0 | 0 | 0 |
		+---+	  +---+---+---+---+
		| 4 |
		+---+   O slice pode ser sempre copiado e a unica coisa que
		        precisaria ser copiado/estar na heap seria o backing
				array (não é o caso agora)
	*/
	fruits[0] = "Apple"  //
	fruits[1] = "Orange" //
	fruits[2] = "Banana" //
	/*				        +----------------------------------+
				        +---|----------------------+           |
	                +---|---|------------|         |           |
	                |   |   |           |A|p|p|l|e|O|r|a|n|g|e|B|a|n|a|n|a|
	   fruits       |   |   |            0 1 2 3 4 0 1 2 3 4 5 0 1 2 3 4 5
	   +---+      +-|-+-|-+-|-+---+     Não são, necessariamente, contíguos
	   | * -------> * | * | * |nil|
	   +---+	  +---+---+---+---+
	   | 3 |	  | 5 | 6 | 6 | 0 |
	   +---+	  +---+---+---+---+
	   | 4 |
	   +---+
	*/

	inspectSlice(fruits)
}

func inspectSlice(slice []string) {
	fmt.Printf("Length[%d] Capacity[%d]\n", len(slice), cap(slice))
	for i, s := range slice { //semantica de valor
		fmt.Printf("[%d] %p %s\n", i, &slice[i], s)
	}
}
