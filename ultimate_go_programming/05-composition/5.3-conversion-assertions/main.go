package main

import "fmt"

type Mover interface {
	Move()
}

type Locker interface {
	Lock()
	Unlock()
}

type MoveLocker interface {
	Mover
	Locker
}

type bike struct{}

//satisfaz à interface fmt.Stringer
func (b *bike) String() string {
	return "I'm a bike"
}

func (bike) Move() {
	fmt.Println("Moving the bike")
}

func (bike) Lock() {
	fmt.Println("Locking the bike")
}

func (bike) Unlock() {
	fmt.Println("Unlocking the bike")
}

func main() {
	var ml MoveLocker
	var m Mover

	ml = bike{} //ml aponta para o valor do tipo concreto bike

	// se o valor apontado por ml também satisfaz o tipo m(Mover), m também ira
	//apontar para bike
	//MoveLocker satisfaz Move
	m = ml

	/*
		var m2 Mover
		m2 = bike{}
		ml = m2 //erro: Mover não satisfaz MoveLocker
	*/

	//type assertio
	//se "dentro" de ml tiver um tipo bike, a atribuição funciona, se não tiver,
	//panic
	b := ml.(bike)
	m = b

	b, ok := m.(bike) // mesma coisa, mas se não funcionar, b = bike zero value

	fmt.Println(ok) //true

	//type assertions são usadas, por exemplo para usar implementações de
	//de metodos num determinado tipo se existir, e usar metodo padrão se não.
	//o fmt.Println recebe argumentos que, se satisfazem a interface Stringer,
	//serão impressos como definidos no proprio tipo, se não satisfizer, será
	//impresso de forma padronizada

	//o metodo String() em bike foi implementado só com ponteiros, então passar
	//o valor de bike causa uma impressão padrão.
	fmt.Println(b)
	fmt.Println(&b)
}
