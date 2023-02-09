package main

// um processador -> 6 cores (processador lógico) -> 2 threads de hardware por core
// o SO acopla (schedule) uma thread de SO em cada thread de hardware
// quando um programa go executa, uma goroutine é criada para a função main,
// e ao invés de usar uma thread de SO, a goroutine é quem é executada. Como
// se fosse uma Thread, mas bem menor no tamanho da stack inicial(2k)
// minimiza a quantidade de memória que é necessária
func main() {
	// inicio do pedaço(frame) da stack da goroutine
	count := 10
	println("Count:\tValue Of [", count, "]\tAddr Of[", &count, "]")
	// chama a função increment e goroutine agora tem um outro "pedaço"
	// relativo à função increment
	increment(count) // semantica de valor - minimiza efeitos colaterais

	//permite isolamento de cada frame para mutação dos dados, impede
	//mutações indesejadas, porém valores muito grandes seriam sempre
	//copiados, duplicando, podendo ocasionar em uso excessivo de recurso.
	// Além disso, ás vezes é útil alterar os valores
	println("Count:\tValue Of [", count, "]\tAddr Of[", &count, "]")
}

func increment(c int) {
	// increment espera um inteiro, então a variável c vai receber uma cópia
	// do valor passado como argumento ao chamar a função
	// por isso é um novo endereço de memória, então qualquer alteração
	// fica restrita ao stack frame da função e não altera
	// a memória original
	c++
	println("Inside increment: Count:\tValue Of [", c, "]\tAddr Of[", &c, "]")
}
