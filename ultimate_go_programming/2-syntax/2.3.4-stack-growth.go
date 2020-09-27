package main

//o tamamnho da stack da goroutine é 2kb. Quando a stack fica sem espaço,
//uma nova stack com 25% a mais do tamanho anterior é disponibilizada e
//o conteúdo é copiado. Várias stacks podem ser criadas, por isso a stack do go
//começa com 2kb (a padrão do OS é 1mb)
//Valores numa stack não podem ser compartilhados com outras stacks.

//const size = 10  // 80 bytes * 10 (chamdas) = 800 bytes. Não precisa nova stack
const size = 100 // 800 bytes * 10 = 8000 bytse. Precisa de novas stacks

func main6() {
	s := "HELLO" //       |-------> tamanho total = size * 8 bytes
	stackCopy(&s, 0, [size]int{})
}

func stackCopy(s *string, c int, a [size]int) {
	println(c, s, *s)
	//o valor do endereço da string muda quando uma nova stack é usada

	c++
	if c == 10 { // cria size * 8 bytes a cada chamada, total = 10 * size * 8
		return
	}

	stackCopy(s, c, a)
}
