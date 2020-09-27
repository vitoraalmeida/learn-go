package main

/*
interface definida no arquivo anterior
type notifier interface  {
	notify()
}

tipo definido no primeiro arquivo
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending User Email to %s<%s>\n", u.name, u.email)
}
*/

type admin struct {
	person user
	level  string
}

// emdedded type - cria uma relação entre o tipo interno e externo
// o tipo externo agora possui os métodos que o interno tem
//está imbutindo um valor do tipo user diretamente em admin2.
//admin2 é promovido (inner type promotion) para poder usar os métodos
//NÃO QUER DIZER QUE admin2 AGORA É UM user. NÃO TEM HERANÇA
type admin2 struct {
	user  //embedded type
	level string
}

/*
se o tipo também implementar, é o seu método que sera usado
func (a *admin2) notify() {
	fmt.Printf("Sending admin Email to %s<%s>\n", u.name, u.email)
}
*/

func main6() {
	ad := admin{
		person: user{
			name:  "john",
			email: "john@mail.com",
		},
		level: "super",
	}
	//ad não possui o método notify, mas seu campo person possui
	//ad.notify() -> erro
	ad.person.notify()

	ad2 := admin2{
		user: user{
			name:  "john",
			email: "john@mail.com",
		},
		level: "super",
	}
	//acessando notify pelo valor user dentro de add2
	ad2.user.notify()
	//ad2 é do tipo admin2 que sofreu inner type promotion
	ad2.notify()

	//admin agora também satisfaz o contrato que o tipo
	//interno satisfaz.
	sendNotification(ad2)

}

/*
definida no arquivo anterior
func sendNotification(n notifier) {
	n.notify()
}
*/
