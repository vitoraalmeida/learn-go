package main

import "fmt"

/*
	Usar semantica de valor:
		- Quando usando tipos built-in, incluindo campos struct
			* Exceção: struct usado para interagir com DB - marshal/un-marshal

		- Quando usando referece types. Já são ponteiros, pra que fazer cópia
		  do endereço de um endereço?
			* Exceção: compartilhando map e slices para baixo da call stack e
					   para funções de decode e un-marshal.

		- Ainda que o método seja de mutação, usar a semantica de valor e retornar
		  o novo. Permite alterar de forma segura, sandbox da função, evita side
		  effects

	Usar semantica de ponteiro:
		- Quando trabalhando com structs, deve-se escolher baseado na necessidade
		  Se não tiver certeza do que usar, usar semantica de ponteiro

	Escolha a semantica e seja consistente, se estiver ruim/errado, refatore.
*/

/*
semantica de ponteiro, pois "não faz sentido" fazer copias de pessoas.

tipo definido no primeiro arquivo:
type user struct {
	name  string
	email string
}
Method:
changeEmail - pointer semantic
*/

func main2() {
	users := []user{
		{"ed", "ed@email.com"},
		{"erick", "erick@email.com"},
	}

	//semantica de valor no for range para um método que tem semantica de
	//ponteiro. Não altera.
	for _, u := range users {
		u.changeEmail("it@wontmatter.com")
	}

	//não surte efeito
	for _, u := range users {
		fmt.Println(u)
	}

	//semantica de pointeiro no for range. Agora altera
	for i := range users {
		users[i].changeEmail("it@wontmatter.com")
	}

	for _, u := range users {
		fmt.Println(u)
	}
}
