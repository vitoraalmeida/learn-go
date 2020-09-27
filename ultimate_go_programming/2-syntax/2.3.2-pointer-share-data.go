package main

import "fmt"

//Cada vez que uma função é chamada, um frame de memória é fornecido para conter
//suas variáveis, seus dados. Cada função só tem acesso à seu frame. Não é per-
//mitido atravessar esse limite. Quando a função retorna, seu stack frame é
//invalidado e disponibilizado para outra chamada de função.
func main4() {
	count := 10

	fmt.Println("count:\tValue Of:[", count, "]\tAddr Of[", &count, "]")

	//uma nova stack frame é criada
	increment(&count) //passagem por referência
	//a stack frame é desmpilhada após o fim da função e o frame é invalidado

	//é preciso saber que quando ocorre essa indireção, ocorrem efeitos
	//colaterais que podem não ser compativeis com o estado anterior. Data races

	fmt.Println("count:\tValue Of:[", count, "]\tAddr Of[", &count, "]")
}

func increment(value *int) { //value recebe cópia do endereço de count
	//value é um ponteiro para um inteiro, local à increment
	*value++ //mesmo sendo a cópia local do endereço, o operador * vai até o endereço
	//dado e modifica o valor naquele edereço dentro de main. Acesso indereto.
	fmt.Println("count:\tValue Of:[", value, "]\tAddr Of[", &value, "]")
}
