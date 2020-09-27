package main

type user struct {
	name  string
	email string
}

func main5() {
	u1 := createUserV1() //stackframe v1 entra
	//v1 sai
	u2 := createUserV2() //stackframe v2 entra
	//v2 sai
	//                        |------> ponteiro para um objeto na heap
	println("u1", &u1, "u2", u2) // compartilhamento do valor, atravez do ptr
}

//go:noinline
func createUserV1() user { //semantica de valor
	//a sintaxe na construção de user não indica se será alocado na heap
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	//u será criado na stack frame pelo fato de o retorno ser por valor

	println("V1", &u) //compartilha o valor de u com println, não escapa do
	//espaço ativo de memória.
	return u //retorna copia do valor, não escapa do espaço ativo
}

//go:noinline
func createUserV2() *user { //semantica de ponteiro, indica alocação na heap
	//a sintaxe na construção de user não indica se será alocado na heap
	u := user{ //alocação na heap, por ter escape no retorno
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	println("V2", &u)
	return &u // escapa o valor criado no stack frame da função.
	//			 é um problema pelo fato de que esse valor estará inacessível,
	//			 já que a stack frame em que foi criada estará inacessível, após
	//			 o retorno da função. Isso indica que o valor precisa ser alocado
	//			 na heap.
}
