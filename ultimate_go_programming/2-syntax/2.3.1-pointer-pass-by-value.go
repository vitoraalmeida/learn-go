package main

import "fmt"

//Cada vez que uma função é chamada, um frame de memória é fornecido para conter
//suas variáveis, seus dados. Cada função só tem acesso à seu frame. Não é per-
//mitido atravessar esse limite. Quando a função retorna, seu stack frame é
//invalidado e disponibilizado para outra chamada de função.
func main3() {
	count := 10

	fmt.Println("count:\tValue Of:[", count, "]\tAddr Of[", &count, "]")

	//uma nova stack frame é criada
	increment1(count) //passagem por valor. Desvantagem: cópia da dados
	//a stack frame é desmpilhada após o fim da função

	fmt.Println("count:\tValue Of:[", count, "]\tAddr Of[", &count, "]")
}

func increment1(value int) { //value local à função|recebe cópia do valor passado
	value++ //não causa efeitos na variavel em main, pois é cópia do valor dela
	fmt.Println("count:\tValue Of:[", value, "]\tAddr Of[", &value, "]")
}
