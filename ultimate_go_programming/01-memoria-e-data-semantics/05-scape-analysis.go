package main

import "fmt"

type user struct {
	user  string
	email string
}

func createUserV1() user {
	u := user{
		user:  "vitor",
		email: "vitor@email.com",
	}

	println("Address inside createUserV1", &u)

	return u
}

func createUserV2() *user {
	// usar semantica de valor para construir e ponteiro só quando for passar
	// pois o operador de referencia & quando passarmos o valor, seja no retorno
	// seja passando em funções, deixa claro que estamos passando uma referência
	// e fica explicito
	u := user{
		user:  "vitor",
		email: "vitor@email.com",
	}

	// compartilhando o valor com a stack frame da função println
	/*
		não tem problema, pois quando finalizar println o stack frame de
		será limpo da memória, mas o valor que gerou a referência (u)
		continuará disponível, pois ainda está na stack frame de createUserV2

			|    main    |                   |    main    |
			+------------+                   +------------+
			|createUserV2|                   |createUserV2|
			|           u|<-+                |   u        |
			+------------+  |                +------------+
			| println  &u|--+

	*/
	println("Address inside createUserV2", &u)

	// compartilhando o valor com a stack frame da função main que é quem
	// recebe &u
	return &u
	/*
			 o problema é que o frame de  createUserV2 vai sumir quando terminar a
			 execução, de forma que main ficaria com um endereço inválido.

			|    main  &u|--+                |    main  &u|--+
			+------------+  |                +------------+  |
			|createUserV2|  |                                |
			|           u|<-+                           ?  <-+
			+------------+

		então o que é feito, na verdade é uma análise de escape e então o valor é alocado na heap

		   |    main  &u|--------+              |    main  &u|--------+
		   +------------+    +---|----+         +------------+    +---|----+
		   |            |    |   \/   |                           |   \/   |
		   |createUserV2|    |  user  | heap					  |  user  | heap
		   |         u  |    |  /\    |							  |        |
		   |         |  |    |   |    |							  |        |
		   |         +-----------+    |              			  |        |
		   +------------+    |        |							  |        |
							 +--------+                           +--------+
	*/
}

func main() {
	user1 := createUserV1()
	fmt.Printf("user: valor[%+v]\taddress = [%p]\n", user1, &user1)

	// createUserV2 compartilha valor para cima da pilha de chamadas de função
	// na goroutine, então o valor vai ser construído na memória heap, então
	// só será desalocado pelo GC
	user2 := createUserV2()
	fmt.Printf("user2: valor[%p]\taddress = [%p]\n", user2, &user2)
}
