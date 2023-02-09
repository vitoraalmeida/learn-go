package main

import "fmt"

func main() {
	fruits1 := make([]string, 3, 4)
	fruits1[0] = "Apple"  //
	fruits1[1] = "Orange" //
	fruits1[2] = "Banana" //

	fmt.Println("\nFruits 1")
	inspectSlice(fruits1)

	//cria outro slice apontando para o elemento de indice 2 até (não incluso)
	//o indice 4 do antigo slice. Ou seja os elem nos indices 2 e 3.
	fruits2 := fruits1[2:4]

	//como o elemento no indice 3 era o ultimo, que estava para crescimento
	//futuro, ele é inicializado com string vazia, pois o fruits2 acessa ele.
	fmt.Println("\nFruits 2 - mesmos endereços de fruits1 1")
	inspectSlice(fruits2)

	fruits2[1] = "Blue Berry"
	fmt.Println("\nFruits 2 - string no ultimo indice modificada")
	inspectSlice(fruits2)
	//alterar a ultima string de fruits2 significa alterar o backing array
	//usado por ambos, mas o fruits1 não vê a mudança, pois seu tamanho não foi
	//modificado
	fmt.Println("\nFruits 1 - não acessa o ultimo elemento, pois len continua 3")
	inspectSlice(fruits1)

	//fazer append em fruits1 vai aumentar seu len de 3 para 4, agora ele tem
	//acesso ao ultimo elemento do array, e sobrescreve o valor anterior
	//adicionado por fruits2.
	fruits1 = append(fruits1, "Pineapple")
	fmt.Println("\nFruits 1 - Append em fruits 1 - aumenta len e acessa o ultimo")
	fmt.Println("modificando também fuits2")
	inspectSlice(fruits1)

	fmt.Println("\nFruits 2 - ultimo elemento alterado pelo append de fruits1")
	inspectSlice(fruits2)
	//append no slice de slice pode fazer com que o valor no slice original
	//seja sobrescrito.
	/*						   +->um novo append em fruits1 age aqui
	                           |  modificando também fruits2
	                           |
		 fruits1<----+-------+ |           +----------+--->fruits1
			         |    +--|--+          |    +----+|
			         |[][]|[]|[]|  ==>     |[][]|[][]||
		             |    +--|--+          |    +----+|
		             +-------+  |          +----|-----+
	                            |               |
	                 fruits2<---+     fruits2<--+
	*/

	//para evitar esse comportamento, pode-se criar o slice determinando também
	//a capacidade, de forma que se houver outro append, obrigatoriamente será
	//criado outro backing array para ele
	fruits3 := fruits2[0:2:2] //do 0 ao 2º (não incluso), cap = len = 2

	/*   ANTES DO APPEND          | DEPOIS DE APPEND EM FRUITS3
	                              |
	                              |
		      +-----fruits3       | fruits3<---+          +----------+--->fruits1
		      |                   |            |          |    +----+|
		 +----+                   |            +------+   |[][]|[][]||
		+|----|------+            |            |[][][]|   |    +----+|
		||    |+----+--->fruits2  |            +------+   +----|-----+
		||[][]||[][]||            |                            |
		||    |+----+|            |                  fruits2<--+
		+|----|------+            |
		 +----+      |___> fruits1|
	*/

	fmt.Println("\nFruits3 - mesmos endereços de fruits1 e fruits2")
	inspectSlice(fruits3)

	fruits3 = append(fruits3, "AMORA")
	fmt.Println("\nFruits3 - append em 3 - novo array - endereços != fruits1 e 2")
	inspectSlice(fruits3)

	fruits1 = append(fruits1, "ABACATE")
	fmt.Println("\nFruits1 - append em 1 - novo array - endereços != fruits3 e 3")
	inspectSlice(fruits1)

	fmt.Println("\nFruits2 - array anterior - endereços != fruits1 e 3")
	inspectSlice(fruits2)
}
