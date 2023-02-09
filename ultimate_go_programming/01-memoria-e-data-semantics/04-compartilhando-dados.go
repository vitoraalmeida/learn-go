package main

func main() {
	// inicio do pedaço(frame) da stack da goroutine
	count := 10
	println("Count:\tValue Of [", count, "]\tAddr Of[", &count, "]")
	// chama a função increment e goroutine agora tem um outro "pedaço"
	// relativo à função increment
	increment(&count) // adiciona um acesso indireto à memória do stack frame de main
	//semantica de ponteiro

	println("Count:\tValue Of [", count, "]\tAddr Of[", &count, "]")
}

func increment(c *int) {
	// agora recebemos um endereço de memória e o valor passado será copiado
	// para a variável interna 'c' do stack frame de increment
	*c++ // cruza os limites de frames, ocasionando em side-effects
	println("Inside increment: Count:\tValue Of [", c, "]\tAddr Of[", &c, "]")
}
