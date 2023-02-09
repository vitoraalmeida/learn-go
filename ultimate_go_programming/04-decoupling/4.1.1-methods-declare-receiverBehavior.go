package main

import "fmt"

type user struct {
	name  string
	email string
}

//Métodos permitem que os dados tenham comportamentos. São funções com
//receivers determinados.
//Em Go, diferente de OOP, é bom que nem todos os tipos, dados (objetos)
//tenham comportamento. Separar estado de comportamento.

//                 semantica de valor
//       |-------> indica método que age sobre uma cópia um user.
//    |--+--|      user é um receiver do método
func (u user) notify() { //quando é chamado, Go faz: notify(u)
	fmt.Printf("Sending User Email To %s<%s>\n", u.name, u.email)
}

//                 semantica de ponteiro
//       |-------> indica método que age sobre o proprio user chamador.
//    |--+--|      *user é um receiver do metodo. Ponteiro
func (u *user) changeEmail(email string) { //Go faz: changeEmail(&u, email)
	u.email = email
}

func main1() {
	bill := user{"bill", "bill@mail.com"} //cria um valor do tipo user
	//(*bill).changeEmail() - Go ajusta a call para bater com o tipo do receiver
	bill.changeEmail("bill@gmail.com")
	bill.notify() //atua na sua cópia de bill

	joe := &user{"joe", "joe@email.com"} // cria um ponteiro para um user
	joe.changeEmail("joe@gmail.com")     //não precisa ajustar, ja é ponteiro
	joe.notify()
	// (*joe).notify() -- Go ajusta a chamada usando ponteiro, pois joe foi
	//declarado como ponteiro. Notify tem semantica de valor, então vai copiar
	//o valor que *joe aponta. Se desviou da semantica. BAAAAD
}
