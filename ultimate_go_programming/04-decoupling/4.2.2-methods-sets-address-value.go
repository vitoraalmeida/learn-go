package main

/*
  Method Sets: descrevem quais métodos, baseado no value receiver, estão dispo-
			   níveis para tipos T e *T.

   | Tipo | métodos                 |
   +------+-------------------------+
   |   T  | value reveiver          |
   |  *T  | value e pointer reveiver|
*/

type notifier interface {
	notify()
}

/* tipo definido no primeiro arquivo
type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending User Email to %s<%s>\n", u.name, u.email)
}
*/

func main5() {
	u := user{name: "bill", email: "buill@email.com"}
	//user criado com semantica de valor, não tem acesso ao
	//método defindo para pointer receiver
	sendNotification(u) //deve-se passar o endereço, pois o método espera pointer

}

func sendNotification(n notifier) {
	n.notify()
}
