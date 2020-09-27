package main // strings literais não são criadas na heap. quando aqui se diz
//              sobre criação na heap, é meramente ilustrativo
import "fmt"

//								  0   1   2   3   4
func main2() { //    |-----------+---+---+---+---+---+
	var fruits [5]string //     |nil|nil|nil|nil|nil|
	fruits[0] = "Apple"  //--+  +---+---+---+---+---+
	fruits[1] = "Orange" //--|  | 0 | 0 | 0 | 0 | 0 |
	fruits[2] = "Banana" //--|  +---+---+---+---+---+
	fruits[3] = "Grape"  //--|  |___|
	fruits[4] = "Plum"   //--|    |____ string = ponteiro para array + tamanho
	/*   +-------------------+
	     |           +------------------------------------------+
	+----+       +---|------------------------------+           |
	|        +---|---|------------------+           |           |
	|    +---|---|---|--------|         |           |           |
	|    |   |   |   |       |A|p|p|l|e|O|r|a|n|g|e|B|a|n|a|n|a|G|r|a|p|e|P|l|u|m|
	|    |   |   |   |        0 1 2 3 4 0 1 2 3 4 5 0 1 2 3 4 5 0 1 2 3 4 0 1 2 3
	|  +-|-+-|-+-|-+-|-+---+	Arrays na heap(ñ proximos necessariamente)|							              |
	|  | * | * | * | * | * -----------------------------------------------+
	+--+---+---+---+---+---+
	   | 5 | 6 | 6 | 5 | 4 |      Para se criar uma string, uma array de bytes
	   +---+---+---+---+---+      (caracteres) é criado na heap e, na stack, a
	Após a criação, a string      a string vai ser criada com um ponteiro para
	é copiada para o array de     o array criado e o tamanho.
	de strings

	Então para inicializar cada elemento do array de strings, uma é criada
	e copiada para o array de strings, totalizando 2 copias do valor (string)

	*/
	//       |---> fruit é uma cópia da string na posição i. 3ªcópia
	for i, fruit := range fruits { //Semantica de valor.
		fmt.Println(i, fruit, &fruit, "mesmo endereço")
		// o endereço de fruit será o mesmo sempre, pois o valor da string é
		// copiada p/ a mesma variável. Não pode haver alteração no orignal
		if i == 1 {
			fruit = "Laranja" // não surte efeito fora desse contexto
		}
	}

	//Até aqui, 4 valores na stack e um na heap para cada string. Uma alocação
	//somente para cada string. Mais eficiente

	println("----------------\n")
	println("fruits:", &fruits, " -> mesmo endereço do primeiro elemento!")
	for i := range fruits { //semantica de ponteiro
		fmt.Println(i, fruits[i], &fruits[i]) //vai mostrar endereços diferentes
		// aqui podem ocorrer alterações da string original
		if i == 1 {
			fruits[i] = "Laranja" // não surte efeito fora desse contexto
		}
	}

	println("----------------\n")

	for i := range fruits {
		if i == 1 {
			fmt.Println(i, fruits[i], " (valor alterado) ", &fruits[i])
			continue
		}
		fmt.Println(i, fruits[i], &fruits[i])
	}

	println("----------------\n")

	numbers := [4]int{10, 20, 30, 40}

	for i := 0; i < len(numbers); i++ {
		fmt.Println(i, numbers[i], &numbers[i])
	}
}
